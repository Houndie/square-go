package objects

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/peterhellberg/duration"
)

type orderFulfillmentPickupDetails struct {
	*orderFulfillmentPickupDetailsAlias
	AutoCompleteDuration  string                                              `json:"auto_complete_duration,omitempty"`
	PickupWindowDuration  string                                              `json:"pickup_window_duration,omitempty"`
	PrepTimeDuration      string                                              `json:"prep_time_duration,omitempty"`
	IsCurbsidePickup      bool                                                `json:"is_curbside_pickup,omitempty"`
	CurbsidePickupDetails *OrderFulfillmentPickupDetailsCurbsidePickupDetails `json:"curbside_pickup_details,omitempty"`
}

type orderFulfillmentPickupDetailsAlias OrderFulfillmentPickupDetails

type OrderFulfillmentPickupDetails struct {
	Recipient            *OrderFulfillmentRecipient                   `json:"recipient,omitempty"`
	ExpiresAt            *time.Time                                   `json:"expires_at,omitempty"`
	AutoCompleteDuration *time.Duration                               `json:"-"`
	PickupAt             *time.Time                                   `json:"pickup_at,omitempty"`
	PickupWindowDuration *time.Duration                               `json:"-"`
	PrepTimeDuration     *time.Duration                               `json:"-"`
	Note                 string                                       `json:"note,omitempty"`
	PlacedAt             *time.Time                                   `json:"placed_at,omitempty"`
	AcceptedAt           *time.Time                                   `json:"accepted_at,omitempty"`
	RejectedAt           *time.Time                                   `json:"rejected_at,omitempty"`
	ReadyAt              *time.Time                                   `json:"ready_at,omitempty"`
	ExpiredAt            *time.Time                                   `json:"expired_at,omitempty"`
	PickedUpAt           *time.Time                                   `json:"picked_up_at,omitempty"`
	CanceledAt           *time.Time                                   `json:"canceled_at,omitempty"`
	CancelReason         string                                       `json:"cancel_reason,omitempty"`
	CurbsidePickup       *OrderFulfillmentPickupDetailsCurbsidePickup `json:"-"`
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
		orderFulfillmentPickupDetailsAlias: (*orderFulfillmentPickupDetailsAlias)(o),
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
	jsonType := orderFulfillmentPickupDetails{
		orderFulfillmentPickupDetailsAlias: (*orderFulfillmentPickupDetailsAlias)(o),
	}
	err := json.Unmarshal(b, &jsonType)
	if err != nil {
		return fmt.Errorf("Error unmarshaling OrderFulfillmentPickupDetails: %w", err)
	}

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
