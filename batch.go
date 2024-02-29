package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/marcozac/hubspot-go/util"
)

type BatchClient[
	ReadInput BatchReadInputEmbedder,
	CreateInput BatchCreateInputEmbedder,
	UpdateInput BatchUpdateInputEmbedder,
	ArchiveInput BatchArchiveInputEmbedder,
] struct {
	endpoint string
	hc       *http.Client
}

func (c *BatchClient[ReadInput, CreateInput, UpdateInput, ArchiveInput]) Archive(ctx context.Context, input *ArchiveInput) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, buf)
	if err != nil {
		return err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

type BatchInput[I any] struct {
	Inputs []I `json:"inputs"`
}

func (BatchInput[I]) embedBatchReadInput()    {}
func (BatchInput[I]) embedBatchCreateInput()  {}
func (BatchInput[I]) embedBatchUpdateInput()  {}
func (BatchInput[I]) embedBatchArchiveInput() {}

type BatchReadInputEmbedder interface {
	embedBatchReadInput()
}

type BatchCreateInputEmbedder interface {
	embedBatchCreateInput()
}

type BatchUpdateInputEmbedder interface {
	embedBatchUpdateInput()
}

type BatchArchiveInputEmbedder interface {
	embedBatchArchiveInput()
}
