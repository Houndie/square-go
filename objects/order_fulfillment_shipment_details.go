package objects

import "time"

type OrderFulfillmentShipmentDetails struct {
	CancelReason      string                     `json:"cancel_reason,omitempty"`
	CanceledAt        *time.Time                 `json:"canceled_at,omitempty"`
	Carrier           string                     `json:"carrier,omitempty"`
	ExpectedShippedAt *time.Time                 `json:"expected_shipped_at,omitempty"`
	FailedAt          *time.Time                 `json:"failed_at,omitempty"`
	FailureReason     string                     `json:"failure_reason,omitempty"`
	InProgressAt      *time.Time                 `json:"in_progress_at,omitempty"`
	PackagedAt        *time.Time                 `json:"packaged_at,omitempty"`
	PlacedAt          *time.Time                 `json:"placed_at,omitempty"`
	Recipient         *OrderFulfillmentRecipient `json:"recipient,omitempty"`
	ShippedAt         *time.Time                 `json:"shipped_at,omitempty"`
	ShippingNote      string                     `json:"shipping_note,omitempty"`
	ShippingType      string                     `json:"shipping_type,omitempty"`
	TrackingNumber    string                     `json:"tracking_number,omitempty"`
	TrackingURL       string                     `json:"tracking_url,omitempty"`
}
