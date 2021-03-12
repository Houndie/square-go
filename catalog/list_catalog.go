package catalog

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type ListIterator interface {
	Value() *objects.CatalogObject
	Error() error
	Next() bool
}

type listIterator struct {
	values []*objects.CatalogObject
	iter   *internal.Iterator
}

func (i *listIterator) Value() *objects.CatalogObject {
	return i.values[i.iter.Value()]
}

func (i *listIterator) Error() error {
	return i.iter.Error()
}

func (i *listIterator) Next() bool {
	return i.iter.Next()
}

func (c *client) List(ctx context.Context, types []objects.CatalogObjectType) ListIterator {
	iter := &listIterator{}
	iter.iter = internal.NewIterator(func(cursor string) (int, string, error) {
		stringTypes := make([]string, len(types))
		for i, oneType := range types {
			stringTypes[i] = string(oneType)
		}
		req := struct {
			Types  string `schema:"types,omitempty"`
			Cursor string `schema:"cursor,omitempty"`
		}{
			Types:  strings.Join(stringTypes, ","),
			Cursor: cursor,
		}
		res := struct {
			internal.WithErrors
			Cursor  string                   `json:"cursor,omitempty"`
			Objects []*objects.CatalogObject `json:"objects,omitempty"`
		}{}

		if err := c.i.Do(ctx, http.MethodGet, "catalog/list", req, &res); err != nil {
			return 0, "", fmt.Errorf("error performing http request: %w", err)
		}

		iter.values = res.Objects
		return len(res.Objects), res.Cursor, nil
	})
	return iter
}
