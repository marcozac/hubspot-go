package hubspot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/marcozac/hubspot-go/limiter"

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
	// Use the default limiter to avoid hitting the rate limit during tests.
	client, err := NewTestClient(ts, WithContext(ctx), WithLimiter(limiter.NewDefault10()))
	require.NoError(t, err, "expected no error when creating client")

	t.Run("Properties", func(t *testing.T) {
		var group *PropertyGroup

		t.Run("Groups", func(t *testing.T) {
			t.Run("Create", func(t *testing.T) {
				var err error
				group, err = client.Properties.Contact.Groups.Create(ctx, &PropertyGroup{
					Name:  "test_group",
					Label: "Test Group",
				})
				assert.NoError(t, err, "expected no error when creating property group")
			})

			// Do not run the next tests if the group was not created.
			require.NotNil(t, group, "expected the group to be created")

			t.Run("Update", func(t *testing.T) {
				group.Label = "Test Group 2"
				_, err := client.Properties.Contact.Groups.Update(ctx, group)
				assert.NoError(t, err, "expected no error when updating property group")
				require.NotNil(t, group, "expected property group to be returned")
				assert.Equal(t, "Test Group 2", group.Label, "expected property group label to match")
			})

			t.Run("Read", func(t *testing.T) {
				v, err := client.Properties.Contact.Groups.Read(ctx, group.Name)
				assert.NoError(t, err, "expected no error when reading property group")
				require.NotNil(t, v, "expected property group to be returned")
				assert.Equal(t, group.Name, v.Name, "expected property group name to match")
			})

			t.Run("List", func(t *testing.T) {
				gs, err := client.Properties.Contact.Groups.List(ctx)
				assert.NoError(t, err, "expected no error when listing property groups")
				assert.NotEmpty(t, gs, "expected property groups to be returned")
				var found bool
				for _, g := range gs {
					if g.Name == group.Name {
						found = true
						break
					}
				}
				assert.True(t, found, "expected property group to be found")
			})
		})

		// Do not run the next tests if the group was not created.
		require.NotNil(t, group, "expected the group to be created")

		var prop *Property
		t.Run("Create", func(t *testing.T) {
			var err error
			prop, err = client.Properties.Contact.Create(ctx, &Property{
				Name:      "test_property",
				Label:     "Test Property",
				GroupName: group.Name,
				Type:      PropertyTypeString,
				FieldType: PropertyFieldTypeText,
			})
			assert.NoError(t, err, "expected no error when creating contact property")
		})

		// Do not run the next tests if the property was not found.
		require.NotNil(t, prop, "expected the property to be created")

		t.Run("Update", func(t *testing.T) {
			prop.Label = "Test Property 2"
			p, err := client.Properties.Contact.Update(ctx, prop)
			assert.NoError(t, err, "expected no error when updating contact property")
			require.NotNil(t, p, "expected property to be returned")
			assert.Equal(t, "Test Property 2", p.Label, "expected property label to match")
		})

		t.Run("Read", func(t *testing.T) {
			p, err := client.Properties.Contact.Read(ctx, prop.Name)
			assert.NoError(t, err, "expected no error when reading contact property")
			require.NotNil(t, p, "expected property to be returned")
			assert.Equal(t, prop.Name, p.Name, "expected property name to match")
		})

		t.Run("List", func(t *testing.T) {
			ps, err := client.Properties.Contact.List(ctx)
			assert.NoError(t, err, "expected no error when listing contact properties")
			require.NotNil(t, ps, "expected properties to be returned")
			var found bool
			for _, p := range ps {
				if p.Name == prop.Name {
					found = true
					break
				}
			}
			assert.True(t, found, "expected property to be found")
		})

		t.Run("Batch", func(t *testing.T) {
			var props []Property
			t.Run("Create", func(t *testing.T) {
				out, err := client.Properties.Contact.Batch.Create(ctx, &PropertiesBatchCreateInput{
					Inputs: []Property{
						{
							Name:      "test_batch_property_1",
							Label:     "Test Batch Property 1",
							GroupName: group.Name,
							Type:      PropertyTypeString,
							FieldType: PropertyFieldTypeText,
						},
						{
							Name:      "test_batch_property_2",
							Label:     "Test Batch Property 2",
							GroupName: group.Name,
							Type:      PropertyTypeString,
							FieldType: PropertyFieldTypeText,
						},
					},
				})
				assert.NoError(t, err, "expected no error when creating contact properties in batch")
				require.NotEmpty(t, out.Results, "expected results to be returned")
				props = out.Results.Results
			})

			t.Run("Read", func(t *testing.T) {
				inputs := make([]PropertiesBatchNameInput, 0, len(props))
				for _, p := range props {
					inputs = append(inputs, PropertiesBatchNameInput{
						Name: p.Name,
					})
				}
				out, err := client.Properties.Contact.Batch.Read(ctx, &PropertiesBatchReadInput{
					Inputs: inputs,
				})
				assert.NoError(t, err, "expected no error when reading contact properties in batch")
				require.NotEmpty(t, out.Results, "expected results to be returned")
				assert.Equal(t, 2, len(out.Results.Results), "expected 2 results to be returned")
			})

			t.Run("Archive", func(t *testing.T) {
				inputs := make([]PropertiesBatchNameInput, 0, len(props))
				for _, p := range props {
					inputs = append(inputs, PropertiesBatchNameInput{
						Name: p.Name,
					})
				}
				err := client.Properties.Contact.Batch.Archive(ctx, &PropertiesBatchArchiveInput{
					Inputs: inputs,
				})
				assert.NoError(t, err, "expected no error when archiving contact properties in batch")
			})
		})

		// Archive the test property
		t.Run("Archive", func(t *testing.T) {
			err := client.Properties.Contact.Archive(ctx, prop.Name)
			assert.NoError(t, err, "expected no error when archiving contact property")
		})

		// Archive the test group
		t.Run("ArchiveGroup", func(t *testing.T) {
			err := client.Properties.Contact.Groups.Archive(ctx, group.Name)
			assert.NoError(t, err, "expected no error when archiving property group")
		})
	})

	t.Run("Contacts", func(t *testing.T) {
		var createdContact *ObjectMutation[ContactPropertiesTest]
		t.Run("Create", func(t *testing.T) {
			c, err := client.Contacts.Create(ctx, &ContactPropertiesTest{
				ContactDefaultProperties: ContactDefaultProperties{
					Email:     "foo@example.com",
					Firstname: "Foo",
					Lastname:  "Bar",
				},
			})
			assert.NoError(t, err, "expected no error when creating contact")
			require.NotNil(t, c, "expected contact to be returned")
			assert.Equal(t, "Foo", c.Properties.Firstname, "expected contact firstname to match")
			createdContact = c
		})

		// Do not run the next tests if the contact was not created.
		require.NotNil(t, createdContact, "expected the contact to be created")

		t.Run("Update", func(t *testing.T) {
			c, err := client.Contacts.Update(ctx, createdContact.ID, &ContactPropertiesTest{
				ContactDefaultProperties: ContactDefaultProperties{
					Firstname: "John2",
					Lastname:  "Doe2",
				},
			})
			assert.NoError(t, err, "expected no error when updating contact")
			require.NotNil(t, c, "expected contact to be returned")
			assert.Equal(t, "John2", c.Properties.Firstname, "expected contact firstname to match")
		})

		t.Run("Read", func(t *testing.T) {
			c, err := client.Contacts.Read(ctx, createdContact.ID)
			assert.NoError(t, err, "expected no error when reading contact")
			assert.NotNil(t, c, "expected contact to be returned")
			assert.Equal(t, createdContact.ID, c.ID, "expected contact ID to match")
		})

		t.Run("List", func(t *testing.T) {
			cs, err := client.Contacts.List(ctx, WithProperties(
				"email",
				"firstname",
				"lastname",
				"my_custom_prop_from_ui",
			))
			assert.NoError(t, err, "expected no error when listing contacts")
			require.NotNil(t, cs, "expected contacts to be returned")
			// The first contact is the example contact, and the second is the
			// created contact.
			assert.GreaterOrEqual(t, len(cs.Results), 2, "expected at least 2 contacts to be returned")

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

		t.Run("Archive", func(t *testing.T) {
			err := client.Contacts.Archive(ctx, createdContact.ID)
			assert.NoError(t, err, "expected no error when deleting contact")
		})

		t.Run("Batch", func(t *testing.T) {
			var contacts []ObjectMutation[ContactPropertiesTest]
			t.Run("Create", func(t *testing.T) {
				out, err := client.Contacts.Batch.Create(ctx, &ObjectBatchCreateInput[ContactPropertiesTest]{
					Inputs: []ObjectMutationRequestBody[ContactPropertiesTest]{
						{
							Properties: &ContactPropertiesTest{
								ContactDefaultProperties: ContactDefaultProperties{
									Email: "foo_batch_1@example.com",
								},
							},
						},
						{
							Properties: &ContactPropertiesTest{
								ContactDefaultProperties: ContactDefaultProperties{
									Email: "foo_batch_2@example.com",
								},
							},
						},
					},
				})

				require.NoError(t, err, "expected no error when creating contacts in batch")
				contacts = out.Results.Results
			})

			// Do not run the next tests if the contacts were not created.
			require.NotEmpty(t, contacts, "expected contacts to be created")

			t.Run("Update", func(t *testing.T) {
				inputs := make([]ObjectBatchUpdateInputStruct[ContactPropertiesTest], 0, len(contacts))
				for _, c := range contacts {
					inputs = append(inputs, ObjectBatchUpdateInputStruct[ContactPropertiesTest]{
						ID: c.ID,
						Properties: ContactPropertiesTest{
							ContactDefaultProperties: ContactDefaultProperties{
								Firstname: "John",
								Lastname:  "Doe",
							},
						},
					})
				}
				ucs, err := client.Contacts.Batch.Update(ctx, &ObjectBatchUpdateInput[ContactPropertiesTest]{
					Inputs: inputs,
				})
				require.NoError(t, err, "expected no error when updating contacts in batch")
				require.NotNil(t, ucs.Results, "expected results to be returned")
				assert.Equal(t, 2, len(ucs.Results.Results), "expected 2 results to be returned")
				for _, c := range ucs.Results.Results {
					assert.Equal(t, "John", c.Properties.Firstname, "expected contact firstname to match")
					assert.Equal(t, "Doe", c.Properties.Lastname, "expected contact lastname to match")
				}
				contacts = ucs.Results.Results
			})

			// Do not run the next tests if the contacts were not created.
			require.NotEmpty(t, contacts, "expected contacts to be created")

			t.Run("Read", func(t *testing.T) {
				inputs := make([]ObjectBatchIDInput, 0, len(contacts))
				for _, c := range contacts {
					inputs = append(inputs, ObjectBatchIDInput{
						ID: c.ID,
					})
				}
				rcs, err := client.Contacts.Batch.Read(ctx, &ObjectBatchReadInput[ContactPropertiesTest]{
					Properties: []string{"email", "firstname", "lastname"},
					BatchInput: BatchInput[ObjectBatchIDInput]{
						Inputs: inputs,
					},
				})
				require.NoError(t, err, "expected no error when reading contacts in batch")
				require.NotNil(t, rcs.Results, "expected results to be returned")
				assert.Equal(t, 2, len(rcs.Results.Results), "expected 2 results to be returned")
				for _, c := range rcs.Results.Results {
					assert.Equal(t, "John", c.Properties.Firstname, "expected contact firstname to match")
					assert.Equal(t, "Doe", c.Properties.Lastname, "expected contact lastname to match")
				}
			})
			t.Run("Archive", func(t *testing.T) {
				inputs := make([]ObjectBatchIDInput, 0, len(contacts))
				for _, c := range contacts {
					inputs = append(inputs, ObjectBatchIDInput{
						ID: c.ID,
					})
				}
				err := client.Contacts.Batch.Archive(ctx, &ObjectBatchArchiveInput{
					Inputs: inputs,
				})
				assert.NoError(t, err, "expected no error when archiving contacts in batch")
			})
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
