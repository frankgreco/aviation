package main

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/frankgreco/aviation"
	"github.com/frankgreco/aviation/api"
	"github.com/frankgreco/aviation/internal/db"
)

var (
	psq   = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	dbURL = fmt.Sprintf("postgres://%s:%s@%s:5432/aviation", os.Getenv("RDS_USERNAME"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_ENDPOINT"))
)

func main() {
	lambda.Start(do)

	// uncomment this code to do a quick test locally without using sam
	//
	// if err := do(context.Background()); err != nil {
	// 	log.WithFields(log.Fields{
	// 		"errors": err.Error(),
	// 	}).Info("routine did not run successfully")
	// 	os.Exit(1)
	// }
}

func do(ctx context.Context) error {
	log.Info("creating aws session")
	sesh, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(aviation.AwsRegion),
		},
	})
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("could not create aws session")
		return err
	}
	log.Info("successfully created aws session")

	runnable := db.Prepare(&db.Options{
		ConnectionString: dbURL,
		Logger:           log.WithFields(log.Fields{}),
	})
	defer runnable.Close(nil)

	dbase := runnable.(*db.DB)

	log.WithFields(log.Fields{
		"time": aviation.Now,
	}).Info("now")

	var wg sync.WaitGroup
	wg.Add(3)

	errChan := make(chan error)

	go func() {
		defer wg.Done()

		if err := process(
			"registrations",
			"MASTER.txt",
			aviation.Now,
			dbase,
			sesh,
			func(items []api.RowBuilder) squirrel.SelectBuilder {
				return psq.
					Select("sub.id").
					From("aviation.registration r").
					JoinClause(fmt.Sprintf("RIGHT JOIN (VALUES %s) AS sub (id) ON sub.id = r.id", toValues(items))).
					Where(squirrel.Eq{
						"r.id": nil,
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
			errChan <- err
		}
		errChan <- nil
	}()

	go func() {
		defer wg.Done()

		if err := process(
			"aircraft",
			"ACFTREF.txt",
			aviation.Now,
			dbase,
			sesh,
			func(items []api.RowBuilder) squirrel.SelectBuilder {
				return psq.
					Select("id").
					From("aviation.aircraft a").
					JoinClause(fmt.Sprintf("NATURAL RIGHT JOIN (VALUES %s) AS sub (id)", toValues(items))).
					Where(squirrel.Eq{
						"a.id": nil,
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
			errChan <- err
		}
		errChan <- nil
	}()

	go func() {
		defer wg.Done()

		if err := process(
			"engines",
			"ENGINE.txt",
			aviation.Now,
			dbase,
			sesh,
			func(items []api.RowBuilder) squirrel.SelectBuilder {
				return psq.
					Select("id").
					From("aviation.engine e").
					JoinClause(fmt.Sprintf("NATURAL RIGHT JOIN (VALUES %s) AS sub (id)", toValues(items))).
					Where(squirrel.Eq{
						"e.id": nil,
					}).
					OrderBy("id ASC")
			},
			psq.Insert("aviation.engine").Columns(strings.Join((api.Engine{}).Columns(), ", ")),
			func(data string) api.RowBuilder {
				return api.NewEngine(data)
			},
		); err != nil {
			log.WithFields(log.Fields{
				"resource": "engine",
				"error":    err.Error(),
			}).Error("failed to process resource")
			errChan <- err
		}
		errChan <- nil
	}()

	e := []string{}
	go func() {
		for range errChan {
			if err := <-errChan; err != nil {
				e = append(e, err.Error())
			}
		}
	}()

	wg.Wait()
	close(errChan)

	if len(e) == 0 {
		return nil
	}
	return errors.New(strings.Join(e, ", "))
}

func getDataFromS3(file string, sesh *session.Session) (io.ReadCloser, error) {
	key := fmt.Sprintf("%s/%s", aviation.Date, file)
	out, err := s3.New(sesh).GetObject(&s3.GetObjectInput{
		Bucket: aws.String(aviation.AwsS3BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"file": key,
			"err":  err.Error(),
		}).Error("could not get file from aws s3")
		return nil, err
	}

	return out.Body, nil
}

func process(resource, fileLocation string, now time.Time, dbase *db.DB, sesh *session.Session, diffQuery func([]api.RowBuilder) squirrel.SelectBuilder, insertQuery squirrel.InsertBuilder, f func(string) api.RowBuilder) error {
	data, err := getDataFromS3(fileLocation, sesh)
	if err != nil {
		return err
	}
	defer data.Close()

	items := unmarshal(data, f)

	n, err := new(resource, items, diffQuery(items), dbase)
	if err != nil {
		return err
	}

	_, err = dbase.QueryRowsTx(context.Background(), nil, buildQueries(
		resource,
		150,
		insertQuery,
		n,
		now,
	)...)
	return err
}

func new(resource string, existing []api.RowBuilder, query squirrel.SelectBuilder, dbase *db.DB) ([]api.RowBuilder, error) {
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
		// dbase.Close()
		return nil, err
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

	return netRb, nil
}

func toValues(items []api.RowBuilder) string {
	if items == nil || len(items) < 1 {
		return "()"
	}

	values := []string{}

	for _, item := range items {
		values = append(values, fmt.Sprintf("('%s')", item.ID()))
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

func buildQueries(resource string, batchSize int, base squirrel.InsertBuilder, items []api.RowBuilder, now time.Time) []db.QueryScan {
	queries := make([]db.QueryScan, int(math.Ceil(float64(len(items)/batchSize)))+1)
	remaining := len(items)
	i, j := 0, 0

	for remaining > 0 {
		size := int(math.Min(float64(batchSize), float64(remaining)))
		query := base

		for _, item := range items[i : i+size] {
			query = query.Values(item.Values(now)...)
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
