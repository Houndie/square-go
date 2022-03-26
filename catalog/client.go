package catalog

import (
	"context"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/options"
)

type Client interface {
	BatchDelete(ctx context.Context, req *BatchDeleteRequest) (*BatchDeleteResponse, error)
	BatchRetrieve(ctx context.Context, req *BatchRetrieveRequest) (*BatchRetrieveResponse, error)
	BatchUpsert(ctx context.Context, req *BatchUpsertRequest) (*BatchUpsertResponse, error)
	CreateImage(ctx context.Context, req *CreateImageRequest) (*CreateImageResponse, error)
	DeleteObject(ctx context.Context, req *DeleteObjectRequest) (*DeleteObjectResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	RetrieveObject(ctx context.Context, req *RetrieveObjectRequest) (*RetrieveObjectResponse, error)
	SearchItems(ctx context.Context, req *SearchItemsRequest) (*SearchItemsResponse, error)
	SearchObjects(ctx context.Context, req *SearchObjectsRequest) (*SearchObjectsResponse, error)
	UpdateItemModifierLists(ctx context.Context, req *UpdateItemModifierListsRequest) (*UpdateItemModifierListsResponse, error)
	UpdateItemTaxes(ctx context.Context, req *UpdateItemTaxesRequest) (*UpdateItemTaxesResponse, error)
	UpsertObject(ctx context.Context, req *UpsertObjectRequest) (*UpsertObjectResponse, error)
}

type client struct {
	i *internal.Client
}

func NewClient(apiKey string, environment objects.Environment, options ...options.ClientOption) (Client, error) {
	opts := make([]internal.ClientOption, len(options))
	for i, o := range options {
		opts[i] = internal.ClientOption(o)
	}

	c, err := internal.NewClient(apiKey, environment, opts...)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return &client{i: c}, nil
}
