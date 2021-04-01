package catalog

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
)

type DeleteObjectRequest struct {
	ObjectID string
}

type DeleteObjectResponse struct {
	DeletedObjectIDs []string   `json:"deleted_object_ids,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

func (c *client) DeleteObject(ctx context.Context, req *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	externalRes := &DeleteObjectResponse{}
	res := struct {
		internal.WithErrors
		*DeleteObjectResponse
	}{
		DeleteObjectResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodDelete, "/catalog/object/"+req.ObjectID, nil, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
