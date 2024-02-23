package hubspot

type PublicObject[P any] struct {
	// ID is the HubSpot object ID.
	ID string `json:"id"`

	// Properties is the object's properties.
	Properties P `json:"properties"`
}

type SimplePublicObject = PublicObject[map[string]string]
