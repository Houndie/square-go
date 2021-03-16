package objects

type OrderReturnTax struct {
	UID             string                `json:"uid,omitempty"`
	SourceTaxUID    string                `json:"source_tax_uid,omitempty"`
	CatalogObjectID string                `json:"catalog_object_id,omitempty"`
	Name            string                `json:"name,omitempty"`
	Type            OrderLineItemTaxType  `json:"type,omitempty"`
	Percentage      string                `json:"percentage,omitempty"`
	AppliedMoney    *Money                `json:"applied_money,omitempty"`
	Scope           OrderLineItemTaxScope `json:"scope,omitempty"`
}
