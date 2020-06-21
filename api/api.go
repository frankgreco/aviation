package api

type RowBuilder interface {
	Columns() []string
	Values() []interface{}
}

type Unmarshaler interface {
	Unmarshal(string) RowBuilder
}
