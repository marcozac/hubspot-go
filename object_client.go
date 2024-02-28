package hubspot

import (
	"context"
	"encoding/json"
	"net/http"
)

// NewObjectClient returns a new object client for the given object type.
//
// The object type may be provided as a type param or inferred from the
// optional pe parameter. In the latter case, the given value is used only
// for type inference and has no effect on the client.
//
// Example:
//
//	type ContactPropertiesTest struct {
//		// Embed the default properties to ensure that the custom properties are
//		// added to the default ones.
//		ContactDefaultProperties
//		// MyCustomPropertyFromUI is a custom property added from the HubSpot UI
//		// for testing purposes.
//		MyCustomPropertyFromUI string `json:"my_custom_prop_from_ui,omitempty"`
//	}
//
//	func main() {
//		NewObjectClient[ContactPropertiesTest](endpoint.Contacts, http.DefaultClient)
//		// or
//		NewObjectClient(endpoint.Contacts, http.DefaultClient, ContactPropertiesTest{})
//	}
func NewObjectClient[PE ObjectPropertiesEmbedder](endpoint string, httpClient *http.Client, pe ...PE) *ObjectClient[PE] {
	return &ObjectClient[PE]{
		endpoint: endpoint,
		hc:       httpClient,
	}
}

type ObjectClient[PE ObjectPropertiesEmbedder] struct {
	endpoint string
	hc       *http.Client
}

// List returns a paginated list of objects.
//
// Allowed options:
//   - WithLimit: the number of objects to return per page
//   - WithAfter: the cursor token to use to get the next page of results
//   - WithArchived: include only archived objects in the response
//   - WithProperties: include only the specified properties in the response
//   - WithPropertiesWithHistory: include the history of the specified
//     properties in the response
//   - WithAssociations: include only the specified associations in the response
//
// Any other option is ignored.
//
// NOTE
// The HubSpot's default limit is 10 objects per page.
func (oc *ObjectClient[PE]) List(ctx context.Context, opts ...RequestOption) (*ObjectListResults[PE], error) {
	cfg := applyRequestOptions(nil, opts...)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, oc.endpoint, nil)
	if err != nil {
		return nil, err
	}
	applyQueryOptions(cfg, req.URL,
		applyLimitQuery,
		applyAfterQuery,
		applyArchivedQuery,
		applyPropertiesQuery,
		applyPropertiesWithHistoryQuery,
		applyAssociationsQuery,
	)
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var results ObjectListResults[PE]
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}
	return &results, nil
}
