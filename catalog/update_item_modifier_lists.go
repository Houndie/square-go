package catalog

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
)

type UpdateItemModifierListsRequest struct {
	ItemIDs                []string `json:"item_ids,omitempty"`
	ModifierListsToEnable  []string `json:"modifier_lists_to_enable,omitempty"`
	ModifierListsToDisable []string `json:"modifier_lists_to_disable,omitempty"`
}

type UpdateItemModifierListsResponse struct {
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c *client) UpdateItemModifierLists(ctx context.Context, req *UpdateItemModifierListsRequest) (*UpdateItemModifierListsResponse, error) {
	externalRes := &UpdateItemModifierListsResponse{}
	res := struct {
		internal.WithErrors
		*UpdateItemModifierListsResponse
	}{
		UpdateItemModifierListsResponse: externalRes,
	}

	if err := c.i.Do(ctx, http.MethodPost, "/catalog/update-item-modifier-lists", req, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}

	return externalRes, nil
}
