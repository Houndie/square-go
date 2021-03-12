package objects

import "time"

type TransactionProduct string

const (
	TransactionProductRegister     TransactionProduct = "REGISTER"
	TransactionProductExternalApi  TransactionProduct = "EXTERNAL_API"
	TransactionProductBilling      TransactionProduct = "BILLING"
	TransactionProductAppointments TransactionProduct = "APPOINTMENTS"
	TransactionProductInvoices     TransactionProduct = "INVOICES"
	TransactionProductOnlineStore  TransactionProduct = "ONLINE_STORE"
	TransactionProductPayroll      TransactionProduct = "PAYROLL"
	TransactionProductOther        TransactionProduct = "OTHER"
)

type Transaction struct {
	ID              string             `json:"id,omitempty"`
	LocationID      string             `json:"location_id,omitempty"`
	CreatedAt       *time.Time         `json:"created_at,omitempty"`
	Tenders         []*Tender          `json:"tenders,omitempty"`
	Refunds         []*Refund          `json:"refunds,omitempty"`
	ReferenceID     string             `json:"reference_id,omitempty"`
	Product         TransactionProduct `json:"product,omitempty"`
	ClientID        string             `json:"client_id,omitempty"`
	ShippingAddress *Address           `json:"shipping_address,omitempty"`
	OrderID         string             `json:"order_id,omitempty"`
}
