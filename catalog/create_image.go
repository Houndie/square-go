package catalog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Houndie/square-go/internal"
	"github.com/Houndie/square-go/objects"
)

type CreateImageRequest struct {
	IdempotencyKey string                 `json:"idempotency_key,omitempty"`
	ObjectID       string                 `json:"object_id,omitempty"`
	Image          *objects.CatalogObject `json:"image,omitempty"`
	ImageBytes     io.Reader              `json:"-"`
	ImageFilename  string                 `json:"-"`
}

type CreateImageResponse struct {
	Image *objects.CatalogObject `json:"image"`
}

func (c *client) CreateImage(ctx context.Context, req *CreateImageRequest) (*CreateImageResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var (
		filepart io.Writer
		err      error
	)

	if req.ImageFilename != "" {
		filepart, err = writer.CreateFormFile("file", req.ImageFilename)
	} else {
		filepart, err = writer.CreateFormField("file")
	}

	if err != nil {
		return nil, fmt.Errorf("error creating file form field: %w", err)
	}

	if _, err := io.Copy(filepart, req.ImageBytes); err != nil {
		return nil, fmt.Errorf("error copying image to request body: %w", err)
	}

	reqBodyBytes, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	jsonpart, err := writer.CreateFormField("request")
	if err != nil {
		return nil, fmt.Errorf("error creating request form field: %w", err)
	}

	if _, err := jsonpart.Write(reqBodyBytes); err != nil {
		return nil, fmt.Errorf("error writing request body: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.i.Endpoint("/catalog/images").String(), body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpReq = httpReq.WithContext(ctx)

	externalRes := &CreateImageResponse{}
	res := struct {
		*CreateImageResponse
		internal.WithErrors
	}{
		CreateImageResponse: externalRes,
	}

	err = c.i.Requestor(c.i.HTTPClient, httpReq, res)
	if err != nil {
		return nil, fmt.Errorf("error with http request: %w", err)
	}

	return externalRes, nil
}
