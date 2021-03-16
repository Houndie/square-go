package objects

type CatalogProductSet struct {
	AllProducts   bool     `json:"all_products,omitempty"`
	Name          string   `json:"name,omitempty"`
	ProductIDsAll []string `json:"product_ids_all,omitempty"`
	ProductIDsAny []string `json:"product_ids_any,omitempty"`
	QuantityExact int      `json:"quantity_exact,omitempty"`
	QuantityMax   int      `json:"quantity_max,omitempty"`
	QuantityMin   int      `json:"quantity_min,omitempty"`
}

func (*CatalogProductSet) isCatalogObjectType() {}
