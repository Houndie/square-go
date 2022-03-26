//nolint: goconst
package catalog

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/options"
	"github.com/gofrs/uuid"
)

func testToken(t *testing.T) string {
	t.Helper()

	token := os.Getenv("TEST_TOKEN")
	if token == "" {
		t.Skip()
	}

	return token
}

func TestUpsert(t *testing.T) {
	t.Parallel()
	token := testToken(t)

	itemName := "name"

	client, err := NewClient(token, objects.Sandbox, options.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}))
	if err != nil {
		t.Fatalf("error creating new client: %v", err)
	}

	idempotencyKey := uuid.Must(uuid.NewV4())

	upsertRes, err := client.UpsertObject(context.Background(), &UpsertObjectRequest{
		IdempotencyKey: idempotencyKey.String(),
		Object: &objects.CatalogObject{
			ID: "#id",
			Type: &objects.CatalogItem{
				Name: itemName,
			},
		},
	})
	if err != nil {
		t.Fatalf("error creating new remote object: %v", err)
	}

	t.Cleanup(func() {
		_, err = client.DeleteObject(context.Background(), &DeleteObjectRequest{
			ObjectID: upsertRes.CatalogObject.ID,
		})

		if err != nil {
			t.Fatalf("error deleting remote object: %v", err)
		}
	})

	retrieveRes, err := client.RetrieveObject(context.Background(), &RetrieveObjectRequest{
		ObjectID: upsertRes.CatalogObject.ID,
	})
	if err != nil {
		t.Fatalf("error retrieving object: %v", err)
	}

	if retrieveRes.Object.ID != upsertRes.CatalogObject.ID {
		t.Fatalf("expected object ID %s, found %s", upsertRes.CatalogObject.ID, retrieveRes.Object.ID)
	}

	item, ok := retrieveRes.Object.Type.(*objects.CatalogItem)
	if !ok {
		t.Fatalf("catalog item expected but not found")
	}

	if item.Name != itemName {
		t.Fatalf("expected item name %s, found %s", itemName, item.Name)
	}
}

func TestUpsertVariation(t *testing.T) {
	t.Parallel()
	token := testToken(t)

	itemName := "name"

	client, err := NewClient(token, objects.Sandbox, options.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}))
	if err != nil {
		t.Fatalf("error creating new client: %v", err)
	}

	upsertIdempotencyKey := uuid.Must(uuid.NewV4())

	upsertRes, err := client.UpsertObject(context.Background(), &UpsertObjectRequest{
		IdempotencyKey: upsertIdempotencyKey.String(),
		Object: &objects.CatalogObject{
			ID: "#id",
			Type: &objects.CatalogItem{
				Name: itemName,
			},
		},
	})
	if err != nil {
		t.Fatalf("error creating new remote object: %v", err)
	}

	t.Cleanup(func() {
		_, err = client.DeleteObject(context.Background(), &DeleteObjectRequest{
			ObjectID: upsertRes.CatalogObject.ID,
		})

		if err != nil {
			t.Fatalf("error deleting remote object: %v", err)
		}
	})

	variationIdempotencyKey := uuid.Must(uuid.NewV4())

	variationUpsertRes, err := client.UpsertObject(context.Background(), &UpsertObjectRequest{
		IdempotencyKey: variationIdempotencyKey.String(),
		Object: &objects.CatalogObject{
			ID: "#id",
			Type: &objects.CatalogItemVariation{
				Name:        itemName,
				PricingType: objects.CatalogPricingTypeFixed,
				PriceMoney: &objects.Money{
					Amount:   5,
					Currency: "USD",
				},
				ItemID: upsertRes.CatalogObject.ID,
			},
		},
	})
	if err != nil {
		t.Fatalf("error creating new remote object: %v", err)
	}

	t.Cleanup(func() {
		_, err = client.DeleteObject(context.Background(), &DeleteObjectRequest{
			ObjectID: variationUpsertRes.CatalogObject.ID,
		})

		if err != nil {
			t.Fatalf("error deleting remote object: %v", err)
		}
	})

	retrieveRes, err := client.RetrieveObject(context.Background(), &RetrieveObjectRequest{
		ObjectID: upsertRes.CatalogObject.ID,
	})
	if err != nil {
		t.Fatalf("error retrieving object: %v", err)
	}

	catalogItem, ok := retrieveRes.Object.Type.(*objects.CatalogItem)
	if !ok {
		t.Fatalf("catalog object is not catalog item")
	}

	if len(catalogItem.Variations) != 2 {
		t.Fatalf("variation not added to item, %d variations found", len(catalogItem.Variations))
	}
}

func TestUpsertDiscount(t *testing.T) {
	t.Parallel()
	token := testToken(t)

	itemName := "name"

	client, err := NewClient(token, objects.Sandbox, options.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}))
	if err != nil {
		t.Fatalf("error creating new client: %v", err)
	}

	upsertIdempotencyKey := uuid.Must(uuid.NewV4())

	upsertRes, err := client.UpsertObject(context.Background(), &UpsertObjectRequest{
		IdempotencyKey: upsertIdempotencyKey.String(),
		Object: &objects.CatalogObject{
			ID: "#id",
			Type: &objects.CatalogDiscount{
				Name: itemName,
				DiscountType: &objects.CatalogDiscountFixedAmount{
					AmountMoney: &objects.Money{
						Amount:   5,
						Currency: "USD",
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("error creating new remote object: %v", err)
	}

	t.Cleanup(func() {
		_, err = client.DeleteObject(context.Background(), &DeleteObjectRequest{
			ObjectID: upsertRes.CatalogObject.ID,
		})

		if err != nil {
			t.Fatalf("error deleting remote object: %v", err)
		}
	})

	retrieveRes, err := client.RetrieveObject(context.Background(), &RetrieveObjectRequest{
		ObjectID: upsertRes.CatalogObject.ID,
	})
	if err != nil {
		t.Fatalf("error retrieving object: %v", err)
	}

	catalogDiscount, ok := retrieveRes.Object.Type.(*objects.CatalogDiscount)
	if !ok {
		t.Fatalf("catalog object is not catalog item")
	}

	if catalogDiscount.Name != itemName {
		t.Fatalf(catalogDiscount.Name)
	}
}
