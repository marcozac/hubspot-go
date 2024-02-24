package hubspot

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/marcozac/hubspot-go/endpoint"
)

func NewClient(ts oauth2.TokenSource, opts ...ClientOption) (*Client, error) {
	if ts == nil {
		return nil, ErrTokenSourceRequired
	}
	cfg := &clientConfig{
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(cfg)
	}
	hc := oauth2.NewClient(cfg.ctx, ts)
	return &Client{
		hc:         hc,
		Properties: NewPropertiesClient(hc),
	}, nil
}

type Client struct {
	hc         *http.Client
	Properties *PropertiesClient
}

func NewPropertiesClient(hc *http.Client) *PropertiesClient {
	return &PropertiesClient{
		Contact: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.ContactProperties,
		},
		Company: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.CompanyProperties,
		},
		Deal: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.DealProperties,
		},
		FeedbackSubmission: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.FeedbackSubmissionProperties,
		},
		LineItem: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.LineItemProperties,
		},
		Product: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.ProductProperties,
		},
		Quote: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.QuoteProperties,
		},
		Discount: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.DiscountProperties,
		},
		Fee: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.FeeProperties,
		},
		Tax: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.TaxProperties,
		},
		Ticket: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.TicketProperties,
		},
		Goal: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.GoalProperties,
		},
	}
}

type PropertiesClient struct {
	Contact            *PropertiesObjectClient
	Company            *PropertiesObjectClient
	Deal               *PropertiesObjectClient
	FeedbackSubmission *PropertiesObjectClient
	LineItem           *PropertiesObjectClient
	Product            *PropertiesObjectClient
	Quote              *PropertiesObjectClient
	Discount           *PropertiesObjectClient
	Fee                *PropertiesObjectClient
	Tax                *PropertiesObjectClient
	Ticket             *PropertiesObjectClient
	Goal               *PropertiesObjectClient
}

type PropertiesObjectClient struct {
	endpoint string
	hc       *http.Client
}

// List returns a list of properties for the object type.
func (poc *PropertiesObjectClient) List(ctx context.Context, opts ...RequestOption) ([]Property, error) {
	cfg := applyRequestOptions(nil, opts...)
	for _, opt := range opts {
		opt(cfg)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, poc.endpoint, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	applyArchivedOption(cfg, q)
	applyPropertiesOption(cfg, q)
	resp, err := poc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var results Results[Property]
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return results.Results, nil
}

// Results is a generic struct that contains a list of results of type T
// returned by the HubSpot API.
type Results[T any] struct {
	Results []T `json:"results"`
}
