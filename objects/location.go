package objects

import "time"

type LocationCapability string

const LocationCapabilityCreditCardProcessing LocationCapability = "CREDIT_CARD_PROCESSING"

type LocationStatus string

const (
	LocationStatusActive   LocationStatus = "ACTIVE"
	LocationStatusInactive LocationStatus = "INACTIVE"
)

type LocationType string

const (
	LocationTypePhysical LocationType = "PHYSICAL"
	LocationTypeMobile   LocationType = "MOBILE"
)

type Location struct {
	ID                string               `json:"id,omitempty"`
	Name              string               `json:"name,omitempty"`
	Address           *Address             `json:"address,omitempty"`
	Timezone          string               `json:"timezone,omitempty"`
	Capabilities      []LocationCapability `json:"capabilities,omitempty"`
	Status            LocationStatus       `json:"status,omitempty"`
	CreatedAt         *time.Time           `json:"created_at,omitempty"`
	MerchantID        string               `json:"merchant_id,omitempty"`
	Country           string               `json:"country,omitempty"`
	LanguageCode      string               `json:"language_code,omitempty"`
	Currency          string               `json:"currency,omitempty"`
	PhoneNumber       string               `json:"phone_number,omitempty"`
	BusinessName      string               `json:"business_name,omitempty"`
	Type              LocationType         `json:"type,omitempty"`
	WebsiteUrl        string               `json:"website_url,omitempty"`
	BusinessHours     *BusinessHours       `json:"business_hours,omitempty"`
	BusinessEmail     string               `json:"business_email,omitempty"`
	Description       string               `json:"description",omitempty`
	TwitterUsername   string               `json:"twitter_username,omitempty"`
	InstagramUsername string               `json:"instagram_username,omitempty"`
	FacebookUrl       string               `json:"facebook_url,omitempty"`
	Coordinates       *Coordinates         `json:"coordinates,omitempty"`
}
