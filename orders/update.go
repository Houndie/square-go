package orders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type UpdateRequest struct {
	LocationID     string         `json:"-"`
	OrderID        string         `json:"-"`
	Order          *objects.Order `json:"order"`
	FieldsToClear  []string       `json:"fields_to_clear"`
	IdempotencyKey string         `json:"idempotency_key"`
}

type UpdateResponse struct {
	Order *objects.Order `json:"order"`
}

func (c *client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	externalRes := &UpdateResponse{}
	res := struct {
		internal.WithErrors
		*UpdateResponse
	}{
		UpdateResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPut, c.i.Endpoint("/locations/"+req.LocationID+"/orders/"+req.OrderID).String(), req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}
	return externalRes, nil
}
