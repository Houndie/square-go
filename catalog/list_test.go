//nolint:goconst
package catalog

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"errors"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/internal/test"
	"github.com/Houndie/square-go/objects"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestListCatalog(t *testing.T) {
	t.Parallel()

	apiKey := "some api key"
	types := []objects.CatalogObjectType{
		objects.CatalogObjectTypeTax,
		objects.CatalogObjectTypeModifier,
	}

	updatedAt := time.Unix(1235634, 0)
	updatedAt2 := time.Unix(2363094, 0)
	expectedObjects := []*objects.CatalogObject{
		&objects.CatalogObject{
			ID:                    "some id",
			UpdatedAt:             &updatedAt,
			Version:               7,
			IsDeleted:             true,
			CatalogV1IDs:          nil,
			PresentAtAllLocations: true,
			PresentAtLocationIDs:  nil,
			ImageID:               "some image id",
			Type: &objects.CatalogTax{
				Name:                   "tax",
				CalculationPhase:       "phase",
				InclusionType:          "inclusion",
				Percentage:             "6",
				AppliesToCustomAmounts: true,
				Enabled:                true,
			},
		},
		&objects.CatalogObject{
			ID:                    "some other id",
			UpdatedAt:             &updatedAt2,
			Version:               2,
			IsDeleted:             false,
			CatalogV1IDs:          nil,
			PresentAtAllLocations: false,
			PresentAtLocationIDs:  []string{"location 1", "location 2"},
			ImageID:               "some other image id",
			Type: &objects.CatalogModifier{
				Name: "modifier",
				PriceMoney: &objects.Money{
					Amount:   3,
					Currency: "pesos",
				},
			},
		},
	}

	cursors := []string{"", "cursor", ""}

	callCount := 0
	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Authorization") != "Bearer "+apiKey {
					t.Fatalf("Found incorrect authorization header %s", r.Header.Get("Authorization"))
				}

				if r.Header.Get("Accept") != "application/json" {
					t.Fatalf("Found incorrect accept header %s", r.Header.Get("Accept"))
				}

				urlParams := struct {
					Types  string `schema:"types"`
					Cursor string `schema:"cursor"`
				}{}

				err := internal.Decoder.Decode(&urlParams, r.URL.Query())
				if err != nil {
					t.Fatalf("unexpected error when decoding url params")
				}

				for _, controlType := range types {
					found := false
					for _, testType := range strings.Split(urlParams.Types, ",") {
						if controlType == objects.CatalogObjectType(testType) {
							found = true
							break
						}
					}
					if !found {
						t.Fatalf("could not find expected type %s in params", controlType)
					}
				}

				if cursors[callCount] != urlParams.Cursor {
					t.Fatalf("found unexpected cursor %s, expected %s", cursors[callCount], urlParams.Cursor)
				}

				body, err := json.Marshal(&struct {
					Cursor  string                   `json:"cursor,omitempty"`
					Objects []*objects.CatalogObject `json:"objects,omitempty"`
				}{
					Cursor:  cursors[callCount+1],
					Objects: []*objects.CatalogObject{expectedObjects[callCount]},
				})

				if err != nil {
					t.Fatalf("found error while marshaling json response: %v", err)
				}

				callCount++

				return &http.Response{
					Status:        http.StatusText(http.StatusOK),
					StatusCode:    http.StatusOK,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					ContentLength: 0,
					Body:          ioutil.NopCloser(bytes.NewReader(body)),
					Request:       r,
				}, nil
			},
		},
	}

	catalogClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	catalogObjects, err := catalogClient.List(context.Background(), &ListRequest{Types: types})
	if err != nil {
		t.Fatalf("found unexpected error: %v", err)
	}

	idx := 0
	for catalogObjects.Objects.Next() {
		if !cmp.Equal(catalogObjects.Objects.Value().Object, expectedObjects[idx], cmpopts.IgnoreUnexported()) {
			t.Fatalf("found unexpected catalog item %s, expected %s", spew.Sdump(catalogObjects.Objects.Value()), spew.Sdump(expectedObjects[idx]))
		}
		idx++
	}

	if catalogObjects.Objects.Error() != nil {
		t.Fatalf("found unexpected error: %v", catalogObjects.Objects.Error())
	}

	if callCount != 2 {
		t.Fatalf("found %d http calls, expected 2", callCount)
	}

	if idx != 2 {
		t.Fatalf("found %d response items, expected 2", idx)
	}
}

func TestListCatalogClientError(t *testing.T) {
	t.Parallel()

	apiKey := "some api key"
	types := []objects.CatalogObjectType{
		objects.CatalogObjectTypeTax,
		objects.CatalogObjectTypeModifier,
	}

	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("some error")
			},
		},
	}

	catalogClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	catalogObjects, err := catalogClient.List(context.Background(), &ListRequest{Types: types})
	if err != nil {
		return
	}

	idx := 0
	for catalogObjects.Objects.Next() {
		idx++
	}

	if catalogObjects.Objects.Error() == nil {
		t.Fatal("Expected error, found none")
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}

func TestListCatalogHttpError(t *testing.T) {
	t.Parallel()

	apiKey := "some api key"
	types := []objects.CatalogObjectType{
		objects.CatalogObjectTypeTax,
		objects.CatalogObjectTypeModifier,
	}

	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					Status:        http.StatusText(http.StatusInternalServerError),
					StatusCode:    http.StatusInternalServerError,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					ContentLength: 0,
					Body:          ioutil.NopCloser(bytes.NewReader([]byte{})),
					Request:       r,
				}, nil
			},
		},
	}

	catalogClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	catalogObjects, err := catalogClient.List(context.Background(), &ListRequest{Types: types})

	validateError := func(err error) {
		var uerr objects.UnexpectedCodeError
		if !errors.As(err, &uerr) {
			t.Fatalf("error was not of type unexpectedCodeError")
		}

		if int(uerr) != http.StatusInternalServerError {
			t.Fatalf("error code was not internal server error, found %v", int(uerr))
		}
	}

	if err != nil {
		validateError(err)
		return
	}

	idx := 0
	for catalogObjects.Objects.Next() {
		idx++
	}

	if catalogObjects.Objects.Error() == nil {
		t.Fatal("Expected error, found none")
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}

func TestListCatalogErrorMessage(t *testing.T) {
	t.Parallel()

	apiKey := "some api key"
	types := []objects.CatalogObjectType{
		objects.CatalogObjectTypeTax,
		objects.CatalogObjectTypeModifier,
	}

	testError := &objects.Error{
		Category: objects.ErrorCategoryAPIError,
		Code:     objects.ErrorCodeInternalServerError,
		Detail:   "some detail",
		Field:    "some field",
	}

	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				resp := struct {
					Errors []*objects.Error
				}{
					Errors: []*objects.Error{testError},
				}

				respJSON, err := json.Marshal(&resp)
				if err != nil {
					t.Fatalf("error marshaling response body: %v", err)
				}
				return &http.Response{
					Status:        http.StatusText(http.StatusInternalServerError),
					StatusCode:    http.StatusInternalServerError,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					ContentLength: 0,
					Body:          ioutil.NopCloser(bytes.NewReader(respJSON)),
					Request:       r,
				}, nil
			},
		},
	}

	catalogClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	catalogObjects, err := catalogClient.List(context.Background(), &ListRequest{Types: types})

	validateError := func(err error) {
		serr := &objects.ErrorList{}
		if !errors.As(err, &serr) {
			t.Fatalf("error not of type square.ErrorList")
		}

		if !cmp.Equal(serr.Errors[0], testError, cmpopts.IgnoreUnexported()) {
			t.Fatalf("expected error %s, found %s", spew.Sdump(serr.Errors[0]), spew.Sdump(testError))
		}
	}

	if err != nil {
		validateError(err)
		return
	}

	idx := 0
	for catalogObjects.Objects.Next() {
		idx++
	}

	if catalogObjects.Objects.Error() == nil {
		t.Fatal("Expected error, found none")
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}
