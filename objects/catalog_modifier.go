package objects

type CatalogModifier struct {
	Name           string `json:"name,omitempty"`
	PriceMoney     *Money `json:"price_money,omitempty"`
	Ordinal        int    `json:"ordinal,omitempty"`
	ModifierListID string `json:"modifier_list_id,omitempty"`
}

func (*CatalogModifier) isCatalogObjectType() {}
