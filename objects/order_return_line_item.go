package objects

type OrderReturnLineItem struct {
	Uid                      string                         `json:"uid,omitempty"`
	SourceLineItemUid        string                         `json:"source_line_item_uid,omitempty"`
	Name                     string                         `json:"name,omitempty"`
	Quantity                 string                         `json:"quantity,omitempty"`
	QuantityUnit             *OrderQuantityUnit             `json:"quantity_unit,omitempty"`
	Note                     string                         `json:"note,omitempty"`
	CatalogObjectID          string                         `json:"catalog_object_id,omitempty"`
	VariationName            string                         `json:"variation_name,omitempty"`
	ReturnModifiers          []*OrderReturnLineItemModifier `json:"return_modifiers,omitempty"`
	ReturnTaxes              []*OrderReturnTax              `json:"return_taxes,omitempty"`
	ReturnDiscounts          []*OrderReturnDiscount         `json:"return_discounts,omitempty"`
	BasePriceMoney           *Money                         `json:"base_price_money,omitempty"`
	VariationTotalPriceMoney *Money                         `json:"variation_total_price_money,omitempty"`
	GrossReturnMoney         *Money                         `json:"gross_return_money,omitempty"`
	TotalTaxMoney            *Money                         `json:"total_tax_money,omitempty"`
	TotalDiscountMoney       *Money                         `json:"total_discount_money,omitempty"`
	TotalMoney               *Money                         `json:"total_money,omitempty"`
}
