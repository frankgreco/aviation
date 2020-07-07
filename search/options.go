package search

import (
	"database/sql"
	"strings"
	//"github.com/frankgreco/aviation/utils/db"
)

type Filters struct {
	TailNumber *string
	Airline    *string
}

type Config struct {
	Filters
	Database *sql.DB
}

type Option func(*Config)

func TailNumber(tailNumber string) Option {
	return Option(func(cfg *Config) {
		if strings.HasPrefix(strings.ToUpper(tailNumber), "N") {
			tailNumber = tailNumber[1:]
		}
		cfg.Filters.TailNumber = &tailNumber
	})
}

func Airline(airline string) Option {
	return Option(func(cfg *Config) {
		cfg.Filters.Airline = &airline
	})
}

func Database(db *sql.DB) Option {
	return Option(func(cfg *Config) {
		cfg.Database = db
	})
}

func (cfg Config) Options() []Option {
	ops := []Option{}

	if cfg.Filters.TailNumber != nil {
		ops = append(ops, TailNumber(*cfg.Filters.TailNumber))
	}

	return ops
}
