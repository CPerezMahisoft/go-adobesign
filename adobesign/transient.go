package adobesign

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
)

const transientDocumentsPath = "transientDocuments"

// TransientDocumentService handles operations related to agreement documents.
//
// ref: https://secure.na1.echosign.com/public/docs/restapi/v6#!/transientDocuments
type TransientDocumentService service

type TransientDocument struct {
	TransientDocumentId string `json:"transientDocumentId"`
}

func (s *TransientDocumentService) UploadTransientDocument(ctx context.Context, file []byte, filename string) (*TransientDocument, error) {

	// Create the multi-part form request
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	part1, err := writer.CreateFormFile("File", filename)
	b, err := io.Copy(part1, bytes.NewReader(file))
	if err != nil {
		fmt.Println(b)
		return nil, err
	}
	_ = writer.WriteField("File-Name", filename)
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := s.client.NewMultiPartRequest(transientDocumentsPath, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	var response *TransientDocument
	if _, err := s.client.Do(ctx, req, &response); err != nil {
		return nil, err
	}

	return response, nil

}
