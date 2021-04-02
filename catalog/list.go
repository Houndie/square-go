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
	Value() *ListIteratorValue
	Error() error
	Next() bool
}

type listIterator struct {
	values []*ListIteratorValue
	iter   *internal.Iterator
}

func (i *listIterator) Value() *ListIteratorValue {
	return i.values[i.iter.Value()]
}

func (i *listIterator) Error() error {
	return i.iter.Error()
}

func (i *listIterator) Next() bool {
	return i.iter.Next()
}

type ListRequest struct {
	Types []objects.CatalogObjectEnumType
}

type ListIteratorValue struct {
	Object *objects.CatalogObject
}

type ListResponse struct {
	Objects ListIterator
}

func (c *client) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	iter := &listIterator{}
	iter.iter = internal.NewIterator(func(cursor string) (int, string, error) {
		stringTypes := make([]string, len(req.Types))
		for i, oneType := range req.Types {
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

		iter.values = make([]*ListIteratorValue, len(res.Objects))
		for i, object := range res.Objects {
			iter.values[i] = &ListIteratorValue{Object: object}
		}
		return len(res.Objects), res.Cursor, nil
	})

	return &ListResponse{
		Objects: iter,
	}, nil
}
