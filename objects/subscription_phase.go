package objects

type SubscriptionCadence string

const (
	SubscriptionCadenceDaily           SubscriptionCadence = "DAILY"
	SubscriptionCadenceWeekly          SubscriptionCadence = "WEEKLY"
	SubscriptionCadenceEveryTwoWeeks   SubscriptionCadence = "EVERY_TWO_WEEKS"
	SubscriptionCadenceThirtyDays      SubscriptionCadence = "THIRTY_DAYS"
	SubscriptionCadenceSixtyDays       SubscriptionCadence = "SIXTY_DAYS"
	SubscriptionCadenceNinetyDays      SubscriptionCadence = "NINETY_DAYS"
	SubscriptionCadenceMonthly         SubscriptionCadence = "MONTHLY"
	SubscriptionCadenceEveryTwoMonths  SubscriptionCadence = "EVERY_TWO_MONTHS"
	SubscriptionCadenceQuarterly       SubscriptionCadence = "QUARTERLY"
	SubscriptionCadenceEveryFourMonths SubscriptionCadence = "EVERY_FOUR_MONTHS"
	SubscriptionCadenceEverySixMonths  SubscriptionCadence = "EVERY_SIX_MONTHS"
	SubscriptionCadenceAnnual          SubscriptionCadence = "ANNUAL"
	SubscriptionCadenceEveryTwoYears   SubscriptionCadence = "EVERY_TWO_YEARS"
)

type SubscriptionPhase struct {
	Cadence             SubscriptionCadence `json:"cadence,omitempty"`
	RecurringPriceMoney *Money              `json:"recurring_price_money,omitempty"`
	Ordinal             int                 `json:"ordinal,omitempty"`
	Periods             int                 `json:"periods,omitempty"`
	UID                 string              `json:"uid,omitempty"`
}
