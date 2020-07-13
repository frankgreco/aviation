package search

import (
	"database/sql"
	"strings"
)

type Filters struct {
	TailNumber *string
	Airline    *string
	Make       *string
	Model      *string
}

type Config struct {
	Filters
	Database *sql.DB
	Limit    *int
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

func Make(make string) Option {
	return Option(func(cfg *Config) {
		make = strings.ToUpper(make)
		cfg.Filters.Make = &make
	})
}

func Model(model string) Option {
	return Option(func(cfg *Config) {
		model = strings.ToUpper(model)
		cfg.Filters.Model = &model
	})
}

func Database(db *sql.DB) Option {
	return Option(func(cfg *Config) {
		cfg.Database = db
	})
}

func Limit(limit int) Option {
	return Option(func(cfg *Config) {
		if limit > 0 {
			cfg.Limit = &limit
		}
	})
}

func (cfg Config) Options() []Option {
	ops := []Option{}

	if cfg.Filters.TailNumber != nil {
		ops = append(ops, TailNumber(*cfg.Filters.TailNumber))
	}

	return ops
}
