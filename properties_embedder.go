// Code generated, DO NOT EDIT.

package hubspot

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
