package objects

import (
	"encoding/json"
	"fmt"

	"errors"
)

type MeasurementUnit struct {
	Type MeasurementUnitType
}

type MeasurementUnitType interface {
	isMeasurementUnit()
}

type MeasurementUnitCustom struct {
	Name         string `json:"name,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
}

type MeasurementUnitArea string

type MeasurementUnitLength string

type MeasurementUnitVolume string

type MeasurementUnitWeight string

type MeasurementUnitGeneric string

type MeasurementUnitTime string

func (*MeasurementUnitCustom) isMeasurementUnit() {}
func (MeasurementUnitArea) isMeasurementUnit()    {}
func (MeasurementUnitLength) isMeasurementUnit()  {}
func (MeasurementUnitVolume) isMeasurementUnit()  {}
func (MeasurementUnitWeight) isMeasurementUnit()  {}
func (MeasurementUnitGeneric) isMeasurementUnit() {}
func (MeasurementUnitTime) isMeasurementUnit()    {}

type measurementUnitType string

const (
	measurementUnitTypeCustom  measurementUnitType = "TYPE_CUSTOM"
	measurementUnitTypeArea    measurementUnitType = "TYPE_AREA"
	measurementUnitTypeLength  measurementUnitType = "TYPE_LENGTH"
	measurementUnitTypeVolume  measurementUnitType = "TYPE_VOLUME"
	measurementUnitTypeWeight  measurementUnitType = "TYPE_WEIGHT"
	measurementUnitTypeGeneric measurementUnitType = "TYPE_GENERIC"
	measurementUnitTypeTime    measurementUnitType = "TYPE_TIME"
)

type measurementUnit struct {
	CustomUnit  *MeasurementUnitCustom `json:"custom_unit,omitempty"`
	AreaUnit    string                 `json:"area_unit,omitempty"`
	LengthUnit  string                 `json:"length_unit,omitempty"`
	VolumeUnit  string                 `json:"volume_unit,omitempty"`
	WeightUnit  string                 `json:"weight_unit,omitempty"`
	GenericUnit string                 `json:"generic_unit,omitempty"`
	TimeUnit    string                 `json:"time_unit,omitempty"`
	Type        measurementUnitType    `json:"type,omitempty"`
}

func (m *MeasurementUnit) MarshalJSON() ([]byte, error) {
	mJSON := measurementUnit{}
	switch t := m.Type.(type) {
	case *MeasurementUnitCustom:
		mJSON.CustomUnit = t
		mJSON.Type = measurementUnitTypeLength
	case MeasurementUnitArea:
		mJSON.AreaUnit = string(t)
		mJSON.Type = measurementUnitTypeArea
	case MeasurementUnitLength:
		mJSON.LengthUnit = string(t)
		mJSON.Type = measurementUnitTypeLength
	case MeasurementUnitVolume:
		mJSON.VolumeUnit = string(t)
		mJSON.Type = measurementUnitTypeVolume
	case MeasurementUnitWeight:
		mJSON.WeightUnit = string(t)
		mJSON.Type = measurementUnitTypeWeight
	case MeasurementUnitGeneric:
		mJSON.GenericUnit = string(t)
		mJSON.Type = measurementUnitTypeGeneric
	case MeasurementUnitTime:
		mJSON.TimeUnit = string(t)
		mJSON.Type = measurementUnitTypeTime
	default:
		return nil, errors.New("found unknown measurement unit type when marshaling")
	}

	b, err := json.Marshal(mJSON)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json measurement unit: %w", err)
	}

	return b, nil
}

func (m *MeasurementUnit) UnmarshalJSON(b []byte) error {
	mJSON := measurementUnit{}
	if err := json.Unmarshal(b, &mJSON); err != nil {
		return fmt.Errorf("error unmarshaling json measurement unit: %w", err)
	}

	switch mJSON.Type {
	case measurementUnitTypeCustom:
		m.Type = mJSON.CustomUnit
		return nil
	case measurementUnitTypeArea:
		m.Type = MeasurementUnitArea(mJSON.AreaUnit)
		return nil
	case measurementUnitTypeLength:
		m.Type = MeasurementUnitLength(mJSON.LengthUnit)
		return nil
	case measurementUnitTypeVolume:
		m.Type = MeasurementUnitVolume(mJSON.VolumeUnit)
		return nil
	case measurementUnitTypeWeight:
		m.Type = MeasurementUnitWeight(mJSON.WeightUnit)
		return nil
	case measurementUnitTypeGeneric:
		m.Type = MeasurementUnitGeneric(mJSON.GenericUnit)
		return nil
	case measurementUnitTypeTime:
		m.Type = MeasurementUnitTime(mJSON.TimeUnit)
		return nil
	}

	return errors.New("no unit types found in json measurement unit")
}
