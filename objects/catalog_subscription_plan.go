package objects

type CatalogSubscriptionPlan struct {
	Name   string               `json:"name,omitempty"`
	Phases []*SubscriptionPhase `json:"phases,omitempty"`
}

func (*CatalogSubscriptionPlan) isCatalogObjectType() {}
