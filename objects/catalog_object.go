package objects

import (
	"encoding/json"
	"fmt"
	"time"

	"errors"
)

type CatalogObjectEnumType string

const (
	CatalogObjectEnumTypeItem                      CatalogObjectEnumType = "ITEM"
	CatalogObjectEnumTypeImage                     CatalogObjectEnumType = "IMAGE"
	CatalogObjectEnumTypeCategory                  CatalogObjectEnumType = "CATEGORY"
	CatalogObjectEnumTypeItemVariation             CatalogObjectEnumType = "ITEM_VARIATION"
	CatalogObjectEnumTypeTax                       CatalogObjectEnumType = "TAX"
	CatalogObjectEnumTypeDiscount                  CatalogObjectEnumType = "DISCOUNT"
	CatalogObjectEnumTypeModifierList              CatalogObjectEnumType = "MODIFIER_LIST"
	CatalogObjectEnumTypeModifier                  CatalogObjectEnumType = "MODIFIER"
	CatalogObjectEnumTypePricingRule               CatalogObjectEnumType = "PRICING_RULE"
	CatalogObjectEnumTypeProductSet                CatalogObjectEnumType = "PRODUCT_SET"
	CatalogObjectEnumTypeTimePeriod                CatalogObjectEnumType = "TIME_PERIOD"
	CatalogObjectEnumTypeMeasurementUnit           CatalogObjectEnumType = "MEASUREMENT_UNIT"
	CatalogObjectEnumTypeSubscriptionPlan          CatalogObjectEnumType = "SUBSCRIPTION_PLAN"
	CatalogObjectEnumTypeItemOption                CatalogObjectEnumType = "ITEM_OPTION"
	CatalogObjectEnumTypeItemOptionVal             CatalogObjectEnumType = "ITEM_OPTION_VAL"
	CatalogObjectEnumTypeCustomAttributeDefinition CatalogObjectEnumType = "CUSTOM_ATTRIBUTE_DEFINITION"
	CatalogObjectEnumTypeQuickAmountsSettings      CatalogObjectEnumType = "QUICK_AMOUNTS_SETTINGS"
)

type catalogObjectAlias CatalogObject

type catalogObject struct {
	*catalogObjectAlias
	Type                          CatalogObjectEnumType             `json:"type,omitempty"`
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

type CatalogObjectType interface {
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
	Type                  CatalogObjectType                       `json:"-"`
}

func (c *CatalogObject) MarshalJSON() ([]byte, error) {
	cJSON := catalogObject{
		catalogObjectAlias: (*catalogObjectAlias)(c),
	}
	switch t := c.Type.(type) {
	case *CatalogItem:
		cJSON.ItemData = t
		cJSON.Type = CatalogObjectEnumTypeItem
	case *CatalogCategory:
		cJSON.CategoryData = t
		cJSON.Type = CatalogObjectEnumTypeCategory
	case *CatalogItemVariation:
		cJSON.ItemVariationData = t
		cJSON.Type = CatalogObjectEnumTypeItemVariation
	case *CatalogTax:
		cJSON.TaxData = t
		cJSON.Type = CatalogObjectEnumTypeTax
	case *CatalogDiscount:
		cJSON.DiscountData = t
		cJSON.Type = CatalogObjectEnumTypeDiscount
	case *CatalogModifierList:
		cJSON.ModifierListData = t
		cJSON.Type = CatalogObjectEnumTypeModifierList
	case *CatalogModifier:
		cJSON.ModifierData = t
		cJSON.Type = CatalogObjectEnumTypeModifier
	case *CatalogImage:
		cJSON.ImageData = t
		cJSON.Type = CatalogObjectEnumTypeImage
	case *CatalogMeasurementUnit:
		cJSON.MeasurementUnitData = t
		cJSON.Type = CatalogObjectEnumTypeMeasurementUnit
	case *CatalogTimePeriod:
		cJSON.TimePeriodData = t
		cJSON.Type = CatalogObjectEnumTypeTimePeriod
	case *CatalogProductSet:
		cJSON.ProductSetData = t
		cJSON.Type = CatalogObjectEnumTypeProductSet
	case *CatalogPricingRule:
		cJSON.PricingRuleData = t
		cJSON.Type = CatalogObjectEnumTypePricingRule
	case *CatalogSubscriptionPlan:
		cJSON.SubscriptionPlanData = t
		cJSON.Type = CatalogObjectEnumTypeSubscriptionPlan
	case *CatalogItemOption:
		cJSON.ItemOptionData = t
		cJSON.Type = CatalogObjectEnumTypeItemOption
	case *CatalogItemOptionValue:
		cJSON.ItemOptionValueData = t
		cJSON.Type = CatalogObjectEnumTypeItemOptionVal
	case *CatalogCustomAttributeDefinition:
		cJSON.CustomAttributeDefinitionData = t
		cJSON.Type = CatalogObjectEnumTypeCustomAttributeDefinition
	case *CatalogQuickAmountsSettings:
		cJSON.QuickAmountsSettingsData = t
		cJSON.Type = CatalogObjectEnumTypeQuickAmountsSettings
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
	case CatalogObjectEnumTypeItem:
		c.Type = cJSON.ItemData
	case CatalogObjectEnumTypeCategory:
		c.Type = cJSON.CategoryData
	case CatalogObjectEnumTypeItemVariation:
		c.Type = cJSON.ItemVariationData
	case CatalogObjectEnumTypeTax:
		c.Type = cJSON.TaxData
	case CatalogObjectEnumTypeDiscount:
		c.Type = cJSON.DiscountData
	case CatalogObjectEnumTypeModifierList:
		c.Type = cJSON.ModifierListData
	case CatalogObjectEnumTypeModifier:
		c.Type = cJSON.ModifierData
	case CatalogObjectEnumTypeImage:
		c.Type = cJSON.ImageData
	case CatalogObjectEnumTypeMeasurementUnit:
		c.Type = cJSON.MeasurementUnitData
	case CatalogObjectEnumTypeTimePeriod:
		c.Type = cJSON.TimePeriodData
	case CatalogObjectEnumTypeProductSet:
		c.Type = cJSON.ProductSetData
	case CatalogObjectEnumTypePricingRule:
		c.Type = cJSON.PricingRuleData
	case CatalogObjectEnumTypeSubscriptionPlan:
		c.Type = cJSON.SubscriptionPlanData
	case CatalogObjectEnumTypeItemOption:
		c.Type = cJSON.ItemOptionData
	case CatalogObjectEnumTypeItemOptionVal:
		c.Type = cJSON.ItemOptionValueData
	case CatalogObjectEnumTypeCustomAttributeDefinition:
		c.Type = cJSON.CustomAttributeDefinitionData
	case CatalogObjectEnumTypeQuickAmountsSettings:
		c.Type = cJSON.QuickAmountsSettingsData
	default:
		return fmt.Errorf("found unknown catalog object type %s", cJSON.Type)
	}

	return nil
}
