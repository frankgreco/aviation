package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib" // this import registers the pgx driver
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}

type DB struct {
	*sqlx.DB
	logger logger
}

type ScanFunc func(*sqlx.Rows) ([]interface{}, error)

type QueryFunc func() (sq.Sqlizer, error)

type QueryScan struct {
	Name      string     // Explicit name of query used for logging purposes.
	Query     sq.Sqlizer // An alternative to Query in cases where queries are dynamically constructed using closure.
	QueryFunc QueryFunc
	Callback  ScanFunc
}

type QueryDetails struct {
	Duration time.Duration `json:"duration"`
}

func New(url string, logger logger) (*DB, error) {
	return Open("pgx", url, logger)
}

func Open(driverName, url string, logger logger) (*DB, error) {
	sqlxDB, err := sqlx.Open(driverName, url)
	if err != nil {
		return nil, err
	}

	if logger == nil {
		logger = log.New()
	}

	db := &DB{sqlxDB, logger}
	go db.logStats()
	return db, nil
}

func (db *DB) logStats() {
	t := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-t.C:
			stats := db.DB.Stats()
			db.logger.Debug(
				"msg", "db stats",
				"idle", stats.Idle,
				"open conns", stats.OpenConnections,
				"in use", stats.InUse,
				"wait count", stats.WaitCount,
				"wait duration", stats.WaitDuration,
				"max idle closed", stats.MaxIdleClosed,
				"max lifetime closed", stats.MaxLifetimeClosed,
			)
		}
	}
}

func (db *DB) QueryRows(ctx context.Context, scanFunc ScanFunc, query sq.Sqlizer) (out []interface{}, err error) {
	return db.QueryRowsTx(ctx, nil, QueryScan{
		Query:    query,
		Callback: scanFunc,
	})
}

func (db *DB) QueryRowsTx(ctx context.Context, existingTx *sqlx.Tx, queryScans ...QueryScan) (out []interface{}, err error) {
	out, _, err = db.QueryRowsTxDetails(ctx, existingTx, queryScans...)
	return out, err
}

// QueryRowsTx will execute one or more queries and envoke the provided called for each.
// If more than one query is provided, a transaction will be used. The callback is envoked
// after the first row has been "loaded". This means that the callback should end with a
// check to rows.Next() rather than begin.
func (db *DB) QueryRowsTxDetails(ctx context.Context, existingTx *sqlx.Tx, queryScans ...QueryScan) (out []interface{}, details QueryDetails, err error) {

	// Perform input validation and silenty fail.
	if queryScans == nil || len(queryScans) < 1 {
		return nil, details, nil
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
		db.logger.Debug("beginning transaction")
		tx, err := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			return nil, details, fmt.Errorf("could not begin transaction: %s", err.Error())
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
		db.logger.Debug(fmt.Sprintf("processing query %s", qs.Name))
		// Compile the query.
		var providedQuery sq.Sqlizer
		{
			providedQuery = qs.Query
			if qs.QueryFunc != nil {
				q, err := qs.QueryFunc()
				if err != nil {
					db.logger.Error(err.Error())
					return nil, details, err
				}
				providedQuery = q
			}
		}
		if providedQuery == nil {
			db.logger.Warn(fmt.Sprintf("query %s was not provided", qs.Name))
			continue
		}
		q, args, err := providedQuery.ToSql()
		if err != nil {
			msg := fmt.Sprintf("failed to generate sql for query %s", qs.Name)
			db.logger.Error(msg)
			return nil, details, fmt.Errorf("%s: %s", msg, err.Error())
		}
		db.logger.Debug("query", q)
		// Execute the query.
		queryTime := time.Now()
		rows, err := driver.QueryxContext(ctx, q, args...)
		details.Duration = time.Duration(int64(details.Duration) + time.Now().Sub(queryTime).Nanoseconds())
		if err != nil {
			msg := fmt.Sprintf("failed to query %s", qs.Name)
			db.logger.Error(msg)
			return nil, details, fmt.Errorf("%s: %s", msg, err.Error())
		}
		defer rows.Close()
		// Load first row now so that we can check for the final batch of errors.
		hasNext := rows.Next()
		// Second place where errors for the above query can be returned.
		if err := rows.Err(); err != nil {
			msg := fmt.Sprintf("failed to query %s", qs.Name)
			db.logger.Error(msg)
			return nil, details, fmt.Errorf("%s: %s", msg, err.Error())
		}
		// The query was successful. This doesn't mean rows were returned.
		// We can't blindly return ErrNotFound as the query may not have requested rows.
		if !hasNext {
			// TODO: is this robust enough?
			if strings.HasPrefix(q, "SELECT") || strings.Contains(q, "RETURNING") {
				return nil, details, ErrNotFound
			}
		}
		// If no callback was provided, continue to the next query.
		if qs.Callback != nil {
			// Envoke callback.
			out, err = qs.Callback(rows)
			if err != nil {
				return nil, details, fmt.Errorf("error scanning rows: %s", err.Error())
			}
		}
		rows.Close()
	}
	// If we used a transaction, attempt to commit.
	if len(queryScans) > 1 && existingTx == nil {
		db.logger.Debug("committing transaction")
		if err := driver.maybeTx.(*sqlx.Tx).Commit(); err != nil {
			err = fmt.Errorf("could not commit transaction: %s", err.Error())
			db.logger.Error(err)
			return nil, details, err
		}
	}
	return
}

func When(conj sq.Sqlizer, condition bool) sq.Sqlizer {
	if !condition {
		return sq.Eq{} // really doesn't matter what type it is, just as long as it implements sq.Sqlizer
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

// NullInt is a helper for inserting a int which can be null.
func NullInt(s int) sql.NullInt64 {
	if s == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: int64(s),
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
