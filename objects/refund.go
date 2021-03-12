package objects

import "time"

type RefundStatus string

const (
	RefundStatusPending  RefundStatus = "PENDING"
	RefundStatusApproved RefundStatus = "APPROVED"
	RefundStatusRejected RefundStatus = "REJECTED"
	RefundStatusFailed   RefundStatus = "FAILED"
)

type Refund struct {
	ID                   string                 `json:"id,omitempty"`
	LocationID           string                 `json:"location_id,omitempty"`
	TransactionID        string                 `json:"transaction_id,omitempty"`
	TenderID             string                 `json:"tender_id,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	Reason               string                 `json:"reason,omitempty"`
	AmountMoney          *Money                 `json:"amount_money,omitempty"`
	Status               RefundStatus           `json:"status,omitempty"`
	ProcessingFeeMoney   *Money                 `json:"processing_fee_money,omitempty"`
	AdditionalRecipients []*AdditionalRecipient `json:"additional_recipients,omitempty"`
}
