package search

import "strings"

// Parse unmarshales a query string into a list of options
// Example: tail number="N959AN" AND make="Boeing" AND model="777" AND airline="American"
func Parse(s string) []Option {
	ops := []Option{}

	filters := strings.Split(s, " AND ")
	for _, f := range filters {
		filter := strings.Split(f, "=")
		if len(filter) != 2 {
			continue
		}
		switch filter[0] {
		case "tail number":
			ops = append(ops, TailNumber(unwrapQuotes(filter[1])))
		case "airline":
			ops = append(ops, Airline(unwrapQuotes(filter[1])))
		case "make":
			ops = append(ops, Make(unwrapQuotes(filter[1])))
		case "model":
			ops = append(ops, Model(unwrapQuotes(filter[1])))
		}
	}

	return ops
}

func unwrapQuotes(in string) (out string) {
	out = in
	if len(out) > 0 && out[0] == '"' {
		out = out[1:]
	}
	if len(out) > 0 && out[len(out)-1] == '"' {
		out = out[:len(out)-1]
	}
	return
}
