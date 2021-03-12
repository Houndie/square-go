package objects

type OrderLineItem struct {
	Uid                      string                          `json:"uid,omitempty"`
	Name                     string                          `json:"name,omitempty"`
	Quantity                 string                          `json:"quantity,omitempty"`
	QuantityUnit             *OrderQuantityUnit              `json:"quantity_unit,omitempty"`
	Note                     string                          `json:"note,omitempty"`
	CatalogObjectID          string                          `json:"catalog_object_id,omitempty"`
	VariationName            string                          `json:"variation_name,omitempty"`
	Modifiers                []*OrderLineItemModifier        `json:"modifiers,omitempty"`
	Taxes                    []*OrderLineItemTax             `json:"taxes,omitempty"`
	BasePriceMoney           *Money                          `json:"base_price_money,omitempty"`
	VariationTotalPriceMoney *Money                          `json:"variation_total_price_money,omitempty"`
	GrossSalesMoney          *Money                          `json:"gross_sales_money,omitempty"`
	TotalTaxMoney            *Money                          `json:"total_tax_money,omitempty"`
	TotalDiscountMoney       *Money                          `json:"total_discount_money,omitempty"`
	TotalMoney               *Money                          `json:"total_money,omitempty"`
	AppliedDiscounts         []*OrderLineItemAppliedDiscount `json:"applied_discounts,omitempty"`
	AppliedTaxes             []*OrderLineItemAppliedTax      `json:"applied_taxes,omitempty"`
	Metadata                 map[string]string               `json:"metadata,omitempty"`
	PricingBlocklists        *OrderLineItemPricingBlocklists `json:"pricing_blocklists,omitempty"`
}
