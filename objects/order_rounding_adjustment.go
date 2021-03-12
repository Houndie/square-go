package objects

type OrderRoundingAdjustment struct {
	Uid         string `json:"uid,omitempty"`
	Name        string `json:"name,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
}
