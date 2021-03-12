package test

import "net/http"

type RoundTripper struct {
	RoundTripFunc func(r *http.Request) (*http.Response, error)
}

func (t *RoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return t.RoundTripFunc(r)
}
