package objects

type TaxCalculationPhase string

const (
	TaxCalculationPhaseSubtotalPhase TaxCalculationPhase = "TAX_SUBTOTAL_PHASE"
	TaxCalculationPhaseTotalPhase    TaxCalculationPhase = "TAX_TOTAL_PHASE"
)

type TaxInclusionType string

const (
	TaxInclusionTypeAdditive  TaxInclusionType = "ADDITIVE"
	TaxInclusionTypeInclusive TaxInclusionType = "INCLUSIVE"
)

type CatalogTax struct {
	Name                   string              `json:"name,omitempty"`
	CalculationPhase       TaxCalculationPhase `json:"calculation_phase,omitempty"`
	InclusionType          TaxInclusionType    `json:"inclusion_type,omitempty"`
	Percentage             string              `json:"percentage,omitempty"`
	AppliesToCustomAmounts bool                `json:"applies_to_custom_amounts,omitempty"`
	Enabled                bool                `json:"enabled,omitempty"`
}

func (*CatalogTax) isCatalogObjectType() {}
