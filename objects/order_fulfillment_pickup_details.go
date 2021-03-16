package objects

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/peterhellberg/duration"
)

type orderFulfillmentPickupDetails struct {
	Recipient             *OrderFulfillmentRecipient                          `json:"recipient,omitempty"`
	ExpiresAt             *time.Time                                          `json:"expires_at,omitempty"`
	AutoCompleteDuration  string                                              `json:"auto_complete_duration,omitempty"`
	PickupAt              *time.Time                                          `json:"pickup_at,omitempty"`
	PickupWindowDuration  string                                              `json:"pickup_window_duration,omitempty"`
	PrepTimeDuration      string                                              `json:"prep_time_duration,omitempty"`
	Note                  string                                              `json:"note,omitempty"`
	PlacedAt              *time.Time                                          `json:"placed_at,omitempty"`
	AcceptedAt            *time.Time                                          `json:"accepted_at,omitempty"`
	RejectedAt            *time.Time                                          `json:"rejected_at,omitempty"`
	ReadyAt               *time.Time                                          `json:"ready_at,omitempty"`
	ExpiredAt             *time.Time                                          `json:"expired_at,omitempty"`
	PickedUpAt            *time.Time                                          `json:"picked_up_at,omitempty"`
	CanceledAt            *time.Time                                          `json:"canceled_at,omitempty"`
	CancelReason          string                                              `json:"cancel_reason,omitempty"`
	IsCurbsidePickup      bool                                                `json:"is_curbside_pickup,omitempty"`
	CurbsidePickupDetails *OrderFulfillmentPickupDetailsCurbsidePickupDetails `json:"curbside_pickup_details,omitempty"`
}

type OrderFulfillmentPickupDetails struct {
	Recipient            *OrderFulfillmentRecipient
	ExpiresAt            *time.Time
	AutoCompleteDuration *time.Duration
	PickupAt             *time.Time
	PickupWindowDuration *time.Duration
	PrepTimeDuration     *time.Duration
	Note                 string
	PlacedAt             *time.Time
	AcceptedAt           *time.Time
	RejectedAt           *time.Time
	ReadyAt              *time.Time
	ExpiredAt            *time.Time
	PickedUpAt           *time.Time
	CanceledAt           *time.Time
	CancelReason         string
	CurbsidePickup       *OrderFulfillmentPickupDetailsCurbsidePickup
}

type OrderFulfillmentPickupDetailsCurbsidePickup struct {
	Details *OrderFulfillmentPickupDetailsCurbsidePickupDetails
}

type OrderFulfillmentPickupDetailsCurbsidePickupDetails struct {
	BuyerArrivedAt  *time.Time `json:"buyer_arrived_at,omitempty"`
	CurbsideDetails string     `json:"curbside_details,omitempty"`
}

func (o *OrderFulfillmentPickupDetails) MarshalJSON() ([]byte, error) {
	jsonType := orderFulfillmentPickupDetails{
		Recipient:    o.Recipient,
		ExpiresAt:    o.ExpiresAt,
		PickupAt:     o.PickupAt,
		Note:         o.Note,
		PlacedAt:     o.PlacedAt,
		AcceptedAt:   o.AcceptedAt,
		RejectedAt:   o.RejectedAt,
		ReadyAt:      o.ReadyAt,
		ExpiredAt:    o.ExpiredAt,
		PickedUpAt:   o.PickedUpAt,
		CanceledAt:   o.CanceledAt,
		CancelReason: o.CancelReason,
	}
	if o.AutoCompleteDuration != nil {
		jsonType.AutoCompleteDuration = fmt.Sprintf("%vS", o.AutoCompleteDuration.Seconds())
	}
	if o.PickupWindowDuration != nil {
		jsonType.PickupWindowDuration = fmt.Sprintf("%vS", o.PickupWindowDuration.Seconds())
	}
	if o.PrepTimeDuration != nil {
		jsonType.PrepTimeDuration = fmt.Sprintf("%vS", o.PrepTimeDuration.Seconds())
	}
	if o.CurbsidePickup != nil {
		jsonType.IsCurbsidePickup = true
		if o.CurbsidePickup.Details != nil {
			jsonType.CurbsidePickupDetails = o.CurbsidePickup.Details
		}
	}
	return json.Marshal(&jsonType)
}

func (o *OrderFulfillmentPickupDetails) UnmarshalJSON(b []byte) error {
	jsonType := orderFulfillmentPickupDetails{}
	err := json.Unmarshal(b, &jsonType)
	if err != nil {
		return fmt.Errorf("Error unmarshaling OrderFulfillmentPickupDetails: %w", err)
	}

	o.Recipient = jsonType.Recipient
	o.ExpiresAt = jsonType.ExpiresAt
	o.PickupAt = jsonType.PickupAt
	o.Note = jsonType.Note
	o.PlacedAt = jsonType.PlacedAt
	o.AcceptedAt = jsonType.AcceptedAt
	o.RejectedAt = jsonType.RejectedAt
	o.ReadyAt = jsonType.ReadyAt
	o.ExpiredAt = jsonType.ExpiredAt
	o.PickedUpAt = jsonType.PickedUpAt
	o.CanceledAt = jsonType.CanceledAt
	o.CancelReason = jsonType.CancelReason

	if jsonType.IsCurbsidePickup {
		o.CurbsidePickup = &OrderFulfillmentPickupDetailsCurbsidePickup{}
		if jsonType.CurbsidePickupDetails != nil {
			o.CurbsidePickup.Details = jsonType.CurbsidePickupDetails
		}
	}

	if jsonType.AutoCompleteDuration != "" {
		d, err := duration.Parse(jsonType.AutoCompleteDuration)
		if err != nil {
			return fmt.Errorf("Error unmarshaling OrderFulfillmentPickupDetails.AutoCompleteDuration: %w", err)
		}
		o.AutoCompleteDuration = &d
	}

	if jsonType.PickupWindowDuration != "" {
		d, err := duration.Parse(jsonType.PickupWindowDuration)
		if err != nil {
			return fmt.Errorf("Error unmarshaling OrderFulfillmentPickupDetails.PickupWindowDuration: %w", err)
		}
		o.PickupWindowDuration = &d
	}

	if jsonType.PrepTimeDuration != "" {
		d, err := duration.Parse(jsonType.PrepTimeDuration)
		if err != nil {
			return fmt.Errorf("Error unmarshaling OrderFulfillmentPickupDetails.PrepTimeDuration: %w", err)
		}
		o.PrepTimeDuration = &d
	}
	return nil
}
