package catalog

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
)

type BatchDeleteRequest struct {
	ObjectIDs []string `json:"object_ids,omitempty"`
}

type BatchDeleteResponse struct {
	DeletedObjectIDs []string   `json:"deleted_object_ids,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

func (c *client) BatchDelete(ctx context.Context, req *BatchDeleteRequest) (*BatchDeleteResponse, error) {
	externalRes := &BatchDeleteResponse{}
	res := struct {
		internal.WithErrors
		*BatchDeleteResponse
	}{
		BatchDeleteResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "/catalog/batch-delete", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
