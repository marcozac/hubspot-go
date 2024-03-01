package endpoint

const (
	BaseURL      = "https://api.hubapi.com"
	AuthorizeURL = "https://app.hubspot.com/oauth/authorize"

	CRMV3 = BaseURL + "/crm/v3"
	CRMV4 = BaseURL + "/crm/v4"
)

const (
	// OAuth

	OAuth = BaseURL + "/oauth/v1"

	OAuthToken         = OAuth + "/token"
	OAuthAccessTokens  = OAuth + "/access-tokens"
	OAuthRefreshTokens = OAuth + "/refresh-tokens"
)

const (
	// Objects

	Objects = CRMV3 + "/objects"

	Contacts             = Objects + "/contacts"
	ContactsBatch        = Contacts + Batch
	ContactsBatchArchive = ContactsBatch + Archive
	ContactsBatchCreate  = ContactsBatch + Create
	ContactsBatchRead    = ContactsBatch + Read
	ContactsBatchUpdate  = ContactsBatch + Update
	ContactsMerge        = Contacts + Merge
	ContactsGDPRDelete   = Contacts + GDPRDelete
	ContactsSearch       = Contacts + Search

	Companies             = Objects + "/companies"
	CompaniesBatch        = Companies + Batch
	CompaniesBatchArchive = CompaniesBatch + Archive
	CompaniesBatchCreate  = CompaniesBatch + Create
	CompaniesBatchRead    = CompaniesBatch + Read
	CompaniesBatchUpdate  = CompaniesBatch + Update
	CompaniesMerge        = Companies + Merge
	CompaniesGDPRDelete   = Companies + GDPRDelete
	CompaniesSearch       = Companies + Search

	Deals             = Objects + "/deals"
	DealsBatch        = Deals + Batch
	DealsBatchArchive = DealsBatch + Archive
	DealsBatchCreate  = DealsBatch + Create
	DealsBatchRead    = DealsBatch + Read
	DealsBatchUpdate  = DealsBatch + Update
	DealsMerge        = Deals + Merge
	DealsGDPRDelete   = Deals + GDPRDelete
	DealsSearch       = Deals + Search

	FeedbackSubmissions          = Objects + "/feedback_submissions"
	FeedbackSubmissionsBatch     = FeedbackSubmissions + Batch
	FeedbackSubmissionsBatchRead = FeedbackSubmissionsBatch + Read

	LineItems             = Objects + "/line_items"
	LineItemsBatch        = LineItems + Batch
	LineItemsBatchArchive = LineItemsBatch + Archive
	LineItemsBatchCreate  = LineItemsBatch + Create
	LineItemsBatchRead    = LineItemsBatch + Read
	LineItemsBatchUpdate  = LineItemsBatch + Update
	LineItemsMerge        = LineItems + Merge
	LineItemsGDPRDelete   = LineItems + GDPRDelete
	LineItemsSearch       = LineItems + Search

	Products             = Objects + "/products"
	ProductsBatch        = Products + Batch
	ProductsBatchArchive = ProductsBatch + Archive
	ProductsBatchCreate  = ProductsBatch + Create
	ProductsBatchRead    = ProductsBatch + Read
	ProductsBatchUpdate  = ProductsBatch + Update
	ProductsMerge        = Products + Merge
	ProductsGDPRDelete   = Products + GDPRDelete
	ProductsSearch       = Products + Search

	Quotes             = Objects + "/quotes"
	QuotesBatch        = Quotes + Batch
	QuotesBatchArchive = QuotesBatch + Archive
	QuotesBatchCreate  = QuotesBatch + Create
	QuotesBatchRead    = QuotesBatch + Read
	QuotesBatchUpdate  = QuotesBatch + Update
	QuotesMerge        = Quotes + Merge
	QuotesGDPRDelete   = Quotes + GDPRDelete
	QuotesSearch       = Quotes + Search

	Discounts             = Objects + "/discounts"
	DiscountsBatch        = Discounts + Batch
	DiscountsBatchArchive = DiscountsBatch + Archive
	DiscountsBatchCreate  = DiscountsBatch + Create
	DiscountsBatchRead    = DiscountsBatch + Read
	DiscountsBatchUpdate  = DiscountsBatch + Update
	DiscountsMerge        = Discounts + Merge
	DiscountsGDPRDelete   = Discounts + GDPRDelete
	DiscountsSearch       = Discounts + Search

	Fees             = Objects + "/fees"
	FeesBatch        = Fees + Batch
	FeesBatchArchive = FeesBatch + Archive
	FeesBatchCreate  = FeesBatch + Create
	FeesBatchRead    = FeesBatch + Read
	FeesBatchUpdate  = FeesBatch + Update
	FeesMerge        = Fees + Merge
	FeesGDPRDelete   = Fees + GDPRDelete
	FeesSearch       = Fees + Search

	Taxes             = Objects + "/taxes"
	TaxesBatch        = Taxes + Batch
	TaxesBatchArchive = TaxesBatch + Archive
	TaxesBatchCreate  = TaxesBatch + Create
	TaxesBatchRead    = TaxesBatch + Read
	TaxesBatchUpdate  = TaxesBatch + Update
	TaxesMerge        = Taxes + Merge
	TaxesGDPRDelete   = Taxes + GDPRDelete
	TaxesSearch       = Taxes + Search

	Tickets             = Objects + "/tickets"
	TicketsBatch        = Tickets + Batch
	TicketsBatchArchive = TicketsBatch + Archive
	TicketsBatchCreate  = TicketsBatch + Create
	TicketsBatchRead    = TicketsBatch + Read
	TicketsBatchUpdate  = TicketsBatch + Update
	TicketsMerge        = Tickets + Merge
	TicketsGDPRDelete   = Tickets + GDPRDelete
	TicketsSearch       = Tickets + Search

	Goals             = Objects + "/goal_targets"
	GoalsBatch        = Goals + Batch
	GoalsBatchArchive = GoalsBatch + Archive
	GoalsBatchCreate  = GoalsBatch + Create
	GoalsBatchRead    = GoalsBatch + Read
	GoalsBatchUpdate  = GoalsBatch + Update
	GoalsMerge        = Goals + Merge
	GoalsGDPRDelete   = Goals + GDPRDelete
	GoalsSearch       = Goals + Search
)

const (
	// Properties

	Properties = CRMV3 + "/properties"

	ContactProperties             = Properties + "/contact"
	ContactPropertiesBatch        = ContactProperties + Batch
	ContactPropertiesBatchArchive = ContactPropertiesBatch + Archive
	ContactPropertiesBatchCreate  = ContactPropertiesBatch + Create
	ContactPropertiesBatchRead    = ContactPropertiesBatch + Read
	ContactPropertiesGroups       = ContactProperties + Groups

	CompanyProperties             = Properties + "/company"
	CompanyPropertiesBatch        = CompanyProperties + Batch
	CompanyPropertiesBatchArchive = CompanyPropertiesBatch + Archive
	CompanyPropertiesBatchCreate  = CompanyPropertiesBatch + Create
	CompanyPropertiesBatchRead    = CompanyPropertiesBatch + Read
	CompanyPropertiesGroups       = CompanyProperties + Groups

	DealProperties             = Properties + "/deal"
	DealPropertiesBatch        = DealProperties + Batch
	DealPropertiesBatchArchive = DealPropertiesBatch + Archive
	DealPropertiesBatchCreate  = DealPropertiesBatch + Create
	DealPropertiesBatchRead    = DealPropertiesBatch + Read
	DealPropertiesGroups       = DealProperties + Groups

	FeedbackSubmissionProperties             = Properties + "/feedback_submission"
	FeedbackSubmissionPropertiesBatch        = FeedbackSubmissionProperties + Batch
	FeedbackSubmissionPropertiesBatchArchive = FeedbackSubmissionPropertiesBatch + Archive
	FeedbackSubmissionPropertiesBatchCreate  = FeedbackSubmissionPropertiesBatch + Create
	FeedbackSubmissionPropertiesBatchRead    = FeedbackSubmissionPropertiesBatch + Read
	FeedbackSubmissionPropertiesGroups       = FeedbackSubmissionProperties + Groups

	LineItemProperties             = Properties + "/line_item"
	LineItemPropertiesBatch        = LineItemProperties + Batch
	LineItemPropertiesBatchArchive = LineItemPropertiesBatch + Archive
	LineItemPropertiesBatchCreate  = LineItemPropertiesBatch + Create
	LineItemPropertiesBatchRead    = LineItemPropertiesBatch + Read
	LineItemPropertiesGroups       = LineItemProperties + Groups

	ProductProperties             = Properties + "/product"
	ProductPropertiesBatch        = ProductProperties + Batch
	ProductPropertiesBatchArchive = ProductPropertiesBatch + Archive
	ProductPropertiesBatchCreate  = ProductPropertiesBatch + Create
	ProductPropertiesBatchRead    = ProductPropertiesBatch + Read
	ProductPropertiesGroups       = ProductProperties + Groups

	QuoteProperties             = Properties + "/quote"
	QuotePropertiesBatch        = QuoteProperties + Batch
	QuotePropertiesBatchArchive = QuotePropertiesBatch + Archive
	QuotePropertiesBatchCreate  = QuotePropertiesBatch + Create
	QuotePropertiesBatchRead    = QuotePropertiesBatch + Read
	QuotePropertiesGroups       = QuoteProperties + Groups

	DiscountProperties             = Properties + "/discount"
	DiscountPropertiesBatch        = DiscountProperties + Batch
	DiscountPropertiesBatchArchive = DiscountPropertiesBatch + Archive
	DiscountPropertiesBatchCreate  = DiscountPropertiesBatch + Create
	DiscountPropertiesBatchRead    = DiscountPropertiesBatch + Read
	DiscountPropertiesGroups       = DiscountProperties + Groups

	FeeProperties             = Properties + "/fee"
	FeePropertiesBatch        = FeeProperties + Batch
	FeePropertiesBatchArchive = FeePropertiesBatch + Archive
	FeePropertiesBatchCreate  = FeePropertiesBatch + Create
	FeePropertiesBatchRead    = FeePropertiesBatch + Read
	FeePropertiesGroups       = FeeProperties + Groups

	TaxProperties             = Properties + "/tax"
	TaxPropertiesBatch        = TaxProperties + Batch
	TaxPropertiesBatchArchive = TaxPropertiesBatch + Archive
	TaxPropertiesBatchCreate  = TaxPropertiesBatch + Create
	TaxPropertiesBatchRead    = TaxPropertiesBatch + Read
	TaxPropertiesGroups       = TaxProperties + Groups

	TicketProperties             = Properties + "/ticket"
	TicketPropertiesBatch        = TicketProperties + Batch
	TicketPropertiesBatchArchive = TicketPropertiesBatch + Archive
	TicketPropertiesBatchCreate  = TicketPropertiesBatch + Create
	TicketPropertiesBatchRead    = TicketPropertiesBatch + Read
	TicketPropertiesGroups       = TicketProperties + Groups

	GoalProperties             = Properties + "/goal_target"
	GoalPropertiesBatch        = GoalProperties + Batch
	GoalPropertiesBatchArchive = GoalPropertiesBatch + Archive
	GoalPropertiesBatchCreate  = GoalPropertiesBatch + Create
	GoalPropertiesBatchRead    = GoalPropertiesBatch + Read
	GoalPropertiesGroups       = GoalProperties + Groups
)

const (
	// Common paths

	Groups = "/groups"
	Batch  = "/batch"
)

const (
	// Common Operations

	Archive = "/archive"
	Create  = "/create"
	Read    = "/read"
	Update  = "/update"

	Merge      = "/merge"
	GDPRDelete = "/gdpr-delete"
	Search     = "/search"
)
