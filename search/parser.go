package search

import "strings"

// tail_number: N949AN, airline: AA, make: Boeing

func Parse(s string) []Option {
	return []Option{
		TailNumber(strings.TrimSpace(strings.SplitAfter(s, ":")[1])),
	}
}
