package objects

import (
	"encoding/json"
	"fmt"
)

type InventoryChangeType interface {
	isInventoryChangeType()
}

type InventoryChange struct {
	Type InventoryChangeType
}

type inventoryChangeType string

const (
	inventoryChangeTypePhysicalCount inventoryChangeType = "PHYSICAL_COUNT"
	inventoryChangeTypeAdjustment    inventoryChangeType = "ADJUSTMENT"
	inventoryChangeTypeTransfer      inventoryChangeType = "TRANSFER"
)

type inventoryChange struct {
	Adjustment    *InventoryAdjustment    `json:"adjustment,omitempty"`
	PhysicalCount *InventoryPhysicalCount `json:"physical_count,omitempty"`
	Transfer      *InventoryTransfer      `json:"inventory_transfer,omitempty"`
	Type          inventoryChangeType     `json:"type,omitempty"`
}

func (c *InventoryChange) MarshalJSON() ([]byte, error) {
	toJson := &inventoryChange{}
	switch t := c.Type.(type) {
	case *InventoryAdjustment:
		toJson.Adjustment = t
		toJson.Type = inventoryChangeTypeAdjustment
	case *InventoryPhysicalCount:
		toJson.PhysicalCount = t
		toJson.Type = inventoryChangeTypePhysicalCount
	case *InventoryTransfer:
		toJson.Transfer = t
		toJson.Type = inventoryChangeTypeTransfer
	default:
		return nil, fmt.Errorf("unknown inventory type: %T", c.Type)
	}

	b, err := json.Marshal(&toJson)
	if err != nil {
		return nil, fmt.Errorf("error marshaling inventory change: %w", err)
	}
	return b, nil
}

func (c *InventoryChange) UnmarshalJSON(b []byte) error {
	fromJson := &inventoryChange{}
	if err := json.Unmarshal(b, &fromJson); err != nil {
		return fmt.Errorf("error unmarshaling inventory change: %w", err)
	}
	switch fromJson.Type {
	case inventoryChangeTypePhysicalCount:
		c.Type = fromJson.PhysicalCount
		return nil
	case inventoryChangeTypeAdjustment:
		c.Type = fromJson.Adjustment
		return nil
	case inventoryChangeTypeTransfer:
		c.Type = fromJson.Transfer
		return nil
	}
	return fmt.Errorf("unknown type: %s", fromJson.Type)
}
