package objects

type CatalogModifier struct {
	Name       string `json:"name,omitempty"`
	PriceMoney *Money `json:"price_money,omitempty"`
}

func (*CatalogModifier) isCatalogObjectType() {}
