package objects

type OrderLineItemTaxType string

const (
	OrderLineItemTaxTypeUnknownTax OrderLineItemTaxType = "UNKNOWN_TAX"
	OrderLineItemTaxTypeAdditive   OrderLineItemTaxType = "ADDITIVE"
	OrderLineItemTaxTypeInclusive  OrderLineItemTaxType = "INCLUSIVE"
)

type OrderLineItemTaxScope string

const (
	OrderLineItemTaxScopeOtherTaxScope OrderLineItemTaxScope = "OTHER_TAX_SCOPE"
	OrderLineItemTaxScopeLineItem      OrderLineItemTaxScope = "LINE_ITEM"
	OrderLineItemTaxScopeOrder         OrderLineItemTaxScope = "ORDER"
)

type OrderLineItemTax struct {
	Uid             string                `json:"uid,omitempty"`
	CatalogObjectID string                `json:"catalog_object_id,omitempty"`
	Name            string                `json:"name,omitempty"`
	Type            OrderLineItemTaxType  `json:"type,omitempty"`
	Percentage      string                `json:"percentage,omitempty"`
	Metadata        map[string]string     `json:"metadata,omitempty"`
	AppliedMoney    *Money                `json:"applied_money,omitempty"`
	Scope           OrderLineItemTaxScope `json:"scope,omitempty"`
}
