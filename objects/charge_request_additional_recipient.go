package objects

type ChargeRequestAdditionalRecipient struct {
	LocationID  string `json:"location_id,omitempty"`
	Description string `json:"description,omitempty"`
	AmountMoney *Money `json:"amount_money,omitempty"`
}
