package objects

type CatalogInfoResponseLimits struct {
	BatchDeleteMaxObjectIDs                         int `json:"batch_delete_max_object_ids,omitempty"`
	BatchRetrieveMaxObjectIDs                       int `json:"batch_retrieve_max_object_ids,omitempty"`
	BatchUpsertMaxObjectsPerBatch                   int `json:"batch_upsert_max_objects_per_batch,omitempty"`
	BatchUpsertMaxTotalObjects                      int `json:"batch_upsert_max_total_objects,omitempty"`
	SearchMaxPageLimit                              int `json:"search_max_page_limit,omitempty"`
	UpdateItemModifierListsMaxItemIDs               int `json:"update_item_modifier_lists_max_item_ids,omitempty"`
	UpdateItemModifierListsMaxModiferListsToDisable int `json:"update_item_modifier_lists_max_modifier_lists_to_disable"`
	UpdateItemModifierListsMaxModiferListsToEnable  int `json:"update_item_modifier_lists_max_modifier_lists_to_enable"`
	UpdateItemTaxesMaxItemIDs                       int `json:"update_item_taxes_max_item_ids,omitempty"`
	UpdateItemTaxesMaxTaxesToDisable                int `json:"update_item_taxes_max_taxes_to_disable,omitempty"`
	UpdateItemTaxesMaxTaxesToEnable                 int `json:"update_item_taxes_max_taxes_to_enable,omitempty"`
}
