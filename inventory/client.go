package inventory

import (
	"context"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type Client interface {
	BatchChange(ctx context.Context, idempotencyKey string, changes []*objects.InventoryChange, opts ...BatchChangeOption) ([]*objects.InventoryCount, error)
	BatchRetrieveCounts(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) BatchRetrieveCountsIterator
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
