package objects

import (
	"encoding/json"
	"fmt"
	"time"

	"errors"
)

type CatalogObjectType string

const (
	CatalogObjectTypeItem                      CatalogObjectType = "ITEM"
	CatalogObjectTypeImage                     CatalogObjectType = "IMAGE"
	CatalogObjectTypeCategory                  CatalogObjectType = "CATEGORY"
	CatalogObjectTypeItemVariation             CatalogObjectType = "ITEM_VARIATION"
	CatalogObjectTypeTax                       CatalogObjectType = "TAX"
	CatalogObjectTypeDiscount                  CatalogObjectType = "DISCOUNT"
	CatalogObjectTypeModifierList              CatalogObjectType = "MODIFIER_LIST"
	CatalogObjectTypeModifier                  CatalogObjectType = "MODIFIER"
	CatalogObjectTypePricingRule               CatalogObjectType = "PRICING_RULE"
	CatalogObjectTypeProductSet                CatalogObjectType = "PRODUCT_SET"
	CatalogObjectTypeTimePeriod                CatalogObjectType = "TIME_PERIOD"
	CatalogObjectTypeMeasurementUnit           CatalogObjectType = "MEASUREMENT_UNIT"
	CatalogObjectTypeSubscriptionPlan          CatalogObjectType = "SUBSCRIPTION_PLAN"
	CatalogObjectTypeItemOption                CatalogObjectType = "ITEM_OPTION"
	CatalogObjectTypeItemOptionVal             CatalogObjectType = "ITEM_OPTION_VAL"
	CatalogObjectTypeCustomAttributeDefinition CatalogObjectType = "CUSTOM_ATTRIBUTE_DEFINITION"
	CatalogObjectTypeQuickAmountsSettings      CatalogObjectType = "QUICK_AMOUNTS_SETTINGS"
)

type catalogObjectAlias CatalogObject

type catalogObject struct {
	*catalogObjectAlias
	Type                          CatalogObjectType                 `json:"type,omitempty"`
	ItemData                      *CatalogItem                      `json:"item_data,omitempty"`
	CategoryData                  *CatalogCategory                  `json:"category_data,omitempty"`
	ItemVariationData             *CatalogItemVariation             `json:"item_variation_data,omitempty"`
	TaxData                       *CatalogTax                       `json:"tax_data,omitempty"`
	DiscountData                  *CatalogDiscount                  `json:"discount_data,omitempty"`
	ModifierListData              *CatalogModifierList              `json:"modifier_list_data,omitempty"`
	ModifierData                  *CatalogModifier                  `json:"modifier_data,omitempty"`
	TimePeriodData                *CatalogTimePeriod                `json:"time_period_data,omitempty"`
	ProductSetData                *CatalogProductSet                `json:"product_set_data,omitempty"`
	PricingRuleData               *CatalogPricingRule               `json:"pricing_rule_data,omitempty"`
	ImageData                     *CatalogImage                     `json:"image_data,omitempty"`
	MeasurementUnitData           *CatalogMeasurementUnit           `json:"catalog_measurement_unit,omitempty"`
	SubscriptionPlanData          *CatalogSubscriptionPlan          `json:"catalog_subscription_plan,omitempty"`
	ItemOptionData                *CatalogItemOption                `json:"item_option_data,omitempty"`
	ItemOptionValueData           *CatalogItemOptionValue           `json:"item_option_value_data,omitempty"`
	CustomAttributeDefinitionData *CatalogCustomAttributeDefinition `json:"custom_attribute_definition_data,omitempty"`
	QuickAmountsSettingsData      *CatalogQuickAmountsSettings      `json:"quick_amounts_settings_data,omitempty"`
}

type catalogObjectType interface {
	isCatalogObjectType()
}

type CatalogObject struct {
	ID                    string                                  `json:"id,omitempty"`
	UpdatedAt             *time.Time                              `json:"updated_at,omitempty"`
	Version               int                                     `json:"version,omitempty"`
	IsDeleted             bool                                    `json:"is_deleted,omitempty"`
	CustomAttributeValues map[string]*CatalogCustomAttributeValue `json:"custom_attribute_values,omitempty"`
	CatalogV1IDs          []*CatalogV1ID                          `json:"catalog_v1_ids,omitempty"`
	PresentAtAllLocations bool                                    `json:"present_at_all_locations,omitempty"`
	PresentAtLocationIDs  []string                                `json:"present_at_location_ids,omitempty"`
	AbsentAtLocationIDs   []string                                `json:"absent_at_location_ids,omitempty"`
	ImageID               string                                  `json:"image_id,omitempty"`
	CatalogObjectType     catalogObjectType                       `json:"-"`
}

func (c *CatalogObject) MarshalJSON() ([]byte, error) {
	cJson := catalogObject{
		catalogObjectAlias: (*catalogObjectAlias)(c),
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
	case *CatalogTimePeriod:
		cJson.TimePeriodData = t
		cJson.Type = CatalogObjectTypeTimePeriod
	case *CatalogProductSet:
		cJson.ProductSetData = t
		cJson.Type = CatalogObjectTypeProductSet
	case *CatalogPricingRule:
		cJson.PricingRuleData = t
		cJson.Type = CatalogObjectTypePricingRule
	case *CatalogSubscriptionPlan:
		cJson.SubscriptionPlanData = t
		cJson.Type = CatalogObjectTypeSubscriptionPlan
	case *CatalogItemOption:
		cJson.ItemOptionData = t
		cJson.Type = CatalogObjectTypeItemOption
	case *CatalogItemOptionValue:
		cJson.ItemOptionValueData = t
		cJson.Type = CatalogObjectTypeItemOptionVal
	case *CatalogCustomAttributeDefinition:
		cJson.CustomAttributeDefinitionData = t
		cJson.Type = CatalogObjectTypeCustomAttributeDefinition
	case *CatalogQuickAmountsSettings:
		cJson.QuickAmountsSettingsData = t
		cJson.Type = CatalogObjectTypeQuickAmountsSettings
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
	cJson := catalogObject{
		catalogObjectAlias: (*catalogObjectAlias)(c),
	}
	err := json.Unmarshal(data, &cJson)
	if err != nil {
		return fmt.Errorf("Error unmarshaling catalog object: %w", err)
	}

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
	case CatalogObjectTypeTimePeriod:
		c.CatalogObjectType = cJson.TimePeriodData
	case CatalogObjectTypeProductSet:
		c.CatalogObjectType = cJson.ProductSetData
	case CatalogObjectTypePricingRule:
		c.CatalogObjectType = cJson.PricingRuleData
	case CatalogObjectTypeSubscriptionPlan:
		c.CatalogObjectType = cJson.SubscriptionPlanData
	case CatalogObjectTypeItemOption:
		c.CatalogObjectType = cJson.ItemOptionData
	case CatalogObjectTypeItemOptionVal:
		c.CatalogObjectType = cJson.ItemOptionValueData
	case CatalogObjectTypeCustomAttributeDefinition:
		c.CatalogObjectType = cJson.CustomAttributeDefinitionData
	case CatalogObjectTypeQuickAmountsSettings:
		c.CatalogObjectType = cJson.QuickAmountsSettingsData
	default:
		return fmt.Errorf("Found unknown catalog object type %s", cJson.Type)
	}
	return nil
}
