package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/marcozac/hubspot-go/endpoint"
	"github.com/marcozac/hubspot-go/limiter"
	"github.com/marcozac/hubspot-go/util"
)

// NewHTTPClient returns a new HTTP client that uses the given token source to
// authenticate requests to the HubSpot API.
//
// The client is also wrapped with a rate limiter that limits the number of
// requests per second to the HubSpot API. By default, the rate limiter allows
// infinite requests per second. You can use the WithLimiter option to set a
// different rate limit.
func NewHTTPClient(ts oauth2.TokenSource, opts ...ClientOption) *http.Client {
	cfg := applyClientOptions(nil, opts...)
	oc := oauth2.NewClient(cfg.ctx, ts)
	// Wrap the HTTP client with a rate limiter.
	return limiter.WrapHTTPClient(oc, cfg.limiter)
}

// NewDefaultClient returns a new HubSpot client with the given token source and
// options. It is a wrapper around NewTypedClient that uses the default
// properties for each object type.
func NewDefaultClient(ts oauth2.TokenSource, opts ...ClientOption) (*DefaultClient, error) {
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
		HTTPClient:          hc,
		Properties:          NewPropertiesClient(hc),
		Contacts:            NewContactClient[ContactProperties](hc),
		Companies:           NewCompanyClient[CompanyProperties](hc),
		Deals:               NewDealClient[DealProperties](hc),
		FeedbackSubmissions: NewFeedbackSubmissionClient[FeedbackSubmissionProperties](hc),
		LineItems:           NewLineItemClient[LineItemProperties](hc),
		Products:            NewProductClient[ProductProperties](hc),
		Quotes:              NewQuoteClient[QuoteProperties](hc),
		Discounts:           NewDiscountClient[DiscountProperties](hc),
		Fees:                NewFeeClient[FeeProperties](hc),
		Taxes:               NewTaxClient[TaxProperties](hc),
		Tickets:             NewTicketClient[TicketProperties](hc),
		Goals:               NewGoalClient[GoalProperties](hc),
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

	Contacts            *ObjectClient[ContactProperties]
	Companies           *ObjectClient[CompanyProperties]
	Deals               *ObjectClient[DealProperties]
	FeedbackSubmissions *ObjectClient[FeedbackSubmissionProperties]
	LineItems           *ObjectClient[LineItemProperties]
	Products            *ObjectClient[ProductProperties]
	Quotes              *ObjectClient[QuoteProperties]
	Discounts           *ObjectClient[DiscountProperties]
	Fees                *ObjectClient[FeeProperties]
	Taxes               *ObjectClient[TaxProperties]
	Tickets             *ObjectClient[TicketProperties]
	Goals               *ObjectClient[GoalProperties]

	Properties *PropertiesClient
	OAuth      OAuthClient
}

// DefaultClient is a client for the HubSpot API with the default properties for
// each object type.
type DefaultClient = Client[
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
]

// NewObjectClient returns a new object client for the given object type.
//
// The object type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
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
//	func main() {
//		NewObjectClient[ContactPropertiesTest](endpoint.Contacts, http.DefaultClient)
//		// or
//		NewObjectClient(endpoint.Contacts, http.DefaultClient, ContactPropertiesTest{})
//	}
func NewObjectClient[PE ObjectPropertiesEmbedder](baseEndpoint string, hc *http.Client, pe ...PE) *ObjectClient[PE] {
	return &ObjectClient[PE]{
		endpoint: baseEndpoint,
		hc:       hc,
		Batch:    NewObjectBatchClient[PE](baseEndpoint+endpoint.Batch, hc),
	}
}

type ObjectClient[PE ObjectPropertiesEmbedder] struct {
	endpoint string
	hc       *http.Client

	Batch *BatchClient[
		ObjectBatchReadInput,
		ObjectBatchCreateInput[PE],
		ObjectBatchUpdateInput[PE],
		ObjectBatchArchiveInput,
		ObjectMutation[PE],
	]
}

// List returns a paginated list of objects.
//
// Allowed options:
//   - WithLimit: the number of objects to return per page
//   - WithAfter: the cursor token to use to get the next page of results
//   - WithArchived: include only archived objects in the response
//   - WithProperties: include only the specified properties in the response
//   - WithPropertiesWithHistory: include the history of the specified
//     properties in the response
//   - WithAssociations: include only the specified associations in the response
//
// Any other option is ignored.
//
// NOTE
// The HubSpot's default limit is 10 objects per page.
func (oc *ObjectClient[PE]) List(ctx context.Context, opts ...RequestOption) (*ObjectListResults[PE], error) {
	cfg := applyRequestOptions(nil, opts...)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, oc.endpoint, nil)
	if err != nil {
		return nil, err
	}
	applyQueryOptions(cfg, req.URL,
		applyLimitQuery,
		applyAfterQuery,
		applyArchivedQuery,
		applyPropertiesQuery,
		applyPropertiesWithHistoryQuery,
		applyAssociationsQuery,
	)
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	results := new(ObjectListResults[PE])
	if err := json.NewDecoder(resp.Body).Decode(results); err != nil {
		return nil, err
	}
	return results, nil
}

// Read returns the object with the given ID.
//
// Object ID is usually a number. If it is declared as a [Int], [Int64], etc.
// it is possible to use the String method to convert it to a string.
//
// Allowed options:
//   - WithArchived: include only archived objects in the response
//   - WithProperties: include only the specified properties in the response
//   - WithPropertiesWithHistory: include the history of the specified
//     properties in the response
//   - WithAssociations: include only the specified associations in the response
//
// Any other option is ignored.
func (oc *ObjectClient[PE]) Read(ctx context.Context, id string, opts ...RequestOption) (*ObjectRead[PE], error) {
	cfg := applyRequestOptions(nil, opts...)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, oc.endpoint+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	applyQueryOptions(cfg, req.URL,
		applyArchivedQuery,
		applyPropertiesQuery,
		applyPropertiesWithHistoryQuery,
		applyAssociationsQuery,
		applyIDPropertyQuery,
	)
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := new(ObjectRead[PE])
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new object with the given properties and, optionally, the
// given associations.
func (oc *ObjectClient[PE]) Create(ctx context.Context, properties *PE, associations ...*AssociationForCreate) (*ObjectMutation[PE], error) {
	rb := &ObjectMutationRequestBody[PE]{
		Properties:   properties,
		Associations: associations,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(rb); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, oc.endpoint, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := new(ObjectMutation[PE])
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Update updates the properties of the object with the given ID.
//
// According to the HubSpot API, read only and non-writable properties are
// ignored. To clear a property, set its value to an empty string.
func (oc *ObjectClient[PE]) Update(ctx context.Context, id string, properties *PE) (*ObjectMutation[PE], error) {
	rb := &ObjectMutationRequestBody[PE]{
		Properties: properties,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(rb); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, oc.endpoint+"/"+id, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := new(ObjectMutation[PE])
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Archive archives the object with the given ID.
func (oc *ObjectClient[PE]) Archive(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, oc.endpoint+"/"+id, nil)
	if err != nil {
		return err
	}
	resp, err := oc.hc.Do(req)
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

func NewObjectBatchClient[PE ObjectPropertiesEmbedder](baseEndpoint string, hc *http.Client) *BatchClient[
	ObjectBatchReadInput,
	ObjectBatchCreateInput[PE],
	ObjectBatchUpdateInput[PE],
	ObjectBatchArchiveInput,
	ObjectMutation[PE],
] {
	return &BatchClient[
		ObjectBatchReadInput,
		ObjectBatchCreateInput[PE],
		ObjectBatchUpdateInput[PE],
		ObjectBatchArchiveInput,
		ObjectMutation[PE],
	]{
		baseEndpoint: baseEndpoint,
		hc:           hc,
	}
}

type (
	ObjectBatchReadInput struct {
		IDProperty            string          `json:"idProperty,omitempty"`
		PropertiesWithHistory []string        `json:"propertiesWithHistory,omitempty"`
		Properties            []string        `json:"properties,omitempty"`
		Inputs                []ObjectBatchID `json:"inputs"`
	}

	ObjectBatchCreateInput[PE ObjectPropertiesEmbedder] struct {
		Inputs []ObjectMutationRequestBody[PE] `json:"inputs"`
	}

	ObjectBatchUpdateInput[PE ObjectPropertiesEmbedder] struct {
		Inputs []ObjectBatchIDProperties[PE] `json:"inputs"`
	}

	ObjectBatchArchiveInput struct {
		Inputs []ObjectBatchID `json:"inputs"`
	}

	ObjectBatchID struct {
		ID string `json:"id,omitempty"`
	}

	ObjectBatchIDProperties[PE ObjectPropertiesEmbedder] struct {
		ID         string `json:"id,omitempty"`
		IDProperty string `json:"idProperty,omitempty"`
		Properties PE     `json:"properties,omitempty"`
	}
)

func (ObjectBatchReadInput) embedBatchReadInput()         {}
func (ObjectBatchCreateInput[PE]) embedBatchCreateInput() {}
func (ObjectBatchUpdateInput[PE]) embedBatchUpdateInput() {}
func (ObjectBatchArchiveInput) embedBatchArchiveInput()   {}

func NewPropertiesClient(hc *http.Client) *PropertiesClient {
	return &PropertiesClient{
		Contact: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.ContactProperties,
			Batch:    NewPropertiesBatchClient(endpoint.ContactPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.ContactPropertiesGroups, hc),
		},
		Company: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.CompanyProperties,
			Batch:    NewPropertiesBatchClient(endpoint.CompanyPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.CompanyPropertiesGroups, hc),
		},
		Deal: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.DealProperties,
			Batch:    NewPropertiesBatchClient(endpoint.DealPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.DealPropertiesGroups, hc),
		},
		FeedbackSubmission: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.FeedbackSubmissionProperties,
			Batch:    NewPropertiesBatchClient(endpoint.FeedbackSubmissionPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.FeedbackSubmissionPropertiesGroups, hc),
		},
		LineItem: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.LineItemProperties,
			Batch:    NewPropertiesBatchClient(endpoint.LineItemPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.LineItemPropertiesGroups, hc),
		},
		Product: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.ProductProperties,
			Batch:    NewPropertiesBatchClient(endpoint.ProductPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.ProductPropertiesGroups, hc),
		},
		Quote: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.QuoteProperties,
			Batch:    NewPropertiesBatchClient(endpoint.QuotePropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.QuotePropertiesGroups, hc),
		},
		Discount: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.DiscountProperties,
			Batch:    NewPropertiesBatchClient(endpoint.DiscountPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.DiscountPropertiesGroups, hc),
		},
		Fee: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.FeeProperties,
			Batch:    NewPropertiesBatchClient(endpoint.FeePropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.FeePropertiesGroups, hc),
		},
		Tax: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.TaxProperties,
			Batch:    NewPropertiesBatchClient(endpoint.TaxPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.TaxPropertiesGroups, hc),
		},
		Ticket: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.TicketProperties,
			Batch:    NewPropertiesBatchClient(endpoint.TicketPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.TicketPropertiesGroups, hc),
		},
		Goal: &PropertiesObjectClient{
			hc:       hc,
			endpoint: endpoint.GoalProperties,
			Batch:    NewPropertiesBatchClient(endpoint.GoalPropertiesBatch, hc),
			Groups:   NewPropertyGroupClient(endpoint.GoalPropertiesGroups, hc),
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

	Batch  *PropertiesBatchClient
	Groups *PropertyGroupClient
}

// List returns a list of properties for the object type.
//
// Allowed options:
//   - WithArchived: include only archived objects in the response
//   - WithProperties: include only the specified properties in the response
//
// Any other option is ignored.
func (poc *PropertiesObjectClient) List(ctx context.Context, opts ...RequestOption) ([]*Property, error) {
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
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var results Results[Property]
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return results.Results, nil
}

// Read returns the property with the given name.
//
// Allowed options:
//   - WithArchived: include only archived objects in the response
//   - WithProperties: include only the specified properties in the response
//
// Any other option is ignored.
func (poc *PropertiesObjectClient) Read(ctx context.Context, name string, opts ...RequestOption) (*Property, error) {
	cfg := applyRequestOptions(nil, opts...)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, poc.endpoint+"/"+name, nil)
	if err != nil {
		return nil, err
	}
	applyQueryOptions(cfg, req.URL, applyArchivedQuery, applyPropertiesQuery)
	resp, err := poc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	prop := new(Property)
	if err := json.NewDecoder(resp.Body).Decode(prop); err != nil {
		return nil, err
	}
	return prop, nil
}

// Create creates a new property.
//
// The given property is modified in place with the response from the API
// and returned. If you need to keep the original property, you should create
// a copy of it before calling this method.
//
// At the moment of writing, the required fields in [HubSpot's docs] are:
//   - Name
//   - Label
//   - GroupName
//   - Type
//   - FieldType
//
// In options:
//   - Hidden
//   - Label
//   - Value
//
// Missing fields are not checked, to avoid breaking changes in the future
// versions of the HubSpot API, and may cause errors in the response.
//
// [HubSpot's docs]: https://developers.hubspot.com/docs/api/crm/properties
func (poc *PropertiesObjectClient) Create(ctx context.Context, prop *Property) (*Property, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(prop); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, poc.endpoint, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := poc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(prop); err != nil {
		return nil, err
	}
	return prop, nil
}

// Update updates the given property. The property name must be set in the
// property struct, that cannot be nil.
//
// The given property is modified in place with the response from the API
// and returned. If you need to keep the original property, you should create
// a copy of it before calling this method.
func (poc *PropertiesObjectClient) Update(ctx context.Context, prop *Property) (*Property, error) {
	if prop == nil {
		return nil, fmt.Errorf("property: %w", ErrNilParam)
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(prop); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, poc.endpoint+"/"+prop.Name, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := poc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(prop); err != nil {
		return nil, err
	}
	return prop, nil
}

// Archive archives the property with the given name.
func (poc *PropertiesObjectClient) Archive(ctx context.Context, name string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, poc.endpoint+"/"+name, nil)
	if err != nil {
		return err
	}
	resp, err := poc.hc.Do(req)
	if err != nil {
		return err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func NewPropertiesBatchClient(baseEndpoint string, hc *http.Client) *PropertiesBatchClient {
	return &PropertiesBatchClient{
		baseEndpoint: baseEndpoint,
		hc:           hc,
	}
}

type (
	PropertiesBatchClient = BatchClient[
		PropertiesBatchReadInput,
		PropertiesBatchCreateInput,
		BatchInput[any], // not available
		PropertiesBatchArchiveInput,
		Property,
	]

	PropertiesBatchReadInput    = BatchInput[PropertiesBatchNameInput]
	PropertiesBatchArchiveInput = PropertiesBatchReadInput
	PropertiesBatchCreateInput  = BatchInput[Property]

	PropertiesBatchNameInput struct {
		Name string `json:"name"`
	}
)

// NewPropertyGroupClient returns a new property group client that uses the
// given HTTP client to make requests to the endpoint.
func NewPropertyGroupClient(endpoint string, hc *http.Client) *PropertyGroupClient {
	return &PropertyGroupClient{
		endpoint: endpoint,
		hc:       hc,
	}
}

type PropertyGroupClient struct {
	endpoint string
	hc       *http.Client
}

// List returns a list of property groups for the object type.
func (pgc *PropertyGroupClient) List(ctx context.Context) ([]*PropertyGroup, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pgc.endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := pgc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var results Results[PropertyGroup]
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return results.Results, nil
}

// Read returns the property group with the given name.
func (pgc *PropertyGroupClient) Read(ctx context.Context, name string) (*PropertyGroup, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pgc.endpoint+"/"+name, nil)
	if err != nil {
		return nil, err
	}
	resp, err := pgc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	pg := new(PropertyGroup)
	if err := json.NewDecoder(resp.Body).Decode(pg); err != nil {
		return nil, err
	}
	return pg, nil
}

// Create creates a new property group.
//
// The given property group is modified in place with the response from the API
// and returned. If you need to keep the original property group, you should
// create a copy of it before calling this method.
//
// At the moment of writing, the required fields in [HubSpot's docs] are:
//   - Name
//   - Label
//
// Missing fields are not checked, to avoid breaking changes in the future
// versions of the HubSpot API, and may cause errors in the response.
func (pgc *PropertyGroupClient) Create(ctx context.Context, group *PropertyGroup) (*PropertyGroup, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(group); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, pgc.endpoint, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := pgc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(group); err != nil {
		return nil, err
	}
	return group, nil
}

// Update updates the given property group. The property group name must be set
// in the property group struct, that cannot be nil.
//
// The given property group is modified in place with the response from the API
// and returned. If you need to keep the original property group, you should
// create a copy of it before calling this method.
func (pgc *PropertyGroupClient) Update(ctx context.Context, group *PropertyGroup) (*PropertyGroup, error) {
	if group == nil {
		return nil, fmt.Errorf("property group: %w", ErrNilParam)
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(group); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, pgc.endpoint+"/"+group.Name, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := pgc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(group); err != nil {
		return nil, err
	}
	return group, nil
}

// Archive archives the property group with the given name.
func (pgc *PropertyGroupClient) Archive(ctx context.Context, name string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, pgc.endpoint+"/"+name, nil)
	if err != nil {
		return err
	}
	resp, err := pgc.hc.Do(req)
	if err != nil {
		return err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// OAuthClient is a client for the HubSpot OAuth API.
//
// It can be used to get access tokens, refresh tokens, and delete refresh tokens.
//
// Token exchange and refresh are not implemented and should be handled by the
// oauth2 package from golang.org/x/oauth2. For example, you can configure an
// oauth2.Config with the HubSpot OAuth 2.0 endpoint and use the Exchange method
// to get an access token from an authorization code or a refresh token.
type OAuthClient struct{}

// GetAccessToken returns the HubSpot [AccessToken] from the given token.
func (OAuthClient) GetAccessToken(ctx context.Context, token string) (*AccessToken, error) {
	url := endpoint.OAuthAccessTokens + "/" + token
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	at := new(AccessToken)
	if err := json.NewDecoder(resp.Body).Decode(at); err != nil {
		return nil, err
	}
	return at, nil
}

// GetRefreshToken returns the HubSpot [RefreshToken] from the given token.
func (OAuthClient) GetRefreshToken(ctx context.Context, token string) (*RefreshToken, error) {
	url := endpoint.OAuthRefreshTokens + "/" + token
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rt := new(RefreshToken)
	if err := json.NewDecoder(resp.Body).Decode(rt); err != nil {
		return nil, err
	}
	return rt, nil
}

// DeleteRefreshToken deletes the given refresh token.
func (OAuthClient) DeleteRefreshToken(ctx context.Context, token string) error {
	url := endpoint.OAuthRefreshTokens + "/" + token
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if err := HubSpotResponseError(resp); err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// Results is a generic struct that contains a list of results of type T
// returned by the HubSpot API.
type Results[T any] struct {
	Results []*T `json:"results"`
}

// results is a type alias for Results for embedding without repeating the
// Results field.
type results[T any] Results[T]

type PaginatedResults[T any] struct {
	Paging  Paging `json:"paging"`
	Results []*T   `json:"results"`
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
type ObjectListResults[PE ObjectPropertiesEmbedder] PaginatedResults[PublicObject[
	PE,
	PaginatedResults[PropertyWithHistory],
	PaginatedResults[AssociationEdge],
]]

// ObjectRead is a generic struct that contains a single object returned by the
// HubSpot API. Associations are also paginated.
type ObjectRead[PE ObjectPropertiesEmbedder] PublicObject[
	PE,
	PropertyWithHistory,
	PaginatedResults[AssociationEdge],
]

// ObjectMutation is a generic struct that contains a single object returned by
// the HubSpot API after a create or update operation. Associations are usually
// not present in the response.
type ObjectMutation[PE ObjectPropertiesEmbedder] PublicObject[
	PE,
	PropertyWithHistory,
	AssociationEdge,
]

// ObjectMutationRequestBody is a generic struct that contains the request body
// for creating or updating an object in the HubSpot API. Associations should
// be set only when creating an object.
type ObjectMutationRequestBody[PE ObjectPropertiesEmbedder] struct {
	Properties   *PE                     `json:"properties,omitempty"`
	Associations []*AssociationForCreate `json:"associations,omitempty"`
}
