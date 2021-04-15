package objects

type CatalogQuickAmountsSettingsOption string

const (
	CatalogQuickAmountsSettingsOptionDisabled CatalogQuickAmountsSettingsOption = "DISABLED"
	CatalogQuickAmountsSettingsOptionAuto     CatalogQuickAmountsSettingsOption = "AUTO"
	CatalogQuickAmountsSettingsOptionManual   CatalogQuickAmountsSettingsOption = "MANUAL"
)

type CatalogQuickAmountsSettings struct {
	Option                 CatalogQuickAmountsSettingsOption `json:"option,omitempty"`
	Amounts                []*CatalogQuickAmount             `json:"amounts,omitempty"`
	EligibleForAutoAmounts bool                              `json:"eligible_for_auto_amounts,omitempty"`
}

func (*CatalogQuickAmountsSettings) isCatalogObjectType() {}
