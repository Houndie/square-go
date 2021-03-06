package objects

import (
	"encoding/json"
	"fmt"

	"errors"
)

type CatalogDiscountType interface {
	isCatalogDiscountType()
}

type CatalogDiscountFixedPercentage struct {
	Percentage string
}

type CatalogDiscountVariablePercentage struct {
	Percentage string
}

type CatalogDiscountFixedAmount struct {
	AmountMoney *Money
}

type CatalogDiscountVariableAmount struct {
	AmountMoney *Money
}

func (*CatalogDiscountFixedPercentage) isCatalogDiscountType()    {}
func (*CatalogDiscountFixedAmount) isCatalogDiscountType()        {}
func (*CatalogDiscountVariablePercentage) isCatalogDiscountType() {}
func (*CatalogDiscountVariableAmount) isCatalogDiscountType()     {}

type CatalogDiscount struct {
	Name           string              `json:"name,omitempty"`
	DiscountType   CatalogDiscountType `json:"-"`
	PinRequired    bool                `json:"pin_required,omitempty"`
	LabelColor     string              `json:"label_color,omitempty"`
	ModifyTaxBasis string              `json:"modify_tax_basis,omitempty"`
}

func (*CatalogDiscount) isCatalogObjectType() {}

type catalogDiscountType string

const (
	catalogDiscountTypeFixedPercentage    catalogDiscountType = "FIXED_PERCENTAGE"
	catalogDiscountTypeFixedAmount        catalogDiscountType = "FIXED_AMOUNT"
	catalogDiscountTypeVariablePercentage catalogDiscountType = "VARIABLE_PERCENTAGE"
	catalogDiscountTypeVariableAmount     catalogDiscountType = "VARIABLE_AMOUNT"
)

type catalogDiscountAlias CatalogDiscount

type catalogDiscount struct {
	*catalogDiscountAlias
	DiscountType catalogDiscountType `json:"discount_type,omitempty"`
	Percentage   string              `json:"percentage,omitempty"`
	AmountMoney  *Money              `json:"amount_money,omitempty"`
}

func (c *CatalogDiscount) MarshalJSON() ([]byte, error) {
	jsonType := catalogDiscount{
		catalogDiscountAlias: (*catalogDiscountAlias)(c),
	}
	switch d := c.DiscountType.(type) {
	case *CatalogDiscountFixedPercentage:
		jsonType.DiscountType = catalogDiscountTypeFixedPercentage
		jsonType.Percentage = d.Percentage
	case *CatalogDiscountVariablePercentage:
		jsonType.DiscountType = catalogDiscountTypeVariablePercentage
		jsonType.Percentage = d.Percentage
	case *CatalogDiscountFixedAmount:
		jsonType.DiscountType = catalogDiscountTypeFixedAmount
		jsonType.AmountMoney = d.AmountMoney
	case *CatalogDiscountVariableAmount:
		jsonType.DiscountType = catalogDiscountTypeVariableAmount
		jsonType.AmountMoney = d.AmountMoney
	default:
		return nil, errors.New("could not marshal catalog discount, unknown catalog discount type found")
	}

	b, err := json.Marshal(&jsonType)
	if err != nil {
		return nil, fmt.Errorf("error marshing catalog discount: %w", err)
	}

	return b, nil
}

func (c *CatalogDiscount) UnmarshalJSON(b []byte) error {
	jsonType := catalogDiscount{
		catalogDiscountAlias: (*catalogDiscountAlias)(c),
	}
	if err := json.Unmarshal(b, &jsonType); err != nil {
		return fmt.Errorf("error unmarshaling catalog discount: %w", err)
	}

	switch jsonType.DiscountType {
	case catalogDiscountTypeFixedPercentage:
		c.DiscountType = &CatalogDiscountFixedPercentage{
			Percentage: jsonType.Percentage,
		}
	case catalogDiscountTypeVariablePercentage:
		c.DiscountType = &CatalogDiscountVariablePercentage{
			Percentage: jsonType.Percentage,
		}
	case catalogDiscountTypeFixedAmount:
		c.DiscountType = &CatalogDiscountFixedAmount{
			AmountMoney: jsonType.AmountMoney,
		}
	case catalogDiscountTypeVariableAmount:
		c.DiscountType = &CatalogDiscountVariableAmount{
			AmountMoney: jsonType.AmountMoney,
		}
	default:
		return fmt.Errorf("could not unmarshal catalog discount, unknown catalog discount type %s found", jsonType.DiscountType)
	}

	return nil
}
