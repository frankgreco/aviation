package api

type RowBuilder interface {
	Columns() []string
	Values() []interface{}
	ID() string
	DBValue() string
}

type Unmarshaler interface {
	Unmarshal(string) RowBuilder
}
