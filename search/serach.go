package search

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/frankgreco/aviation/api"
)

var (
	psq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func New(ops ...Option) ([]api.SearchResult, error) {
	cfg := &Config{}
	for _, op := range ops {
		op(cfg)
	}

	items := []api.SearchResult{}

	rows, err := cfg.Database.QueryContext(
		context.Background(),
		"SELECT r.tail_number, r.year_manufactured, a.make, a.model, a.num_engines, a.num_seats, a.cruising_speed FROM aviation.registration r JOIN aviation.aircraft a ON r.aircraft_id = a.id WHERE r.tail_number LIKE $1 LIMIT $2",
		strings.ToUpper(fmt.Sprintf("%s%%", *cfg.Filters.TailNumber)),
		cfg.Limit,
	)
	if err != nil {
		return nil, api.WrapErr(err, "could not execute query")
	}
	defer rows.Close()
	hasNext := rows.Next()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if hasNext {
		for {
			var result api.SearchResult
			var cruisingSpeed, numEngines, numSeats sql.NullInt64
			if err = rows.Scan(
				&result.NNumber,
				&result.YearManufactured,
				&result.Make,
				&result.Model,
				&numEngines,
				&numSeats,
				&cruisingSpeed,
			); err == sql.ErrNoRows {
				return nil, nil
			}
			if err != nil {
				return nil, api.WrapErr(err, "could not process row")
			}
			if cruisingSpeed.Valid {
				result.CruisingSpeed = int(cruisingSpeed.Int64)
			}

			if numEngines.Valid {
				result.NumEngines = int(numEngines.Int64)
			}

			if numSeats.Valid {
				result.NumSeats = int(numSeats.Int64)
			}
			items = append(items, result)
			if !rows.Next() {
				break
			}
		}
	}

	return items, nil
}
