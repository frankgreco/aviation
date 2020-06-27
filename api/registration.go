package api

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Classification int32
type StatusCode int32
type ApprovedOperation int32
type RegistrantType int32

const (
	Classification_Standard Classification = iota
	Classification_Limited
	Classification_Restricted
	Classification_Experimental
	Classification_Provisional
	Classification_Multiple
	Classification_Primary
	Classification_SpecialFlightPermit
	Classification_LightSport
	Classification_Unknown

	StatusCode_TheTriennialAircraftRegistrationFormWasMailedAndHasNotBeenReceivedByThePostOffice StatusCode = iota
	StatusCode_ExpiredDealer
	StatusCode_TheCertificateOfAircraftRegistrationWasRevokedByEnforcementAction
	StatusCode_AircraftRegisteredToTheManufacturerUnderTheirDealerCertificate
	StatusCode_NonCitizenCorporationsWhichHaveNotReturnedTheirFlightHourReports
	StatusCode_RegistrationPending
	StatusCode_SecondTriennialAircraftRegistrationFormHasBeenMailedAndHasNotBeenReturnedByThePostOffice
	StatusCode_ValidRegistrationFromATrainee
	StatusCode_ValidRegistration
	StatusCode_CertificateOfRegistrationHasBeenDeemedIneffectiveOrInvalid
	StatusCode_EnforcementLetter
	StatusCode_PermanentReserved
	StatusCode_TriennialAircraftRegistrationFormWasReturnedByThePostOfficeAsUndeliverable
	StatusCode_NNumberAssignedButHasNotYetBeenRegistered
	StatusCode_NNumberAssignedAsANonTypeCertificatedAircraftButHasNotYetBeenRegistered
	StatusCode_NNumberAssignedAsImportButHasNotYetBeenRegistered
	StatusCode_ReservedNNumber
	StatusCode_AdministrativelyCanceled
	StatusCode_SaleReported
	StatusCode_ASecondAttemptHasBeenMadeAtMailingATriennialAircraftRegistrationFormToTheOwnerWithNoResponse
	StatusCode_CertificateOfRegistrationHasBeenRevoked
	StatusCode_NNumberAssignedHasNotBeenRegisteredAndIsPendingCancellation
	StatusCode_NNumberAssignedAsANonTypeCertificatedAmateurButHasNotBeenRegisteredThatIsPendingCancellation
	StatusCode_NNumberAssignedAsImportButHasNotBeenRegisteredThatIsPendingCancellation
	StatusCode_RegistrationExpired
	StatusCode_FirstNoticeForReRegistrationRenewal
	StatusCode_SecondNoticeForReRegistrationRenewal
	StatusCode_RegistrationExpiredPendingCancellation
	StatusCode_SaleReportedPendingCancellation
	StatusCode_SaleReportedCanceled
	StatusCode_RegistrationPendingPendingCancellation
	StatusCode_RegistrationPendingCanceled
	StatusCode_RevokedPendingCancellation
	StatusCode_RevokedCanceled
	StatusCode_ExpiredDealerPendingCancellation
	StatusCode_ThirdNoticeForReRegistrationRenewal
	StatusCode_FirstNoticeForRegistrationRenewal
	StatusCode_SecondNoticeForRegistrationRenewal
	StatusCode_ThirdNoticeForRegistrationRenewal
	StatusCode_Unknown

	ApprovedOperation_Normal ApprovedOperation = iota
	ApprovedOperation_Utility
	ApprovedOperation_Acrobatic
	ApprovedOperation_Transport
	ApprovedOperation_Glider
	ApprovedOperation_Balloon
	ApprovedOperation_Commuter
	ApprovedOperation_Other
	ApprovedOperation_AgricultureAndPestControl
	ApprovedOperation_AerialSurveying
	ApprovedOperation_AerialAdvertising
	ApprovedOperation_Forest
	ApprovedOperation_Patrolling
	ApprovedOperation_WeatherControl
	ApprovedOperation_CarriageOfCargo
	ApprovedOperation_ToShowComplianceWithFAR
	ApprovedOperation_ResearchAndDevelopment
	ApprovedOperation_AmateurBuilt
	ApprovedOperation_Exhibition
	ApprovedOperation_Racing
	ApprovedOperation_CrewTraining
	ApprovedOperation_MarketSurvey
	ApprovedOperation_OperatingKitBuiltAircraft
	ApprovedOperation_RegisteredPriorTo_01_31_08
	ApprovedOperation_OperatingLightSportKitBuilt
	ApprovedOperation_OperatingLightSportPreviouslyIssuedCertUnder_21_190
	ApprovedOperation_UnmannedAircraftResearchAndDevelopment
	ApprovedOperation_UnmannedAircraftMarketSurvey
	ApprovedOperation_UnmannedAircraftCrewTraining
	ApprovedOperation_UnmannedAircraftExhibition
	ApprovedOperation_UnmannedAircraftComplianceWithCFR
	ApprovedOperation_ClassOne
	ApprovedOperation_ClassTwo
	ApprovedOperation_Standard
	ApprovedOperation_Limited
	ApprovedOperation_Restricted
	ApprovedOperation_FerryFlightForRepairsAlterationsMaintenanceOrStorage
	ApprovedOperation_EvacuateFromAreaOfImpendingDanger
	ApprovedOperation_OperationInExcessOfMaximumCertificated
	ApprovedOperation_DeliveryOrExport
	ApprovedOperation_ProductionFlightTesting
	ApprovedOperation_CustomerDemo
	ApprovedOperation_Airplane
	ApprovedOperation_LighterThanAir
	ApprovedOperation_PowerParachute
	ApprovedOperation_WeightShiftControl

	RegistrantType_Individual RegistrantType = iota
	RegistrantType_Partnership
	RegistrantType_Corporation
	RegistrantType_CoOwned
	RegistrantType_Government
	RegistrantType_LLC
	RegistrantType_NonCitizenCorporation
	RegistrantType_NonCitizenCoOwned
	RegistrantType_Unknown
)

var (
	RegistrantType_name = map[string]RegistrantType{
		"1": RegistrantType_Individual,
		"2": RegistrantType_Partnership,
		"3": RegistrantType_Corporation,
		"4": RegistrantType_CoOwned,
		"5": RegistrantType_Government,
		"7": RegistrantType_LLC,
		"8": RegistrantType_NonCitizenCorporation,
		"9": RegistrantType_NonCitizenCoOwned,
	}
	RegistrantType_value = map[RegistrantType]string{
		RegistrantType_Individual:            "INDIVIDUAL",
		RegistrantType_Partnership:           "PARTNERSHIP",
		RegistrantType_Corporation:           "CORPORATION",
		RegistrantType_CoOwned:               "CO_OWNED",
		RegistrantType_Government:            "GOVERNMENT",
		RegistrantType_LLC:                   "LLC",
		RegistrantType_NonCitizenCorporation: "NON_CITIZEN_CORPORATION",
		RegistrantType_NonCitizenCoOwned:     "NON_CITIZEN_CO_OWNED",
		RegistrantType_Unknown:               "UNKNOWN",
	}
	Classification_name = map[string]Classification{
		"1": Classification_Standard,
		"2": Classification_Limited,
		"3": Classification_Restricted,
		"4": Classification_Experimental,
		"5": Classification_Provisional,
		"6": Classification_Multiple,
		"7": Classification_Primary,
		"8": Classification_SpecialFlightPermit,
		"9": Classification_LightSport,
	}
	Classification_value = map[Classification]string{
		Classification_Standard:            "STANDARD",
		Classification_Limited:             "LIMITED",
		Classification_Restricted:          "RESTRICTED",
		Classification_Experimental:        "EXPERIMENTAL",
		Classification_Provisional:         "PROVISIONAL",
		Classification_Multiple:            "MULTIPLE",
		Classification_Primary:             "PRIMARY",
		Classification_SpecialFlightPermit: "SPECIAL_FLIGHT_PERMIT",
		Classification_LightSport:          "LIGHT_SPORT",
		Classification_Unknown:             "UNKNOWN",
	}
	StatusCode_name = map[string]StatusCode{
		"A":  StatusCode_TheTriennialAircraftRegistrationFormWasMailedAndHasNotBeenReceivedByThePostOffice,
		"D":  StatusCode_ExpiredDealer,
		"E":  StatusCode_TheCertificateOfAircraftRegistrationWasRevokedByEnforcementAction,
		"M":  StatusCode_AircraftRegisteredToTheManufacturerUnderTheirDealerCertificate,
		"N":  StatusCode_NonCitizenCorporationsWhichHaveNotReturnedTheirFlightHourReports,
		"R":  StatusCode_RegistrationPending,
		"S":  StatusCode_SecondTriennialAircraftRegistrationFormHasBeenMailedAndHasNotBeenReturnedByThePostOffice,
		"T":  StatusCode_ValidRegistrationFromATrainee,
		"V":  StatusCode_ValidRegistration,
		"W":  StatusCode_CertificateOfRegistrationHasBeenDeemedIneffectiveOrInvalid,
		"X":  StatusCode_EnforcementLetter,
		"Z":  StatusCode_PermanentReserved,
		"1":  StatusCode_TriennialAircraftRegistrationFormWasReturnedByThePostOfficeAsUndeliverable,
		"2":  StatusCode_NNumberAssignedButHasNotYetBeenRegistered,
		"3":  StatusCode_NNumberAssignedAsANonTypeCertificatedAircraftButHasNotYetBeenRegistered,
		"4":  StatusCode_NNumberAssignedAsImportButHasNotYetBeenRegistered,
		"5":  StatusCode_ReservedNNumber,
		"6":  StatusCode_AdministrativelyCanceled,
		"7":  StatusCode_SaleReported,
		"8":  StatusCode_ASecondAttemptHasBeenMadeAtMailingATriennialAircraftRegistrationFormToTheOwnerWithNoResponse,
		"9":  StatusCode_CertificateOfRegistrationHasBeenRevoked,
		"10": StatusCode_NNumberAssignedHasNotBeenRegisteredAndIsPendingCancellation,
		"11": StatusCode_NNumberAssignedAsANonTypeCertificatedAmateurButHasNotBeenRegisteredThatIsPendingCancellation,
		"12": StatusCode_NNumberAssignedAsImportButHasNotBeenRegisteredThatIsPendingCancellation,
		"13": StatusCode_RegistrationExpired,
		"14": StatusCode_FirstNoticeForReRegistrationRenewal,
		"15": StatusCode_SecondNoticeForReRegistrationRenewal,
		"16": StatusCode_RegistrationExpiredPendingCancellation,
		"17": StatusCode_SaleReportedPendingCancellation,
		"18": StatusCode_SaleReportedCanceled,
		"19": StatusCode_RegistrationPendingPendingCancellation,
		"20": StatusCode_RegistrationPendingCanceled,
		"21": StatusCode_RevokedPendingCancellation,
		"22": StatusCode_RevokedCanceled,
		"23": StatusCode_ExpiredDealerPendingCancellation,
		"24": StatusCode_ThirdNoticeForReRegistrationRenewal,
		"25": StatusCode_FirstNoticeForRegistrationRenewal,
		"26": StatusCode_SecondNoticeForRegistrationRenewal,
		"27": StatusCode_RegistrationExpired,
		"28": StatusCode_ThirdNoticeForRegistrationRenewal,
		"29": StatusCode_RegistrationExpiredPendingCancellation,
	}
	StatusCode_value = map[StatusCode]string{
		StatusCode_TheTriennialAircraftRegistrationFormWasMailedAndHasNotBeenReceivedByThePostOffice: "THE_TRIENNIAL_AIRCRAFT_REGISTRATION_FORM_WAS_MAILED_AND_HAS_NOT_BEEN_RECEIVED_BY_THE_POST_OFFICE",
		StatusCode_ExpiredDealer: "EXPIRED_DEALER",
		StatusCode_TheCertificateOfAircraftRegistrationWasRevokedByEnforcementAction:                            "THE_CERTIFICATE_OF_AIRCRAFT_REGISTRATION_WAS_REVOKED_BY_ENFORCEMENT_ACTION",
		StatusCode_AircraftRegisteredToTheManufacturerUnderTheirDealerCertificate:                               "AIRCRAFT_REGISTERED_TO_THE_MANUFACTURER_UNDER_THEIR_DEALER_CERTIFICATE",
		StatusCode_NonCitizenCorporationsWhichHaveNotReturnedTheirFlightHourReports:                             "NON_CITIZEN_CORPORATIONS_WHICH_HAVE_NOT_RETURNED_THEIR_FLIGHT_HOUR_REPORTS",
		StatusCode_RegistrationPending:                                                                          "REGISTRATION_PENDING",
		StatusCode_SecondTriennialAircraftRegistrationFormHasBeenMailedAndHasNotBeenReturnedByThePostOffice:     "SECOND_TRIENNIAL_AIRCRAFT_REGISTRATION_FORM_HAS_BEEN_MAILED_AND_HAS_NOT_BEEN_RETURNED_BY_THE_POST_OFFICE",
		StatusCode_ValidRegistrationFromATrainee:                                                                "VALID_REGISTRATION_FROM_A_TRAINEE",
		StatusCode_ValidRegistration:                                                                            "VALID_REGISTRATION",
		StatusCode_CertificateOfRegistrationHasBeenDeemedIneffectiveOrInvalid:                                   "CERTIFICATE_OF_REGISTRATION_HAS_BEEN_DEEMED_INEFFECTIVE_OR_INVALID",
		StatusCode_EnforcementLetter:                                                                            "ENFORCEMENT_LETTER",
		StatusCode_PermanentReserved:                                                                            "PERMANENT_RESERVED",
		StatusCode_TriennialAircraftRegistrationFormWasReturnedByThePostOfficeAsUndeliverable:                   "TRIENNIAL_AIRCRAFT_REGISTRATION_FORM_WAS_RETURNED_BY_THE_POST_OFFICE_AS_UNDELIVERABLE",
		StatusCode_NNumberAssignedButHasNotYetBeenRegistered:                                                    "N_NUMBER_ASSIGNED_BUT_HAS_NOT_YET_BEEN_REGISTERED",
		StatusCode_NNumberAssignedAsANonTypeCertificatedAircraftButHasNotYetBeenRegistered:                      "N_NUMBER_ASSIGNED_AS_A_NON_TYPE_CERTIFICATED_AIRCRAFT_BUT_HAS_NOT_YET_BEEN_REGISTERED",
		StatusCode_NNumberAssignedAsImportButHasNotYetBeenRegistered:                                            "N_NUMBER_ASSIGNED_AS_IMPORT_BUT_HAS_NOT_YET_BEEN_REGISTERED",
		StatusCode_ReservedNNumber:                                                                              "RESERVED_NUMBER",
		StatusCode_AdministrativelyCanceled:                                                                     "ADMINISTRATIVELY_CANCELED",
		StatusCode_SaleReported:                                                                                 "SALE_REPORTED",
		StatusCode_ASecondAttemptHasBeenMadeAtMailingATriennialAircraftRegistrationFormToTheOwnerWithNoResponse: "A_SECOND_ATTEMPT_HAS_BEEN_MADE_AT_MAILING_A_TRIENNIAL_AIRCRAFT_REGISTRATION_FORM_TO_THE_OWNER_WITH_NO_RESPONSE",
		StatusCode_CertificateOfRegistrationHasBeenRevoked:                                                      "CERTIFICATE_OF_REGISTRATION_HAS_BEEN_REVOKED",
		StatusCode_NNumberAssignedHasNotBeenRegisteredAndIsPendingCancellation:                                  "N_NUMBER_ASSIGNED_HAS_NOT_BEEN_REGISTERED_AND_IS_PENDING_CANCELLATION",
		StatusCode_NNumberAssignedAsANonTypeCertificatedAmateurButHasNotBeenRegisteredThatIsPendingCancellation: "N_NUMBER_ASSIGNED_AS_A_NON_TYPE_CERTIFICATED_AMATEUR_BUT_HAS_NOT_BEEN_REGISTERED_THAT_IS_PENDING_CANCELLATION",
		StatusCode_NNumberAssignedAsImportButHasNotBeenRegisteredThatIsPendingCancellation:                      "N_NUMBER_ASSIGNED_AS_IMPORT_BUT_HAS_NOT_BEEN_REGISTERED_THAT_IS_PENDING_CANCELLATION",
		StatusCode_RegistrationExpired:                                                                          "REGISTRATION_EXPIRED",
		StatusCode_FirstNoticeForReRegistrationRenewal:                                                          "FIRST_NOTICE_FOR_REREGISTRATION_RENEWAL",
		StatusCode_SecondNoticeForReRegistrationRenewal:                                                         "SECOND_NOTICE_FOR_REREGISTRATION_RENEWAL",
		StatusCode_RegistrationExpiredPendingCancellation:                                                       "REGISTRATION_EXPIRED_PENDING_CANCELLATION",
		StatusCode_SaleReportedPendingCancellation:                                                              "SALE_REPORTED_PENDING_CANCELLATION",
		StatusCode_SaleReportedCanceled:                                                                         "SALE_REPORTED_CANCELLED",
		StatusCode_RegistrationPendingPendingCancellation:                                                       "REGISTRATION_PENDING_PENDING_CANCELLATION",
		StatusCode_RegistrationPendingCanceled:                                                                  "REGISTRATION_PENDING_CANCELED",
		StatusCode_RevokedPendingCancellation:                                                                   "REVOKED_PENDING_CANCELLATION",
		StatusCode_RevokedCanceled:                                                                              "REVOKED_CANCELLATION",
		StatusCode_ExpiredDealerPendingCancellation:                                                             "EXPIRED_DEALTER_PENDING_CANCELLATION",
		StatusCode_ThirdNoticeForReRegistrationRenewal:                                                          "THIRD_NOTICE_FOR_REREGISTRATION_RENEWAL",
		StatusCode_FirstNoticeForRegistrationRenewal:                                                            "FIRST_NOTICE_FOR_REGISTRATION_RENEWAL",
		StatusCode_SecondNoticeForRegistrationRenewal:                                                           "SECOND_NOTICE_FOR_REGISTRATION_RENEWAL",
		StatusCode_ThirdNoticeForRegistrationRenewal:                                                            "THIRD_NOTICE_FOR_REGISTRATION_RENEWAL",
	}
	ApprovedOperations_value = map[ApprovedOperation]string{
		ApprovedOperation_Normal:                                               "NORMAL",
		ApprovedOperation_Utility:                                              "UTILITY",
		ApprovedOperation_Acrobatic:                                            "ACROBATIC",
		ApprovedOperation_Transport:                                            "TRANSPORT",
		ApprovedOperation_Glider:                                               "GLIDER",
		ApprovedOperation_Balloon:                                              "BALLOON",
		ApprovedOperation_Commuter:                                             "COMMUTER",
		ApprovedOperation_Other:                                                "OTHER",
		ApprovedOperation_AgricultureAndPestControl:                            "AGRICULTURE_AND_PEST_CONTROL",
		ApprovedOperation_AerialSurveying:                                      "AERIAL_SERVEYING",
		ApprovedOperation_AerialAdvertising:                                    "AERIAL_ADVERTISING",
		ApprovedOperation_Forest:                                               "FOREST",
		ApprovedOperation_Patrolling:                                           "PATROLLING",
		ApprovedOperation_WeatherControl:                                       "WEATHER_CONTROL",
		ApprovedOperation_CarriageOfCargo:                                      "CARRIAGE_OF_CARGO",
		ApprovedOperation_ToShowComplianceWithFAR:                              "TO_SHOW_COMPLIANCE_WITH_FAR",
		ApprovedOperation_ResearchAndDevelopment:                               "RESEARCH_AND_DEVELOPMENT",
		ApprovedOperation_AmateurBuilt:                                         "AMATEUR_BUILT",
		ApprovedOperation_Exhibition:                                           "EXHIBITION",
		ApprovedOperation_Racing:                                               "RACING",
		ApprovedOperation_CrewTraining:                                         "CREW_TRAINING",
		ApprovedOperation_MarketSurvey:                                         "MARKET_SURVEY",
		ApprovedOperation_OperatingKitBuiltAircraft:                            "OPERATING_KIT_BUILT_AIRCRAFT",
		ApprovedOperation_RegisteredPriorTo_01_31_08:                           "REGISTERED_PRIOR_TO_01_31_08",
		ApprovedOperation_OperatingLightSportKitBuilt:                          "OPERATING_LIGHT_SPORT_KIT_BUILT",
		ApprovedOperation_OperatingLightSportPreviouslyIssuedCertUnder_21_190:  "OPERATING_LIGHT_SPORT_PREVIOUSLY_ISSUED_CERT_UNDER_21_190",
		ApprovedOperation_UnmannedAircraftResearchAndDevelopment:               "UNMANNED_AIRCRAFT_RESEARCH_AND_DEVELOPEMENT",
		ApprovedOperation_UnmannedAircraftMarketSurvey:                         "UNMANNED_AIRCRAFT_MARKET_SURVEY",
		ApprovedOperation_UnmannedAircraftCrewTraining:                         "UNMANNED_AIRCRAFT_CREW_TRAINING",
		ApprovedOperation_UnmannedAircraftExhibition:                           "UNMANNED_AIRCRAFT_EXHIBITION",
		ApprovedOperation_UnmannedAircraftComplianceWithCFR:                    "UNMANNED_AIRCRAFT_COMPLIANCE_WITH_FAR",
		ApprovedOperation_ClassOne:                                             "CLASS_ONE",
		ApprovedOperation_ClassTwo:                                             "CLASS_TWO",
		ApprovedOperation_Standard:                                             "STANDARD",
		ApprovedOperation_Limited:                                              "LIMITED",
		ApprovedOperation_Restricted:                                           "RESTRICTED",
		ApprovedOperation_FerryFlightForRepairsAlterationsMaintenanceOrStorage: "FERRY_FLIGHT_FOR_REPAIRS_ALTERATIONS_MAINTENANCE_OR_STORAGE",
		ApprovedOperation_EvacuateFromAreaOfImpendingDanger:                    "EVACTUATE_FROM_AREA_OF_IMPENDING_DANGER",
		ApprovedOperation_OperationInExcessOfMaximumCertificated:               "OPERATION_IN_EXCESS_OF_MAXIMUM_CERTIFICATED",
		ApprovedOperation_DeliveryOrExport:                                     "DELIVERY_OF_EXPORT",
		ApprovedOperation_ProductionFlightTesting:                              "PRODUCTION_FLIGHT_TESTING",
		ApprovedOperation_CustomerDemo:                                         "CUSTOMER_DEMO",
		ApprovedOperation_Airplane:                                             "AIRPLANE",
		ApprovedOperation_LighterThanAir:                                       "LIGHER_THAN_AIR",
		ApprovedOperation_PowerParachute:                                       "POWER_PARACHUTE",
		ApprovedOperation_WeightShiftControl:                                   "WEIGHT_SHIFT_CONTROL",
	}
	ApprovedOperations_name = map[Classification]map[string]ApprovedOperation{
		Classification_Standard: {
			"N": ApprovedOperation_Normal,
			"U": ApprovedOperation_Utility,
			"A": ApprovedOperation_Acrobatic,
			"T": ApprovedOperation_Transport,
			"G": ApprovedOperation_Glider,
			"B": ApprovedOperation_Balloon,
			"C": ApprovedOperation_Commuter,
			"O": ApprovedOperation_Other,
		},
		Classification_Limited: {},
		Classification_Restricted: {
			"0": ApprovedOperation_Other,
			"1": ApprovedOperation_AgricultureAndPestControl,
			"2": ApprovedOperation_AerialSurveying,
			"3": ApprovedOperation_AerialAdvertising,
			"4": ApprovedOperation_Forest,
			"5": ApprovedOperation_Patrolling,
			"6": ApprovedOperation_WeatherControl,
			"7": ApprovedOperation_CarriageOfCargo,
		},

		Classification_Experimental: {
			"0":  ApprovedOperation_ToShowComplianceWithFAR,
			"1":  ApprovedOperation_ResearchAndDevelopment,
			"2":  ApprovedOperation_AmateurBuilt,
			"3":  ApprovedOperation_Exhibition,
			"4":  ApprovedOperation_Racing,
			"5":  ApprovedOperation_CrewTraining,
			"6":  ApprovedOperation_MarketSurvey,
			"7":  ApprovedOperation_OperatingKitBuiltAircraft,
			"8A": ApprovedOperation_RegisteredPriorTo_01_31_08,
			"8B": ApprovedOperation_OperatingLightSportKitBuilt,
			"8C": ApprovedOperation_OperatingLightSportPreviouslyIssuedCertUnder_21_190,
			"9A": ApprovedOperation_UnmannedAircraftResearchAndDevelopment,
			"9B": ApprovedOperation_UnmannedAircraftMarketSurvey,
			"9C": ApprovedOperation_UnmannedAircraftCrewTraining,
			"9D": ApprovedOperation_UnmannedAircraftExhibition,
			"9E": ApprovedOperation_UnmannedAircraftComplianceWithCFR,
		},
		Classification_Provisional: {
			"1": ApprovedOperation_ClassOne,
			"2": ApprovedOperation_ClassTwo,
		},
		Classification_Multiple: {
			"P1_1": ApprovedOperation_Standard,
			"P1_2": ApprovedOperation_Limited,
			"P1_3": ApprovedOperation_Restricted,
			"P2_0": ApprovedOperation_Other,
			"P2_1": ApprovedOperation_AgricultureAndPestControl,
			"P2_2": ApprovedOperation_AerialSurveying,
			"P2_3": ApprovedOperation_AerialAdvertising,
			"P2_4": ApprovedOperation_Forest,
			"P2_5": ApprovedOperation_Patrolling,
			"P2_6": ApprovedOperation_WeatherControl,
			"P2_7": ApprovedOperation_CarriageOfCargo,
		},
		Classification_Primary: {},
		Classification_SpecialFlightPermit: {
			"1": ApprovedOperation_FerryFlightForRepairsAlterationsMaintenanceOrStorage,
			"2": ApprovedOperation_EvacuateFromAreaOfImpendingDanger,
			"3": ApprovedOperation_OperationInExcessOfMaximumCertificated,
			"4": ApprovedOperation_DeliveryOrExport,
			"5": ApprovedOperation_ProductionFlightTesting,
			"6": ApprovedOperation_CustomerDemo,
		},
		Classification_LightSport: {
			"A": ApprovedOperation_Airplane,
			"G": ApprovedOperation_Glider,
			"L": ApprovedOperation_LighterThanAir,
			"P": ApprovedOperation_PowerParachute,
			"W": ApprovedOperation_WeightShiftControl,
		},
	}
)

type Address struct {
	Street1 string `json:"street_one"`
	Street2 string `json:"street_two"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Region  string `json:"region"`
	County  string `json:"county"`
	Country string `json:"country"`
}

type Kit struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}

type Registrant struct {
	Type RegistrantType `json:"type"`
	Name string         `json:"name"`
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
	Registrant               *Registrant
	Address                  *Address
	LastActivityDate         *time.Time
	CertificationIssueDate   *time.Time
	Classification           Classification
	ApprovedOperations       string
	AircraftType             AircraftType
	EngineType               EngineType
	StatusCode               StatusCode
	ModeSCode                string
	FractionalOwnership      string
	AirworthinessDate        *time.Time
	OtherNameOne             string
	OtherNameTwo             string
	OtherNameThree           string
	OtherNameFour            string
	OtherNameFive            string
	ExpirationDate           *time.Time
	UniqueID                 string
	Kit                      *Kit
}

func (ac Classification) String() string {
	val, ok := Classification_value[ac]
	if !ok {
		return "UNKNOWN"
	}
	return val
}

func (ac ApprovedOperation) String() string {
	val, ok := ApprovedOperations_value[ac]
	if !ok {
		return "UNKNOWN"
	}
	return val
}

func (ac StatusCode) String() string {
	val, ok := StatusCode_value[ac]
	if !ok {
		return "UNKNOWN"
	}
	return val
}

func (s Address) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s RegistrantType) MarshalJSON() ([]byte, error) {
	registrantType, ok := RegistrantType_value[s]
	if !ok {
		return nil, errors.New("unknown registrant type")
	}
	return []byte(fmt.Sprintf("\"%s\"", registrantType)), nil
}

func (s Kit) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s Registrant) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (reg Registration) ID() string {
	return reg.UniqueID
}

func (reg Registration) GetApprovedOperations() []string {
	toReturn := []string{}

	switch reg.Classification {
	case Classification_Standard:
		if reg.ApprovedOperations != "" {
			val, ok := ApprovedOperations_name[Classification_Standard][reg.ApprovedOperations[0:1]]
			if ok {
				toReturn = append(toReturn, val.String())
			}
		}
	case Classification_Limited, Classification_Primary:
	case Classification_Restricted:
		for _, op := range reg.ApprovedOperations {
			val, ok := ApprovedOperations_name[Classification_Restricted][fmt.Sprintf("%c", op)]
			if !ok {
				continue
			}
			toReturn = append(toReturn, val.String())
		}
	case Classification_Experimental:
		i := 0
		for i < len(reg.ApprovedOperations) {
			var key string
			if (reg.ApprovedOperations[i] == '8' || reg.ApprovedOperations[i] == '9') && len(reg.ApprovedOperations) >= i+2 {
				key = reg.ApprovedOperations[i : i+2]
				i += 2
			} else {
				key = fmt.Sprintf("%c", reg.ApprovedOperations[i:i+1])
				i++
			}
			val, ok := ApprovedOperations_name[Classification_Experimental][key]
			if ok {
				toReturn = append(toReturn, val.String())
			}
		}
	case Classification_Provisional:
		if reg.ApprovedOperations != "" {
			val, ok := ApprovedOperations_name[Classification_Provisional][reg.ApprovedOperations[0:1]]
			if !ok {
				return toReturn
			}
			toReturn = append(toReturn, val.String())
		}
	case Classification_Multiple:
		switch l := len(reg.ApprovedOperations); {
		case l > 0:
			left := math.Min(2.0, float64(len(reg.ApprovedOperations)))
			for _, op := range reg.ApprovedOperations[:int(left)] {
				val, ok := ApprovedOperations_name[Classification_Multiple][fmt.Sprintf("P1_%c", op)]
				if ok {
					toReturn = append(toReturn, val.String())
				}
			}
			fallthrough
		case l > 2:
			for _, op := range reg.ApprovedOperations[2:] {
				val, ok := ApprovedOperations_name[Classification_Multiple][fmt.Sprintf("P2_%c", op)]
				if ok {
					toReturn = append(toReturn, val.String())
				}
			}
		}
	case Classification_SpecialFlightPermit:
		for _, op := range reg.ApprovedOperations {
			val, ok := ApprovedOperations_name[Classification_SpecialFlightPermit][fmt.Sprintf("%c", op)]
			if ok {
				toReturn = append(toReturn, val.String())
			}
		}
	case Classification_LightSport:
		if reg.ApprovedOperations != "" {
			val, ok := ApprovedOperations_name[Classification_LightSport][reg.ApprovedOperations[0:1]]
			if ok {
				toReturn = append(toReturn, val.String())
			}
		}
	}
	return toReturn
}

func (reg *Registration) Unmarshal(data string) RowBuilder {
	return NewRegistration(data)
}

func (ac Registration) Columns() []string {
	return []string{
		"id",
		"tail_number",
		"serial_number",
		"year_manufactured",
		"aircraft_id",
		"registrant",
		"address",
		"last_activity_date",
		"certificate_issue_date",
		"classification",
		"approved_operations",
		"type",
		"engine_type",
		"status_code",
		"model_s_code",
		"is_fractionally_owned",
		"airworthiness_date",
		"other_names",
		"expiration_date",
		"kit",
		"created",
	}
}

func (reg Registration) Values(now time.Time) []interface{} {
	var otherNames []string
	{
		for _, name := range []string{
			reg.OtherNameOne,
			reg.OtherNameTwo,
			reg.OtherNameThree,
			reg.OtherNameFour,
			reg.OtherNameFive,
		} {
			if name != "" {
				otherNames = append(otherNames, name)
			}
		}
	}

	return []interface{}{
		reg.UniqueID,
		reg.Id,
		reg.SerialNumber,
		reg.YearManufactured,
		fmt.Sprintf("%s-%s-%s", reg.AircraftManufacturerCode, reg.AircraftModelCode, reg.AircraftSeriesCode),
		reg.Registrant,
		reg.Address,
		reg.LastActivityDate,
		reg.CertificationIssueDate,
		reg.Classification.String(),
		pq.Array(reg.GetApprovedOperations()),
		reg.AircraftType.String(),
		reg.EngineType.String(),
		reg.StatusCode.String(),
		reg.ModeSCode,
		reg.FractionalOwnership == "Y",
		reg.AirworthinessDate,
		pq.Array(otherNames),
		reg.ExpirationDate,
		reg.Kit,
		now.UTC(),
	}
}

func toTime(y, m, d string) (*time.Time, error) {
	month, err := strconv.Atoi(m)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse("2006-Jan-2", fmt.Sprintf("%s-%s-%s", y, time.Month(month).String(), d))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func NewRegistration(data string) Registration {
	var lastActivityDate *time.Time
	{
		if t, err := toTime(strings.TrimSpace(data[219:223]), strings.TrimSpace(data[223:225]), strings.TrimSpace(data[225:227])); err == nil {
			lastActivityDate = t
		}
	}

	var certificationIssueDate *time.Time
	{
		if t, err := toTime(strings.TrimSpace(data[228:232]), strings.TrimSpace(data[232:234]), strings.TrimSpace(data[234:236])); err == nil {
			certificationIssueDate = t
		}
	}

	var expirationDate *time.Time
	{
		if t, err := toTime(strings.TrimSpace(data[531:535]), strings.TrimSpace(data[535:537]), strings.TrimSpace(data[537:539])); err == nil {
			expirationDate = t
		}
	}

	var airworthinessDate *time.Time
	{
		if t, err := toTime(strings.TrimSpace(data[267:271]), strings.TrimSpace(data[271:273]), strings.TrimSpace(data[273:275])); err == nil {
			airworthinessDate = t
		}
	}

	aircraftType, ok := AircraftType_name[strings.TrimSpace(data[248:249])]
	if !ok {
		aircraftType = AircraftType_Unknown
	}

	engineType, ok := EngineType_name[strings.TrimSpace(data[250:252])]
	if !ok {
		engineType = EngineType_Unknown
	}

	classification, ok := Classification_name[strings.TrimSpace(data[237:238])]
	if !ok {
		classification = Classification_Unknown
	}

	statusCode, ok := StatusCode_name[strings.TrimSpace(data[253:255])]
	if !ok {
		statusCode = StatusCode_Unknown
	}

	registrantType, ok := RegistrantType_name[strings.TrimSpace(data[56:57])]
	if !ok {
		registrantType = RegistrantType_Unknown
	}

	var kit *Kit
	{
		manufacturer := strings.TrimSpace(data[549:579])
		model := strings.TrimSpace(data[580:600])
		if manufacturer != "" || model != "" {
			kit = &Kit{
				Manufacturer: manufacturer,
				Model:        model,
			}
		}
	}

	return Registration{
		Id:                       strings.TrimSpace(data[0:5]),
		SerialNumber:             strings.TrimSpace(data[6:36]),
		AircraftManufacturerCode: strings.TrimSpace(data[37:40]),
		AircraftModelCode:        strings.TrimSpace(data[40:42]),
		AircraftSeriesCode:       strings.TrimSpace(data[42:44]),
		EngineManufacturerCode:   strings.TrimSpace(data[45:48]),
		EngineModelCode:          strings.TrimSpace(data[48:50]),
		YearManufactured:         strings.TrimSpace(data[51:55]),
		Registrant: &Registrant{
			Type: registrantType,
			Name: strings.TrimSpace(data[58:108]),
		},
		Address: &Address{
			Street1: strings.TrimSpace(data[109:142]),
			Street2: strings.TrimSpace(data[143:176]),
			City:    strings.TrimSpace(data[177:195]),
			State:   strings.TrimSpace(data[196:198]),
			ZipCode: strings.TrimSpace(data[199:209]),
			Region:  strings.TrimSpace(data[210:211]),
			County:  strings.TrimSpace(data[212:215]),
			Country: strings.TrimSpace(data[216:218]),
		},
		LastActivityDate:       lastActivityDate,
		CertificationIssueDate: certificationIssueDate,
		Classification:         classification,
		ApprovedOperations:     strings.TrimSpace(data[238:247]),
		AircraftType:           aircraftType,
		EngineType:             engineType,
		StatusCode:             statusCode,
		ModeSCode:              strings.TrimSpace(data[256:264]),
		FractionalOwnership:    strings.TrimSpace(data[265:266]),
		AirworthinessDate:      airworthinessDate,
		OtherNameOne:           strings.TrimSpace(data[276:326]),
		OtherNameTwo:           strings.TrimSpace(data[327:377]),
		OtherNameThree:         strings.TrimSpace(data[378:428]),
		OtherNameFour:          strings.TrimSpace(data[429:479]),
		OtherNameFive:          strings.TrimSpace(data[480:530]),
		ExpirationDate:         expirationDate,
		UniqueID:               strings.TrimSpace(data[540:548]),
		Kit:                    kit,
	}
}
