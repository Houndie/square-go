package objects

type CatalogQuickAmountSettingsOption string

const (
	CatalogQuickAmountSettingsOptionDisabled CatalogQuickAmountSettingsOption = "DISABLED"
	CatalogQuickAmountSettingsOptionAuto     CatalogQuickAmountSettingsOption = "AUTO"
	CatalogQuickAmountSettingsOptionManual   CatalogQuickAmountSettingsOption = "MANUAL"
)

type CatalogQuickAmountsSettings struct {
	Option                 CatalogQuickAmountSettingsOption `json:"option,omitempty"`
	Amounts                []*CatalogQuickAmount            `json:"amounts,omitempty"`
	EligibleForAutoAmounts bool                             `json:"eligible_for_auto_amounts,omitempty"`
}

func (*CatalogQuickAmountsSettings) isCatalogObjectType() {}
