package objects

type CatalogQuickAmountType string

const (
	CatalogQuickAmountTypeManual CatalogQuickAmountType = "QUICK_AMOUNT_TYPE_MANUAL"
	CatalogQuickAmountTypeAuto   CatalogQuickAmountType = "QUICK_AMOUNT_TYPE_AUTO"
)

type CatalogQuickAmount struct {
	Amount  *Money                 `json:"amount,omitempty"`
	Type    CatalogQuickAmountType `json:"type,omitempty"`
	Ordinal int                    `json:"ordinal,omitempty"`
	Score   int                    `json:"score,omitempty"`
}
