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
	UID             string                     `json:"uid,omitempty"`
	CatalogObjectID string                     `json:"catalog_object_id,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Type            orderLineItemDiscountType  `json:"type,omitempty"`
	Percentage      string                     `json:"percentage,omitempty"`
	AmountMoney     *Money                     `json:"amount_money,omitempty"`
	AppliedMoney    *Money                     `json:"applied_money,omitempty"`
	Scope           OrderLineItemDiscountScope `json:"scope,omitempty"`
}

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
	UID             string
	CatalogObjectID string
	Name            string
	Type            OrderLineItemDiscountType
	AppliedMoney    *Money
	Scope           OrderLineItemDiscountScope
}

func (o *OrderLineItemDiscount) MarshalJSON() ([]byte, error) {
	jsonData := orderLineItemDiscount{
		UID:             o.UID,
		CatalogObjectID: o.CatalogObjectID,
		Name:            o.Name,
		AppliedMoney:    o.AppliedMoney,
		Scope:           o.Scope,
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
	jsonData := orderLineItemDiscount{}
	err := json.Unmarshal(input, &jsonData)
	if err != nil {
		return err
	}

	o.UID = jsonData.UID
	o.CatalogObjectID = jsonData.CatalogObjectID
	o.Name = jsonData.Name
	o.AppliedMoney = jsonData.AppliedMoney
	o.Scope = jsonData.Scope
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
