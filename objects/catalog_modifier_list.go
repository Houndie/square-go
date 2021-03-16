package objects

type CatalogModifierList struct {
	Name          string           `json:"name,omitempty"`
	Ordinal       int              `json:"ordinal,omitempty"`
	SelectionType string           `json:"selection_type,omitempty"`
	Modifiers     []*CatalogObject `json:"modifiers,omitempty"`
}

func (*CatalogModifierList) isCatalogObjectType() {}
