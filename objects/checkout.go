package objects

import "time"

type Checkout struct {
	ID                         string                 `json:"id,omitempty"`
	CheckoutPageUrl            string                 `json:"checkout_page_url,omitempty"`
	AskForShippingAddress      bool                   `json:"ask_for_shipping_address,omitempty"`
	MerchantSupportEmail       string                 `json:"merchant_support_email,omitempty"`
	PrePopulateBuyerEmail      string                 `json:"pre_populate_buyer_email,omitempty"`
	PrePopulateShippingAddress *Address               `json:"pre_populate_shipping_address,omitempty"`
	RedirectUrl                string                 `json:"redirect_url,omitempty"`
	Order                      *Order                 `json:"order,omitempty"`
	CreatedAt                  *time.Time             `json:"created_at,omitempty"`
	AdditionalRecipients       []*AdditionalRecipient `json:"additional_recipients,omitempty"`
}
