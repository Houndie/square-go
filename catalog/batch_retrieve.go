package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type BatchRetrieveRequest struct {
	ObjectIDs             []string `json:"object_ids,omitempty"`
	IncludeRelatedObjects bool     `json:"include_related_objects,omitempty"`
	CatalogVersion        int      `json:"catalog_version,omitempty"`
}

type BatchRetrieveResponse struct {
	Objects        []*objects.CatalogObject `json:"objects,omitempty"`
	RelatedObjects []*objects.CatalogObject `json:"related_objects"`
}

func (c *client) BatchRetrieve(ctx context.Context, req *BatchRetrieveRequest) (*BatchRetrieveResponse, error) {
	externalRes := &BatchRetrieveResponse{}
	res := struct {
		internal.WithErrors
		*BatchRetrieveResponse
	}{
		BatchRetrieveResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "/catalog/batch-retrieve", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
