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
	toJSON := &inventoryChange{}
	switch t := c.Type.(type) {
	case *InventoryAdjustment:
		toJSON.Adjustment = t
		toJSON.Type = inventoryChangeTypeAdjustment
	case *InventoryPhysicalCount:
		toJSON.PhysicalCount = t
		toJSON.Type = inventoryChangeTypePhysicalCount
	case *InventoryTransfer:
		toJSON.Transfer = t
		toJSON.Type = inventoryChangeTypeTransfer
	default:
		return nil, fmt.Errorf("unknown inventory type: %T", c.Type)
	}

	b, err := json.Marshal(&toJSON)
	if err != nil {
		return nil, fmt.Errorf("error marshaling inventory change: %w", err)
	}

	return b, nil
}

func (c *InventoryChange) UnmarshalJSON(b []byte) error {
	fromJSON := &inventoryChange{}
	if err := json.Unmarshal(b, &fromJSON); err != nil {
		return fmt.Errorf("error unmarshaling inventory change: %w", err)
	}

	switch fromJSON.Type {
	case inventoryChangeTypePhysicalCount:
		c.Type = fromJSON.PhysicalCount
		return nil
	case inventoryChangeTypeAdjustment:
		c.Type = fromJSON.Adjustment
		return nil
	case inventoryChangeTypeTransfer:
		c.Type = fromJSON.Transfer
		return nil
	}

	return fmt.Errorf("unknown type: %s", fromJSON.Type)
}
