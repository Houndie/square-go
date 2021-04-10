package objects

type CatalogModifierListSelectionType string

const (
	CatalogModifierListSelectionTypeSingle   CatalogModifierListSelectionType = "SINGLE"
	CatalogModifierListSelectionTypeMultiple CatalogModifierListSelectionType = "MULTIPLE"
)

type CatalogModifierList struct {
	Name          string                           `json:"name,omitempty"`
	Ordinal       int                              `json:"ordinal,omitempty"`
	SelectionType CatalogModifierListSelectionType `json:"selection_type,omitempty"`
	Modifiers     []*CatalogObject                 `json:"modifiers,omitempty"`
}

func (*CatalogModifierList) isCatalogObjectType() {}
