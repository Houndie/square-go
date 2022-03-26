package options

import (
	"net/http"
	"time"

	"github.com/Houndie/square-go/internal"
)

type ClientOption internal.ClientOption

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return ClientOption(internal.WithHTTPClient(httpClient))
}

func WithRateLimit(maxTime time.Duration) ClientOption {
	return ClientOption(internal.WithRateLimit(maxTime))
}
