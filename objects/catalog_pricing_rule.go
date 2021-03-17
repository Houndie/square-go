package objects

import (
	"time"
)

type ExcludeStrategy string

const (
	ExcludeStrategyLeastExpensive ExcludeStrategy = "LEAST_EXPENSIVE"
	ExcludeStrategyMostExpensive  ExcludeStrategy = "MOST_EXPENSIVE"
)

type CatalogPricingRule struct {
	DiscountID          string          `json:"discount_id,omitempty"`
	ExcludeProductsID   string          `json:"exclude_products_id,omitempty"`
	ExcludeStrategy     ExcludeStrategy `json:"exclude_strategy,omitempty"`
	MatchProductsID     string          `json:"match_products_id,omitempty"`
	Name                string          `json:"name,omitempty"`
	TimePeriodIDs       []string        `json:"time_period_ids,omitempty"`
	ValidFromDate       *time.Time      `json:"valid_from_date,omitempty"`
	ValidFromLocalTime  string          `json:"valid_from_local_time,omitempty"`
	ValidUntilDate      *time.Time      `json:"valid_until_date,omitempty"`
	ValidUntilLocalTime string          `json:"valid_until_local_time,omitempty"`
}

func (*CatalogPricingRule) isCatalogObjectType() {}
