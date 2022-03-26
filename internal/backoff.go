package internal

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/Houndie/square-go/objects"
)

type Requestor func(client *http.Client, req *http.Request, res interface{ GetErrors() []*objects.Error }) error

func DefaultRequestor() Requestor {
	return func(client *http.Client, req *http.Request, res interface{ GetErrors() []*objects.Error }) error {
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error with http request: %w", err)
		}
		defer resp.Body.Close()

		if err := ParseResponse(resp, res); err != nil {
			return fmt.Errorf("error parsing response body: %w", err)
		}

		return nil
	}
}

var ErrMaxBackoff = errors.New("rate limited, hit max backoff amount")

func BackoffRequestor(maxTime time.Duration, r Requestor) Requestor {
	return func(client *http.Client, req *http.Request, res interface{ GetErrors() []*objects.Error }) error {
		done := time.Now().Add(maxTime)

		waitExp := 0
		nextWait := 0 * time.Second

		for time.Now().Add(nextWait).Before(done) {
			if nextWait > 0 {
				time.Sleep(nextWait)
			}

			err := r(client, req, res)
			if err != nil {
				var errs *objects.ErrorList
				if !errors.As(err, &errs) {
					return err
				}

				for _, e := range errs.Errors {
					if e.Category != objects.ErrorCategoryRateLimitError || e.Code != objects.ErrorCodeRateLimited {
						return err
					}
				}
			}

			nextWait = time.Duration(math.Pow(2, float64(waitExp))) * time.Second //nolint:gomnd
			waitExp++                                                             //nolint:wastedassign
		}

		return ErrMaxBackoff
	}
}
