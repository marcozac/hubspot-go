package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/marcozac/hubspot-go/endpoint"
	"github.com/marcozac/hubspot-go/util"
)

type BatchClient[
	ReadInput BatchReadInputEmbedder,
	CreateInput BatchCreateInputEmbedder,
	UpdateInput BatchUpdateInputEmbedder,
	ArchiveInput BatchArchiveInputEmbedder,
	Out any,
] struct {
	baseEndpoint string
	hc           *http.Client
}

// Read retrieves a batch of objects from HubSpot as specified by the input.
func (c *BatchClient[RI, CI, UI, AI, O]) Read(ctx context.Context, input *RI) (*BatchOutput[O], error) {
	endpoint := c.baseEndpoint + endpoint.Read
	return c.do(ctx, endpoint, input)
}

// Create creates a new batch of objects in HubSpot as specified by the input.
func (c *BatchClient[RI, CI, UI, AI, O]) Create(ctx context.Context, input *CI) (*BatchOutput[O], error) {
	endpoint := c.baseEndpoint + endpoint.Create
	return c.do(ctx, endpoint, input)
}

// Update updates a batch of objects in HubSpot as specified by the input.
func (c *BatchClient[RI, CI, UI, AI, O]) Update(ctx context.Context, input *UI) (*BatchOutput[O], error) {
	endpoint := c.baseEndpoint + endpoint.Update
	return c.do(ctx, endpoint, input)
}

// Archive archives a batch of objects in HubSpot as specified by the input.
func (c *BatchClient[ReadInput, CreateInput, UpdateInput, ArchiveInput, Out]) Archive(ctx context.Context, input *ArchiveInput) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseEndpoint+endpoint.Archive, buf)
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

func (c *BatchClient[RI, CI, UI, AI, O]) do(ctx context.Context, endpoint string, input any) (*BatchOutput[O], error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(input); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, buf)
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
	results := new(BatchOutput[O])
	if err := json.NewDecoder(resp.Body).Decode(results); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return results, nil
}

type BatchInput[I any] struct {
	Inputs []I `json:"inputs"`
}

func (BatchInput[I]) embedBatchReadInput()    {}
func (BatchInput[I]) embedBatchCreateInput()  {}
func (BatchInput[I]) embedBatchUpdateInput()  {}
func (BatchInput[I]) embedBatchArchiveInput() {}

type BatchOutput[R any] struct {
	results[R] // embed the results

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
