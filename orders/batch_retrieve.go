package orders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

func (c *client) BatchRetrieve(ctx context.Context, locationID string, orderIDs []string) ([]*objects.Order, error) {
	req := struct {
		OrderIDs []string `json:"order_ids"`
	}{
		OrderIDs: orderIDs,
	}

	res := struct {
		internal.WithErrors
		Orders []*objects.Order `json:"orders"`
	}{}

	if err := c.i.Do(ctx, http.MethodPost, c.i.Endpoint("locations/"+locationID+"/orders/batch-retrieve").String(), req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}
	return res.Orders, nil
}
