package objects

type CatalogTimePeriod struct {
	Event string `json:"event,omitempty"`
}

func (*CatalogTimePeriod) isCatalogObjectType() {}
