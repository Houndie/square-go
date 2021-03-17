package inventory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type BatchChangeRequest struct {
	IdempotencyKey        string                     `json:"idempotency_key,omitempty"`
	Changes               []*objects.InventoryChange `json:"changes,omitempty"`
	IgnoreUnchangedCounts *bool                      `json:"ignore_unchanged_counts,omitempty"`
}

type BatchChangeResponse struct {
	Counts []*objects.InventoryCount `json:"counts"`
}

func (c *client) BatchChange(ctx context.Context, req *BatchChangeRequest) (*BatchChangeResponse, error) {
	externalRes := &BatchChangeResponse{}
	res := struct {
		internal.WithErrors
		*BatchChangeResponse
	}{
		BatchChangeResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "inventory/batch-change", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}
	return externalRes, nil
}
