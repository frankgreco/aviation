package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/frankgreco/aviation/api"
	"github.com/frankgreco/aviation/utils/db"
)

var (
	psq   = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	dbURL = "postgres:///aviation"
)

func main() {
	dbase, err := db.New(dbURL, log.New().WithFields(log.Fields{
		"app": "database",
	}))
	if err != nil {
		panic(err)
	}
	defer dbase.Close()

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()

		if err := process(
			"registrations",
			"master.txt",
			dbase,
			func(items []api.RowBuilder) squirrel.SelectBuilder {
				return psq.
					Select("sub.id").
					From("aviation.registration r").
					JoinClause(fmt.Sprintf("RIGHT JOIN (VALUES %s) AS sub (id) ON sub.id = r.unique_id", toValues(items))).
					Where(squirrel.Eq{
						"r.unique_id": nil,
					}).
					OrderBy("sub.id ASC")
			},
			psq.Insert("aviation.registration").Columns(strings.Join((api.Registration{}).Columns(), ", ")),
			func(data string) api.RowBuilder {
				return api.NewRegistration(data)
			},
		); err != nil {
			log.WithFields(log.Fields{
				"resource": "registrations",
				"error":    err.Error(),
			}).Error("failed to process resource")
		}

	}()

	go func() {
		defer wg.Done()

		if err := process(
			"aircraft",
			"aircraft.txt",
			dbase,
			func(items []api.RowBuilder) squirrel.SelectBuilder {
				return psq.
					Select("sub.manufacturer || '-' ||  sub.model || '-' || sub.series as id").
					From("aviation.aircraft a").
					JoinClause(fmt.Sprintf("NATURAL RIGHT JOIN (VALUES %s) AS sub (manufacturer, model, series)", toValues(items))).
					Where(squirrel.Eq{
						"a.manufactuer_name": nil,
					}).
					OrderBy("id ASC")
			},
			psq.Insert("aviation.aircraft").Columns(strings.Join((api.Aircraft{}).Columns(), ", ")),
			func(data string) api.RowBuilder {
				return api.NewAircraft(data)
			},
		); err != nil {
			log.WithFields(log.Fields{
				"resource": "aircraft",
				"error":    err.Error(),
			}).Error("failed to process resource")
		}
	}()

	wg.Wait()
}

func process(resource, fileLocation string, dbase *db.DB, diffQuery func([]api.RowBuilder) squirrel.SelectBuilder, insertQuery squirrel.InsertBuilder, f func(string) api.RowBuilder) error {
	file, err := os.Open(fileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	items := unmarshal(file, f)

	_, err = dbase.QueryRowsTx(context.Background(), nil, buildQueries(
		resource,
		150,
		insertQuery,
		new(resource, items, diffQuery(items), dbase),
	)...)
	return err
}

func new(resource string, existing []api.RowBuilder, query squirrel.SelectBuilder, dbase *db.DB) []api.RowBuilder {
	netIDs := []string{}

	if _, err := dbase.QueryRowsTx(context.Background(), nil, db.QueryScan{
		Name:  fmt.Sprintf("get all %s not currently in the database", resource),
		Query: query,
		Callback: db.ScanFunc(func(rows *sqlx.Rows) (arr []interface{}, err error) {
			for {
				var id string
				if err = rows.Scan(&id); err == sql.ErrNoRows {
					return nil, nil
				}
				if err != nil {
					return
				}
				netIDs = append(netIDs, id)
				if !rows.Next() {
					break
				}
			}
			return
		}),
	}); err != nil && err != db.ErrNotFound {
		dbase.Close()
		panic(err)
	}

	netRb := make([]api.RowBuilder, len(netIDs))
	j := 0
	for i := 0; i < len(existing); i++ {
		if j == len(netIDs) {
			break
		}
		if existing[i].ID() < netIDs[j] {
			continue
		}
		netRb[j] = existing[i]
		j++
	}

	if len(netIDs) == 0 {
		log.Info(fmt.Sprintf("no new %s", resource))
	} else {
		log.WithFields(log.Fields{
			resource: strings.Join(netIDs, ", "),
		}).Info(fmt.Sprintf("new %s", resource))
	}

	return netRb
}

func toValues(items []api.RowBuilder) string {
	if items == nil || len(items) < 1 {
		return "()"
	}

	values := []string{}

	for _, item := range items {
		values = append(values, item.DBValue())
	}

	return strings.Join(values, ", ")
}

func unmarshal(r io.Reader, f func(string) api.RowBuilder) []api.RowBuilder {
	scanner := bufio.NewScanner(r)
	// ditch the header row
	scanner.Scan()
	data := []api.RowBuilder{}
	for scanner.Scan() {
		data = append(data, f(scanner.Text()))
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID() < data[j].ID()
	})
	return data
}

func buildQueries(resource string, batchSize int, base squirrel.InsertBuilder, items []api.RowBuilder) []db.QueryScan {
	queries := make([]db.QueryScan, int(math.Ceil(float64(len(items)/batchSize)))+1)
	remaining := len(items)
	i, j := 0, 0

	for remaining > 0 {
		size := int(math.Min(float64(batchSize), float64(remaining)))
		query := base

		for _, item := range items[i : i+size] {
			query = query.Values(item.Values()...)
		}
		queries[j] = db.QueryScan{
			Name:  fmt.Sprintf("inserting %s [%d,%d)", resource, i, i+size),
			Query: query,
		}

		remaining -= size
		i += size
		j++
	}

	return queries
}
