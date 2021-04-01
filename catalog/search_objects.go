package catalog

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type SearchObjectsRequest struct {
	ObjectTypes           []objects.CatalogObjectType `json:"object_types,omitempty"`
	IncludeDeletedObjects bool                        `json:"include_deleted_objects,omitempty"`
	BeginTime             *time.Time                  `json:"begin_time,omitempty"`
	Query                 *objects.CatalogQuery       `json:"query,omitempty"`
	Limit                 int                         `json:"limit,omitempty"`
}

type SearchObjectsIteratorValue struct {
	Object         *objects.CatalogObject
	RelatedObjects []*objects.CatalogObject
}

type SearchObjectsIterator interface {
	Value() *SearchObjectsIteratorValue
	Error() error
	Next() bool
}

type searchObjectsIterator struct {
	values []*SearchObjectsIteratorValue
	iter   *internal.Iterator
}

func (i *searchObjectsIterator) Value() *SearchObjectsIteratorValue {
	return i.values[i.iter.Value()]
}

func (i *searchObjectsIterator) Error() error {
	return i.iter.Error()
}

func (i *searchObjectsIterator) Next() bool {
	return i.iter.Next()
}

type SearchObjectsResponse struct {
	Objects    SearchObjectsIterator
	LatestTime *time.Time
}

func (c *client) SearchObjects(ctx context.Context, req *SearchObjectsRequest) (*SearchObjectsResponse, error) {
	var latestTime *time.Time

	iter := &searchObjectsIterator{}
	refreshFunc := func(cursor string) (int, string, error) {
		req := struct {
			*SearchObjectsRequest
			Cursor string `schema:"cursor,omitempty"`
		}{
			SearchObjectsRequest: req,
			Cursor:               cursor,
		}
		res := struct {
			internal.WithErrors
			Cursor         string                   `json:"cursor,omitempty"`
			Objects        []*objects.CatalogObject `json:"objects,omitempty"`
			RelatedObjects []*objects.CatalogObject `json:"related_objects,omitempty"`
			LatestTime     *time.Time               `json:"latest_time,omitempty"`
		}{}

		if err := c.i.Do(ctx, http.MethodGet, "catalog/search", req, &res); err != nil {
			return 0, "", fmt.Errorf("error performing http request: %w", err)
		}

		relatedObjectsMap := map[string]*objects.CatalogObject{}
		for _, object := range res.RelatedObjects {
			relatedObjectsMap[object.ID] = object
		}

		iter.values = make([]*SearchObjectsIteratorValue, len(res.Objects))

		for i, object := range res.Objects {
			relatedObjects := []*objects.CatalogObject{}

			catalogItem, ok := object.Type.(*objects.CatalogItem)
			if !ok {
				continue
			}

			for _, variation := range catalogItem.Variations {
				if _, found := relatedObjectsMap[variation.ID]; found {
					relatedObjects = append(relatedObjects, relatedObjectsMap[variation.ID])
				}
			}

			for _, id := range catalogItem.TaxIDs {
				if _, found := relatedObjectsMap[id]; found {
					relatedObjects = append(relatedObjects, relatedObjectsMap[id])
				}
			}

			if catalogItem.CategoryID != "" {
				if _, found := relatedObjectsMap[catalogItem.CategoryID]; found {
					relatedObjects = append(relatedObjects, relatedObjectsMap[catalogItem.CategoryID])
				}
			}

			iter.values[i] = &SearchObjectsIteratorValue{
				Object:         object,
				RelatedObjects: relatedObjects,
			}
		}

		latestTime = res.LatestTime

		return len(res.Objects), res.Cursor, nil
	}

	length, cursor, err := refreshFunc("")
	if err != nil {
		return nil, err
	}

	iter.iter = internal.NewIteratorWithValues(refreshFunc, length, cursor)

	// Copy latest time so we don't keep touching the one we pass back out of the function
	timeCopy := *latestTime

	return &SearchObjectsResponse{
		Objects:    iter,
		LatestTime: &timeCopy,
	}, nil
}
