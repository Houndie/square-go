package objects

type CatalogMeasurementUnit struct {
	MeasurementUnit *MeasurementUnit `json:"measurement_unit,omitempty"`
	Precision       int              `json:"precision,omitempty"`
}

func (*CatalogMeasurementUnit) isCatalogObjectType() {}
