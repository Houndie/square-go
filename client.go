package square

import (
	"fmt"
	"net/http"

	"github.com/Houndie/square-go/catalog"
	"github.com/Houndie/square-go/checkout"
	"github.com/Houndie/square-go/inventory"
	"github.com/Houndie/square-go/locations"
	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/orders"
)

type Client struct {
	Catalog   catalog.Client
	Checkout  checkout.Client
	Inventory inventory.Client
	Locations locations.Client
	Orders    orders.Client
}

func NewClient(apiKey string, environment objects.Environment, httpClient *http.Client) (*Client, error) {
	catalog, err := catalog.NewClient(apiKey, environment, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error constructing catalog client: %w", err)
	}

	checkout, err := checkout.NewClient(apiKey, environment, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error constructing catalog client: %w", err)
	}

	inventory, err := inventory.NewClient(apiKey, environment, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error constructing catalog client: %w", err)
	}

	locations, err := locations.NewClient(apiKey, environment, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error constructing catalog client: %w", err)
	}

	orders, err := orders.NewClient(apiKey, environment, httpClient)
	if err != nil {
		return nil, fmt.Errorf("error constructing catalog client: %w", err)
	}

	return &Client{
		Catalog:   catalog,
		Checkout:  checkout,
		Inventory: inventory,
		Locations: locations,
		Orders:    orders,
	}, nil
}
