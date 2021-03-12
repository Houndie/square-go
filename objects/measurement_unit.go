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

func (*MeasurementUnitCustom) isMeasurementUnit() {}
func (MeasurementUnitArea) isMeasurementUnit()    {}
func (MeasurementUnitLength) isMeasurementUnit()  {}
func (MeasurementUnitVolume) isMeasurementUnit()  {}
func (MeasurementUnitWeight) isMeasurementUnit()  {}
func (MeasurementUnitGeneric) isMeasurementUnit() {}

type measurementUnit struct {
	CustomUnit  *MeasurementUnitCustom `json:"custom_unit,omitempty"`
	AreaUnit    string                 `json:"area_unit,omitempty"`
	LengthUnit  string                 `json:"length_unit,omitempty"`
	VolumeUnit  string                 `json:"volume_unit,omitempty"`
	WeightUnit  string                 `json:"weight_unit,omitempty"`
	GenericUnit string                 `json:"generic_unit,omitempty"`
}

func (m *MeasurementUnit) MarshalJSON() ([]byte, error) {
	mJson := measurementUnit{}
	switch t := m.Type.(type) {
	case *MeasurementUnitCustom:
		mJson.CustomUnit = t
	case MeasurementUnitArea:
		mJson.AreaUnit = string(t)
	case MeasurementUnitLength:
		mJson.LengthUnit = string(t)
	case MeasurementUnitVolume:
		mJson.VolumeUnit = string(t)
	case MeasurementUnitWeight:
		mJson.WeightUnit = string(t)
	case MeasurementUnitGeneric:
		mJson.GenericUnit = string(t)
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

	switch {
	case mJson.CustomUnit != nil:
		m.Type = mJson.CustomUnit
		return nil
	case mJson.AreaUnit != "":
		m.Type = MeasurementUnitArea(mJson.AreaUnit)
		return nil
	case mJson.LengthUnit != "":
		m.Type = MeasurementUnitLength(mJson.LengthUnit)
		return nil
	case mJson.VolumeUnit != "":
		m.Type = MeasurementUnitVolume(mJson.VolumeUnit)
		return nil
	case mJson.WeightUnit != "":
		m.Type = MeasurementUnitWeight(mJson.WeightUnit)
		return nil
	case mJson.GenericUnit != "":
		m.Type = MeasurementUnitGeneric(mJson.GenericUnit)
		return nil
	}
	return errors.New("No unit types found in json measurement unit")
}
