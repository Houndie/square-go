package orders

import (
	"context"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/options"
)

type Client interface {
	BatchRetrieve(ctx context.Context, req *BatchRetrieveRequest) (*BatchRetrieveResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
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
