package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type UpsertObjectRequest struct {
	IdempotencyKey string                 `json:"idempotency_key,omitempty"`
	Object         *objects.CatalogObject `json:"object,omitempty"`
}

type UpsertObjectResponse struct {
	CatalogObject *objects.CatalogObject      `json:"catalog_object,omitempty"`
	IDMappings    []*objects.CatalogIDMapping `json:"id_mappings,omitempty"`
}

func (c *client) UpsertObject(ctx context.Context, req *UpsertObjectRequest) (*UpsertObjectResponse, error) {
	externalRes := &UpsertObjectResponse{}
	res := struct {
		internal.WithErrors
		*UpsertObjectResponse
	}{
		UpsertObjectResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "/catalog/object", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
