package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type StockLevel string

const (
	StockLevelOut StockLevel = "OUT"
	StockLevelIn  StockLevel = "LOW"
)

type SearchItemsRequest struct {
	TextFilter             string                           `json:"text_filter,omitempty"`
	CategoryIDs            []string                         `json:"category_ids,omitempty"`
	StockLevels            []StockLevel                     `json:"stock_levels,omitempty"`
	EnabledLocationIDs     []string                         `json:"enabled_location_ids,omitempty"`
	Limit                  int                              `json:"limit,omitempty"`
	SortOrder              int                              `json:"sort_order,omitempty"`
	ProductTypes           []string                         `json:"product_types,omitempty"`
	CustomAttributeFilters []*objects.CustomAttributeFilter `json:"custom_attribute_filters,omitempty"`
}

type SearchItemsIteratorValue struct {
	Item                *objects.CatalogObject
	MatchedVariationIDs []string
}

type SearchItemsIterator interface {
	Value() *SearchItemsIteratorValue
	Error() error
	Next() bool
}

type searchItemsIterator struct {
	values []*SearchItemsIteratorValue
	iter   *internal.Iterator
}

func (i *searchItemsIterator) Value() *SearchItemsIteratorValue {
	return i.values[i.iter.Value()]
}

func (i *searchItemsIterator) Error() error {
	return i.iter.Error()
}

func (i *searchItemsIterator) Next() bool {
	return i.iter.Next()
}

type SearchItemsResponse struct {
	Items SearchItemsIterator
}

func (c *client) SearchItems(ctx context.Context, req *SearchItemsRequest) (*SearchItemsResponse, error) {
	iter := &searchItemsIterator{}
	iter.iter = internal.NewIterator(func(cursor string) (int, string, error) {
		req := struct {
			*SearchItemsRequest
			Cursor string `schema:"cursor,omitempty"`
		}{
			SearchItemsRequest: req,
			Cursor:             cursor,
		}
		res := struct {
			internal.WithErrors
			Cursor              string                   `json:"cursor,omitempty"`
			Items               []*objects.CatalogObject `json:"items,omitempty"`
			MatchedVariationIDs []string                 `json:"match_variation_ids,omitempty"`
		}{}

		if err := c.i.Do(ctx, http.MethodGet, "catalog/search-catalog-items", req, &res); err != nil {
			return 0, "", fmt.Errorf("error performing http request: %w", err)
		}

		variationMap := map[string]struct{}{}
		for _, variationID := range res.MatchedVariationIDs {
			variationMap[variationID] = struct{}{}
		}

		iter.values = make([]*SearchItemsIteratorValue, len(res.Items))
		for i, item := range res.Items {
			matchedVariationIDs := []string{}
			if catalogItem, ok := item.Type.(*objects.CatalogItem); ok {
				for _, variation := range catalogItem.Variations {
					if _, found := variationMap[variation.ID]; found {
						matchedVariationIDs = append(matchedVariationIDs, variation.ID)
					}
				}
			}
			iter.values[i] = &SearchItemsIteratorValue{
				Item:                item,
				MatchedVariationIDs: matchedVariationIDs,
			}
		}
		return len(res.Items), res.Cursor, nil
	})

	return &SearchItemsResponse{
		Items: iter,
	}, nil
}
