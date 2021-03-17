package checkout

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type CreateRequest struct {
	LocationID                 string                                      `json:"-"`
	IdempotencyKey             string                                      `json:"idempotency_key,omitempty"`
	Order                      *objects.CreateOrderRequest                 `json:"order,omitempty"`
	AskForShippingAddress      bool                                        `json:"ask_for_shipping_address,omitempty"`
	MerchantSupportEmail       string                                      `json:"merchant_support_email,omitempty"`
	PrePopulateBuyerEmail      string                                      `json:"pre_populate_buyer_email,omitempty"`
	PrePopulateShippingAddress *objects.Address                            `json:"pre_populate_shipping_address,omitempty"`
	RedirectURL                string                                      `json:"redirect_url,omitempty"`
	AdditionalRecipients       []*objects.ChargeRequestAdditionalRecipient `json:"additional_recipients,omitempty"`
	Note                       string                                      `json:"note,omitempty"`
}

type CreateResponse struct {
	Checkout *objects.Checkout `json:"checkout"`
}

func (c *client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	externalRes := &CreateResponse{}
	res := struct {
		internal.WithErrors
		*CreateResponse
	}{
		CreateResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, c.i.Endpoint("locations/"+req.LocationID+"/checkouts").String(), req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
