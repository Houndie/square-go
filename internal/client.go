package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/Houndie/square-go/objects"
)

type Client struct {
	HTTPClient   *http.Client
	endpointBase *url.URL
	Requestor    Requestor
}

func NewClient(apiKey string, environment objects.Environment, options ...ClientOption) (*Client, error) {
	endpoint, err := endpointToURL(environment)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint: %w", err)
	}

	c := &Client{
		HTTPClient: &http.Client{
			Transport: &middleware{
				apiKey: apiKey,
				wrap:   http.DefaultTransport,
			},
		},
		endpointBase: endpoint,
		Requestor:    DefaultRequestor(),
	}

	for _, o := range options {
		err := o(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) Endpoint(e string) *url.URL {
	return &url.URL{
		Scheme: c.endpointBase.Scheme,
		User:   c.endpointBase.User,
		Host:   c.endpointBase.Host,
		Path:   path.Join(c.endpointBase.Path, e),
	}
}

func (c *Client) Do(ctx context.Context, method, path string, req interface{}, res interface{ GetErrors() []*objects.Error }) error {
	var (
		endpoint string
		bodyBuf  io.Reader
	)

	if method == http.MethodGet {
		u := c.Endpoint(path)

		if req != nil {
			q := u.Query()

			if err := Encoder.Encode(req, q); err != nil {
				return fmt.Errorf("error populating url parameters: %w", err)
			}

			u.RawQuery = q.Encode()
		}

		endpoint = u.String()
		bodyBuf = nil
	} else {
		endpoint = c.Endpoint(path).String()
		reqBodyBytes, err := json.Marshal(&req)
		if err != nil {
			return fmt.Errorf("error marshaling request body: %w", err)
		}

		if string(reqBodyBytes) != "null" {
			bodyBuf = bytes.NewBuffer(reqBodyBytes)
		}
	}

	httpReq, err := http.NewRequest(method, endpoint, bodyBuf)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)

	return c.Requestor(c.HTTPClient, httpReq, res)
}
