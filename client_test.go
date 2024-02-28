package hubspot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/joho/godotenv/autoload"
)

var NewTestClient = NewTypedClient[
	ContactPropertiesTest,
	CompanyDefaultProperties,
	DealDefaultProperties,
	FeedbackSubmissionDefaultProperties,
	LineItemDefaultProperties,
	ProductDefaultProperties,
	QuoteDefaultProperties,
	DiscountDefaultProperties,
	FeeDefaultProperties,
	TaxDefaultProperties,
	TicketDefaultProperties,
	GoalDefaultProperties,
]

// TestClient tests the NewClient function and the returned client's methods.
//
// WARNING
// This test requires a valid access token to be present in the environment
// variable "TEST_CLIENT_HUBSPOT_ACCESS_TOKEN".
// It also requires an high permission level to access the HubSpot API and
// will perform real read and write requests to the API. It is recommended to
// generate the access token in a [HubSpot test account].
//
// The environment variable "TEST_CLIENT_HUBSPOT_ACCESS_TOKEN" may be provided
// in a .env file in the root of the project.
//
// [HubSpot test account]: https://developers.hubspot.com/docs/api/account-types#developer-test-accounts
func TestClient(t *testing.T) {
	// use a different key to avoid conflicts with auth tests
	const envTokenKey = "TEST_CLIENT_HUBSPOT_ACCESS_TOKEN"
	ts, err := NewEnvTokenSource(envTokenKey)
	require.NoError(t, err, "expected no error when creating token source")

	_, err = NewTestClient(nil)
	assert.ErrorIs(t, err, ErrTokenSourceRequired)

	ctx := context.Background()
	client, err := NewTestClient(ts, WithContext(ctx))
	require.NoError(t, err, "expected no error when creating client")

	t.Run("Properties", func(t *testing.T) {
		t.Run("List", func(t *testing.T) {
			ps, err := client.Properties.Contact.List(ctx,
				WithArchived(false), WithProperties("name"),
			)
			assert.NoError(t, err, "expected no error when listing contact properties")
			assert.NotEmpty(t, ps, "expected properties to be returned")
		})
	})

	t.Run("Contacts", func(t *testing.T) {
		t.Run("List", func(t *testing.T) {
			cs, err := client.Contacts.List(ctx,
				WithArchived(false),
				WithProperties(
					"email",
					"firstname",
					"lastname",
					"my_custom_prop_from_ui",
				),
			)
			assert.NoError(t, err, "expected no error when listing contacts")
			require.NotNil(t, cs, "expected contacts to be returned")

			for _, r := range cs.Results {
				// HubSpot uses the email address as the unique identifier for
				// contacts, so we can use it to find the example contact.
				if r.Properties.Email == "johndoe@example.com" {
					assert.Equal(t, "John", r.Properties.Firstname)
					assert.Equal(t, "Doe", r.Properties.Lastname)
					assert.Equal(t, "foo", r.Properties.MyCustomPropertyFromUI)
				}
			}
		})
	})
}

type ContactPropertiesTest struct {
	// Embed the default properties to ensure that the custom properties are
	// added to the default ones.
	ContactDefaultProperties

	// MyCustomPropertyFromUI is a custom property added from the HubSpot UI
	// for testing purposes.
	MyCustomPropertyFromUI string `json:"my_custom_prop_from_ui,omitempty"`
}
