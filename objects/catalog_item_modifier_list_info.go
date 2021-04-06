package objects

type CatalogItemModifierListInfo struct {
	ModifierListID       string                     `json:"modifier_list_id,omitempty"`
	ModifierOverrides    []*CatalogModifierOverride `json:"modifier_overrides,omitempty"`
	MinSelectedModifiers int                        `json:"min_selected_modifiers,omitempty"`
	MaxSelectedModifiers int                        `json:"max_selected_modifiers,omitempty"`
	Enabled              *bool                      `json:"enabled,omitempty"`
}

type CatalogModifierOverride struct {
	ModifierID  string `json:"modifier_id,omitempty"`
	OnByDefault bool   `json:"on_by_default,omitempty"`
}
