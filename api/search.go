package api

type SearchRequest struct {
	Query string
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
	// Registrant         Registrant     `json:"registrant"`
	// Classification     Classification `json:"classificaiotn"`
	// ApprovedOperations []string       `json:"approved_operations"`
	// Type               AircraftType   `json:"type"`
	Make  string `json:"make,omitempty"`
	Model string `json:""model,omitempty`
	// EngineType         EngineType     `json:"engine_type"`
	NumEngines int `json:"num_engines,omitempty"`
	NumSeats   int `json:"num_seats,omitempty"`
	// Weight        int `json:"weight,omitempty"`
	CruisingSpeed int `json:"cruising_speed,omitempty"`
}

func (resp *SearchResponse) GetError() *Error {
	return resp.Error
}
