package square

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type Environment int

const (
	productionEndpoint = "https://connect.squareup.com/v2"
	sandboxEndpoint    = "https://connect.squareupsandbox.com/v2"

	Production Environment = iota
	Sandbox
)

type Client struct {
	httpClient   *http.Client
	endpointBase *url.URL
}

func NewClient(apiKey string, environment Environment, httpClient *http.Client) (*Client, error) {
	var endpoint string
	switch environment {
	case Production:
		endpoint = productionEndpoint
	case Sandbox:
		endpoint = sandboxEndpoint
	default:
		return nil, fmt.Errorf("unknown environment")
	}
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint url: %w", err)
	}
	if httpClient == nil {
		return &Client{
			endpointBase: u,
			httpClient: &http.Client{
				Transport: &middleware{
					apiKey: apiKey,
					wrap:   http.DefaultTransport,
				},
			},
		}, nil
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
	return &Client{
		endpointBase: u,
		httpClient: &http.Client{
			Transport:     transport,
			CheckRedirect: httpClient.CheckRedirect,
			Jar:           httpClient.Jar,
			Timeout:       httpClient.Timeout,
		},
	}, nil
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

func (c *Client) endpoint(e string) *url.URL {
	u := &url.URL{
		Scheme: c.endpointBase.Scheme,
		User:   c.endpointBase.User,
		Host:   c.endpointBase.Host,
		Path:   path.Join(c.endpointBase.Path, e),
	}
	return u
}
