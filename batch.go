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
	Out any,
] struct {
	endpoint string
	hc       *http.Client
}

// Create creates a new batch of objects in HubSpot as specified by the input.
func (c *BatchClient[ReadInput, CreateInput, UpdateInput, ArchiveInput, Out]) Create(ctx context.Context, input *CreateInput) (*BatchOutput[Out], error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	results := new(BatchOutput[Out])
	if err := json.NewDecoder(resp.Body).Decode(results); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return results, nil
}

// Archive archives a batch of objects in HubSpot as specified by the input.
func (c *BatchClient[ReadInput, CreateInput, UpdateInput, ArchiveInput, Out]) Archive(ctx context.Context, input *ArchiveInput) error {
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

type BatchOutput[R any] struct {
	Results[R] // embed the results

	RequestedAt *DateTime         `json:"requestedAt,omitempty"`
	StartedAt   *DateTime         `json:"startedAt,omitempty"`
	CompletedAt *DateTime         `json:"completedAt,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
	Status      string            `json:"status,omitempty"`
}

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
