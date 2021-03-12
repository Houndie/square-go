package objects

import "time"

type InventoryState string

const (
	InventoryStateCustom             InventoryState = "CUSTOM"
	InventoryStateInStock            InventoryState = "IN_STOCK"
	InventoryStateSold               InventoryState = "SOLD"
	InventoryStateReturnedByCustomer InventoryState = "RETURNED_BY_CUSTOMER"
	InventoryStateReservedForSale    InventoryState = "RESERVED_FOR_SALE"
	InventoryStateSoldOnline         InventoryState = "SOLD_ONLINE"
	InventoryStateOrderedFromVendor  InventoryState = "ORDERED_FROM_VENDOR"
	InventoryStateReceivedFromVendor InventoryState = "RECEIVED_FROM_VENDOR"
	InventoryStateInTransitTo        InventoryState = "IN_TRANSIT_TO"
	InventoryStateNone               InventoryState = "NONE"
	InventoryStateWaste              InventoryState = "WASTE"
	InventoryStateUnlinkedReturn     InventoryState = "UNLINKED_RETURN"
)

type InventoryCount struct {
	CatalogObjectID   string            `json:"catalog_object_id,omitempty"`
	CatalogObjectType CatalogObjectType `json:"catalog_object_type,omitempty"`
	State             InventoryState    `json:"state,omitempty"`
	LocationID        string            `json:"location_id,omitempty"`
	Quantity          string            `json:"quantity,omitempty"`
	CalculatedAt      *time.Time        `json:"calculated_at,omitempty"`
}
