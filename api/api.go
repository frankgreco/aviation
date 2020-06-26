package api

import "time"

type RowBuilder interface {
	Columns() []string
	Values(time.Time) []interface{}
	ID() string
}

type Unmarshaler interface {
	Unmarshal(string) RowBuilder
}
