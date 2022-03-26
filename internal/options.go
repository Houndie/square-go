package internal

import (
	"errors"
	"net/http"
	"time"
)

type ClientOption func(*Client) error

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		var transport http.RoundTripper
		if httpClient.Transport == nil {
			transport = c.HTTPClient.Transport
		} else {
			m, ok := c.HTTPClient.Transport.(*middleware)
			if !ok {
				return errors.New("http client transport was not our own middlware?")
			}
			transport = &middleware{
				apiKey: m.apiKey,
				wrap:   httpClient.Transport,
			}
		}

		c.HTTPClient = &http.Client{
			Transport:     transport,
			CheckRedirect: httpClient.CheckRedirect,
			Jar:           httpClient.Jar,
			Timeout:       httpClient.Timeout,
		}

		return nil
	}
}

func WithRateLimit(maxTime time.Duration) ClientOption {
	return func(c *Client) error {
		c.Requestor = BackoffRequestor(maxTime, c.Requestor)
		return nil
	}
}
