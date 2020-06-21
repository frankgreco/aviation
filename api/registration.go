package api

import "strings"

type Address struct {
	Street1 string
	Street2 string
	City    string
	State   string
	ZipCode string
	Region  string
	County  string
	Country string
}

type Certification struct {
	IssueDate                       string
	AirworthinessClassificationCode string
}

type Registration struct {
	Id                       string
	SerialNumber             string
	AircraftManufacturerCode string
	AircraftModelCode        string
	AircraftSeriesCode       string
	EngineManufacturerCode   string
	EngineModelCode          string
	YearManufactured         string
	RegistrantType           string
	RegistrantName           string
	Address                  Address
	LastActivityDate         string
	Certification            Certification
	AircraftType             string
	EngineType               string
	StatusCode               string
	ModeSCode                string
	FractionalOwnership      string
	AirworthinessDate        string
	ExpirationDate           string
	UniqueID                 string
	KitManufacturer          string
	KitModel                 string
}

func (reg *Registration) Unmarshal(data string) RowBuilder {
	return NewRegistration(data)
}

func (ac Registration) Columns() []string {
	return []string{
		"unique_id",
		"id",
		"serial_number",
		"year_manufactured",
		"manufacturer",
		"model",
		"series",
	}
}

func (reg Registration) Values() []interface{} {
	return []interface{}{
		reg.UniqueID,
		reg.Id,
		reg.SerialNumber,
		reg.YearManufactured,
		reg.AircraftManufacturerCode,
		reg.AircraftModelCode,
		reg.AircraftSeriesCode,
	}
}

func NewRegistration(data string) Registration {
	return Registration{
		Id:                       strings.TrimSpace(data[0:5]),
		SerialNumber:             strings.TrimSpace(data[6:36]),
		AircraftManufacturerCode: strings.TrimSpace(data[37:40]),
		AircraftModelCode:        strings.TrimSpace(data[40:42]),
		AircraftSeriesCode:       strings.TrimSpace(data[42:44]),
		EngineManufacturerCode:   strings.TrimSpace(data[45:48]),
		EngineModelCode:          strings.TrimSpace(data[48:50]),
		YearManufactured:         strings.TrimSpace(data[51:55]),
		RegistrantType:           strings.TrimSpace(data[56:57]),
		RegistrantName:           strings.TrimSpace(data[58:108]),
		Address: Address{
			Street1: strings.TrimSpace(data[109:142]),
			Street2: strings.TrimSpace(data[143:176]),
			City:    strings.TrimSpace(data[177:195]),
			State:   strings.TrimSpace(data[196:198]),
			ZipCode: strings.TrimSpace(data[199:209]),
			Region:  strings.TrimSpace(data[210:211]),
			County:  strings.TrimSpace(data[212:215]),
			Country: strings.TrimSpace(data[216:218]),
		},
		LastActivityDate: strings.TrimSpace(data[219:227]),
		Certification: Certification{
			IssueDate:                       strings.TrimSpace(data[228:236]),
			AirworthinessClassificationCode: strings.TrimSpace(data[238:238]),
		},
		AircraftType:        strings.TrimSpace(data[248:248]),
		EngineType:          strings.TrimSpace(data[250:252]),
		StatusCode:          strings.TrimSpace(data[253:255]),
		ModeSCode:           strings.TrimSpace(data[256:264]),
		FractionalOwnership: strings.TrimSpace(data[265:266]),
		AirworthinessDate:   strings.TrimSpace(data[267:275]),
		ExpirationDate:      strings.TrimSpace(data[531:539]),
		UniqueID:            strings.TrimSpace(data[540:548]),
		KitManufacturer:     strings.TrimSpace(data[549:579]),
		KitModel:            strings.TrimSpace(data[580:600]),
	}
}
