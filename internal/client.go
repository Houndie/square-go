package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/Houndie/square-go/objects"
)

type Client struct {
	HttpClient   *http.Client
	EndpointBase *url.URL
}

func NewClient(apiKey string, environment objects.Environment, httpClient *http.Client) (*Client, error) {
	endpoint, err := endpointToURL(environment)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint: %w", err)
	}
	return &Client{
		HttpClient:   makeHTTPClient(apiKey, httpClient),
		EndpointBase: endpoint,
	}, nil
}

func (c *Client) Endpoint(e string) *url.URL {
	return &url.URL{
		Scheme: c.EndpointBase.Scheme,
		User:   c.EndpointBase.User,
		Host:   c.EndpointBase.Host,
		Path:   path.Join(c.EndpointBase.Path, e),
	}
}

func (c *Client) Do(ctx context.Context, method, path string, req interface{}, res interface{ GetErrors() []*objects.Error }) error {
	var endpoint string
	var bodyBuf io.Reader
	if method == http.MethodGet {
		u := c.Endpoint(path)
		q := u.Query()
		err := Encoder.Encode(req, q)
		if err != nil {
			return fmt.Errorf("error populating url paramters: %w", err)
		}
		u.RawQuery = q.Encode()
		endpoint = u.String()
		bodyBuf = nil
	} else {
		endpoint = c.Endpoint(path).String()
		reqBodyBytes, err := json.Marshal(&req)
		if err != nil {
			return fmt.Errorf("error marshaling request body: %w", err)
		}

		bodyBuf = bytes.NewBuffer(reqBodyBytes)
	}

	httpReq, err := http.NewRequest(method, endpoint, bodyBuf)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)

	resp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error with http request: %w", err)
	}
	defer resp.Body.Close()
	var codeErr error
	if resp.StatusCode != http.StatusOK {
		codeErr = objects.UnexpectedCodeError(resp.StatusCode)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if codeErr != nil {
			return codeErr
		}
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(bytes, res)
	if err != nil {
		if codeErr != nil {
			return codeErr
		}
		return fmt.Errorf("error unmarshalling json response: %w", err)
	}
	if errs := res.GetErrors(); len(errs) != 0 {
		return &objects.ErrorList{Errors: errs}
	}
	if codeErr != nil {
		return codeErr
	}
	return nil
}
