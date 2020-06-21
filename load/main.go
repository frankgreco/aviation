package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/Masterminds/squirrel"

	"github.com/frankgreco/aviation/api"
	"github.com/frankgreco/aviation/utils/db"
)

var (
	psq   = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	dbURL = "postgres:///aviation"
)

func main() {
	registrationFile, err := os.Open("master.txt")
	if err != nil {
		panic(err)
	}
	defer registrationFile.Close()

	aircraftFile, err := os.Open("aircraft.txt")
	if err != nil {
		panic(err)
	}
	defer aircraftFile.Close()

	var registrations, aircraft []api.RowBuilder
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		registrations = unmarshal(registrationFile, new(api.Registration))
	}()
	go func() {
		defer wg.Done()
		aircraft = unmarshal(aircraftFile, new(api.Aircraft))
	}()

	wg.Wait()

	dbase, err := db.New(dbURL, nil)
	if err != nil {
		panic(err)
	}
	defer dbase.Close()

	wg.Add(2)
	go func() {
		defer wg.Done()
		if _, err := dbase.QueryRowsTx(context.Background(), nil, buildQueries(
			150,
			psq.Insert("aviation.aircraft").Columns(strings.Join((api.Aircraft{}).Columns(), ", ")),
			aircraft,
		)...); err != nil {
			dbase.Close()
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := dbase.QueryRowsTx(context.Background(), nil, buildQueries(
			150,
			psq.Insert("aviation.registration").Columns(strings.Join((api.Registration{}).Columns(), ", ")),
			registrations,
		)...); err != nil {
			dbase.Close()
			panic(err)
		}
	}()

	wg.Wait()
}

func unmarshal(r io.Reader, rb api.Unmarshaler) []api.RowBuilder {
	scanner := bufio.NewScanner(r)
	// ditch the header row
	scanner.Scan()
	data := []api.RowBuilder{}
	for scanner.Scan() {
		data = append(data, rb.Unmarshal(scanner.Text()))
	}
	return data
}

func buildQueries(batchSize int, base squirrel.InsertBuilder, items []api.RowBuilder) []db.QueryScan {
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
			Name:  fmt.Sprintf("inserting registrations [%d,%d)", i, i+size),
			Query: query,
		}

		remaining -= size
		i += size
		j++
	}

	return queries
}
