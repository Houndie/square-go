//nolint:goconst
package inventory

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"errors"

	"github.com/Houndie/square-go/internal/test"
	"github.com/Houndie/square-go/objects"
)

func TestBatchRetrieveInventoryCounts(t *testing.T) {
	t.Parallel()

	catalogObjectIDs := []string{"id1", "id2"}
	locationIDs := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

	cursors := []string{"", "CURSOR", ""}
	callCount := 0

	time1 := time.Unix(1234567, 0)
	time2 := time.Unix(3446678, 0)
	expectedCounts := []*objects.InventoryCount{
		&objects.InventoryCount{
			CatalogObjectID:   catalogObjectIDs[0],
			CatalogObjectType: objects.CatalogObjectEnumTypeItemVariation,
			State:             "OH",
			LocationID:        locationIDs[0],
			Quantity:          "7",
			CalculatedAt:      &time1,
		},
		&objects.InventoryCount{
			CatalogObjectID:   catalogObjectIDs[1],
			CatalogObjectType: objects.CatalogObjectEnumTypeItemVariation,
			State:             "PA",
			LocationID:        locationIDs[1],
			Quantity:          "3.4",
			CalculatedAt:      &time2,
		},
	}
	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Authorization") != "Bearer "+apiKey {
					t.Fatalf("Found incorrect authorization header %s, expected %s", r.Header.Get("Authorization"), "Bearer "+apiKey)
				}
				if r.Header.Get("Accept") != "application/json" {
					t.Fatalf("found incorrect accept header %s, expected application/json", r.Header.Get("Accept"))
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Fatalf("found incorrect content-type %s, expected application/json", r.Header.Get("Content-Type"))
				}

				if r.URL.String() != "https://connect.squareup.com/v2/inventory/batch-retrieve-counts" {
					t.Fatalf("Found incorrect request url %s", r.URL.String())
				}

				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("Error reading request body: %v", err)
				}

				jsonRequest := struct {
					CatalogObjectIDs []string   `json:"catalog_object_ids"`
					LocationIDs      []string   `json:"location_ids"`
					UpdatedAfter     *time.Time `json:"updated_after"`
					Cursor           string     `json:"cursor"`
				}{}
				err = json.Unmarshal(body, &jsonRequest)
				if err != nil {
					t.Fatalf("Error unmarshalling request body: %v", err)
				}

				if len(catalogObjectIDs) != len(jsonRequest.CatalogObjectIDs) {
					t.Fatalf("wrong number of catalog object ids (found %v, expected %v)", len(jsonRequest.CatalogObjectIDs), len(catalogObjectIDs))
				}

				for _, controlID := range catalogObjectIDs {
					found := false
					for _, testID := range jsonRequest.CatalogObjectIDs {
						if testID == controlID {
							found = true
							break
						}
					}
					if !found {
						t.Fatalf("Could not find control catalog id %s", controlID)
					}
				}

				if len(locationIDs) != len(jsonRequest.LocationIDs) {
					t.Fatalf("wrong number of location ids (found %v, expected %v)", len(jsonRequest.LocationIDs), len(locationIDs))
				}

				for _, controlID := range locationIDs {
					found := false
					for _, testID := range jsonRequest.LocationIDs {
						if testID == controlID {
							found = true
							break
						}
					}
					if !found {
						t.Fatalf("Could not find control location id %s", controlID)
					}
				}

				if !updatedAfter.Equal(*jsonRequest.UpdatedAfter) {
					t.Fatalf("Wrong updated after (found %v, expected %v)", jsonRequest.UpdatedAfter, updatedAfter)
				}

				if jsonRequest.Cursor != cursors[callCount] {
					t.Fatalf("incorrect cursor found %s, expected %s", jsonRequest.Cursor, cursors[callCount])
				}

				resp := struct {
					Cursor string
					Counts []*objects.InventoryCount
				}{
					Cursor: cursors[callCount+1],
					Counts: []*objects.InventoryCount{expectedCounts[callCount]},
				}

				jsonResp, err := json.Marshal(&resp)
				if err != nil {
					t.Fatalf("unxpected error marshalling response: %v", err)
				}

				header := http.Header{}
				header.Set("Content-Type", "application/json")

				callCount++
				return &http.Response{
					Status:        http.StatusText(http.StatusOK),
					StatusCode:    http.StatusOK,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					Header:        header,
					Body:          ioutil.NopCloser(bytes.NewReader(jsonResp)),
					ContentLength: -1,
					Request:       r,
				}, nil
			},
		},
	}

	squareClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	res, err := squareClient.BatchRetrieveCounts(context.Background(), &BatchRetrieveCountsRequest{
		CatalogObjectIDs: catalogObjectIDs,
		LocationIDs:      locationIDs,
		UpdatedAfter:     &updatedAfter,
	})
	if err != nil {
		t.Fatalf("unexpected error from batch retrieve counts: %v", err)
	}

	idx, inventoryCounts := 0, res.Counts
	for inventoryCounts.Next() {
		if inventoryCounts.Value().Count.CatalogObjectID != expectedCounts[idx].CatalogObjectID {
			t.Fatalf("found catalog object id %s, expected %s", inventoryCounts.Value().Count.CatalogObjectID, expectedCounts[idx].CatalogObjectID)
		}

		if inventoryCounts.Value().Count.CatalogObjectType != expectedCounts[idx].CatalogObjectType {
			t.Fatalf("found catalog object type %s, expected %s", inventoryCounts.Value().Count.CatalogObjectType, expectedCounts[idx].CatalogObjectType)
		}

		if inventoryCounts.Value().Count.State != expectedCounts[idx].State {
			t.Fatalf("found state %s, expected %s", inventoryCounts.Value().Count.State, expectedCounts[idx].State)
		}

		if inventoryCounts.Value().Count.LocationID != expectedCounts[idx].LocationID {
			t.Fatalf("found location id %s, expected %s", inventoryCounts.Value().Count.LocationID, expectedCounts[idx].LocationID)
		}

		if inventoryCounts.Value().Count.Quantity != expectedCounts[idx].Quantity {
			t.Fatalf("found quantity %s, expected %s", inventoryCounts.Value().Count.Quantity, expectedCounts[idx].Quantity)
		}

		if !inventoryCounts.Value().Count.CalculatedAt.Equal(*expectedCounts[idx].CalculatedAt) {
			t.Fatalf("found calculated time %s, expected %s", inventoryCounts.Value().Count.CalculatedAt, expectedCounts[idx].CalculatedAt)
		}

		idx++
	}

	if inventoryCounts.Error() != nil {
		t.Fatalf("found unexpected error: %v", inventoryCounts.Error())
	}

	if idx != len(expectedCounts) {
		t.Fatalf("found unxepected number of items %v, expected %v", idx, len(expectedCounts))
	}
}

func TestBatchRetrieveInventoryCountsClientError(t *testing.T) {
	t.Parallel()

	catalogObjectIDs := []string{"id1", "id2"}
	locationIDs := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("some error")
			},
		},
	}

	squareClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	inventoryCounts, err := squareClient.BatchRetrieveCounts(context.Background(), &BatchRetrieveCountsRequest{
		CatalogObjectIDs: catalogObjectIDs,
		LocationIDs:      locationIDs,
		UpdatedAfter:     &updatedAfter,
	})
	if err != nil {
		return
	}

	idx := 0
	for inventoryCounts.Counts.Next() {
		idx++
	}

	if inventoryCounts.Counts.Error() == nil {
		t.Fatalf("expected error, found none")
	}

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}

func TestBatchRetrieveInventoryCountsErrorCode(t *testing.T) {
	t.Parallel()

	catalogObjectIDs := []string{"id1", "id2"}
	locationIDs := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

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

	squareClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	inventoryCounts, err := squareClient.BatchRetrieveCounts(context.Background(), &BatchRetrieveCountsRequest{
		CatalogObjectIDs: catalogObjectIDs,
		LocationIDs:      locationIDs,
		UpdatedAfter:     &updatedAfter,
	})

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
	for inventoryCounts.Counts.Next() {
		idx++
	}

	if inventoryCounts.Counts.Error() == nil {
		t.Fatalf("expected error, found none")
	}

	validateError(inventoryCounts.Counts.Error())

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}

func TestBatchRetrieveInventoryCountsErrorMessage(t *testing.T) {
	t.Parallel()

	catalogObjectIDs := []string{"id1", "id2"}
	locationIDs := []string{"id3", "id4", "id5"}
	updatedAfter := time.Unix(1287529, 0)
	apiKey := "apiKey"

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

	squareClient, err := NewClient(apiKey, objects.Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	inventoryCounts, err := squareClient.BatchRetrieveCounts(context.Background(), &BatchRetrieveCountsRequest{
		CatalogObjectIDs: catalogObjectIDs,
		LocationIDs:      locationIDs,
		UpdatedAfter:     &updatedAfter,
	})

	validateError := func(err error) {
		serr := &objects.ErrorList{}
		if !errors.As(err, &serr) {
			t.Fatalf("error not of type square.Error")
		}

		if len(serr.Errors) != 1 {
			t.Fatalf("found %v errors, expected %v", len(serr.Errors), 1)
		}

		if serr.Errors[0].Category != testError.Category {
			t.Fatalf("found error category %s, expected %s", serr.Errors[0].Category, testError.Category)
		}

		if serr.Errors[0].Code != testError.Code {
			t.Fatalf("found error code %s, expected %s", serr.Errors[0].Code, testError.Code)
		}

		if serr.Errors[0].Detail != testError.Detail {
			t.Fatalf("found error detail %s, expected %s", serr.Errors[0].Detail, testError.Detail)
		}

		if serr.Errors[0].Field != testError.Field {
			t.Fatalf("found error field %s, expected %s", serr.Errors[0].Field, testError.Field)
		}
	}

	if err != nil {
		validateError(err)
		return
	}

	idx := 0
	for inventoryCounts.Counts.Next() {
		idx++
	}

	if inventoryCounts.Counts.Error() == nil {
		t.Fatalf("expected error, found none")
	}

	validateError(inventoryCounts.Counts.Error())

	if idx != 0 {
		t.Fatalf("found %v items when 0 was expected", idx)
	}
}
