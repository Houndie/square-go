package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Houndie/square-go/objects"
)

type mockRoundTripper struct {
	RoundTripFunc func(r *http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(r)
}

func TestBackoff(t *testing.T) {
	t.Parallel()

	requestor := BackoffRequestor(10*time.Second, DefaultRequestor())

	count := 0
	client := &http.Client{
		Transport: &mockRoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				count++
				errs := WithErrors{
					Errors: []*objects.Error{
						{
							Category: objects.ErrorCategoryRateLimitError,
							Code:     objects.ErrorCodeRateLimited,
						},
					},
				}
				w := httptest.NewRecorder()
				if err := json.NewEncoder(w).Encode(errs); err != nil {
					t.Fatal(err)
				}

				return w.Result(), nil
			},
		},
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := WithErrors{}

	err = requestor(client, req, &res)
	if err == nil {
		t.Fatal("expected eventual failure")
	}

	if !errors.Is(err, ErrMaxBackoff) {
		t.Fatal(err)
	}

	if count != 4 {
		t.Fatal(count)
	}
}
