package inventory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type BatchChangeOption interface {
	isBatchChangeOption()
}

type IgnoreUnchangedCounts bool

func (*IgnoreUnchangedCounts) isBatchChangeOption() {}

func (c *client) BatchChange(ctx context.Context, idempotencyKey string, changes []*objects.InventoryChange, opts ...BatchChangeOption) ([]*objects.InventoryCount, error) {
	req := struct {
		IdempotencyKey        string                     `json:"idempotency_key,omitempty"`
		Changes               []*objects.InventoryChange `json:"changes,omitempty"`
		IgnoreUnchangedCounts *bool                      `json:"ignore_unchanged_counts,omitempty"`
	}{
		IdempotencyKey: idempotencyKey,
		Changes:        changes,
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case *IgnoreUnchangedCounts:
			val := bool(*o)
			req.IgnoreUnchangedCounts = &val
		}
	}

	res := struct {
		internal.WithErrors
		Counts []*objects.InventoryCount `json:"counts"`
	}{}

	if err := c.i.Do(ctx, http.MethodPost, "inventory/batch-change", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}
	return res.Counts, nil
}
