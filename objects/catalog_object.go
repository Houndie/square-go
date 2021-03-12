package objects

import (
	"encoding/json"
	"fmt"
	"time"

	"errors"
)

type CatalogObjectType string

const (
	CatalogObjectTypeItem            CatalogObjectType = "ITEM"
	CatalogObjectTypeItemVariation   CatalogObjectType = "ITEM_VARIATION"
	CatalogObjectTypeModifier        CatalogObjectType = "MODIFIER"
	CatalogObjectTypeModifierList    CatalogObjectType = "MODIFIER_LIST"
	CatalogObjectTypeCategory        CatalogObjectType = "CATEGORY"
	CatalogObjectTypeDiscount        CatalogObjectType = "DISCOUNT"
	CatalogObjectTypeTax             CatalogObjectType = "TAX"
	CatalogObjectTypeImage           CatalogObjectType = "IMAGE"
	CatalogObjectTypeMeasurementUnit CatalogObjectType = "MEASUREMENT_UNIT"
)

type catalogObject struct {
	Type                  CatalogObjectType       `json:"type,omitempty"`
	ID                    string                  `json:"id,omitempty"`
	UpdatedAt             *time.Time              `json:"updated_at,omitempty"`
	Version               int                     `json:"version,omitempty"`
	IsDeleted             bool                    `json:"is_deleted,omitempty"`
	CatalogV1IDs          []*CatalogV1ID          `json:"catalog_v1_ids,omitempty"`
	PresentAtAllLocations bool                    `json:"present_at_all_locations,omitempty"`
	PresentAtLocationIDs  []string                `json:"present_at_location_ids,omitempty"`
	AbsentAtLocationIDs   []string                `json:"absent_at_location_ids,omitempty"`
	ImageID               string                  `json:"image_id,omitempty"`
	ItemData              *CatalogItem            `json:"item_data,omitempty"`
	CategoryData          *CatalogCategory        `json:"category_data,omitempty"`
	ItemVariationData     *CatalogItemVariation   `json:"item_variation_data,omitempty"`
	TaxData               *CatalogTax             `json:"tax_data,omitempty"`
	DiscountData          *CatalogDiscount        `json:"discount_data,omitempty"`
	ModifierListData      *CatalogModifierList    `json:"modifier_list_data,omitempty"`
	ModifierData          *CatalogModifier        `json:"modifier_data,omitempty"`
	ImageData             *CatalogImage           `json:"image_data,omitempty"`
	MeasurementUnitData   *CatalogMeasurementUnit `json:"catalog_measurement_unit,omitempty"`
}

type catalogObjectType interface {
	isCatalogObjectType()
}

type CatalogObject struct {
	ID                    string
	UpdatedAt             *time.Time
	Version               int
	IsDeleted             bool
	CatalogV1IDs          []*CatalogV1ID
	PresentAtAllLocations bool
	PresentAtLocationIDs  []string
	AbsentAtLocationIDs   []string
	ImageID               string
	CatalogObjectType     catalogObjectType
}

func (c *CatalogObject) MarshalJSON() ([]byte, error) {
	cJson := catalogObject{
		ID:                    c.ID,
		UpdatedAt:             c.UpdatedAt,
		Version:               c.Version,
		IsDeleted:             c.IsDeleted,
		CatalogV1IDs:          c.CatalogV1IDs,
		PresentAtAllLocations: c.PresentAtAllLocations,
		PresentAtLocationIDs:  c.PresentAtLocationIDs,
		AbsentAtLocationIDs:   c.AbsentAtLocationIDs,
		ImageID:               c.ImageID,
	}
	switch t := c.CatalogObjectType.(type) {
	case *CatalogItem:
		cJson.ItemData = t
		cJson.Type = CatalogObjectTypeItem
	case *CatalogCategory:
		cJson.CategoryData = t
		cJson.Type = CatalogObjectTypeCategory
	case *CatalogItemVariation:
		cJson.ItemVariationData = t
		cJson.Type = CatalogObjectTypeItemVariation
	case *CatalogTax:
		cJson.TaxData = t
		cJson.Type = CatalogObjectTypeTax
	case *CatalogDiscount:
		cJson.DiscountData = t
		cJson.Type = CatalogObjectTypeDiscount
	case *CatalogModifierList:
		cJson.ModifierListData = t
		cJson.Type = CatalogObjectTypeModifierList
	case *CatalogModifier:
		cJson.ModifierData = t
		cJson.Type = CatalogObjectTypeModifier
	case *CatalogImage:
		cJson.ImageData = t
		cJson.Type = CatalogObjectTypeImage
	case *CatalogMeasurementUnit:
		cJson.MeasurementUnitData = t
		cJson.Type = CatalogObjectTypeMeasurementUnit
	default:
		return nil, errors.New("Found unknown catalog object data type")
	}
	json, err := json.Marshal(&cJson)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling json catalog object: %w", err)
	}
	return json, nil
}

func (c *CatalogObject) UnmarshalJSON(data []byte) error {
	cJson := &catalogObject{}
	err := json.Unmarshal(data, &cJson)
	if err != nil {
		return fmt.Errorf("Error unmarshaling catalog object: %w", err)
	}
	c.ID = cJson.ID
	c.UpdatedAt = cJson.UpdatedAt
	c.Version = cJson.Version
	c.IsDeleted = cJson.IsDeleted
	c.CatalogV1IDs = cJson.CatalogV1IDs
	c.PresentAtAllLocations = cJson.PresentAtAllLocations
	c.PresentAtLocationIDs = cJson.PresentAtLocationIDs
	c.AbsentAtLocationIDs = cJson.AbsentAtLocationIDs
	c.ImageID = cJson.ImageID

	switch cJson.Type {
	case CatalogObjectTypeItem:
		c.CatalogObjectType = cJson.ItemData
	case CatalogObjectTypeCategory:
		c.CatalogObjectType = cJson.CategoryData
	case CatalogObjectTypeItemVariation:
		c.CatalogObjectType = cJson.ItemVariationData
	case CatalogObjectTypeTax:
		c.CatalogObjectType = cJson.TaxData
	case CatalogObjectTypeDiscount:
		c.CatalogObjectType = cJson.DiscountData
	case CatalogObjectTypeModifierList:
		c.CatalogObjectType = cJson.ModifierListData
	case CatalogObjectTypeModifier:
		c.CatalogObjectType = cJson.ModifierData
	case CatalogObjectTypeImage:
		c.CatalogObjectType = cJson.ImageData
	case CatalogObjectTypeMeasurementUnit:
		c.CatalogObjectType = cJson.MeasurementUnitData
	default:
		return fmt.Errorf("Found unknown catalog object type %s", cJson.Type)
	}
	return nil
}
