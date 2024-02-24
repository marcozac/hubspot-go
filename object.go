package hubspot

const (
	ObjectTypeContact            = "contact"
	ObjectTypeCompany            = "company"
	ObjectTypeDeal               = "deal"
	ObjectTypeFeedbackSubmission = "feedback_submission"
	ObjectTypeLineItem           = "line_item"
	ObjectTypeProduct            = "product"
	ObjectTypeQuote              = "quote"
	ObjectTypeDiscount           = "discount"
	ObjectTypeFee                = "fee"
	ObjectTypeTax                = "tax"
	ObjectTypeTicket             = "ticket"
	ObjectTypeGoal               = "goal_target"

	ObjectTypeIDContact            = "0-1"
	ObjectTypeIDCompany            = "0-2"
	ObjectTypeIDDeal               = "0-3"
	ObjectTypeIDTicket             = "0-5"
	ObjectTypeIDProduct            = "0-7"
	ObjectTypeIDLineItem           = "0-8"
	ObjectTypeIDQuote              = "0-14"
	ObjectTypeIDFeedbackSubmission = "0-19"
	ObjectTypeIDCall               = "0-48"
	ObjectTypeIDEmail              = "0-49"
	ObjectTypeIDMeeting            = "0-47"
	ObjectTypeIDNote               = "0-4"
	ObjectTypeIDTask               = "0-27"
	ObjectTypeIDCommunication      = "0-18"
	ObjectTypeIDPostalMail         = "0-116"
	ObjectTypeIDMarketingEvent     = "0-54"
)

type PublicObject[P any] struct {
	// ID is the HubSpot object ID.
	ID string `json:"id"`

	// Properties is the object's properties.
	Properties P `json:"properties"`
}

type SimplePublicObject = PublicObject[map[string]string]
