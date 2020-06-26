package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Engine struct {
	Manufacturer     string
	Model            string
	ManufacturerName string
	ModelName        string
	Type             EngineType
	Horsepower       int
	Thrust           int
}

func (e Engine) Columns() []string {
	return []string{
		"id",
		"make",
		"model",
		"type",
		"horsepower",
		"thrust",
		"created",
	}
}

func (e Engine) Values(now time.Time) []interface{} {
	return []interface{}{
		fmt.Sprintf("%s-%s", e.Manufacturer, e.Model),
		e.ManufacturerName,
		e.ModelName,
		e.Type.String(),
		e.Horsepower,
		e.Thrust,
		now.UTC(),
	}
}

func (e Engine) ID() string {
	return fmt.Sprintf("%s-%s", e.Manufacturer, e.Model)
}

func NewEngine(data string) Engine {
	engineType, ok := EngineType_name[strings.TrimSpace(data[31:33])]
	if !ok {
		engineType = EngineType_Unknown
	}

	var horsepower int
	{
		i, err := strconv.Atoi(strings.TrimSpace(data[34:39]))
		if err != nil {
			horsepower = 0
		}
		horsepower = i
	}

	var thrust int
	{
		i, err := strconv.Atoi(strings.TrimSpace(data[40:46]))
		if err != nil {
			thrust = 0
		}
		thrust = i
	}

	return Engine{
		Manufacturer:     strings.TrimSpace(data[0:3]),
		Model:            strings.TrimSpace(data[3:5]),
		ManufacturerName: strings.TrimSpace(data[6:16]),
		ModelName:        strings.TrimSpace(data[17:30]),
		Type:             engineType,
		Horsepower:       horsepower,
		Thrust:           thrust,
	}
}
