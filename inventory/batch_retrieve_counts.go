package inventory

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type BatchRetrieveCountsRequest struct {
	CatalogObjectIDs []string   `json:"catalog_object_ids,omitempty"`
	LocationIDs      []string   `json:"location_ids,omitempty"`
	UpdatedAfter     *time.Time `json:"updated_after,omitempty"`
}

type BatchRetrieveCountsIteratorValue struct {
	Count *objects.InventoryCount
}

type BatchRetrieveCountsIterator interface {
	Value() *BatchRetrieveCountsIteratorValue
	Error() error
	Next() bool
}

type batchRetrieveCountsIterator struct {
	values []*BatchRetrieveCountsIteratorValue
	iter   *internal.Iterator
}

func (i *batchRetrieveCountsIterator) Value() *BatchRetrieveCountsIteratorValue {
	return i.values[i.iter.Value()]
}

func (i *batchRetrieveCountsIterator) Error() error {
	return i.iter.Error()
}

func (i *batchRetrieveCountsIterator) Next() bool {
	return i.iter.Next()
}

type BatchRetrieveCountsResponse struct {
	Counts BatchRetrieveCountsIterator
}

func (c *client) BatchRetrieveCounts(ctx context.Context, req *BatchRetrieveCountsRequest) (*BatchRetrieveCountsResponse, error) {
	iter := &batchRetrieveCountsIterator{}
	iter.iter = internal.NewIterator(func(cursor string) (int, string, error) {
		req := struct {
			Cursor string `json:"cursor,omitempty"`
			*BatchRetrieveCountsRequest
		}{
			BatchRetrieveCountsRequest: req,
			Cursor:                     cursor,
		}
		res := struct {
			internal.WithErrors
			Cursor string                    `json:"cursor"`
			Counts []*objects.InventoryCount `json:"counts"`
		}{}
		if err := c.i.Do(ctx, http.MethodPost, "inventory/batch-retrieve-counts", req, &res); err != nil {
			return 0, "", fmt.Errorf("error performing http request: %w", err)
		}
		iter.values = make([]*BatchRetrieveCountsIteratorValue, len(res.Counts))
		for i, count := range res.Counts {
			iter.values[i] = &BatchRetrieveCountsIteratorValue{Count: count}
		}
		return len(res.Counts), res.Cursor, nil
	})

	return &BatchRetrieveCountsResponse{
		Counts: iter,
	}, nil
}
