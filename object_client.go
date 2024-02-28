package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/marcozac/hubspot-go/util"
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
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	results := new(ObjectListResults[PE])
	if err := json.NewDecoder(resp.Body).Decode(results); err != nil {
		return nil, err
	}
	return results, nil
}

// Read returns the object with the given ID.
//
// Object ID is usually a number. If it is declared as a [Int], [Int64], etc.
// it is possible to use the String method to convert it to a string.
//
// Allowed options:
//   - WithArchived: include only archived objects in the response
//   - WithProperties: include only the specified properties in the response
//   - WithPropertiesWithHistory: include the history of the specified
//     properties in the response
//   - WithAssociations: include only the specified associations in the response
//
// Any other option is ignored.
func (oc *ObjectClient[PE]) Read(ctx context.Context, id string, opts ...RequestOption) (*ObjectRead[PE], error) {
	cfg := applyRequestOptions(nil, opts...)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, oc.endpoint+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	applyQueryOptions(cfg, req.URL,
		applyArchivedQuery,
		applyPropertiesQuery,
		applyPropertiesWithHistoryQuery,
		applyAssociationsQuery,
		applyIDPropertyQuery,
	)
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := new(ObjectRead[PE])
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates a new object with the given properties and, optionally, the
// given associations.
func (oc *ObjectClient[PE]) Create(ctx context.Context, properties *PE, associations ...AssociationForCreate) (*ObjectMutation[PE], error) {
	rb := &ObjectMutationRequestBody[PE]{
		Properties:   properties,
		Associations: associations,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(rb); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, oc.endpoint, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := new(ObjectMutation[PE])
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Update updates the properties of the object with the given ID.
//
// According to the HubSpot API, read only and non-writable properties are
// ignored. To clear a property, set its value to an empty string.
func (oc *ObjectClient[PE]) Update(ctx context.Context, id string, properties *PE) (*ObjectMutation[PE], error) {
	rb := &ObjectMutationRequestBody[PE]{
		Properties: properties,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(rb); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, oc.endpoint+"/"+id, buf)
	if err != nil {
		return nil, err
	}
	util.SetJSONHeader(req) // Set the Content-Type header to application/json.
	resp, err := oc.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := new(ObjectMutation[PE])
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Archive archives the object with the given ID.
func (oc *ObjectClient[PE]) Archive(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, oc.endpoint+"/"+id, nil)
	if err != nil {
		return err
	}
	resp, err := oc.hc.Do(req)
	if err != nil {
		return err
	}
	// Check for errors in the response.
	if err := HubSpotResponseError(resp); err != nil {
		return err
	}
	return nil
}
