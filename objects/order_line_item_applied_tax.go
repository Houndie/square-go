package objects

type OrderLineItemAppliedTax struct {
	UID          string `json:"uid,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	TaxUID       string `json:"tax_uid,omitempty"`
}
