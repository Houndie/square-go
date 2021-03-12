package objects

type OrderQuantityUnit struct {
	MeasurementUnit *MeasurementUnit `json:"measurement_unit,omitempty"`
	Precision       int              `json:"precision,omitempty"`
}
