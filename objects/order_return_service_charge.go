package objects

type OrderReturnServiceCharge struct {
	UID                    string                             `json:"uid,omitempty"`
	SourceServiceChargeUID string                             `json:"source_service_charge_uid,omitempty"`
	Name                   string                             `json:"name,omitempty"`
	CatalogObjectID        string                             `json:"catalog_object_id,omitempty"`
	Percentage             string                             `json:"percentage,omitempty"`
	AmountMoney            *Money                             `json:"amount_money,omitempty"`
	AppliedMoney           *Money                             `json:"applied_money,omitempty"`
	TotalMoney             *Money                             `json:"total_money,omitempty"`
	TotalTaxMoney          *Money                             `json:"total_tax_money,omitempty"`
	CalculationPhase       OrderServiceChargeCalculationPhase `json:"calculation_phase,omitempty"`
	Taxable                bool                               `json:"taxable,omitempty"`
	AppliedTaxes           []*OrderLineItemAppliedTax         `json:"applied_taxes,omitempty"`
}
