package checkout

import (
	"context"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type Client interface {
	Create(ctx context.Context, locationID, idempotencyKey string, order *objects.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail string, prePopulateBuyerEmail string, prePopulateShippingAddress *objects.Address, redirectUrl string, additionalRecipients []*objects.ChargeRequestAdditionalRecipient, note string) (*objects.Checkout, error)
}

type client struct {
	i *internal.Client
}

func NewClient(apiKey string, environment objects.Environment, httpClient *http.Client) (Client, error) {
	c, err := internal.NewClient(apiKey, environment, httpClient)
	if err != nil {
		return nil, err
	}
	return &client{i: c}, nil
}
