package inventory

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type BatchRetrieveCountsIterator interface {
	Value() *objects.InventoryCount
	Error() error
	Next() bool
}

type batchRetrieveCountsIterator struct {
	catalogObjectIDs []string
	locationIDs      []string
	updatedAfter     *time.Time
	cursor           string

	counts []*objects.InventoryCount
	idx    int
	done   bool
	err    error
	c      *client
	ctx    context.Context
}

func (i *batchRetrieveCountsIterator) setError(err error) bool {
	i.counts = nil
	i.err = err
	return false
}

func (i *batchRetrieveCountsIterator) Value() *objects.InventoryCount {
	return i.counts[i.idx]
}

func (i *batchRetrieveCountsIterator) Error() error {
	return i.err
}

func (i *batchRetrieveCountsIterator) Next() bool {
	i.idx = i.idx + 1
	if i.idx < len(i.counts) {
		return true
	}

	if i.done {
		return false
	}

	req := struct {
		CatalogObjectIDs []string   `json:"catalog_object_ids,omitempty"`
		LocationIDs      []string   `json:"location_ids,omitempty"`
		UpdatedAfter     *time.Time `json:"updated_after,omitempty"`
		Cursor           string     `json:"cursor,omitempty"`
	}{
		CatalogObjectIDs: i.catalogObjectIDs,
		LocationIDs:      i.locationIDs,
		UpdatedAfter:     i.updatedAfter,
		Cursor:           i.cursor,
	}

	res := struct {
		internal.WithErrors
		Cursor string                    `json:"cursor"`
		Counts []*objects.InventoryCount `json:"counts"`
	}{}

	if err := i.c.i.Do(i.ctx, http.MethodPost, "inventory/batch-retrieve-counts", req, &res); err != nil {
		return i.setError(fmt.Errorf("error performing http request: %w", err))
	}

	if len(res.Counts) == 0 {
		return false
	}
	i.counts = res.Counts
	i.idx = 0
	if res.Cursor == "" {
		i.done = true
	} else {
		i.cursor = res.Cursor
	}
	return true
}

func (c *client) BatchRetrieveCounts(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) BatchRetrieveCountsIterator {
	return &batchRetrieveCountsIterator{
		catalogObjectIDs: catalogObjectIDs,
		locationIDs:      locationIDs,
		updatedAfter:     updatedAfter,
		cursor:           "",
		counts:           nil,
		idx:              -1,
		done:             false,
		err:              nil,
		c:                c,
		ctx:              ctx,
	}
}
