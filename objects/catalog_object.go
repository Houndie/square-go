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
	Type                  catalogObjectType                       `json:"-"`
}

func (c *CatalogObject) MarshalJSON() ([]byte, error) {
	cJSON := catalogObject{
		catalogObjectAlias: (*catalogObjectAlias)(c),
	}
	switch t := c.Type.(type) {
	case *CatalogItem:
		cJSON.ItemData = t
		cJSON.Type = CatalogObjectTypeItem
	case *CatalogCategory:
		cJSON.CategoryData = t
		cJSON.Type = CatalogObjectTypeCategory
	case *CatalogItemVariation:
		cJSON.ItemVariationData = t
		cJSON.Type = CatalogObjectTypeItemVariation
	case *CatalogTax:
		cJSON.TaxData = t
		cJSON.Type = CatalogObjectTypeTax
	case *CatalogDiscount:
		cJSON.DiscountData = t
		cJSON.Type = CatalogObjectTypeDiscount
	case *CatalogModifierList:
		cJSON.ModifierListData = t
		cJSON.Type = CatalogObjectTypeModifierList
	case *CatalogModifier:
		cJSON.ModifierData = t
		cJSON.Type = CatalogObjectTypeModifier
	case *CatalogImage:
		cJSON.ImageData = t
		cJSON.Type = CatalogObjectTypeImage
	case *CatalogMeasurementUnit:
		cJSON.MeasurementUnitData = t
		cJSON.Type = CatalogObjectTypeMeasurementUnit
	case *CatalogTimePeriod:
		cJSON.TimePeriodData = t
		cJSON.Type = CatalogObjectTypeTimePeriod
	case *CatalogProductSet:
		cJSON.ProductSetData = t
		cJSON.Type = CatalogObjectTypeProductSet
	case *CatalogPricingRule:
		cJSON.PricingRuleData = t
		cJSON.Type = CatalogObjectTypePricingRule
	case *CatalogSubscriptionPlan:
		cJSON.SubscriptionPlanData = t
		cJSON.Type = CatalogObjectTypeSubscriptionPlan
	case *CatalogItemOption:
		cJSON.ItemOptionData = t
		cJSON.Type = CatalogObjectTypeItemOption
	case *CatalogItemOptionValue:
		cJSON.ItemOptionValueData = t
		cJSON.Type = CatalogObjectTypeItemOptionVal
	case *CatalogCustomAttributeDefinition:
		cJSON.CustomAttributeDefinitionData = t
		cJSON.Type = CatalogObjectTypeCustomAttributeDefinition
	case *CatalogQuickAmountsSettings:
		cJSON.QuickAmountsSettingsData = t
		cJSON.Type = CatalogObjectTypeQuickAmountsSettings
	default:
		return nil, errors.New("found unknown catalog object data type")
	}

	json, err := json.Marshal(&cJSON)
	if err != nil {
		return nil, fmt.Errorf("error marshaling json catalog object: %w", err)
	}

	return json, nil
}

func (c *CatalogObject) UnmarshalJSON(data []byte) error {
	cJSON := catalogObject{
		catalogObjectAlias: (*catalogObjectAlias)(c),
	}

	err := json.Unmarshal(data, &cJSON)
	if err != nil {
		return fmt.Errorf("error unmarshaling catalog object: %w", err)
	}

	switch cJSON.Type {
	case CatalogObjectTypeItem:
		c.Type = cJSON.ItemData
	case CatalogObjectTypeCategory:
		c.Type = cJSON.CategoryData
	case CatalogObjectTypeItemVariation:
		c.Type = cJSON.ItemVariationData
	case CatalogObjectTypeTax:
		c.Type = cJSON.TaxData
	case CatalogObjectTypeDiscount:
		c.Type = cJSON.DiscountData
	case CatalogObjectTypeModifierList:
		c.Type = cJSON.ModifierListData
	case CatalogObjectTypeModifier:
		c.Type = cJSON.ModifierData
	case CatalogObjectTypeImage:
		c.Type = cJSON.ImageData
	case CatalogObjectTypeMeasurementUnit:
		c.Type = cJSON.MeasurementUnitData
	case CatalogObjectTypeTimePeriod:
		c.Type = cJSON.TimePeriodData
	case CatalogObjectTypeProductSet:
		c.Type = cJSON.ProductSetData
	case CatalogObjectTypePricingRule:
		c.Type = cJSON.PricingRuleData
	case CatalogObjectTypeSubscriptionPlan:
		c.Type = cJSON.SubscriptionPlanData
	case CatalogObjectTypeItemOption:
		c.Type = cJSON.ItemOptionData
	case CatalogObjectTypeItemOptionVal:
		c.Type = cJSON.ItemOptionValueData
	case CatalogObjectTypeCustomAttributeDefinition:
		c.Type = cJSON.CustomAttributeDefinitionData
	case CatalogObjectTypeQuickAmountsSettings:
		c.Type = cJSON.QuickAmountsSettingsData
	default:
		return fmt.Errorf("found unknown catalog object type %s", cJSON.Type)
	}

	return nil
}
