// Code generated by hsc. DO NOT EDIT.

package hubspot

import (
	"net/http"

	endpoint "github.com/marcozac/hubspot-go/endpoint"
)

// ContactPropertiesEmbedder is the interface that must be implemented to
// create a client for the Contact object.
//
// Example:
//
//	type MyContactProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type ContactPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedContactProperties()
}

// NewContactClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewContactClient[PE ContactPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Contacts, httpClient, pe...)
}

// CompanyPropertiesEmbedder is the interface that must be implemented to
// create a client for the Company object.
//
// Example:
//
//	type MyCompanyProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		CompanyDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type CompanyPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedCompanyProperties()
}

// NewCompanyClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewCompanyClient[PE CompanyPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Companies, httpClient, pe...)
}

// DealPropertiesEmbedder is the interface that must be implemented to
// create a client for the Deal object.
//
// Example:
//
//	type MyDealProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		DealDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type DealPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedDealProperties()
}

// NewDealClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewDealClient[PE DealPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Deals, httpClient, pe...)
}

// FeedbackSubmissionPropertiesEmbedder is the interface that must be implemented to
// create a client for the FeedbackSubmission object.
//
// Example:
//
//	type MyFeedbackSubmissionProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		FeedbackSubmissionDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type FeedbackSubmissionPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedFeedbackSubmissionProperties()
}

// NewFeedbackSubmissionClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewFeedbackSubmissionClient[PE FeedbackSubmissionPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.FeedbackSubmissions, httpClient, pe...)
}

// LineItemPropertiesEmbedder is the interface that must be implemented to
// create a client for the LineItem object.
//
// Example:
//
//	type MyLineItemProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		LineItemDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type LineItemPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedLineItemProperties()
}

// NewLineItemClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewLineItemClient[PE LineItemPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.LineItems, httpClient, pe...)
}

// ProductPropertiesEmbedder is the interface that must be implemented to
// create a client for the Product object.
//
// Example:
//
//	type MyProductProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ProductDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type ProductPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedProductProperties()
}

// NewProductClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewProductClient[PE ProductPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Products, httpClient, pe...)
}

// QuotePropertiesEmbedder is the interface that must be implemented to
// create a client for the Quote object.
//
// Example:
//
//	type MyQuoteProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		QuoteDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type QuotePropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedQuoteProperties()
}

// NewQuoteClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewQuoteClient[PE QuotePropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Quotes, httpClient, pe...)
}

// DiscountPropertiesEmbedder is the interface that must be implemented to
// create a client for the Discount object.
//
// Example:
//
//	type MyDiscountProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		DiscountDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type DiscountPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedDiscountProperties()
}

// NewDiscountClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewDiscountClient[PE DiscountPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Discounts, httpClient, pe...)
}

// FeePropertiesEmbedder is the interface that must be implemented to
// create a client for the Fee object.
//
// Example:
//
//	type MyFeeProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		FeeDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type FeePropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedFeeProperties()
}

// NewFeeClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewFeeClient[PE FeePropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Fees, httpClient, pe...)
}

// TaxPropertiesEmbedder is the interface that must be implemented to
// create a client for the Tax object.
//
// Example:
//
//	type MyTaxProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		TaxDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type TaxPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedTaxProperties()
}

// NewTaxClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewTaxClient[PE TaxPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Taxes, httpClient, pe...)
}

// TicketPropertiesEmbedder is the interface that must be implemented to
// create a client for the Ticket object.
//
// Example:
//
//	type MyTicketProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		TicketDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type TicketPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedTicketProperties()
}

// NewTicketClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewTicketClient[PE TicketPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Tickets, httpClient, pe...)
}

// GoalPropertiesEmbedder is the interface that must be implemented to
// create a client for the Goal object.
//
// Example:
//
//	type MyGoalProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		GoalDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
type GoalPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	EmbedGoalProperties()
}

// NewGoalClient returns a new client for the HubSpot API.
//
// The given HTTP client is used to make requests to the HubSpot API and must
// be configured to use the appropriate authentication method. You can use
// [NewHTTPClient] to create a new HTTP client from an OAuth2 token source.
//
// The properties type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
func NewGoalClient[PE GoalPropertiesEmbedder](httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return NewObjectClient(endpoint.Goals, httpClient, pe...)
}

// Default objects' properties embedder.
type ObjectPropertiesEmbedder interface {
	embedProperties()
}
