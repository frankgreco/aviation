package api

type SearchRequest struct {
	Query string
	Limit int
}

type SearchResponse struct {
	Items []SearchResult `json:"items,omitempty""`
	Error *Error         `json:"error,omitempty"`
}

// SearchResult provides a flattened view of the fields that will be used
// the the UI for presentation. It contains fields from both api.Registration
// and api.Aircraft.
type SearchResult struct {
	NNumber          string `json:"n_number,omitempty"`
	YearManufactured string `json:"year_manufactured,omitempty"`
	Make             string `json:"make,omitempty"`
	Model            string `json:""model,omitempty`
	NumEngines       int    `json:"num_engines,omitempty"`
	NumSeats         int    `json:"num_seats,omitempty"`
	CruisingSpeed    int    `json:"cruising_speed,omitempty"`
}

func (resp *SearchResponse) GetError() *Error {
	return resp.Error
}
