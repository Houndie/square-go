package objects

type OrderReturnDiscount struct {
	Uid               string                     `json:"uid,omitempty"`
	SourceDiscountUID string                     `json:"source_discount_uid,omitempty"`
	CatalogObjectID   string                     `json:"catalog_object_id,omitempty"`
	Name              string                     `json:"name,omitempty"`
	Type              OrderLineItemDiscountType  `json:"type,omitempty"`
	Percentage        string                     `json:"percentage,omitempty"`
	AmountMoney       *Money                     `json:"amount_money,omitempty"`
	AppliedMoney      *Money                     `json:"applied_money,omitempty"`
	Scope             OrderLineItemDiscountScope `json:"scope,omitempty"`
}
