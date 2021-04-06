package objects

import (
	"encoding/json"
	"fmt"
)

type CatalogPricingType string

const (
	CatalogPricingTypeFixed    CatalogPricingType = "FIXED_PRICING"
	CatalogPricingTypeVariable CatalogPricingType = "VARIABLE_PRICING"
)

type CatalogItemVariation struct {
	ItemID                  string                                    `json:"item_id,omitempty"`
	Name                    string                                    `json:"name,omitempty"`
	Sku                     string                                    `json:"sku,omitempty"`
	Upc                     string                                    `json:"string,omitempty"`
	Ordinal                 int                                       `json:"ordinal,omitempty"`
	PricingType             CatalogPricingType                        `json:"pricing_type,omitempty"`
	PriceMoney              *Money                                    `json:"price_money,omitempty"`
	LocationOverrides       []*ItemVariationLocationOverrides         `json:"location_overrides,omitempty"`
	TrackInventory          bool                                      `json:"track_inventory,omitempty"`
	InventoryAlertType      string                                    `json:"inventory_alert_type,omitempty"`
	InventoryAlertThreshold int                                       `json:"inventory_alert_threshold,omitempty"`
	UserData                string                                    `json:"user_data,omitempty"`
	ServiceDuration         int                                       `json:"service_duration,omitempty"`
	AvailableForBooking     bool                                      `json:"available_for_booking,omitempty"`
	ItemOptionValues        []*CatalogItemOptionValueForItemVariation `json:"item_option_values,omitempty"`
	MeasurementUnitID       string                                    `json:"measurement_unit_id,omitempty"`
	TeamMemberIDs           []string                                  `json:"team_member_ids,omitempty"`
}

type InventoryAlertType interface {
	isInventoryAlertType()
}

type InventoryAlertTypeNone struct{}
type InventoryAlertTypeLowQuantity struct {
	Threshold int
}

func (*InventoryAlertTypeNone) isInventoryAlertType()        {}
func (*InventoryAlertTypeLowQuantity) isInventoryAlertType() {}

type ItemVariationLocationOverrides struct {
	LocationID         string             `json:"location_id,omitempty"`
	PriceMoney         *Money             `json:"price_money,omitempty"`
	PricingType        string             `json:"pricing_type,omitempty"`
	TrackInventory     bool               `json:"track_inventory,omitempty"`
	InventoryAlertType InventoryAlertType `json:"-"`
}

type itemVariationLocationOverridesAlias ItemVariationLocationOverrides

type inventoryAlertType string

const (
	inventoryAlertTypeNone        inventoryAlertType = "NONE"
	inventoryAlertTypeLowQuantity inventoryAlertType = "LOW_QUANTITY"
)

type itemVariationLocationOverrides struct {
	*itemVariationLocationOverridesAlias
	InventoryAlertType      inventoryAlertType `json:"inventory_alert_type,omitempty"`
	InventoryAlertThreshold int                `json:"inventory_alert_threshold,omitempty"`
}

func (o *ItemVariationLocationOverrides) MarshalJSON() ([]byte, error) {
	jsonType := itemVariationLocationOverrides{
		itemVariationLocationOverridesAlias: (*itemVariationLocationOverridesAlias)(o),
	}

	switch t := o.InventoryAlertType.(type) {
	case *InventoryAlertTypeNone:
		jsonType.InventoryAlertType = inventoryAlertTypeNone
	case *InventoryAlertTypeLowQuantity:
		jsonType.InventoryAlertType = inventoryAlertTypeLowQuantity
		jsonType.InventoryAlertThreshold = t.Threshold
	default:
		return nil, fmt.Errorf("unknown inventory alert type found")
	}

	b, err := json.Marshal(jsonType)
	if err != nil {
		return nil, fmt.Errorf("error marshaling item variation location overrides: %w", err)
	}

	return b, nil
}

func (o *ItemVariationLocationOverrides) UnmarshalJSON(bytes []byte) error {
	jsonType := itemVariationLocationOverrides{
		itemVariationLocationOverridesAlias: (*itemVariationLocationOverridesAlias)(o),
	}

	if err := json.Unmarshal(bytes, &jsonType); err != nil {
		return fmt.Errorf("error unmarshaling item variation location overrides: %w", err)
	}

	switch jsonType.InventoryAlertType {
	case inventoryAlertTypeNone:
		o.InventoryAlertType = &InventoryAlertTypeNone{}
	case inventoryAlertTypeLowQuantity:
		o.InventoryAlertType = &InventoryAlertTypeLowQuantity{
			Threshold: jsonType.InventoryAlertThreshold,
		}
	default:
		return fmt.Errorf("unknown inventory alert type found: %s", jsonType.InventoryAlertType)
	}

	return nil
}

type CatalogItemOptionValueForItemVariation struct {
	ItemOptionID      string `json:"item_option_id,omitempty"`
	ItemOptionValueID string `json:"item_option_value_id,omitempty"`
}

func (*CatalogItemVariation) isCatalogObjectType() {}
