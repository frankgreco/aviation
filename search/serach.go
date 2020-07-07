package search

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	//"github.com/Masterminds/squirrel"
	//"github.com/jmoiron/sqlx"

	"github.com/Masterminds/squirrel"
	"github.com/frankgreco/aviation/api"
	//"github.com/frankgreco/aviation/utils/db"
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

	//n := time.Now()
	rows, err := cfg.Database.QueryContext(context.Background(), "SELECT r.tail_number, r.year_manufactured, a.make, a.model, a.num_engines, a.num_seats, a.cruising_speed FROM aviation.registration r JOIN aviation.aircraft a ON r.aircraft_id = a.id WHERE r.tail_number LIKE $1", strings.ToUpper(fmt.Sprintf("%s%%", *cfg.Filters.TailNumber)))
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
				// &result.Registrant,
				// &result.Classification,
				// &result.ApprovedOperations,
				// &result.Type,
				&result.Make,
				&result.Model,
				// &result.EngineType,
				&numEngines,
				&numSeats,
				// &result.Weight,
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

	// query := psq.Select("a.make, a.model, r.tail_number").
	// 	From("aviation.registration r").
	// 	Join("aviation.aircraft a ON r.aircraft_id = a.id").
	// 	Where(db.When(squirrel.Like{
	// 		"r.tail_number": strings.ToUpper(fmt.Sprintf("%s%%", *cfg.Filters.TailNumber)),
	// 	}, cfg.Filters.TailNumber != nil))

	// _, details, err := cfg.Database.QueryRowsTxDetails(context.Background(), nil, db.QueryScan{
	// 	Name:  "search",
	// 	Query: query,
	// 	Callback: db.ScanFunc(func(rows *sqlx.Rows) (arr []interface{}, err error) {
	// 		for {
	// 			var result Result
	// 			if err = rows.Scan(&result.Aircraft.Manufacturer, &result.Aircraft.Model, &result.Registration.Id); err == sql.ErrNoRows {
	// 				return nil, nil
	// 			}
	// 			if err != nil {
	// 				return
	// 			}
	// 			results.Results = append(results.Results, result)
	// 			if !rows.Next() {
	// 				break
	// 			}
	// 		}
	// 		return
	// 	}),
	// })
	// if err != nil && err != db.ErrNotFound {
	// 	return nil, err
	// }

	// results.Duration = time.Now().Sub(n).String()

	return items, nil
}
