package objects

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestRetrieveTransaction(t *testing.T) {
	locationID := "some location id"
	transactionID := "some transaction id"
	apiKey := "apiKey"

	createdAt := time.Unix(1234567, 0)

	expectedTransaction := &Transaction{
		ID:          transactionID,
		LocationID:  locationID,
		CreatedAt:   &createdAt,
		Tenders:     nil,
		Refunds:     nil,
		ReferenceID: "some reference id",
		Product:     TransactionProductRegister,
		ClientID:    "some client id",
		ShippingAddress: &Address{
			AddressLine1:                 "123 Any St.",
			Locality:                     "New York",
			AdministrativeDistrictLevel1: "New York",
			PostalCode:                   "12345",
			Country:                      CountryTheUnitedStatesOfAmerica,
			FirstName:                    "John",
			LastName:                     "Doe",
			Organization:                 "Crazy People",
		},
		OrderID: "some order id",
	}
	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("Accept") != "application/json" {
					t.Fatalf("incorrect accept header found: %s", r.Header.Get("Accept"))
				}
				if r.Header.Get("Authorization") != "Bearer "+apiKey {
					t.Fatalf("incorrect authorization header found: %s", r.Header.Get("Authorization"))
				}

				resp := struct {
					Transaction *Transaction `json:transaction"`
				}{
					Transaction: expectedTransaction,
				}

				jsonBody, err := json.Marshal(&resp)
				if err != nil {
					t.Fatalf("error marshaling json body: %v", err)
				}

				return &http.Response{
					Status:        http.StatusText(http.StatusOK),
					StatusCode:    http.StatusOK,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					ContentLength: 0,
					Body:          ioutil.NopCloser(bytes.NewReader(jsonBody)),
					Request:       r,
				}, nil

			},
		},
	}
	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}
	transaction, err := squareClient.RetrieveTransaction(context.Background(), locationID, transactionID)
	if err != nil {
		t.Fatalf("unexpected error returned when retrieving transaction: %v", err)
	}

	if !cmp.Equal(transaction, expectedTransaction, cmpopts.IgnoreUnexported()) {
		t.Fatalf("found transaction %s, expected transaction %s", spew.Sdump(transaction), spew.Sdump(expectedTransaction))
	}
}

func TestRetrieveTransactionHttpError(t *testing.T) {
	locationID := "some location id"
	transactionID := "some transaction id"
	apiKey := "apiKey"

	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("some error")
			},
		},
	}
	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}
	_, err = squareClient.RetrieveTransaction(context.Background(), locationID, transactionID)
	if err == nil {
		t.Fatalf("expected error returned, found none")
	}
}

func TestRetrieveTransactionErrorCode(t *testing.T) {
	locationID := "some location id"
	transactionID := "some transaction id"
	apiKey := "apiKey"

	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
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
	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}
	_, err = squareClient.RetrieveTransaction(context.Background(), locationID, transactionID)
	if err == nil {
		t.Fatalf("expected error returned, found none")
	}

	uerr, ok := err.(unexpectedCodeError)
	if !ok {
		t.Fatalf("error was not of type unexpectedCodeError")
	}
	if int(uerr) != http.StatusInternalServerError {
		t.Fatalf("error code was not internal server error, found %v", int(uerr))
	}
}

func TestRetrieveTransactionErrorMessage(t *testing.T) {
	locationID := "some location id"
	transactionID := "some transaction id"
	apiKey := "apiKey"

	testError := &Error{
		Category: ErrorCategoryApiError,
		Code:     ErrorCodeInternalServerError,
		Detail:   "some detail",
		Field:    "some field",
	}

	client := &http.Client{
		Transport: &testRoundTripper{
			roundTripFunc: func(r *http.Request) (*http.Response, error) {
				resp := struct {
					Errors []*Error
				}{
					Errors: []*Error{testError},
				}

				jsonBody, err := json.Marshal(&resp)
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
					Body:          ioutil.NopCloser(bytes.NewReader(jsonBody)),
					Request:       r,
				}, nil
			},
		},
	}
	squareClient, err := NewClient(apiKey, Production, client)
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}
	_, err = squareClient.RetrieveTransaction(context.Background(), locationID, transactionID)
	if err == nil {
		t.Fatalf("expected error returned, found none")
	}

	serr, ok := err.(*ErrorList)
	if !ok {
		t.Fatalf("error was not of type error list")
	}

	if len(serr.Errors) != 1 {
		t.Fatalf("found incorrect number of errors %d, expected 1", len(serr.Errors))
	}

	if !cmp.Equal(serr.Errors[0], testError, cmpopts.IgnoreUnexported()) {
		t.Fatalf("found error %s was different from expected %s", spew.Sdump(serr.Errors[0]), spew.Sdump(testError))
	}
}
