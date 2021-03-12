package objects

type AdditionalRecipient struct {
	LocationID   string `json:"location_id,omitempty"`
	Description  string `json:"description,omitempty"`
	AmountMoney  *Money `json:"amount_money,omitempty"`
	ReceivableID string `json:"receivable_id,omitempty"`
}
