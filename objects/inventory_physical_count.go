package objects

import "time"

type InventoryPhysicalCount struct {
	ID                string             `json:"id,omitempty"`
	CatalogObjectID   string             `json:"catalog_object_id,omitempty"`
	CatalogObjectType CatalogObjectType  `json:"catalog_object_type,omitempty"`
	CreatedAt         *time.Time         `json:"created_at,omitempty"`
	EmployeeID        string             `json:"employee_id,omitempty"`
	LocationID        string             `json:"location_id,omitempty"`
	OccurredAt        *time.Time         `json:"occurred_at,omitempty"`
	Quantity          string             `json:"quanity,omitempty"`
	ReferenceID       string             `json:"reference_id,omitempty"`
	Source            *SourceApplication `json:"source,omitempty"`
	State             InventoryState     `json:"state,omitempty"`
}

func (*InventoryPhysicalCount) isInventoryChangeType() {}
