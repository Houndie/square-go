package objects

type CatalogTax struct {
	Name                   string `json:"name,omitempty"`
	CalculationPhase       string `json:"calculation_phase,omitempty"`
	InclusionType          string `json:"inclusion_type,omitempty"`
	Percentage             string `json:"percentage,omitempty"`
	AppliesToCustomAmounts bool   `json:"applies_to_custom_amounts,omitempty"`
	Enabled                bool   `json:"enabled,omitempty"`
}

func (*CatalogTax) isCatalogObjectType() {}
