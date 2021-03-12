package objects

type OrderLineItemModifier struct {
	Uid             string `json:"uid,omitempty"`
	CatalogObjectID string `json:"catalog_object_id,omitempty"`
	Name            string `json:"name,omitempty"`
	BasePriceMoney  *Money `json:"base_price_money,omitempty"`
	TotalPriceMoney *Money `json:"total_price_money,omitempty"`
}
