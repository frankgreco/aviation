package api

import "strings"

type Aircraft struct {
	Manufacturer     string
	Model            string
	Series           string
	ManufacturerName string
	ModelName        string
	NumEngines       string
	NumSeats         string
	Weight           string
	CruisingSpeed    string
}

func (ac Aircraft) Columns() []string {
	return []string{
		"manufacturer",
		"model",
		"series",
		"manufactuer_name",
		"model_name",
		"num_engines",
		"num_seats",
		"weight",
		"cruising_speed",
	}
}

func (ac Aircraft) Values() []interface{} {
	return []interface{}{
		ac.Manufacturer,
		ac.Model,
		ac.Series,
		ac.ManufacturerName,
		ac.ModelName,
		ac.NumEngines,
		ac.NumSeats,
		ac.Weight,
		ac.CruisingSpeed,
	}
}

func (reg *Aircraft) Unmarshal(data string) RowBuilder {
	return NewAircraft(data)
}

func NewAircraft(data string) Aircraft {
	return Aircraft{
		Manufacturer:     strings.TrimSpace(data[0:3]),
		Model:            strings.TrimSpace(data[3:5]),
		Series:           strings.TrimSpace(data[5:7]),
		ManufacturerName: strings.TrimSpace(data[8:38]),
		ModelName:        strings.TrimSpace(data[39:59]),
		NumEngines:       strings.TrimSpace(data[69:71]),
		NumSeats:         strings.TrimSpace(data[72:75]),
		Weight:           strings.TrimSpace(data[76:83]),
		CruisingSpeed:    strings.TrimSpace(data[84:88]),
	}
}
