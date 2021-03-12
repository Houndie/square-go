package locations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

func (c *client) List(ctx context.Context) ([]*objects.Location, error) {
	res := struct {
		internal.WithErrors
		Locations []*objects.Location `json:"locations,omitempty"`
	}{}

	if err := c.i.Do(ctx, http.MethodGet, "locations", nil, &res); err != nil {
		return nil, fmt.Errorf("error performing http request: %w", err)
	}
	return res.Locations, nil
}
