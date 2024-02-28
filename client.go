package hubspot

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/marcozac/hubspot-go/endpoint"
)

// NewHTTPClient returns a new HTTP client that uses the given token source to
// authenticate requests to the HubSpot API.
func NewHTTPClient(ts oauth2.TokenSource, opts ...ClientOption) *http.Client {
	cfg := applyClientOptions(nil, opts...)
	// TODO
	// Add the limit and retry middlewares to the HTTP client.
	return oauth2.NewClient(cfg.ctx, ts)
}

// NewDefaultClient returns a new HubSpot client with the given token source and
// options. It is a wrapper around NewTypedClient that uses the default
// properties for each object type.
func NewDefaultClient(ts oauth2.TokenSource, opts ...ClientOption) (*Client[
	ContactDefaultProperties,
	CompanyDefaultProperties,
	DealDefaultProperties,
	FeedbackSubmissionDefaultProperties,
	LineItemDefaultProperties,
	ProductDefaultProperties,
	QuoteDefaultProperties,
	DiscountDefaultProperties,
	FeeDefaultProperties,
	TaxDefaultProperties,
	TicketDefaultProperties,
	GoalDefaultProperties,
], error,
) {
	return NewTypedClient[
		ContactDefaultProperties,
		CompanyDefaultProperties,
		DealDefaultProperties,
		FeedbackSubmissionDefaultProperties,
		LineItemDefaultProperties,
		ProductDefaultProperties,
		QuoteDefaultProperties,
		DiscountDefaultProperties,
		FeeDefaultProperties,
		TaxDefaultProperties,
		TicketDefaultProperties,
		GoalDefaultProperties,
	](ts, opts...)
}

// NewTypedClient returns a new HubSpot client with the given token source and
// options. It requires a type param for each object type to be used with the
// client. The type param is used to embed the default properties in the custom
// properties struct.
//
// This function may be useful when many objects have custom properties.
// Otherwise, you may create your own Client struct embedding the default
// client using [NewObjectClient] to create a client for the objects with
// custom properties.
//
// You can also create a variable with the desired type params and use it to
// call this function. This way, you can avoid repeating the type params in
// every call.
//
// Example:
//
//	type ContactPropertiesTest struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomPropertyFromUI is a custom property added from the HubSpot UI
//		// for testing purposes.
//		MyCustomPropertyFromUI string `json:"my_custom_prop_from_ui,omitempty"`
//	}
//
//	var NewClient = NewTypedClient[
//		ContactPropertiesTest,
//		CompanyDefaultProperties,
//		DealDefaultProperties,
//		FeedbackSubmissionDefaultProperties,
//		LineItemDefaultProperties,
//		ProductDefaultProperties,
//		QuoteDefaultProperties,
//		DiscountDefaultProperties,
//		FeeDefaultProperties,
//		TaxDefaultProperties,
//		TicketDefaultProperties,
//		GoalDefaultProperties,
//	]
//
//	func main() {
//		client, err := NewClient(ts, WithContext(ctx))
//		// ...
//	}
func NewTypedClient[
	ContactProperties ContactPropertiesEmbedder,
	CompanyProperties CompanyPropertiesEmbedder,
	DealProperties DealPropertiesEmbedder,
	FeedbackSubmissionProperties FeedbackSubmissionPropertiesEmbedder,
	LineItemProperties LineItemPropertiesEmbedder,
	ProductProperties ProductPropertiesEmbedder,
	QuoteProperties QuotePropertiesEmbedder,
	DiscountProperties DiscountPropertiesEmbedder,
	FeeProperties FeePropertiesEmbedder,
	TaxProperties TaxPropertiesEmbedder,
	TicketProperties TicketPropertiesEmbedder,
	GoalProperties GoalPropertiesEmbedder,
](ts oauth2.TokenSource, opts ...ClientOption) (
	*Client[
		ContactProperties,
		CompanyProperties,
		DealProperties,
		FeedbackSubmissionProperties,
		LineItemProperties,
		ProductProperties,
		QuoteProperties,
		DiscountProperties,
		FeeProperties,
		TaxProperties,
		TicketProperties,
		GoalProperties,
	], error,
) {
	if ts == nil {
		return nil, ErrTokenSourceRequired
	}
	hc := NewHTTPClient(ts, opts...)
	return &Client[
		ContactProperties,
		CompanyProperties,
		DealProperties,
		FeedbackSubmissionProperties,
		LineItemProperties,
		ProductProperties,
		QuoteProperties,
		DiscountProperties,
		FeeProperties,
		TaxProperties,
		TicketProperties,
		GoalProperties,
	]{
		HTTPClient: hc,
		Properties: NewPropertiesClient(hc),
		Contacts:   NewContactsClient[ContactProperties](hc),
		Companies:  NewCompaniesClient[CompanyProperties](hc),
		Deals:      NewDealsClient[DealProperties](hc),
		// FeedbackSubmissions: NewFeedbackSubmissionsClient( endpoint.FeedbackSubmissions,hc),
		LineItems: NewLineItemsClient[LineItemProperties](hc),
		Products:  NewProductsClient[ProductProperties](hc),
		Quotes:    NewQuotesClient[QuoteProperties](hc),
		Discounts: NewDiscountsClient[DiscountProperties](hc),
		Fees:      NewFeesClient[FeeProperties](hc),
		Taxes:     NewTaxesClient[TaxProperties](hc),
		Tickets:   NewTicketsClient[TicketProperties](hc),
		Goals:     NewGoalsClient[GoalProperties](hc),
	}, nil
}

type Client[
	ContactProperties ContactPropertiesEmbedder,
	CompanyProperties CompanyPropertiesEmbedder,
	DealProperties DealPropertiesEmbedder,
	FeedbackSubmissionProperties FeedbackSubmissionPropertiesEmbedder,
	LineItemProperties LineItemPropertiesEmbedder,
	ProductProperties ProductPropertiesEmbedder,
	QuoteProperties QuotePropertiesEmbedder,
	DiscountProperties DiscountPropertiesEmbedder,
	FeeProperties FeePropertiesEmbedder,
	TaxProperties TaxPropertiesEmbedder,
	TicketProperties TicketPropertiesEmbedder,
	GoalProperties GoalPropertiesEmbedder,
] struct {
	HTTPClient *http.Client

	Contacts  *ObjectClient[ContactProperties]
	Companies *ObjectClient[CompanyProperties]
	Deals     *ObjectClient[DealProperties]
	// FeedbackSubmissions *ObjectClient[FeedbackSubmissionProperties]
	LineItems *ObjectClient[LineItemProperties]
	Products  *ObjectClient[ProductProperties]
	Quotes    *ObjectClient[QuoteProperties]
	Discounts *ObjectClient[DiscountProperties]
	Fees      *ObjectClient[FeeProperties]
	Taxes     *ObjectClient[TaxProperties]
	Tickets   *ObjectClient[TicketProperties]
	Goals     *ObjectClient[GoalProperties]

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
//
// Allowed options:
//   - WithArchived: include only archived objects in the response
//   - WithProperties: include only the specified properties in the response
//
// Any other option is ignored.
func (poc *PropertiesObjectClient) List(ctx context.Context, opts ...RequestOption) ([]Property, error) {
	cfg := applyRequestOptions(nil, opts...)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, poc.endpoint, nil)
	if err != nil {
		return nil, err
	}
	applyQueryOptions(cfg, req.URL, applyArchivedQuery, applyPropertiesQuery)
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

type PaginatedResults[T any] struct {
	Paging  Paging `json:"paging"`
	Results []T    `json:"results"`
}

type Paging struct {
	// Next is the paging information for the next page of results.
	Next PagingNext `json:"next,omitempty"`

	// Prev is the paging information for the previous page of results.
	Prev PagingPrev `json:"prev,omitempty"`
}

type PagingNext struct {
	// Link is the URL query string to use to get the next page of results.
	//
	// It starts with a question mark and includes the "after" parameter.
	//
	// Example: "?after=NTI1Cg%3D%3D"
	Link string `json:"link,omitempty"`

	// After is the paging cursor token used as value of the "after"
	// parameter in the "link" field.
	//
	// It may be useful to use this value to construct the next request
	// when other parameters are also present in the query string.
	//
	// Example: "NTI1Cg%3D%3D"
	After string `json:"after,omitempty"`
}

type PagingPrev struct {
	// Link is the URL query string to use to get the previous page of results.
	//
	// It starts with a question mark and includes the "before" parameter.
	//
	// Example: "?before=NTI1Cg%3D%3D"
	Link string `json:"link,omitempty"`

	// Before is the paging cursor token used as value of the "before"
	// parameter in the "link" field.
	//
	// It may be useful to use this value to construct the previous request
	// when other parameters are also present in the query string.
	//
	// Example: "NTI1Cg%3D%3D"
	Before string `json:"before,omitempty"`
}

// ObjectListResults is a generic struct that contains a paginated list of
// objects returned by the HubSpot API. Associations and PropertiesWithHistory
// are also paginated.
type ObjectListResults[PE ObjectPropertiesEmbedder] PaginatedResults[GenericPublicObject[
	PE,
	PaginatedResults[PropertyWithHistory],
	PaginatedResults[AssociationEdge],
]]

// ObjectRead is a generic struct that contains a single object returned by the
// HubSpot API. Associations are also paginated.
type ObjectRead[PE ObjectPropertiesEmbedder] GenericPublicObject[
	PE,
	PropertyWithHistory,
	PaginatedResults[AssociationEdge],
]
