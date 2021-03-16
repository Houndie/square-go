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
	mJson := measurementUnit{}
	switch t := m.Type.(type) {
	case *MeasurementUnitCustom:
		mJson.CustomUnit = t
		mJson.Type = measurementUnitTypeLength
	case MeasurementUnitArea:
		mJson.AreaUnit = string(t)
		mJson.Type = measurementUnitTypeArea
	case MeasurementUnitLength:
		mJson.LengthUnit = string(t)
		mJson.Type = measurementUnitTypeLength
	case MeasurementUnitVolume:
		mJson.VolumeUnit = string(t)
		mJson.Type = measurementUnitTypeVolume
	case MeasurementUnitWeight:
		mJson.WeightUnit = string(t)
		mJson.Type = measurementUnitTypeWeight
	case MeasurementUnitGeneric:
		mJson.GenericUnit = string(t)
		mJson.Type = measurementUnitTypeGeneric
	case MeasurementUnitTime:
		mJson.TimeUnit = string(t)
		mJson.Type = measurementUnitTypeTime
	default:
		return nil, errors.New("found unknown measurement unit type when marshaling")
	}

	b, err := json.Marshal(mJson)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json measurement unit: %w", err)
	}
	return b, nil
}

func (m *MeasurementUnit) UnmarshalJSON(b []byte) error {
	mJson := measurementUnit{}
	if err := json.Unmarshal(b, &mJson); err != nil {
		return fmt.Errorf("error unmarshaling json measurement unit: %w", err)
	}

	switch mJson.Type {
	case measurementUnitTypeCustom:
		m.Type = mJson.CustomUnit
		return nil
	case measurementUnitTypeArea:
		m.Type = MeasurementUnitArea(mJson.AreaUnit)
		return nil
	case measurementUnitTypeLength:
		m.Type = MeasurementUnitLength(mJson.LengthUnit)
		return nil
	case measurementUnitTypeVolume:
		m.Type = MeasurementUnitVolume(mJson.VolumeUnit)
		return nil
	case measurementUnitTypeWeight:
		m.Type = MeasurementUnitWeight(mJson.WeightUnit)
		return nil
	case measurementUnitTypeGeneric:
		m.Type = MeasurementUnitGeneric(mJson.GenericUnit)
		return nil
	case measurementUnitTypeTime:
		m.Type = MeasurementUnitTime(mJson.TimeUnit)
		return nil
	}
	return errors.New("No unit types found in json measurement unit")
}
