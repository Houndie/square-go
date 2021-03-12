package objects

type OrderMoneyAmounts struct {
	TotalMoney         *Money `json:"total_money,omitempty"`
	TaxMoney           *Money `json:"tax_money,omitempty"`
	DiscountMoney      *Money `json:"discount_money,omitempty"`
	TipMoney           *Money `json:"tip_money,omitempty"`
	ServiceChargeMoney *Money `json:"service_charge_money,omitempty"`
}
