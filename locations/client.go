package locations

import (
	"context"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type Client interface {
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
}

type client struct {
	i *internal.Client
}

func NewClient(apiKey string, environment objects.Environment, httpClient *http.Client) (Client, error) {
	c, err := internal.NewClient(apiKey, environment, httpClient)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	return &client{i: c}, nil
}
