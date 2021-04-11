package objects

type CatalogItemOptionValue struct {
	Color        string `json:"color,omitempty"`
	Description  string `json:"description,omitempty"`
	ItemOptionID string `json:"item_option_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Ordinal      int    `json:"ordinal,omitempty"`
}

func (*CatalogItemOptionValue) isCatalogObjectType() {}
