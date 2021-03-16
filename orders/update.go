package orders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

func (c *client) Update(ctx context.Context, locationID, orderID string, order *objects.Order, fieldsToClear []string, idempotencyKey string) (*objects.Order, error) {
	req := struct {
		Order          *objects.Order `json:"order"`
		FieldsToClear  []string       `json:"fields_to_clear"`
		IdempotencyKey string         `json:"idempotency_key"`
	}{
		Order:          order,
		FieldsToClear:  fieldsToClear,
		IdempotencyKey: idempotencyKey,
	}

	res := struct {
		internal.WithErrors
		Order *objects.Order `json:"order"`
	}{}

	if err := c.i.Do(ctx, http.MethodPut, c.i.Endpoint("/locations/"+locationID+"/orders/"+orderID).String(), req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}
	return res.Order, nil
}
