package internal

import (
	"fmt"
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
	return u, err
}

type middleware struct {
	apiKey string
	wrap   http.RoundTripper
}

func (m middleware) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Bearer "+m.apiKey)
	r.Header.Add("Accept", "application/json")
	if r.Method == "POST" {
		r.Header.Add("Content-Type", "application/json")
	}
	return m.wrap.RoundTrip(r)
}

func makeHTTPClient(apiKey string, httpClient *http.Client) *http.Client {
	if httpClient == nil {
		return &http.Client{
			Transport: &middleware{
				apiKey: apiKey,
				wrap:   http.DefaultTransport,
			},
		}
	}

	var transport http.RoundTripper
	if httpClient.Transport == nil {
		transport = &middleware{
			apiKey: apiKey,
			wrap:   http.DefaultTransport,
		}
	} else {
		transport = &middleware{
			apiKey: apiKey,
			wrap:   httpClient.Transport,
		}
	}
	return &http.Client{
		Transport:     transport,
		CheckRedirect: httpClient.CheckRedirect,
		Jar:           httpClient.Jar,
		Timeout:       httpClient.Timeout,
	}
}

type WithErrors struct {
	Errors []*objects.Error `json:"errors"`
}

func (w WithErrors) GetErrors() []*objects.Error {
	return w.Errors
}
