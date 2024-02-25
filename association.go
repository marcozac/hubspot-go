package hubspot

const (
	// Association categories.

	AssociationCategoryHubSpotDefined    = "HUBSPOT_DEFINED"
	AssociationCategoryUserDefined       = "USER_DEFINED"
	AssociationCategotyIntegratorDefined = "INTEGRATOR_DEFINED"
)

const (
	// HubSpot-defined association type IDs.
	//
	// https://developers.hubspot.com/docs/api/crm/associations

	AssociationTypeIDContactToContact        = 449
	AssociationTypeIDContactToCompany        = 279
	AssociationTypeIDContactToCompanyPrimary = 1
	AssociationTypeIDContactToDeal           = 4
	AssociationTypeIDContactToTicket         = 15
	AssociationTypeIDContactToCall           = 193
	AssociationTypeIDContactToEmail          = 197
	AssociationTypeIDContactToMeeting        = 199
	AssociationTypeIDContactToNote           = 201
	AssociationTypeIDContactToTask           = 203
	AssociationTypeIDContactToCommunication  = 82
	AssociationTypeIDContactToPostalMail     = 454

	AssociationTypeIDCompanyToCompany        = 450
	AssociationTypeIDChildToParentCompany    = 14
	AssociationTypeIDParentToChildCompany    = 13
	AssociationTypeIDCompanyToContact        = 280
	AssociationTypeIDCompanyToContactPrimary = 2
	AssociationTypeIDCompanyToDeal           = 342
	AssociationTypeIDCompanyToDealPrimary    = 6
	AssociationTypeIDCompanyToTicket         = 340
	AssociationTypeIDCompanyToTicketPrimary  = 25
	AssociationTypeIDCompanyToCall           = 181
	AssociationTypeIDCompanyToEmail          = 185
	AssociationTypeIDCompanyToMeeting        = 187
	AssociationTypeIDCompanyToNote           = 189
	AssociationTypeIDCompanyToTask           = 191
	AssociationTypeIDCompanyToCommunication  = 88
	AssociationTypeIDCompanyToPostalMail     = 460

	AssociationTypeIDDealToDeal           = 451
	AssociationTypeIDDealToContact        = 3
	AssociationTypeIDDealToCompany        = 341
	AssociationTypeIDDealToCompanyPrimary = 5
	AssociationTypeIDDealToTicket         = 27
	AssociationTypeIDDealToCall           = 205
	AssociationTypeIDDealToEmail          = 209
	AssociationTypeIDDealToMeeting        = 211
	AssociationTypeIDDealToNote           = 213
	AssociationTypeIDDealToTask           = 215
	AssociationTypeIDDealToCommunication  = 86
	AssociationTypeIDDealToPostalMail     = 458

	AssociationTypeIDTicketToTicket         = 452
	AssociationTypeIDTicketToContact        = 16
	AssociationTypeIDTicketToCompany        = 339
	AssociationTypeIDTicketToCompanyPrimary = 26
	AssociationTypeIDTicketToDeal           = 28
	AssociationTypeIDTicketToCall           = 219
	AssociationTypeIDTicketToEmail          = 223
	AssociationTypeIDTicketToMeeting        = 225
	AssociationTypeIDTicketToNote           = 227
	AssociationTypeIDTicketToTask           = 229
	AssociationTypeIDTicketToCommunication  = 84
	AssociationTypeIDTicketToPostalMail     = 456

	AssociationTypeIDCallToContact = 194
	AssociationTypeIDCallToCompany = 182
	AssociationTypeIDCallToDeal    = 206
	AssociationTypeIDCallToTicket  = 220

	AssociationTypeIDEmailToContact = 198
	AssociationTypeIDEmailToCompany = 186
	AssociationTypeIDEmailToDeal    = 210
	AssociationTypeIDEmailToTicket  = 224

	AssociationTypeIDMeetingToContact = 200
	AssociationTypeIDMeetingToCompany = 188
	AssociationTypeIDMeetingToDeal    = 212
	AssociationTypeIDMeetingToTicket  = 226

	AssociationTypeIDNoteToContact = 202
	AssociationTypeIDNoteToCompany = 190
	AssociationTypeIDNoteToDeal    = 214
	AssociationTypeIDNoteToTicket  = 228

	AssociationTypeIDPostalMailToContact = 453
	AssociationTypeIDPostalMailToCompany = 459
	AssociationTypeIDPostalMailToDeal    = 457
	AssociationTypeIDPostalMailToTicket  = 455

	AssociationTypeIDQuoteToContact       = 69
	AssociationTypeIDQuoteToCompany       = 71
	AssociationTypeIDQuoteToDeal          = 64
	AssociationTypeIDQuoteToLineItem      = 67
	AssociationTypeIDQuoteToQuoteTemplate = 286
	AssociationTypeIDQuoteToDiscount      = 362
	AssociationTypeIDQuoteToFee           = 364
	AssociationTypeIDQuoteToTax           = 366
	AssociationTypeIDContactSigner        = 702

	AssociationTypeIDTaskToContact = 204
	AssociationTypeIDTaskToCompany = 192
	AssociationTypeIDTaskToDeal    = 216
	AssociationTypeIDTaskToTicket  = 230

	AssociationTypeIDCommunicationToContact = 81
	AssociationTypeIDCommunicationToCompany = 87
	AssociationTypeIDCommunicationToDeal    = 85
	AssociationTypeIDCommunicationToTicket  = 83
)

type (
	// Association is the association between two objects.
	//
	// Due to the nature of the HubSpot API, which uses different
	// representations of the same data in different methods, this
	// type is used in requests and responses with different behaviors.
	//
	// For example, creating a batch of default associations, the request
	// needs to include the "from" and "to" objects and not the association
	// type. Instead, creating a contact, the request needs to include
	// (under the "associations" key) only the association type and the
	// "to" edge.
	//
	// Always refer to the method-specific documentation for details,
	// required fields and examples.
	Association struct {
		// Types is the list of association types.
		Types []AssociationType `json:"types,omitempty"`

		// To is the object to which the association is being made.
		To AssociationEdge `json:"to,omitempty"`

		// From is the object from which the association is being made.
		From AssociationEdge `json:"from,omitempty"`
	}

	// AssociationType is the type of association.
	//
	// This type is used in requests and responses with different representations.
	// See the method-specific documentation for details.
	AssociationType struct {
		// The ID of the association type.
		ID int `json:"typeId,omitempty"`

		// The ID of the association type. It is the same as typeId,
		// used in some requests and responses.
		AssociationTypeId int `json:"associationTypeId,omitempty"`

		// The label of the association type.
		Label string `json:"label,omitempty"`

		// The label of the association type when viewed from the other
		// side of the association.
		InverseLabel string `json:"inverseLabel,omitempty"`

		// The name of the association type.
		Name string `json:"name,omitempty"`

		// The category of the association type.
		Category string `json:"category,omitempty"`
	}

	// AssociationEdge represents the relationship between two objects.
	// It is used in requests and responses with different representations,
	// see the method-specific documentation for details.
	AssociationEdge struct {
		// The ID of the object to which the association is being made.
		ID int `json:"id"`
	}
)
