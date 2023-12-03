package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/frankgreco/aviation/internal/log"
	"github.com/frankgreco/aviation/internal/run"

	_ "github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type DB struct {
	*sqlx.DB
	Logger log.Logger
	err    error
	cancel context.CancelFunc
	ctx    context.Context
}

type Options struct {
	ConnectionString string
	Logger           log.Logger
}

func Prepare(opts *Options) run.Runnable {
	runnable := &DB{
		Logger: opts.Logger,
	}

	conn, err := sqlx.Open("pgx", opts.ConnectionString)
	if err != nil {
		runnable.Logger.Error(fmt.Sprintf("error opening database connection: %s", err.Error()))
		runnable.err = err
		return runnable
	}

	runnable.ctx, runnable.cancel = context.WithCancel(context.Background())
	runnable.DB = conn

	return runnable
}

func (p *DB) Run() error {
	if p.err != nil {
		err := fmt.Errorf("error starting databas: %s", p.err.Error())
		p.Logger.Error(err.Error())
		return err
	}

	// TODO: investigate why first ping always succeeds.
	p.Ping()

	if err := p.Ping(); err != nil {
		p.Logger.Error(fmt.Sprintf("error connecting to database: %s", err.Error()))
		return nil // We don't have to return the actual error here.
	}

	p.Logger.Info("successfully connected to database")

	<-p.ctx.Done()
	return nil
}

func (p *DB) Close(error) error {
	p.cancel()
	if err := p.DB.Close(); err != nil {
		err := fmt.Errorf("error closing databas: %s", p.err.Error())
		p.Logger.Error(err.Error())
		return err
	}
	p.Logger.Info("successfully closed connection to database")
	return nil
}

type ScanFunc func(*sqlx.Rows) ([]interface{}, error)

type QueryFunc func() (squirrel.Sqlizer, error)

type QueryScanConfig struct {
	// If true, if no rows are returned by a query, it will not be treated as a fatal error.
	// A common use case would be in a transaction when you don't want the ith query that would
	// return no rows to terminte the transaction.
	NoErrorIfNotFound bool
}

type QueryScan struct {
	Name      string           // Explicit name of query used for logging purposes.
	Query     squirrel.Sqlizer // An alternative to Query in cases where queries are dynamically constructed using closure.
	QueryFunc QueryFunc
	Callback  ScanFunc
	Config    QueryScanConfig
}

func (db *DB) QueryRows(ctx context.Context, scanFunc ScanFunc, query squirrel.Sqlizer) (out []interface{}, err error) {
	return db.QueryRowsTx(ctx, nil, QueryScan{
		Query:    query,
		Callback: scanFunc,
	})
}

// QueryRowsTx will execute one or more queries and envoke the provided called for each.
// If more than one query is provided, a transaction will be used. The callback is envoked
// after the first row has been "loaded". This means that the callback should end with a
// check to rows.Next() rather than begin.
func (db *DB) QueryRowsTx(ctx context.Context, existingTx *sqlx.Tx, queryScans ...QueryScan) (out []interface{}, err error) {

	// Perform input validation and silenty fail.
	if queryScans == nil || len(queryScans) < 1 {
		return nil, nil
	}

	// The method we envoke to execute the query lives on different receiver types
	// depending on whether or not we use a transaction.
	type maybeTx interface {
		QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	}

	// Default to no transaction.
	driver := struct {
		maybeTx
	}{db}

	// If transaction, update the driver and begin th transaction.
	if len(queryScans) > 1 && existingTx == nil {
		db.Logger.Debug("beginning transaction")
		tx, err := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			return nil, fmt.Errorf("%s: %s", err.Error(), "could not begin transaction")
		}
		defer tx.Rollback()
		driver.maybeTx = tx
	}

	if existingTx != nil {
		driver.maybeTx = existingTx
	}

	for i, qs := range queryScans {
		if qs.Name == "" {
			qs.Name = fmt.Sprintf("#%d", i)
		}
		db.Logger.Debug(fmt.Sprintf("processing query %s", qs.Name))
		// Compile the query.
		var providedQuery squirrel.Sqlizer
		{
			providedQuery = qs.Query
			if qs.QueryFunc != nil {
				q, err := qs.QueryFunc()
				if err != nil {
					db.Logger.Error(err.Error())
					return nil, err
				}
				providedQuery = q
			}
		}
		q, args, err := providedQuery.ToSql()
		db.Logger.Debug(fmt.Sprintf("query: %v, args: %v", q, args))
		if err != nil {
			msg := fmt.Sprintf("failed to generate sql for query %s", qs.Name)
			db.Logger.Error(msg)
			return nil, fmt.Errorf("%s: %s", err.Error(), msg)
		}
		// Execute the query.
		rows, err := driver.QueryxContext(ctx, q, args...)
		if err != nil {
			msg := fmt.Sprintf("failed to query %s", qs.Name)
			db.Logger.Error(msg)
			return nil, fmt.Errorf("%s: %s", err.Error(), msg)
		}
		defer rows.Close()
		// Load first row now so that we can check for the final batch of errors.
		hasNext := rows.Next()
		// Second place where errors for the above query can be returned.
		if err := rows.Err(); err != nil {
			msg := fmt.Sprintf("failed to query %s", qs.Name)
			db.Logger.Error(msg)
			return nil, fmt.Errorf("%s: %s", err.Error(), msg)
		}
		// The query was successful. This doesn't mean rows were returned.
		// We can't blindly return ErrNotFound as the query may not have requested rows.
		if !hasNext {
			// TODO: is this robust enough?
			if strings.HasPrefix(q, "SELECT") || strings.Contains(q, "RETURNING") {
				if qs.Config.NoErrorIfNotFound {
					rows.Close()
					continue
				} else {
					return nil, ErrNotFound
				}
			}
		}
		// If no callback was provided, continue to the next query.
		if qs.Callback != nil {
			// Envoke callback.
			out, err = qs.Callback(rows)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", err.Error(), "error scanning rows")
			}
		}
		rows.Close()
	}
	// If we used a transaction, attempt to commit.
	if len(queryScans) > 1 && existingTx == nil {
		db.Logger.Debug("committing transaction")
		if err := driver.maybeTx.(*sqlx.Tx).Commit(); err != nil {
			err = fmt.Errorf("%s: could not commit transaction", err.Error())
			db.Logger.Error(err)
			return nil, err
		}
	}
	return
}

func When(conj squirrel.Sqlizer, condition bool) squirrel.Sqlizer {
	if !condition {
		return squirrel.Eq{} // really doesn't matter what type it is, just as long as it implements sq.Sqlizer
	}
	return conj
}

func WithAlias(columns []string, alias string) string {
	if alias != "" {
		alias = alias + "."
	}
	return alias + strings.Join(columns, fmt.Sprintf(", %s", alias))
}

func NullStruct(s interface{}) interface{} {
	if reflect.ValueOf(s).IsNil() {
		return nil
	}
	return s
}

// NullString is a helper for inserting a string which can be null.
func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NullInt is a helper for inserting an int which can be null.
func NullInt(i int) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: int64(i),
		Valid: true,
	}
}

// NullBoolP is a helper for inserting a pointer to a bool which can be null.
func NullBoolP(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{
			Valid: true,
		}
	}
	return sql.NullBool{
		Bool:  *b,
		Valid: true,
	}
}
