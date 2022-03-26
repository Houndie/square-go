package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Houndie/square-go/objects"
	"github.com/gorilla/schema"
)

var (
	Decoder = schema.NewDecoder()
	Encoder = schema.NewEncoder()
)

const (
	productionEndpoint = "https://connect.squareup.com/v2"
	sandboxEndpoint    = "https://connect.squareupsandbox.com/v2"
)

func endpointToURL(environment objects.Environment) (*url.URL, error) {
	var endpoint string

	switch environment {
	case objects.Production:
		endpoint = productionEndpoint
	case objects.Sandbox:
		endpoint = sandboxEndpoint
	default:
		return nil, fmt.Errorf("unknown environment")
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint url: %w", err)
	}

	return u, nil
}

type middleware struct {
	apiKey string
	wrap   http.RoundTripper
}

func (m middleware) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer "+m.apiKey)
	r.Header.Set("Accept", "application/json")

	if r.Method == "POST" {
		r.Header.Set("Content-Type", "application/json")
	}

	return m.wrap.RoundTrip(r)
}

type WithErrors struct {
	Errors []*objects.Error `json:"errors"`
}

func (w WithErrors) GetErrors() []*objects.Error {
	return w.Errors
}

func ParseResponse(resp *http.Response, res interface{ GetErrors() []*objects.Error }) error {
	var codeErr error
	if resp.StatusCode != http.StatusOK {
		codeErr = objects.UnexpectedCodeError(resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if codeErr != nil {
			return codeErr //nolint:wrapcheck
		}

		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(bytes, res)
	if err != nil {
		if codeErr != nil {
			return codeErr //nolint:wrapcheck
		}

		return fmt.Errorf("error unmarshalling json response: %w", err)
	}

	if errs := res.GetErrors(); len(errs) != 0 {
		return &objects.ErrorList{Errors: errs}
	}

	if codeErr != nil {
		return codeErr //nolint:wrapcheck
	}

	return nil
}
