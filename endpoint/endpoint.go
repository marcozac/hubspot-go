package endpoint

const (
	BaseURL = "https://api.hubapi.com"

	CRMV3 = BaseURL + "/crm/v3"
	CRMV4 = BaseURL + "/crm/v4"
)

const (
	// Objects

	Objects = CRMV3 + "/objects"

	Contacts             = Objects + "/contacts"
	ContactsBatchArchive = Contacts + BatchArchive
	ContactsBatchCreate  = Contacts + BatchCreate
	ContactsBatchRead    = Contacts + BatchRead
	ContactsBatchUpdate  = Contacts + BatchUpdate
	ContactsMerge        = Contacts + Merge
	ContactsGDPRDelete   = Contacts + GDPRDelete
	ContactsSearch       = Contacts + Search

	Companies             = Objects + "/companies"
	CompaniesBatchArchive = Companies + BatchArchive
	CompaniesBatchCreate  = Companies + BatchCreate
	CompaniesBatchRead    = Companies + BatchRead
	CompaniesBatchUpdate  = Companies + BatchUpdate
	CompaniesMerge        = Companies + Merge
	CompaniesGDPRDelete   = Companies + GDPRDelete
	CompaniesSearch       = Companies + Search

	Deals             = Objects + "/deals"
	DealsBatchArchive = Deals + BatchArchive
	DealsBatchCreate  = Deals + BatchCreate
	DealsBatchRead    = Deals + BatchRead
	DealsBatchUpdate  = Deals + BatchUpdate
	DealsMerge        = Deals + Merge
	DealsGDPRDelete   = Deals + GDPRDelete
	DealsSearch       = Deals + Search

	FeedbackSubmissions          = Objects + "/feedback_submissions"
	FeedbackSubmissionsBatchRead = FeedbackSubmissions + BatchRead

	LineItems             = Objects + "/line_items"
	LineItemsBatchArchive = LineItems + BatchArchive
	LineItemsBatchCreate  = LineItems + BatchCreate
	LineItemsBatchRead    = LineItems + BatchRead
	LineItemsBatchUpdate  = LineItems + BatchUpdate
	LineItemsMerge        = LineItems + Merge
	LineItemsGDPRDelete   = LineItems + GDPRDelete
	LineItemsSearch       = LineItems + Search

	Products             = Objects + "/products"
	ProductsBatchArchive = Products + BatchArchive
	ProductsBatchCreate  = Products + BatchCreate
	ProductsBatchRead    = Products + BatchRead
	ProductsBatchUpdate  = Products + BatchUpdate
	ProductsMerge        = Products + Merge
	ProductsGDPRDelete   = Products + GDPRDelete
	ProductsSearch       = Products + Search

	Quotes             = Objects + "/quotes"
	QuotesBatchArchive = Quotes + BatchArchive
	QuotesBatchCreate  = Quotes + BatchCreate
	QuotesBatchRead    = Quotes + BatchRead
	QuotesBatchUpdate  = Quotes + BatchUpdate
	QuotesMerge        = Quotes + Merge
	QuotesGDPRDelete   = Quotes + GDPRDelete
	QuotesSearch       = Quotes + Search

	Discounts             = Objects + "/discounts"
	DiscountsBatchArchive = Discounts + BatchArchive
	DiscountsBatchCreate  = Discounts + BatchCreate
	DiscountsBatchRead    = Discounts + BatchRead
	DiscountsBatchUpdate  = Discounts + BatchUpdate
	DiscountsMerge        = Discounts + Merge
	DiscountsGDPRDelete   = Discounts + GDPRDelete
	DiscountsSearch       = Discounts + Search

	Fees             = Objects + "/fees"
	FeesBatchArchive = Fees + BatchArchive
	FeesBatchCreate  = Fees + BatchCreate
	FeesBatchRead    = Fees + BatchRead
	FeesBatchUpdate  = Fees + BatchUpdate
	FeesMerge        = Fees + Merge
	FeesGDPRDelete   = Fees + GDPRDelete
	FeesSearch       = Fees + Search

	Taxes             = Objects + "/taxes"
	TaxesBatchArchive = Taxes + BatchArchive
	TaxesBatchCreate  = Taxes + BatchCreate
	TaxesBatchRead    = Taxes + BatchRead
	TaxesBatchUpdate  = Taxes + BatchUpdate
	TaxesMerge        = Taxes + Merge
	TaxesGDPRDelete   = Taxes + GDPRDelete
	TaxesSearch       = Taxes + Search

	Tickets             = Objects + "/tickets"
	TicketsBatchArchive = Tickets + BatchArchive
	TicketsBatchCreate  = Tickets + BatchCreate
	TicketsBatchRead    = Tickets + BatchRead
	TicketsBatchUpdate  = Tickets + BatchUpdate
	TicketsMerge        = Tickets + Merge
	TicketsGDPRDelete   = Tickets + GDPRDelete
	TicketsSearch       = Tickets + Search

	Goals             = Objects + "/goal_targets"
	GoalsBatchArchive = Goals + BatchArchive
	GoalsBatchCreate  = Goals + BatchCreate
	GoalsBatchRead    = Goals + BatchRead
	GoalsBatchUpdate  = Goals + BatchUpdate
	GoalsMerge        = Goals + Merge
	GoalsGDPRDelete   = Goals + GDPRDelete
	GoalsSearch       = Goals + Search
)

const (
	// Properties

	Properties = CRMV3 + "/properties"

	ContactProperties             = Properties + "/contact"
	ContactPropertiesBatchArchive = ContactProperties + BatchArchive
	ContactPropertiesBatchCreate  = ContactProperties + BatchCreate
	ContactPropertiesBatchRead    = ContactProperties + BatchRead
	ContactPropertiesGroups       = ContactProperties + Groups

	CompanyProperties             = Properties + "/company"
	CompanyPropertiesBatchArchive = CompanyProperties + BatchArchive
	CompanyPropertiesBatchCreate  = CompanyProperties + BatchCreate
	CompanyPropertiesBatchRead    = CompanyProperties + BatchRead
	CompanyPropertiesGroups       = CompanyProperties + Groups

	DealProperties             = Properties + "/deal"
	DealPropertiesBatchArchive = DealProperties + BatchArchive
	DealPropertiesBatchCreate  = DealProperties + BatchCreate
	DealPropertiesBatchRead    = DealProperties + BatchRead
	DealPropertiesGroups       = DealProperties + Groups

	FeedbackSubmissionProperties             = Properties + "/feedback_submission"
	FeedbackSubmissionPropertiesBatchArchive = FeedbackSubmissionProperties + BatchArchive
	FeedbackSubmissionPropertiesBatchCreate  = FeedbackSubmissionProperties + BatchCreate
	FeedbackSubmissionPropertiesBatchRead    = FeedbackSubmissionProperties + BatchRead
	FeedbackSubmissionPropertiesGroups       = FeedbackSubmissionProperties + Groups

	LineItemProperties             = Properties + "/line_item"
	LineItemPropertiesBatchArchive = LineItemProperties + BatchArchive
	LineItemPropertiesBatchCreate  = LineItemProperties + BatchCreate
	LineItemPropertiesBatchRead    = LineItemProperties + BatchRead
	LineItemPropertiesGroups       = LineItemProperties + Groups

	ProductProperties             = Properties + "/product"
	ProductPropertiesBatchArchive = ProductProperties + BatchArchive
	ProductPropertiesBatchCreate  = ProductProperties + BatchCreate
	ProductPropertiesBatchRead    = ProductProperties + BatchRead
	ProductPropertiesGroups       = ProductProperties + Groups

	QuoteProperties             = Properties + "/quote"
	QuotePropertiesBatchArchive = QuoteProperties + BatchArchive
	QuotePropertiesBatchCreate  = QuoteProperties + BatchCreate
	QuotePropertiesBatchRead    = QuoteProperties + BatchRead
	QuotePropertiesGroups       = QuoteProperties + Groups

	DiscountProperties             = Properties + "/discount"
	DiscountPropertiesBatchArchive = DiscountProperties + BatchArchive
	DiscountPropertiesBatchCreate  = DiscountProperties + BatchCreate
	DiscountPropertiesBatchRead    = DiscountProperties + BatchRead
	DiscountPropertiesGroups       = DiscountProperties + Groups

	FeeProperties             = Properties + "/fee"
	FeePropertiesBatchArchive = FeeProperties + BatchArchive
	FeePropertiesBatchCreate  = FeeProperties + BatchCreate
	FeePropertiesBatchRead    = FeeProperties + BatchRead
	FeePropertiesGroups       = FeeProperties + Groups

	TaxProperties             = Properties + "/tax"
	TaxPropertiesBatchArchive = TaxProperties + BatchArchive
	TaxPropertiesBatchCreate  = TaxProperties + BatchCreate
	TaxPropertiesBatchRead    = TaxProperties + BatchRead
	TaxPropertiesGroups       = TaxProperties + Groups

	TicketProperties             = Properties + "/ticket"
	TicketPropertiesBatchArchive = TicketProperties + BatchArchive
	TicketPropertiesBatchCreate  = TicketProperties + BatchCreate
	TicketPropertiesBatchRead    = TicketProperties + BatchRead
	TicketPropertiesGroups       = TicketProperties + Groups

	GoalProperties             = Properties + "/goal_target"
	GoalPropertiesBatchArchive = GoalProperties + BatchArchive
	GoalPropertiesBatchCreate  = GoalProperties + BatchCreate
	GoalPropertiesBatchRead    = GoalProperties + BatchRead
	GoalPropertiesGroups       = GoalProperties + Groups
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

const (
	// Other common paths

	Groups = "/groups"
)
