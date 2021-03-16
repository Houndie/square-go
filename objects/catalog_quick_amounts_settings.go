package objects

type CatalogQuickAmountsSettings struct {
	Option                 string                `json:"option,omitempty"`
	Amounts                []*CatalogQuickAmount `json:"amounts,omitempty"`
	EligibleForAutoAmounts bool                  `json:"eligible_for_auto_amounts,omitempty"`
}

func (*CatalogQuickAmountsSettings) isCatalogObjectType() {}
