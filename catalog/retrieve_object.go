package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type RetrieveObjectRequest struct {
	ObjectID              string `schema:"-"`
	IncludeRelatedObjects bool   `schema:"include_related_objects,omitempty"`
	CatalogVersion        int    `schema:"catalog_version,omitempty"`
}

type RetrieveObjectResponse struct {
	Object         *objects.CatalogObject   `json:"object,omitempty"`
	RelatedObjects []*objects.CatalogObject `json:"related_objects,omitempty"`
}

func (c *client) RetrieveObject(ctx context.Context, req *RetrieveObjectRequest) (*RetrieveObjectResponse, error) {
	externalRes := &RetrieveObjectResponse{}
	res := struct {
		internal.WithErrors
		*RetrieveObjectResponse
	}{
		RetrieveObjectResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodGet, "/catalog/object/"+req.ObjectID, req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
