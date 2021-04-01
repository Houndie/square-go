package catalog

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type BatchUpsertRequest struct {
	IdempotencyKey string                        `json:"idempotency_key,omitempty"`
	Batches        []*objects.CatalogObjectBatch `json:"batches,omitempty"`
}

type BatchUpsertResponse struct {
	Objects    []*objects.CatalogObject    `json:"objects"`
	UpdatedAt  *time.Time                  `json:"updated_at"`
	IDMappings []*objects.CatalogIDMapping `json:"id_mappings"`
}

func (c *client) BatchUpsert(ctx context.Context, req *BatchUpsertRequest) (*BatchUpsertResponse, error) {
	externalRes := &BatchUpsertResponse{}
	res := struct {
		internal.WithErrors
		*BatchUpsertResponse
	}{
		BatchUpsertResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "/catalog/batch-upsert", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
