package objects

type OrderFulfillmentRecipient struct {
	CustomerID   string   `json:"customer_id,omitempty"`
	DisplayName  string   `json:"display_name,omitempty"`
	EmailAddress string   `json:"email_address,omitempty"`
	PhoneNumber  string   `json:"phone_number,omitempty"`
	Address      *Address `json:"address,omitempty"`
}
