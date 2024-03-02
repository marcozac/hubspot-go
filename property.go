package hubspot

const (
	PropertyTypeString      PropertyType = "string"
	PropertyTypePhoneNumber PropertyType = "phone_number" // Alias for "string"
	PropertyTypeNumber      PropertyType = "number"
	PropertyTypeDate        PropertyType = "date"
	PropertyTypeDateTime    PropertyType = "datetime"
	PropertyTypeEnumeration PropertyType = "enumeration"
	PropertyTypeBool        PropertyType = "bool"

	PropertyFieldTypeText                PropertyFieldType = "text"
	PropertyFieldTypeTextarea            PropertyFieldType = "textarea"
	PropertyFieldTypeDate                PropertyFieldType = "date"
	PropertyFieldTypeFile                PropertyFieldType = "file"
	PropertyFieldTypeNumber              PropertyFieldType = "number"
	PropertyFieldTypeSelect              PropertyFieldType = "select"
	PropertyFieldTypeRadio               PropertyFieldType = "radio"
	PropertyFieldTypeCheckbox            PropertyFieldType = "checkbox"
	PropertyFieldTypeBooleanCheckbox     PropertyFieldType = "booleancheckbox"
	PropertyFieldTypeCalculationEquation PropertyFieldType = "calculation_equation"
	PropertyFieldTypeCalculationRollup   PropertyFieldType = "calculation_rollup"

	PropertyReferencedObjectTypeOwner    PropertyReferencedObjectType = "OWNER"
	PropertyReferencedObjectTypeContact  PropertyReferencedObjectType = "CONTACT"
	PropertyReferencedObjectTypeCompany  PropertyReferencedObjectType = "COMPANY"
	PropertyReferencedObjectTypeDeal     PropertyReferencedObjectType = "DEAL"
	PropertyReferencedObjectTypeSequence PropertyReferencedObjectType = "SEQUENCE"
)

type (
	// Property is a HubSpot property as returned by the properties API.
	//
	// Do not confuse it with [Object.Properties], which are the properties of a
	// HubSpot object (e.g. a contact, a company, a deal).
	Property struct {
		// Name is the internal name of the property as used in the API.
		Name string `json:"name"`

		// Label is the label of the property as displayed in the HubSpot UI.
		Label string `json:"label,omitempty"`

		// GroupName is the name of the group the property belongs to.
		GroupName string `json:"groupName,omitempty"`

		// Type is the type of the property, e.g. "string", "number", "enumeration".
		Type PropertyType `json:"type,omitempty"`

		// FieldType is the type of the field, e.g. "text", "select", "checkbox".
		FieldType PropertyFieldType `json:"fieldType,omitempty"`

		// ReferencedObjectType is the type of the object the property references.
		//
		// From HubSpot's documentation:
		//
		// Should be set to 'OWNER' when 'externalOptions' is true, which causes
		// the property to dynamically pull option values from the current
		// HubSpot users.
		ReferencedObjectType PropertyReferencedObjectType `json:"referencedObjectType,omitempty"`

		// Description is a description of the property.
		Description string `json:"description,omitempty"`

		// DisplayOrder is the order in which the property is displayed.
		DisplayOrder int32 `json:"displayOrder,omitempty"`

		// Options is a list of options for enumeration properties.
		Options []PropertyOption `json:"options,omitempty"`

		// ExternalOptions is true if the property has external options.
		ExternalOptions bool `json:"externalOptions,omitempty"`

		// Calculated is true if the property is calculated.
		Calculated bool `json:"calculated,omitempty"`

		// CalculationFormula is the formula used to calculate the property.
		//
		// From HubSpot's documentation:
		//
		// Applicable only for 'enumeration' type properties. Should be set to
		// true in conjunction with a 'referencedObjectType' of 'OWNER'.
		// Otherwise false.
		CalculationFormula string `json:"calculationFormula,omitempty"`

		// HasUniqueValue is true if the property has unique values.
		HasUniqueValue bool `json:"hasUniqueValue,omitempty"`

		// Hidden is true if the property is hidden.
		Hidden bool `json:"hidden,omitempty"`

		// ShowCurrencySymbol is true if the currency symbol is shown.
		ShowCurrencySymbol bool `json:"showCurrencySymbol,omitempty"`

		// FormField is true if the property is a form field.
		FormField bool `json:"formField,omitempty"`

		// Archived is true if the property is archived.
		Archived bool `json:"archived,omitempty"`

		// ModificationMetadata is the modification metadata of the property.
		ModificationMetadata PropertyModificationMetadata `json:"modificationMetadata,omitempty"`

		// CreatedUserId is the ID of the user who created the property.
		CreatedUserId string `json:"createdUserId,omitempty"`

		// UpdatedAt is the time the property was last updated.
		UpdatedAt *DateTime `json:"updatedAt,omitempty"`

		// UpdatedUserId is the ID of the user who last updated the property.
		UpdatedUserId string `json:"updatedUserId,omitempty"`

		// CreatedAt is the time the property was created.
		CreatedAt *DateTime `json:"createdAt,omitempty"`
	}

	// PropertyType is the type of a property.
	PropertyType string

	// PropertyFieldType is the type of the field of a property.
	PropertyFieldType string

	// PropertyReferencedObjectType is the type of the object a property references.
	// It should be set to 'OWNER' when 'externalOptions' is true.
	PropertyReferencedObjectType string

	// PropertyOption is an option for an enumeration property.
	PropertyOption struct {
		// Label is the label of the option.
		Label string `json:"label"`

		// Value is the value of the option.
		Value string `json:"value"`

		// Hidden is true if the option is hidden.
		Hidden bool `json:"hidden"`

		// Description is a description of the option.
		Description string `json:"description,omitempty"`

		// DisplayOrder is the order in which the option is displayed.
		DisplayOrder int `json:"displayOrder,omitempty"`
	}

	// PropertyModificationMetadata is the modification metadata of a property.
	PropertyModificationMetadata struct {
		// Archivable is true if the property is archivable.
		Archivable bool `json:"archivable"`

		// ReadOnlyDefinition is true if the property definition is read-only.
		ReadOnlyDefinition bool `json:"readOnlyDefinition"`

		// ReadOnlyValue is true if the property value is read-only.
		ReadOnlyValue bool `json:"readOnlyValue"`
	}

	PropertyGroup struct {
		// Name is the internal name of the group as used in the API.
		Name string `json:"name"`

		// Label is the label of the group as displayed in the HubSpot UI.
		Label string `json:"label"`

		// DisplayOrder is the order in which the group is displayed.
		DisplayOrder int `json:"displayOrder,omitempty"`

		// Archived is true if the group is archived.
		Archived bool `json:"archived,omitempty"`
	}

	PropertyWithHistory struct {
		SourceID        string    `json:"sourceId,omitempty"`
		SourceLabel     string    `json:"sourceLabel,omitempty"`
		SourceType      string    `json:"sourceType,omitempty"`
		UpdatedByUserID int64     `json:"updatedByUserId,omitempty"`
		Value           string    `json:"value,omitempty"`
		Timestamp       *DateTime `json:"timestamp,omitempty"`
	}
)
