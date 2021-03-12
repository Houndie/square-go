package checkout

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

func (c *client) Create(ctx context.Context, locationID, idempotencyKey string, order *objects.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail string, prePopulateBuyerEmail string, prePopulateShippingAddress *objects.Address, redirectUrl string, additionalRecipients []*objects.ChargeRequestAdditionalRecipient, note string) (*objects.Checkout, error) {
	req := struct {
		IdempotencyKey             string                                      `json:"idempotency_key,omitempty"`
		Order                      *objects.CreateOrderRequest                 `json:"order,omitempty"`
		AskForShippingAddress      bool                                        `json:"ask_for_shipping_address,omitempty"`
		MerchantSupportEmail       string                                      `json:"merchant_support_email,omitempty"`
		PrePopulateBuyerEmail      string                                      `json:"pre_populate_buyer_email,omitempty"`
		PrePopulateShippingAddress *objects.Address                            `json:"pre_populate_shipping_address,omitempty"`
		RedirectUrl                string                                      `json:"redirect_url,omitempty"`
		AdditionalRecipients       []*objects.ChargeRequestAdditionalRecipient `json:"additional_recipients,omitempty"`
		Note                       string                                      `json:"note,omitempty"`
	}{
		IdempotencyKey:             idempotencyKey,
		Order:                      order,
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectUrl:                redirectUrl,
		AdditionalRecipients:       additionalRecipients,
		Note:                       note,
	}
	res := struct {
		internal.WithErrors
		Checkout *objects.Checkout `json:"checkout"`
	}{}

	if err := c.i.Do(ctx, http.MethodPost, c.i.Endpoint("locations/"+locationID+"/checkouts").String(), req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return res.Checkout, nil
}
