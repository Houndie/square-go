package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type InfoRequest struct{}

type InfoResponse struct {
	Limits                       *objects.CatalogInfoResponseLimits    `json:"limits"`
	StandardUnitDescriptionGroup *objects.StandardUnitDescriptionGroup `json:"standard_unit_description_group"`
}

func (c *client) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, error) {
	externalRes := &InfoResponse{}
	res := struct {
		internal.WithErrors
		*InfoResponse
	}{
		InfoResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodGet, "/catalog/info", nil, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
