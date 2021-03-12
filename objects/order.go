package objects

import "time"

type OrderState string

const (
	OrderStateOpen      OrderState = "OPEN"
	OrderStateCompleted OrderState = "COMPLETED"
	OrderStateCanceled  OrderState = "CANCELED"
)

type Order struct {
	ID                      string                   `json:"id,omitempty"`
	LocationID              string                   `json:"location_id,omitempty"`
	ReferenceID             string                   `json:"reference_id,omitempty"`
	Source                  *OrderSource             `json:"source,omitempty"`
	CustomerID              string                   `json:"customer_id,omitempty"`
	LineItems               []*OrderLineItem         `json:"line_items,omitempty"`
	Taxes                   []*OrderLineItemTax      `json:"taxes,omitempty"`
	Discounts               []*OrderLineItemDiscount `json:"discounts,omitempty"`
	ServiceCharges          []*OrderServiceCharge    `json:"service_charges,omitempty"`
	Fulfillments            []*OrderFulfillment      `json:"fulfillments,omitempty"`
	Returns                 []*OrderReturn           `json:"returns,omitempty"`
	ReturnAmounts           *OrderMoneyAmounts       `json:"return_amounts,omitempty"`
	NetAmounts              *OrderMoneyAmounts       `json:"net_amounts,omitempty"`
	RoundingAdjustment      *OrderRoundingAdjustment `json:"rounding_adjustment,omitempty"`
	Tenders                 []*Tender                `json:"tenders,omitempty"`
	Refunds                 []*Refund                `json:"refunds,omitempty"`
	CreatedAt               *time.Time               `json:"created_at,omitempty"`
	UpdatedAt               *time.Time               `json:"updated_at,omitempty"`
	ClosedAt                *time.Time               `json:"closed_at,omitempty"`
	State                   OrderState               `json:"state,omitempty"`
	TotalMoney              *Money                   `json:"total_money,omitempty"`
	TotalTaxMoney           *Money                   `json:"total_tax_money,omitempty"`
	TotalDiscountMoney      *Money                   `json:"total_discount_money,omitempty"`
	TotalServiceChargeMoney *Money                   `json:"total_service_charge_money,omitempty"`
	Version                 int                      `json:"version,omitempty"`
	Metadata                map[string]string        `json:"metadata,omitempty"`
	PricingOptions          *OrderPricingOptions     `json:"pricing_options,omitempty"`
	Rewards                 []*OrderReward           `json:"rewards,omitempty"`
	TotalTipMoney           *Money                   `json:"total_tip_money,omitempty"`
}
