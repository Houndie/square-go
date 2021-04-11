package objects

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CatalogProductSetQuantity interface {
	isCatalogProductSetQuantity()
}

type CatalogProductSetQuantityExact struct {
	Amount int
}

type CatalogProductSetQuantityRange struct {
	Min int
	Max int
}

func (*CatalogProductSetQuantityExact) isCatalogProductSetQuantity() {}
func (*CatalogProductSetQuantityRange) isCatalogProductSetQuantity() {}

type CatalogProductSetProduct interface {
	isCatalogProductSetProduct()
}

type CatalogProductSetAllProducts struct{}

type CatalogProductSetAllIDs struct {
	IDs []string
}

type CatalogProductSetAnyIDs struct {
	IDs []string
}

func (*CatalogProductSetAllProducts) isCatalogProductSetProduct() {}
func (*CatalogProductSetAllIDs) isCatalogProductSetProduct()      {}
func (*CatalogProductSetAnyIDs) isCatalogProductSetProduct()      {}

type catalogProductSet struct {
	*catalogProductSetAlias
	AllProducts   bool     `json:"all_products,omitempty"`
	ProductIDsAll []string `json:"product_ids_all,omitempty"`
	ProductIDsAny []string `json:"product_ids_any,omitempty"`
	QuantityExact int      `json:"quantity_exact,omitempty"`
	QuantityMax   int      `json:"quantity_max,omitempty"`
	QuantityMin   int      `json:"quantity_min,omitempty"`
}

type catalogProductSetAlias CatalogProductSet

type CatalogProductSet struct {
	Name     string                    `json:"name,omitempty"`
	Products CatalogProductSetProduct  `json:"-"`
	Quantity CatalogProductSetQuantity `json:"-"`
}

func (*CatalogProductSet) isCatalogObjectType() {}

func (c *CatalogProductSet) MarshalJSON() ([]byte, error) {
	jsonType := &catalogProductSet{
		catalogProductSetAlias: (*catalogProductSetAlias)(c),
	}

	switch t := c.Products.(type) {
	case *CatalogProductSetAllProducts:
		jsonType.AllProducts = true
	case *CatalogProductSetAllIDs:
		jsonType.ProductIDsAll = t.IDs
	case *CatalogProductSetAnyIDs:
		jsonType.ProductIDsAny = t.IDs
	default:
		return nil, fmt.Errorf("unknown product type found")
	}

	switch t := c.Quantity.(type) {
	case *CatalogProductSetQuantityExact:
		jsonType.QuantityExact = t.Amount
	case *CatalogProductSetQuantityRange:
		jsonType.QuantityMin = t.Min
		jsonType.QuantityMax = t.Max
	}

	b, err := json.Marshal(jsonType)
	if err != nil {
		return nil, fmt.Errorf("error marshaling catalog product set: %w", err)
	}

	return b, nil
}

func (c *CatalogProductSet) UnmarshalJSON(b []byte) error {
	jsonType := &catalogProductSet{
		catalogProductSetAlias: (*catalogProductSetAlias)(c),
	}

	err := json.Unmarshal(b, &jsonType)
	if err != nil {
		return fmt.Errorf("error unmarshaling catalog product set: %w", err)
	}

	if jsonType.QuantityExact != 0 {
		c.Quantity = &CatalogProductSetQuantityExact{
			Amount: jsonType.QuantityExact,
		}
	} else {
		c.Quantity = &CatalogProductSetQuantityRange{
			Min: jsonType.QuantityMin,
			Max: jsonType.QuantityMax,
		}
	}

	switch {
	case jsonType.AllProducts:
		c.Products = &CatalogProductSetAllProducts{}
	case jsonType.ProductIDsAll != nil:
		c.Products = &CatalogProductSetAllIDs{
			IDs: jsonType.ProductIDsAll,
		}
	case jsonType.ProductIDsAny != nil:
		c.Products = &CatalogProductSetAnyIDs{
			IDs: jsonType.ProductIDsAny,
		}
	default:
		return errors.New("no product specifier set")
	}

	return nil
}
