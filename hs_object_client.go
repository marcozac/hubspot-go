// Code generated, DO NOT EDIT.

package hubspot

import (
	"net/http"

	"github.com/marcozac/hubspot-go/endpoint"
)

// NewContactsClient returns a new client for the HubSpot
// Contacts API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [ContactPropertiesEmbedder] example.
func NewContactsClient[PE ContactPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Contacts, httpClient, pe...)
}

// NewCompaniesClient returns a new client for the HubSpot
// Companies API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [CompanyPropertiesEmbedder] example.
func NewCompaniesClient[PE CompanyPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Companies, httpClient, pe...)
}

// NewDealsClient returns a new client for the HubSpot
// Deals API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [DealPropertiesEmbedder] example.
func NewDealsClient[PE DealPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Deals, httpClient, pe...)
}

// NewFeedbackSubmissionsClient returns a new client for the HubSpot
// FeedbackSubmissions API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [FeedbackSubmissionPropertiesEmbedder] example.
func NewFeedbackSubmissionsClient[PE FeedbackSubmissionPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.FeedbackSubmissions, httpClient, pe...)
}

// NewLineItemsClient returns a new client for the HubSpot
// LineItems API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [LineItemPropertiesEmbedder] example.
func NewLineItemsClient[PE LineItemPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.LineItems, httpClient, pe...)
}

// NewProductsClient returns a new client for the HubSpot
// Products API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [ProductPropertiesEmbedder] example.
func NewProductsClient[PE ProductPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Products, httpClient, pe...)
}

// NewQuotesClient returns a new client for the HubSpot
// Quotes API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [QuotePropertiesEmbedder] example.
func NewQuotesClient[PE QuotePropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Quotes, httpClient, pe...)
}

// NewDiscountsClient returns a new client for the HubSpot
// Discounts API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [DiscountPropertiesEmbedder] example.
func NewDiscountsClient[PE DiscountPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Discounts, httpClient, pe...)
}

// NewFeesClient returns a new client for the HubSpot
// Fees API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [FeePropertiesEmbedder] example.
func NewFeesClient[PE FeePropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Fees, httpClient, pe...)
}

// NewTaxesClient returns a new client for the HubSpot
// Taxes API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [TaxPropertiesEmbedder] example.
func NewTaxesClient[PE TaxPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Taxes, httpClient, pe...)
}

// NewTicketsClient returns a new client for the HubSpot
// Tickets API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [TicketPropertiesEmbedder] example.
func NewTicketsClient[PE TicketPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Tickets, httpClient, pe...)
}

// NewGoalsClient returns a new client for the HubSpot
// Goals API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// See the [GoalPropertiesEmbedder] example.
func NewGoalsClient[PE GoalPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Goals, httpClient, pe...)
}
