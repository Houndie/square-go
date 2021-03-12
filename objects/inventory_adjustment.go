package objects

import "time"

type InventoryAdjustment struct {
	ID                string             `json:"id,omitempty"`
	CatalogObjectID   string             `json:"catalog_object_id,omitempty"`
	CatalogObjectType CatalogObjectType  `json:"catalog_object_type,omitempty"`
	CreatedAt         *time.Time         `json:"created_at,omitempty"`
	EmployeeID        string             `json:"employee_id,omitempty"`
	FromState         InventoryState     `json:"from_state,omitempty"`
	GoodsReceiptID    string             `json:"goods_receipt_id,omitempty"`
	LocationID        string             `json:"location_id,omitempty"`
	OccurredAt        *time.Time         `json:"occurred_at,omitempty"`
	PurchaseOrderID   string             `json:"purchase_order_id,omitempty"`
	Quantity          string             `json:"quanity,omitempty"`
	ReferenceID       string             `json:"reference_id",omitempty"`
	RefundID          string             `json:"refund_id",omitempty"`
	Source            *SourceApplication `json:"source,omitempty"`
	ToState           InventoryState     `json:"to_state,omitempty"`
	TotalPriceMoney   *Money             `json:"total_price_money,omitempty"`
	TransactionID     string             `json:"transaction_id,omitempty"`
}

func (*InventoryAdjustment) isInventoryChangeType() {}
