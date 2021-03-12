package objects

type CardBrand string

const (
	CardBrandOtherBrand      CardBrand = "OTHER_BRAND"
	CardBrandVisa            CardBrand = "VISA"
	CardBrandMastercard      CardBrand = "MASTERCARD"
	CardBrandAmericanExpress CardBrand = "AMERICAN_EXPRESS"
	CardBrandDiscover        CardBrand = "DISCOVER"
	CardBrandDiscoverDiners  CardBrand = "DINERS"
	CardBrandJcb             CardBrand = "JCB"
	CardBrandChinaUnionpay   CardBrand = "CHINA_UNIONPAY"
	CardBrandSquareGiftCard  CardBrand = "SQUARE_GIFT_CARD"
)

type Card struct {
	ID             string    `json:"string,omitempty"`
	CardBrand      CardBrand `json:"card_brand,omitempty"`
	Last4          string    `json:"last_4,omitempty"`
	ExpMonth       int       `json:"exp_month,omitempty"`
	ExpYear        int       `json:"exp_year,omitempty"`
	CardholderName string    `json:"cardholder_name,omitempty"`
	BillingAddress Address   `json:"address,omitempty"`
	Fingerprint    string    `json:"fingerprint,omitempty"`
}
