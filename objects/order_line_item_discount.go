package objects

import (
	"encoding/json"

	"errors"
)

type OrderLineItemDiscountScope string

const (
	OrderLineItemDiscountScopeOtherDiscountScope OrderLineItemDiscountScope = "OTHER_DISCOUNT_SCOPE"
	OrderLineItemDiscountScopeLineItem           OrderLineItemDiscountScope = "LINE_ITEM"
	OrderLineItemDiscountScopeOrder              OrderLineItemDiscountScope = "ORDER"
)

type orderLineItemDiscountType string

const (
	orderLineItemDiscountTypeUnknownDiscount    orderLineItemDiscountType = "UNKNOWN_DISCOUNT"
	orderLineItemDiscountTypeFixedPercentage    orderLineItemDiscountType = "FIXED_PERCENTAGE"
	orderLineItemDiscountTypeFixedAmount        orderLineItemDiscountType = "FIXED_AMOUNT"
	orderLineItemDiscountTypeVariablePercentage orderLineItemDiscountType = "VARIABLE_PERCENTAGE"
	orderLineItemDiscountTypeVariableAmount     orderLineItemDiscountType = "VARIABLE_AMOUNT"
)

type orderLineItemDiscount struct {
	*orderLineItemDiscountAlias
	Type        orderLineItemDiscountType `json:"type,omitempty"`
	Percentage  string                    `json:"percentage,omitempty"`
	AmountMoney *Money                    `json:"amount_money,omitempty"`
}

type orderLineItemDiscountAlias OrderLineItemDiscount

type OrderLineItemDiscountType interface {
	isOrderLineItemDiscountType()
}

type OrderLineItemDiscountUnknownDiscount struct{}

type OrderLineItemDiscountFixedPercentage struct {
	Percentage string
}

type OrderLineItemDiscountVariablePercentage struct {
	Percentage string
}

type OrderLineItemDiscountFixedAmount struct {
	AmountMoney *Money
}

type OrderLineItemDiscountVariableAmount struct {
	AmountMoney *Money
}

func (*OrderLineItemDiscountUnknownDiscount) isOrderLineItemDiscountType()    {}
func (*OrderLineItemDiscountFixedPercentage) isOrderLineItemDiscountType()    {}
func (*OrderLineItemDiscountFixedAmount) isOrderLineItemDiscountType()        {}
func (*OrderLineItemDiscountVariablePercentage) isOrderLineItemDiscountType() {}
func (*OrderLineItemDiscountVariableAmount) isOrderLineItemDiscountType()     {}

type OrderLineItemDiscount struct {
	UID             string                     `json:"uid,omitempty"`
	CatalogObjectID string                     `json:"catalog_object_id,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Type            OrderLineItemDiscountType  `json:"-"`
	AppliedMoney    *Money                     `json:"applied_money,omitempty"`
	Scope           OrderLineItemDiscountScope `json:"scope,omitempty"`
	Metadata        map[string]string          `json:"metadata,omitempty"`
	RewardIDs       []string                   `json:"reward_ids,omitempty"`
	PricingRuleID   string                     `json:"pricing_rule_id,omitempty"`
}

func (o *OrderLineItemDiscount) MarshalJSON() ([]byte, error) {
	jsonData := orderLineItemDiscount{
		orderLineItemDiscountAlias: (*orderLineItemDiscountAlias)(o),
	}

	switch t := o.Type.(type) {
	case *OrderLineItemDiscountUnknownDiscount:
		jsonData.Type = orderLineItemDiscountTypeUnknownDiscount
	case *OrderLineItemDiscountFixedAmount:
		jsonData.Type = orderLineItemDiscountTypeFixedAmount
		jsonData.AmountMoney = t.AmountMoney
	case *OrderLineItemDiscountVariableAmount:
		jsonData.Type = orderLineItemDiscountTypeVariableAmount
		jsonData.AmountMoney = t.AmountMoney
	case *OrderLineItemDiscountFixedPercentage:
		jsonData.Type = orderLineItemDiscountTypeFixedPercentage
		jsonData.Percentage = t.Percentage
	case *OrderLineItemDiscountVariablePercentage:
		jsonData.Type = orderLineItemDiscountTypeVariablePercentage
		jsonData.Percentage = t.Percentage
	case nil:
		// Do Nothing
	default:
		return nil, errors.New("unknown discount type found")
	}

	return json.Marshal(&jsonData)
}

func (o *OrderLineItemDiscount) UnmarshalJSON(input []byte) error {
	jsonData := orderLineItemDiscount{
		orderLineItemDiscountAlias: (*orderLineItemDiscountAlias)(o),
	}
	err := json.Unmarshal(input, &jsonData)
	if err != nil {
		return err
	}

	switch jsonData.Type {
	case orderLineItemDiscountTypeUnknownDiscount:
		o.Type = &OrderLineItemDiscountUnknownDiscount{}
	case orderLineItemDiscountTypeFixedAmount:
		o.Type = &OrderLineItemDiscountFixedAmount{
			AmountMoney: jsonData.AmountMoney,
		}
	case orderLineItemDiscountTypeVariableAmount:
		o.Type = &OrderLineItemDiscountVariableAmount{
			AmountMoney: jsonData.AmountMoney,
		}
	case orderLineItemDiscountTypeVariablePercentage:
		o.Type = &OrderLineItemDiscountVariablePercentage{
			Percentage: jsonData.Percentage,
		}
	case orderLineItemDiscountTypeFixedPercentage:
		o.Type = &OrderLineItemDiscountFixedPercentage{
			Percentage: jsonData.Percentage,
		}
	case "":
		// Do Nothing
	default:
		return errors.New("unknown discount type found")
	}

	return nil
}
