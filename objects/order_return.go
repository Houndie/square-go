package objects

type OrderReturn struct {
	UID                  string                      `json:"uid,omitempty"`
	SourceOrderID        string                      `json:"source_order_id,omitempty"`
	ReturnLineItems      []*OrderReturnLineItem      `json:"return_line_items,omitempty"`
	ReturnServiceCharges []*OrderReturnServiceCharge `json:"return_service_charges,omitempty"`
	ReturnTaxes          []*OrderReturnTax           `json:"return_taxes,omitempty"`
	ReturnDiscounts      []*OrderReturnDiscount      `json:"return_discounts,omitempty"`
	RoundingAdjustment   *OrderRoundingAdjustment    `json:"rounding_adjustment,omitempty"`
	ReturnAmounts        *OrderMoneyAmounts          `json:"return_amounts,omitempty"`
}
