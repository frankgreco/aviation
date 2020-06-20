package main

import (
	"bufio"
	"context"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/Masterminds/squirrel"
	"github.com/confluentinc/cc-utils/pg"
)

var (
	psq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

type Registrations []Registration

func (rs *Registrations) Add(r Registration) {
	*rs = append(*rs, r)
}

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
	//ApprovedOperationCodes          string
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
	UniqueID                 string //
	KitManufacturer          string
	KitModel                 string
}

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

func main() {
	registrationFile, err := os.Open("master.txt")
	if err != nil {
		panic(err)
	}
	defer registrationFile.Close()

	aircraftFile, err := os.Open("aircraft.txt")
	if err != nil {
		panic(err)
	}
	defer aircraftFile.Close()

	regScanner := bufio.NewScanner(registrationFile)
	acScanner := bufio.NewScanner(aircraftFile)

	// ditch the header row
	regScanner.Scan()
	acScanner.Scan()

	registrations := new(Registrations)

	db, err := pg.New("postgres:///aviation", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for acScanner.Scan() {
		ac := NewAircraft(acScanner.Text())
		if _, err := db.QueryRowsTx(context.Background(), nil, pg.QueryScan{
			Name:  "inserting aircraft",
			Query: psq.Insert("aviation.aircraft").Columns("manufacturer, model, series, manufactuer_name, model_name, num_engines, num_seats, weight, cruising_speed").Values(ac.Manufacturer, ac.Model, ac.Series, ac.ManufacturerName, ac.ModelName, ac.NumEngines, ac.NumSeats, ac.Weight, ac.CruisingSpeed),
		}); err != nil {
			db.Close()
			panic(err)
		}
	}

	for regScanner.Scan() {
		reg := NewRegistration(regScanner.Text())
		registrations.Add(reg)
	}

	leftToProcess := len(*registrations)
	index := 0
	var wg sync.WaitGroup
	guard := make(chan struct{}, 100)

	for leftToProcess > 0 {
		size := math.Min(150.0, float64(leftToProcess))
		guard <- struct{}{}
		wg.Add(1)
		go func(start, end int, query squirrel.InsertBuilder) {
			for _, reg := range (*registrations)[start:end] {
				query = query.Values(reg.UniqueID, reg.Id, reg.SerialNumber, reg.YearManufactured, reg.AircraftManufacturerCode, reg.AircraftModelCode, reg.AircraftSeriesCode)
			}
			if _, err := db.QueryRowsTx(context.Background(), nil, pg.QueryScan{
				Name:  "inserting registrations",
				Query: query,
			}); err != nil {
				db.Close()
				panic(err)
			}
			<-guard
			wg.Done()
		}(index, index+int(size), psq.Insert("aviation.registration").Columns("unique_id, id, serial_number, year_manufactured, manufacturer, model, series"))
		leftToProcess -= int(size)
		index += int(size)
	}

	wg.Wait()

}
