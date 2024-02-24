package hubspotgo

const (
	BaseURL = "https://api.hubapi.com"

	CRMV3 = BaseURL + "/crm/v3"
	CRMV4 = BaseURL + "/crm/v4"
)

const (
	// Objects

	Contacts             = CRMV3 + "/objects/contacts"
	ContactsBatchArchive = Contacts + BatchArchive
	ContactsBatchCreate  = Contacts + BatchCreate
	ContactsBatchRead    = Contacts + BatchRead
	ContactsBatchUpdate  = Contacts + BatchUpdate
	ContactsMerge        = Contacts + Merge
	ContactsGDPRDelete   = Contacts + GDPRDelete
	ContactsSearch       = Contacts + Search

	Companies             = CRMV3 + "/objects/companies"
	CompaniesBatchArchive = Companies + BatchArchive
	CompaniesBatchCreate  = Companies + BatchCreate
	CompaniesBatchRead    = Companies + BatchRead
	CompaniesBatchUpdate  = Companies + BatchUpdate
	CompaniesMerge        = Companies + Merge
	CompaniesGDPRDelete   = Companies + GDPRDelete
	CompaniesSearch       = Companies + Search

	Deals             = CRMV3 + "/objects/deals"
	DealsBatchArchive = Deals + BatchArchive
	DealsBatchCreate  = Deals + BatchCreate
	DealsBatchRead    = Deals + BatchRead
	DealsBatchUpdate  = Deals + BatchUpdate
	DealsMerge        = Deals + Merge
	DealsGDPRDelete   = Deals + GDPRDelete
	DealsSearch       = Deals + Search

	FeedbackSubmissions          = CRMV3 + "/objects/feedback_submissions"
	FeedbackSubmissionsBatchRead = FeedbackSubmissions + BatchRead

	LineItems             = CRMV3 + "/objects/line_items"
	LineItemsBatchArchive = LineItems + BatchArchive
	LineItemsBatchCreate  = LineItems + BatchCreate
	LineItemsBatchRead    = LineItems + BatchRead
	LineItemsBatchUpdate  = LineItems + BatchUpdate
	LineItemsMerge        = LineItems + Merge
	LineItemsGDPRDelete   = LineItems + GDPRDelete
	LineItemsSearch       = LineItems + Search

	Products             = CRMV3 + "/objects/products"
	ProductsBatchArchive = Products + BatchArchive
	ProductsBatchCreate  = Products + BatchCreate
	ProductsBatchRead    = Products + BatchRead
	ProductsBatchUpdate  = Products + BatchUpdate
	ProductsMerge        = Products + Merge
	ProductsGDPRDelete   = Products + GDPRDelete
	ProductsSearch       = Products + Search

	Quotes             = CRMV3 + "/objects/quotes"
	QuotesBatchArchive = Quotes + BatchArchive
	QuotesBatchCreate  = Quotes + BatchCreate
	QuotesBatchRead    = Quotes + BatchRead
	QuotesBatchUpdate  = Quotes + BatchUpdate
	QuotesMerge        = Quotes + Merge
	QuotesGDPRDelete   = Quotes + GDPRDelete
	QuotesSearch       = Quotes + Search

	Discounts             = CRMV3 + "/objects/discounts"
	DiscountsBatchArchive = Discounts + BatchArchive
	DiscountsBatchCreate  = Discounts + BatchCreate
	DiscountsBatchRead    = Discounts + BatchRead
	DiscountsBatchUpdate  = Discounts + BatchUpdate
	DiscountsMerge        = Discounts + Merge
	DiscountsGDPRDelete   = Discounts + GDPRDelete
	DiscountsSearch       = Discounts + Search

	Fees             = CRMV3 + "/objects/fees"
	FeesBatchArchive = Fees + BatchArchive
	FeesBatchCreate  = Fees + BatchCreate
	FeesBatchRead    = Fees + BatchRead
	FeesBatchUpdate  = Fees + BatchUpdate
	FeesMerge        = Fees + Merge
	FeesGDPRDelete   = Fees + GDPRDelete
	FeesSearch       = Fees + Search

	Taxes             = CRMV3 + "/objects/taxes"
	TaxesBatchArchive = Taxes + BatchArchive
	TaxesBatchCreate  = Taxes + BatchCreate
	TaxesBatchRead    = Taxes + BatchRead
	TaxesBatchUpdate  = Taxes + BatchUpdate
	TaxesMerge        = Taxes + Merge
	TaxesGDPRDelete   = Taxes + GDPRDelete
	TaxesSearch       = Taxes + Search

	Tickets             = CRMV3 + "/objects/tickets"
	TicketsBatchArchive = Tickets + BatchArchive
	TicketsBatchCreate  = Tickets + BatchCreate
	TicketsBatchRead    = Tickets + BatchRead
	TicketsBatchUpdate  = Tickets + BatchUpdate
	TicketsMerge        = Tickets + Merge
	TicketsGDPRDelete   = Tickets + GDPRDelete
	TicketsSearch       = Tickets + Search

	Goals             = CRMV3 + "/objects/goal_targets"
	GoalsBatchArchive = Goals + BatchArchive
	GoalsBatchCreate  = Goals + BatchCreate
	GoalsBatchRead    = Goals + BatchRead
	GoalsBatchUpdate  = Goals + BatchUpdate
	GoalsMerge        = Goals + Merge
	GoalsGDPRDelete   = Goals + GDPRDelete
	GoalsSearch       = Goals + Search
)

const (
	// Common Operations

	BatchArchive = "/batch/archive"
	BatchCreate  = "/batch/create"
	BatchRead    = "/batch/read"
	BatchUpdate  = "/batch/update"

	Merge      = "/merge"
	GDPRDelete = "/gdpr-delete"
	Search     = "/search"
)
