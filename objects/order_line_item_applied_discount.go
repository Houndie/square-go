package objects

type OrderLineItemAppliedDiscount struct {
	UID          string `json:"uid,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	DiscountUID  string `json:"discount_uid,omitempty"`
}
