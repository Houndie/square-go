package objects

type CatalogQuery struct {
	ExactQuery                             *CatalogQueryExact                             `json:"exact_query,omitempty"`
	ItemVariationsForItemOptionValuesQuery *CatalogQueryItemVariationsForItemOptionValues `json:"item_variations_for_item_option_values_query,omitempty"`
	ItemsForItemOptionsQuery               *CatalogQueryItemsForItemOptions               `json:"items_for_item_options_query,omitempty"`
	ItemsForModifierListQuery              *CatalogQueryItemsForModifierList              `json:"items_for_modifier_list_query,omitempty"`
	ItemsForTaxQuery                       *CatalogQueryItemsForTax                       `json:"items_for_tax_query,omitempty"`
	PrefixQuery                            *CatalogQueryPrefix                            `json:"prefix_query,omitempty"`
	RangeQuery                             *CatalogQueryRange                             `json:"range_query,omitempty"`
	SetQuery                               *CatalogQuerySet                               `json:"set_query,omitempty"`
	SortedAttributeQuery                   *CatalogQuerySortedAttribute                   `json:"sorted_attribute_query,omitempty"`
	TextQuery                              *CatalogQueryText                              `json:"text_query,omitempty"`
}

type CatalogQueryExact struct {
	AttributeName  string `json:"attribute_name,omitempty"`
	AttributeValue string `json:"attribute_value,omitempty"`
}

type CatalogQueryItemVariationsForItemOptionValues struct {
	ItemOptionValueIDs []string `json:"item_option_value_ids,omitempty"`
}

type CatalogQueryItemsForItemOptions struct {
	ItemOptionIDs []string `json:"Item_option_ids,omitempty"`
}

type CatalogQueryItemsForModifierList struct {
	ModifierListIDs []string `json:"modifier_list_ids,omitempty"`
}

type CatalogQueryItemsForTax struct {
	TaxIDs []string `json:"tax_ids,omitempty"`
}

type CatalogQueryPrefix struct {
	AttributeName   string `json:"attribute_name,omitempty"`
	AttributePrefix string `json:"attribute_prefix,omitempty"`
}

type CatalogQueryRange struct {
	AttributeName     string `json:"attribute_name,omitempty"`
	AttributeMaxValue int    `json:"attribute_max_value,omitempty"`
	AttributeMinValue int    `json:"attribute_min_value,omitempty"`
}

type CatalogQuerySet struct {
	AttributeName   string   `json:"attribute_name,omitempty"`
	AttributeValues []string `json:"attribute_values,omitempty"`
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

type CatalogQuerySortedAttribute struct {
	AttributeName         string    `json:"attribute_name,omitempty"`
	InitialAttributeValue string    `json:"initial_attribute_value,omitempty"`
	SortOrder             SortOrder `json:"sort_order,omitempty"`
}

type CatalogQueryText struct {
	Keywords []string `json:"keywords,omitempty"`
}
