package api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/frankgreco/aviation/utils/db"
)

type AircraftType int32
type EngineType int32
type AircraftCategoryCode int32
type BuilderCertificationCode int32

const (
	AircraftType_Glider AircraftType = iota
	AircraftType_Balloon
	AircraftType_Blimp
	AircraftType_FixedWingSingleEngine
	AircraftType_FixedWingMultiEngine
	AircraftType_Rotorcraft
	AircraftType_WeightShiftControl
	AircraftType_PoweredParachute
	AircraftType_Gyroplane
	AircraftType_HybridLift
	AircraftType_Other
	AircraftType_Unknown

	EngineType_None EngineType = iota
	EngineType_Reciprocating
	EngineType_TurboProp
	EngineType_TurboShaft
	EngineType_TurboJet
	EngineType_TurboFan
	EngineType_Ramjet
	EngineType_TwoCycle
	EngineType_FourCycle
	EngineType_Unknown
	EngineType_Electric
	EngineType_Rotary

	AircraftCategoryCode_Land AircraftCategoryCode = iota
	AircraftCategoryCode_Sea
	AircraftCategoryCode_Amphibian
	AircraftCategoryCode_Unknown

	BuilderCertificationCode_Certificated BuilderCertificationCode = iota
	BuilderCertificationCode_NotCertificated
	BuilderCertificationCode_LightSport
	BuilderCertificationCode_Unknown
)

var (
	AircraftType_name = map[string]AircraftType{
		"1": AircraftType_Glider,
		"2": AircraftType_Balloon,
		"3": AircraftType_Blimp,
		"4": AircraftType_FixedWingSingleEngine,
		"5": AircraftType_FixedWingMultiEngine,
		"6": AircraftType_Rotorcraft,
		"7": AircraftType_WeightShiftControl,
		"8": AircraftType_PoweredParachute,
		"9": AircraftType_Gyroplane,
		"H": AircraftType_HybridLift,
		"O": AircraftType_Other,
	}
	AircraftType_value = map[AircraftType]string{
		AircraftType_Glider:                "GLIDER",
		AircraftType_Balloon:               "BALOON",
		AircraftType_Blimp:                 "BLIMP",
		AircraftType_FixedWingSingleEngine: "FIXED_WING_SINGLE_ENGINE",
		AircraftType_FixedWingMultiEngine:  "FIXED_WING_MULTI_ENGINE",
		AircraftType_Rotorcraft:            "ROTORCRAFT",
		AircraftType_WeightShiftControl:    "WEIGHT_SHIFT_CONTROL",
		AircraftType_PoweredParachute:      "POWERED_PARACHUTE",
		AircraftType_Gyroplane:             "GYROPLANE",
		AircraftType_HybridLift:            "HYBRID_LIFT",
		AircraftType_Other:                 "OTHER",
		AircraftType_Unknown:               "UNKNOWN",
	}
	EngineType_name = map[string]EngineType{
		"0":  EngineType_None,
		"1":  EngineType_Reciprocating,
		"2":  EngineType_TurboProp,
		"3":  EngineType_TurboShaft,
		"4":  EngineType_TurboJet,
		"5":  EngineType_TurboFan,
		"6":  EngineType_Ramjet,
		"7":  EngineType_TwoCycle,
		"8":  EngineType_FourCycle,
		"9":  EngineType_Unknown,
		"10": EngineType_Electric,
		"11": EngineType_Rotary,
	}
	EngineType_value = map[EngineType]string{
		EngineType_None:          "NONE",
		EngineType_Reciprocating: "RECIPROCATING",
		EngineType_TurboProp:     "TURBO_PROP",
		EngineType_TurboShaft:    "TURBO_SHAFT",
		EngineType_TurboJet:      "TURBO_JET",
		EngineType_TurboFan:      "TURBO_FAN",
		EngineType_Ramjet:        "RAM_JET",
		EngineType_TwoCycle:      "TWO_CYCLE",
		EngineType_FourCycle:     "FOUR_CYCLE",
		EngineType_Unknown:       "UNKNOWN",
		EngineType_Electric:      "ELECTRIC",
		EngineType_Rotary:        "ROTARY",
	}
	AircraftCategoryCode_name = map[string]AircraftCategoryCode{
		"1": AircraftCategoryCode_Land,
		"2": AircraftCategoryCode_Sea,
		"3": AircraftCategoryCode_Amphibian,
	}
	AircraftCategoryCode_value = map[AircraftCategoryCode]string{
		AircraftCategoryCode_Land:      "LAND",
		AircraftCategoryCode_Sea:       "SEA",
		AircraftCategoryCode_Amphibian: "AMPHIBIAN",
		AircraftCategoryCode_Unknown:   "UNKNOWN",
	}
	BuilderCertificationCode_name = map[string]BuilderCertificationCode{
		"0": BuilderCertificationCode_Certificated,
		"1": BuilderCertificationCode_NotCertificated,
		"2": BuilderCertificationCode_LightSport,
	}
	BuilderCertificationCode_value = map[BuilderCertificationCode]string{
		BuilderCertificationCode_Certificated:    "CERTIFICATED",
		BuilderCertificationCode_NotCertificated: "NOT_CERTIFICATED",
		BuilderCertificationCode_LightSport:      "LIGHT_SPORT",
		BuilderCertificationCode_Unknown:         "UNKNOWN",
	}
)

type Aircraft struct {
	Manufacturer             string
	Model                    string
	Series                   string
	ManufacturerName         string
	ModelName                string
	Type                     AircraftType
	EngineType               EngineType
	CategoryCode             AircraftCategoryCode
	BuilderCertificationCode BuilderCertificationCode
	NumEngines               int
	NumSeats                 int
	Weight                   string
	CruisingSpeed            int
}

func (ac Aircraft) Columns() []string {
	return []string{
		"id",
		"make",
		"model",
		"type",
		"engine_type",
		"category_code",
		"builder_certification_code",
		"num_engines",
		"num_seats",
		"weight",
		"cruising_speed",
	}
}

func (ac Aircraft) Values() []interface{} {
	return []interface{}{
		fmt.Sprintf("%s-%s-%s", ac.Manufacturer, ac.Model, ac.Series),
		ac.ManufacturerName,
		ac.ModelName,
		ac.Type.String(),
		ac.EngineType.String(),
		ac.CategoryCode.String(),
		ac.BuilderCertificationCode.String(),
		db.NullInt(ac.NumEngines),
		db.NullInt(ac.NumSeats),
		ac.Weight,
		db.NullInt(ac.CruisingSpeed),
	}
}

func (ac AircraftType) String() string {
	val, ok := AircraftType_value[ac]
	if !ok {
		return "UNKNOWN"
	}
	return val
}

func (ac EngineType) String() string {
	val, ok := EngineType_value[ac]
	if !ok {
		return "UNKNOWN"
	}
	return val
}

func (ac AircraftCategoryCode) String() string {
	val, ok := AircraftCategoryCode_value[ac]
	if !ok {
		return "UNKNOWN"
	}
	return val
}

func (ac BuilderCertificationCode) String() string {
	val, ok := BuilderCertificationCode_value[ac]
	if !ok {
		return "UNKNOWN"
	}
	return val
}

func (reg Aircraft) ID() string {
	return fmt.Sprintf("%s-%s-%s", reg.Manufacturer, reg.Model, reg.Series)
}

func (reg Aircraft) DBValue() string {
	return fmt.Sprintf("('%s')", reg.ID())
}

func (reg *Aircraft) Unmarshal(data string) RowBuilder {
	return NewAircraft(data)
}

func NewAircraft(data string) Aircraft {
	aircraftType, ok := AircraftType_name[strings.TrimSpace(data[60:61])]
	if !ok {
		aircraftType = AircraftType_Unknown
	}

	engineType, ok := EngineType_name[strings.TrimSpace(data[62:64])]
	if !ok {
		engineType = EngineType_Unknown
	}

	aircraftCategoryCode, ok := AircraftCategoryCode_name[strings.TrimSpace(data[65:66])]
	if !ok {
		aircraftCategoryCode = AircraftCategoryCode_Unknown
	}

	builderCertificationCode, ok := BuilderCertificationCode_name[strings.TrimSpace(data[67:68])]
	if !ok {
		builderCertificationCode = BuilderCertificationCode_Unknown
	}

	var cruisingSpeed int
	{
		i, err := strconv.Atoi(strings.TrimSpace(data[84:88]))
		if err != nil {
			cruisingSpeed = 0
		}
		cruisingSpeed = i
	}

	var numSeats int
	{
		i, err := strconv.Atoi(strings.TrimSpace(data[72:75]))
		if err != nil {
			numSeats = 0
		}
		numSeats = i
	}

	var numEngines int
	{
		i, err := strconv.Atoi(strings.TrimSpace(data[69:71]))
		if err != nil {
			numEngines = 0
		}
		numEngines = i
	}

	return Aircraft{
		Manufacturer:             strings.TrimSpace(data[0:3]),
		Model:                    strings.TrimSpace(data[3:5]),
		Series:                   strings.TrimSpace(data[5:7]),
		ManufacturerName:         strings.TrimSpace(data[8:38]),
		ModelName:                strings.TrimSpace(data[39:59]),
		Type:                     aircraftType,
		EngineType:               engineType,
		CategoryCode:             aircraftCategoryCode,
		BuilderCertificationCode: builderCertificationCode,
		NumEngines:               numEngines,
		NumSeats:                 numSeats,
		Weight:                   strings.TrimSpace(data[76:83]),
		CruisingSpeed:            cruisingSpeed,
	}
}
