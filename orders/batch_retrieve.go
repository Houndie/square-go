package orders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type BatchRetrieveRequest struct {
	LocationID string   `json:"-"`
	OrderIDs   []string `json:"order_ids"`
}

type BatchRetrieveResponse struct {
	Orders []*objects.Order `json:"orders"`
}

func (c *client) BatchRetrieve(ctx context.Context, req *BatchRetrieveRequest) (*BatchRetrieveResponse, error) {
	externalRes := &BatchRetrieveResponse{}
	res := struct {
		internal.WithErrors
		*BatchRetrieveResponse
	}{
		BatchRetrieveResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "locations/"+req.LocationID+"/orders/batch-retrieve", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
