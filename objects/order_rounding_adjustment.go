package objects

type OrderRoundingAdjustment struct {
	UID         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
}
