//nolint:goconst
package checkout

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
	"github.com/Houndie/square-go/options"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateCheckout(t *testing.T) {
	t.Parallel()

	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &objects.CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &objects.Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &objects.OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*objects.OrderLineItem{
				&objects.OrderLineItem{
					UID:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &objects.OrderQuantityUnit{
						MeasurementUnit: &objects.MeasurementUnit{
							Type: objects.MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &objects.Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      objects.CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectURL := "www.mywebsite.com"
	additionalRecipients := []*objects.ChargeRequestAdditionalRecipient{
		&objects.ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &objects.Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

	createdAt := time.Unix(1234567, 0)

	expectedCheckout := &objects.Checkout{
		ID:                         "some checkout id",
		CheckoutPageURL:            "www.sqaureup.com/gohere",
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectURL:                redirectURL,
		Order:                      order.Order,
		CreatedAt:                  &createdAt,
		AdditionalRecipients: []*objects.AdditionalRecipient{
			&objects.AdditionalRecipient{
				LocationID:   additionalRecipients[0].LocationID,
				Description:  additionalRecipients[0].Description,
				AmountMoney:  additionalRecipients[0].AmountMoney,
				ReceivableID: "some receivable id",
			},
		},
	}
	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				if r.Method != "POST" {
					t.Fatalf("found non post method %s", r.Method)
				}

				if r.Header.Get("Content-Type") != "application/json" {
					t.Fatalf("found wrong content type header %s", r.Header.Get("Content-Type"))
				}
				if r.Header.Get("Accept") != "application/json" {
					t.Fatalf("found wrong accept header %s", r.Header.Get("Accept"))
				}
				if r.Header.Get("Authorization") != "Bearer "+apiKey {
					t.Fatalf("found wrong authorization header %s", r.Header.Get("Authorization"))
				}

				reqJSON := struct {
					IdempotencyKey             string                                      `json:"idempotency_key,omitempty"`
					Order                      *objects.CreateOrderRequest                 `json:"order,omitempty"`
					AskForShippingAddress      bool                                        `json:"ask_for_shipping_address,omitempty"`
					MerchantSupportEmail       string                                      `json:"merchant_support_email,omitempty"`
					PrePopulateBuyerEmail      string                                      `json:"pre_populate_buyer_email,omitempty"`
					PrePopulateShippingAddress *objects.Address                            `json:"pre_populate_shipping_address,omitempty"`
					RedirectURL                string                                      `json:"redirect_url,omitempty"`
					AdditionalRecipients       []*objects.ChargeRequestAdditionalRecipient `json:"additional_recipients,omitempty"`
					Note                       string                                      `json:"note,omitempty"`
				}{
					IdempotencyKey:             idempotencyKey,
					Order:                      order,
					AskForShippingAddress:      askForShippingAddress,
					MerchantSupportEmail:       merchantSupportEmail,
					PrePopulateBuyerEmail:      prePopulateBuyerEmail,
					PrePopulateShippingAddress: prePopulateShippingAddress,
					RedirectURL:                redirectURL,
					AdditionalRecipients:       additionalRecipients,
					Note:                       note,
				}

				reqBody, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("error reading request body: %v", err)
				}

				err = json.Unmarshal(reqBody, &reqJSON)
				if err != nil {
					t.Fatalf("error unmarshaling request body: %v", err)
				}

				if reqJSON.IdempotencyKey != idempotencyKey {
					t.Fatalf("found idepotency key %s, expected %s", reqJSON.IdempotencyKey, idempotencyKey)
				}

				if !cmp.Equal(reqJSON.Order, order, cmpopts.IgnoreUnexported()) {
					t.Fatalf("found order %s not equal to existing order %s", spew.Sdump(reqJSON.Order), spew.Sdump(order))
				}

				if reqJSON.AskForShippingAddress != askForShippingAddress {
					t.Fatalf("found ask for shipping param %v, expected %v", reqJSON.AskForShippingAddress, askForShippingAddress)
				}

				if reqJSON.MerchantSupportEmail != merchantSupportEmail {
					t.Fatalf("found merchant support email %s, expected %s", reqJSON.MerchantSupportEmail, merchantSupportEmail)
				}

				if reqJSON.PrePopulateBuyerEmail != prePopulateBuyerEmail {
					t.Fatalf("found pre populate buyer email %s, expected %s", reqJSON.PrePopulateBuyerEmail, prePopulateBuyerEmail)
				}

				if !cmp.Equal(reqJSON.PrePopulateShippingAddress, prePopulateShippingAddress, cmpopts.IgnoreUnexported()) {
					t.Fatalf("found wrong pre populate shipping address %s, expected %s", spew.Sdump(reqJSON.PrePopulateShippingAddress), spew.Sdump(prePopulateShippingAddress))
				}

				if reqJSON.RedirectURL != redirectURL {
					t.Fatalf("found redirect url %s, expected %s", reqJSON.RedirectURL, redirectURL)
				}

				if !cmp.Equal(reqJSON.AdditionalRecipients, additionalRecipients, cmpopts.IgnoreUnexported()) {
					t.Fatalf("found additional recipients %s, expected %s", spew.Sdump(reqJSON.AdditionalRecipients), spew.Sdump(additionalRecipients))
				}

				if reqJSON.Note != note {
					t.Fatalf("found note %s, expected %s", reqJSON.Note, note)
				}

				resJSON := struct {
					Checkout *objects.Checkout `json:"checkout"`
				}{
					Checkout: expectedCheckout,
				}

				resBody, err := json.Marshal(&resJSON)
				if err != nil {
					t.Fatalf("error marshaling response body: %v", err)
				}

				header := http.Header{}
				header.Set("Content-Type", "application/json")

				return &http.Response{
					Status:        http.StatusText(http.StatusOK),
					StatusCode:    http.StatusOK,
					Proto:         "HTTP/1.0",
					ProtoMajor:    1,
					ProtoMinor:    0,
					Header:        header,
					Body:          ioutil.NopCloser(bytes.NewReader(resBody)),
					ContentLength: -1,
					Request:       r,
				}, nil
			},
		},
	}

	squareClient, err := NewClient(apiKey, objects.Production, options.WithHTTPClient(client))
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	res, err := squareClient.Create(context.Background(), &CreateRequest{
		LocationID:                 locationID,
		IdempotencyKey:             idempotencyKey,
		Order:                      order,
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectURL:                redirectURL,
		AdditionalRecipients:       additionalRecipients,
		Note:                       note,
	})
	if err != nil {
		t.Fatalf("found unxpected error from CreateCheckout: %v", err)
	}

	if !cmp.Equal(res.Checkout, expectedCheckout, cmpopts.IgnoreUnexported()) {
		t.Fatalf("found checkout %s, expected %s", spew.Sdump(res.Checkout), spew.Sdump(expectedCheckout))
	}
}

func TestCreateCheckoutClientError(t *testing.T) {
	t.Parallel()

	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &objects.CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &objects.Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &objects.OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*objects.OrderLineItem{
				&objects.OrderLineItem{
					UID:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &objects.OrderQuantityUnit{
						MeasurementUnit: &objects.MeasurementUnit{
							Type: objects.MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &objects.Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      objects.CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectURL := "www.mywebsite.com"
	additionalRecipients := []*objects.ChargeRequestAdditionalRecipient{
		&objects.ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &objects.Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

	client := &http.Client{
		Transport: &test.RoundTripper{
			RoundTripFunc: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("some error")
			},
		},
	}

	squareClient, err := NewClient(apiKey, objects.Production, options.WithHTTPClient(client))
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	_, err = squareClient.Create(context.Background(), &CreateRequest{
		LocationID:                 locationID,
		IdempotencyKey:             idempotencyKey,
		Order:                      order,
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectURL:                redirectURL,
		AdditionalRecipients:       additionalRecipients,
		Note:                       note,
	})
	if err == nil {
		t.Fatal("found no error when client returned one?")
	}
}

func TestCreateCheckoutErrorCode(t *testing.T) {
	t.Parallel()

	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &objects.CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &objects.Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &objects.OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*objects.OrderLineItem{
				&objects.OrderLineItem{
					UID:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &objects.OrderQuantityUnit{
						MeasurementUnit: &objects.MeasurementUnit{
							Type: objects.MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &objects.Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      objects.CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectURL := "www.mywebsite.com"
	additionalRecipients := []*objects.ChargeRequestAdditionalRecipient{
		&objects.ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &objects.Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

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

	squareClient, err := NewClient(apiKey, objects.Production, options.WithHTTPClient(client))
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	_, err = squareClient.Create(context.Background(), &CreateRequest{
		LocationID:                 locationID,
		IdempotencyKey:             idempotencyKey,
		Order:                      order,
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectURL:                redirectURL,
		AdditionalRecipients:       additionalRecipients,
		Note:                       note,
	})

	if err == nil {
		t.Fatal("found no error when client returned one?")
	}

	var uerr objects.UnexpectedCodeError
	if !errors.As(err, &uerr) {
		t.Fatalf("error was not of type unexpectedCodeError")
	}

	if int(uerr) != http.StatusInternalServerError {
		t.Fatalf("error code was not internal server error, found %v", int(uerr))
	}
}

func TestCreateCheckoutErrorMessage(t *testing.T) {
	t.Parallel()

	locationID := "some location id"
	idempotencyKey := "some unique key"
	order := &objects.CreateOrderRequest{
		IdempotencyKey: "some other unique key",
		Order: &objects.Order{
			ID:          "some id",
			LocationID:  "some other location id",
			ReferenceID: "some reference id",
			Source: &objects.OrderSource{
				Name: "some source",
			},
			CustomerID: "some customer",
			LineItems: []*objects.OrderLineItem{
				&objects.OrderLineItem{
					UID:      "some unique id",
					Name:     "im a line item",
					Quantity: "70",
					QuantityUnit: &objects.OrderQuantityUnit{
						MeasurementUnit: &objects.MeasurementUnit{
							Type: objects.MeasurementUnitLength("3"),
						},
						Precision: 7,
					},
				},
			},
		},
	}
	askForShippingAddress := true
	merchantSupportEmail := "someemail@whatever.com"
	prePopulateBuyerEmail := "someotheremail@aol.com"
	prePopulateShippingAddress := &objects.Address{
		AddressLine1:                 "1234 Any St.",
		Locality:                     "New York",
		AdministrativeDistrictLevel1: "New York",
		PostalCode:                   "12345",
		Country:                      objects.CountryTheUnitedStatesOfAmerica,
		FirstName:                    "John",
		LastName:                     "Doe",
		Organization:                 "Ninjas",
	}
	redirectURL := "www.mywebsite.com"
	additionalRecipients := []*objects.ChargeRequestAdditionalRecipient{
		&objects.ChargeRequestAdditionalRecipient{
			LocationID:  "more locations",
			Description: "blah blah",
			AmountMoney: &objects.Money{
				Amount:   1234,
				Currency: "Rupies",
			},
		},
	}
	note := "you're breathtaking!"

	apiKey := "some key"

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

	squareClient, err := NewClient(apiKey, objects.Production, options.WithHTTPClient(client))
	if err != nil {
		t.Fatalf("error creating square client: %v", err)
	}

	_, err = squareClient.Create(context.Background(), &CreateRequest{
		LocationID:                 locationID,
		IdempotencyKey:             idempotencyKey,
		Order:                      order,
		AskForShippingAddress:      askForShippingAddress,
		MerchantSupportEmail:       merchantSupportEmail,
		PrePopulateBuyerEmail:      prePopulateBuyerEmail,
		PrePopulateShippingAddress: prePopulateShippingAddress,
		RedirectURL:                redirectURL,
		AdditionalRecipients:       additionalRecipients,
		Note:                       note,
	})

	if err == nil {
		t.Fatal("found no error when client returned one?")
	}

	serr := &objects.ErrorList{}
	if !errors.As(err, &serr) {
		t.Fatalf("error was not of type square.ErrorList")
	}

	if len(serr.Errors) != 1 {
		t.Fatalf("found %v errors, expected %v", len(serr.Errors), 1)
	}

	if !cmp.Equal(serr.Errors[0], testError, cmpopts.IgnoreUnexported()) {
		t.Fatalf("errors were not equal, found %s, expected %s", spew.Sdump(serr.Errors[0]), spew.Sdump(testError))
	}
}
