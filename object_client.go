// Code generated, DO NOT EDIT.

package hubspot

import (
	"net/http"

	"github.com/marcozac/hubspot-go/endpoint"
)

// Default objects' properties embedder.
type ObjectPropertiesEmbedder interface {
	embedProperties()
}

// ContactPropertiesEmbedder is the interface that must be implemented to
// create a client for the Contact object.
//
// [ContactDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
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
//
//	func main() {
//		NewContactsClient[MyContactProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyContactProperties{})
//	}
type ContactPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedContactProperties()
}

func (ContactDefaultProperties) embedProperties()        {}
func (ContactDefaultProperties) embedContactProperties() {}

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

// CompanyPropertiesEmbedder is the interface that must be implemented to
// create a client for the Company object.
//
// [CompanyDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyCompanyProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewCompaniesClient[MyCompanyProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyCompanyProperties{})
//	}
type CompanyPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedCompanyProperties()
}

func (CompanyDefaultProperties) embedProperties()        {}
func (CompanyDefaultProperties) embedCompanyProperties() {}

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

// DealPropertiesEmbedder is the interface that must be implemented to
// create a client for the Deal object.
//
// [DealDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyDealProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewDealsClient[MyDealProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyDealProperties{})
//	}
type DealPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedDealProperties()
}

func (DealDefaultProperties) embedProperties()     {}
func (DealDefaultProperties) embedDealProperties() {}

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

// FeedbackSubmissionPropertiesEmbedder is the interface that must be implemented to
// create a client for the FeedbackSubmission object.
//
// [FeedbackSubmissionDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyFeedbackSubmissionProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewFeedbackSubmissionsClient[MyFeedbackSubmissionProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyFeedbackSubmissionProperties{})
//	}
type FeedbackSubmissionPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedFeedbackSubmissionProperties()
}

func (FeedbackSubmissionDefaultProperties) embedProperties()                   {}
func (FeedbackSubmissionDefaultProperties) embedFeedbackSubmissionProperties() {}

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

// LineItemPropertiesEmbedder is the interface that must be implemented to
// create a client for the LineItem object.
//
// [LineItemDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyLineItemProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewLineItemsClient[MyLineItemProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyLineItemProperties{})
//	}
type LineItemPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedLineItemProperties()
}

func (LineItemDefaultProperties) embedProperties()         {}
func (LineItemDefaultProperties) embedLineItemProperties() {}

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

// ProductPropertiesEmbedder is the interface that must be implemented to
// create a client for the Product object.
//
// [ProductDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyProductProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewProductsClient[MyProductProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyProductProperties{})
//	}
type ProductPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedProductProperties()
}

func (ProductDefaultProperties) embedProperties()        {}
func (ProductDefaultProperties) embedProductProperties() {}

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

// QuotePropertiesEmbedder is the interface that must be implemented to
// create a client for the Quote object.
//
// [QuoteDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyQuoteProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewQuotesClient[MyQuoteProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyQuoteProperties{})
//	}
type QuotePropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedQuoteProperties()
}

func (QuoteDefaultProperties) embedProperties()      {}
func (QuoteDefaultProperties) embedQuoteProperties() {}

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

// DiscountPropertiesEmbedder is the interface that must be implemented to
// create a client for the Discount object.
//
// [DiscountDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyDiscountProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewDiscountsClient[MyDiscountProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyDiscountProperties{})
//	}
type DiscountPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedDiscountProperties()
}

func (DiscountDefaultProperties) embedProperties()         {}
func (DiscountDefaultProperties) embedDiscountProperties() {}

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

// FeePropertiesEmbedder is the interface that must be implemented to
// create a client for the Fee object.
//
// [FeeDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyFeeProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewFeesClient[MyFeeProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyFeeProperties{})
//	}
type FeePropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedFeeProperties()
}

func (FeeDefaultProperties) embedProperties()    {}
func (FeeDefaultProperties) embedFeeProperties() {}

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

// TaxPropertiesEmbedder is the interface that must be implemented to
// create a client for the Tax object.
//
// [TaxDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyTaxProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewTaxesClient[MyTaxProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyTaxProperties{})
//	}
type TaxPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedTaxProperties()
}

func (TaxDefaultProperties) embedProperties()    {}
func (TaxDefaultProperties) embedTaxProperties() {}

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

// TicketPropertiesEmbedder is the interface that must be implemented to
// create a client for the Ticket object.
//
// [TicketDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyTicketProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewTicketsClient[MyTicketProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyTicketProperties{})
//	}
type TicketPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedTicketProperties()
}

func (TicketDefaultProperties) embedProperties()       {}
func (TicketDefaultProperties) embedTicketProperties() {}

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

// GoalPropertiesEmbedder is the interface that must be implemented to
// create a client for the Goal object.
//
// [GoalDefaultProperties] already implements this interface. You can
// embed it in you own properties definition.
//
// Example:
//
//	type MyGoalProperties struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomProperty is a an example custom property.
//		MyCustomProperty string `json:"my_custom_property,omitempty"`
//	}
//
//	func main() {
//		NewGoalsClient[MyGoalProperties](httpClient)
//		// or
//		NewObjectClient(httpClient, MyGoalProperties{})
//	}
type GoalPropertiesEmbedder interface {
	ObjectPropertiesEmbedder
	embedGoalProperties()
}

func (GoalDefaultProperties) embedProperties()     {}
func (GoalDefaultProperties) embedGoalProperties() {}

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
