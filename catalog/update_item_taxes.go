package catalog

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
)

type UpdateItemTaxesRequest struct {
	ItemIDs        []string `json:"item_ids,omitempty"`
	TaxesToEnable  []string `json:"taxes_to_enable,omitempty"`
	TaxesToDisable []string `json:"taxes_to_disable,omitempty"`
}

type UpdateItemTaxesResponse struct {
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c *client) UpdateItemTaxes(ctx context.Context, req *UpdateItemTaxesRequest) (*UpdateItemTaxesResponse, error) {
	externalRes := &UpdateItemTaxesResponse{}
	res := struct {
		internal.WithErrors
		*UpdateItemTaxesResponse
	}{
		UpdateItemTaxesResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "/catalog/update-item-taxes", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
