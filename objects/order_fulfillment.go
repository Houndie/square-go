package objects

type OrderFulfillmentType string

const OrderFulfillmentTypePickup OrderFulfillmentType = "PICKUP"

type OrderFulfillmentState string

const (
	OrderFullfillmentStatProposed  OrderFulfillmentState = "PROPOSED"
	OrderFulfillmentStateReserved  OrderFulfillmentState = "RESERVED"
	OrderFulfillmentStatePrepared  OrderFulfillmentState = "PREPARED"
	OrderFulfillmentStateCompleted OrderFulfillmentState = "COMPLETED"
	OrderFulfillmentStateCanceled  OrderFulfillmentState = "CANCELED"
	OrderFulfillmentStateFailed    OrderFulfillmentState = "FAILED"
)

type OrderFulfillment struct {
	Type          OrderFulfillmentType           `json:"type,omitempty"`
	State         OrderFulfillmentState          `json:"state,omitempty"`
	PickupDetails *OrderFulfillmentPickupDetails `json:"pickup_details,omitempty"`
}
