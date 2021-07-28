package locations

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/Houndie/square-go/objects"
)

func testToken(t *testing.T) string {
	t.Helper()

	token := os.Getenv("TEST_TOKEN")
	if token == "" {
		t.Skip()
	}

	return token
}

func TestList(t *testing.T) {
	t.Parallel()

	client, err := NewClient(testToken(t), objects.Sandbox, &http.Client{})
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Locations) != 1 {
		t.Fatalf("unexpected number of locations: %v", len(res.Locations))
	}

	if res.Locations[0].ID != "LP78X79R1PA0P" {
		t.Fatalf("incorrect location id: %s", res.Locations[0].ID)
	}
}
