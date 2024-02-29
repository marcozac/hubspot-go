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

type PublicObject[P ObjectPropertiesEmbedder] GenericPublicObject[P, PropertyWithHistory, AssociationEdge]

type GenericPublicObject[
	P ObjectPropertiesEmbedder,
	PWH PropertyWithHistory | Results[PropertyWithHistory] | PaginatedResults[PropertyWithHistory],
	AE AssociationEdge | Results[AssociationEdge] | PaginatedResults[AssociationEdge],
] struct {
	// ID is the HubSpot object ID.
	ID string `json:"id,omitempty"`

	// Properties is the object's properties.
	Properties P `json:"properties,omitempty"`

	// PropertiesWithHistory is the object's properties with history.
	//
	// The map key is the property name.
	//
	// Example:
	// 	{
	// 		"firstname": [
	// 			{
	// 				"value": "John",
	// 				"timestamp": "2024-02-25T15:52:36.931Z",
	// 				"sourceType": "CRM_UI",
	// 				"sourceId": "userId:51446676",
	// 				"updatedByUserId": 51446676
	// 			}
	// 		]
	// 	}
	PropertiesWithHistory map[string][]PWH `json:"propertiesWithHistory,omitempty"`

	// Associations is the object's associations.
	//
	// The map key is the association type name.
	//
	// Example:
	// 	{
	// 		"companies": [
	// 			{
	// 				"id": "10259467251",
	// 				"type": "contact_to_company"
	// 			}
	// 		]
	// 	}
	Associations map[string][]AE `json:"associations,omitempty"`

	CreatedAt *DateTime `json:"createdAt,omitempty"`

	UpdatedAt *DateTime `json:"updatedAt,omitempty"`

	Archived bool `json:"archived,omitempty"`

	ArchivedAt *DateTime `json:"archivedAt,omitempty"`
}

type CustomObject struct{}

func (CustomObject) embedProperties() {}
