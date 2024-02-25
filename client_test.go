package hubspot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestClient tests the NewClient function and the returned client's methods.
//
// WARNING
// This test requires a valid access token to be present in the environment.
// It also requires an high permission level to access the HubSpot API and
// will perform real read and write requests to the API. It is recommended to
// generate the access token in a [HubSpot test account].
//
// [HubSpot test account]: https://developers.hubspot.com/docs/api/account-types#developer-test-accounts
func TestClient(t *testing.T) {
	// use a different key to avoid conflicts with auth tests
	const envTokenKey = "TEST_CLIENT_HUBSPOT_ACCESS_TOKEN"
	ts, err := NewEnvTokenSource(envTokenKey)
	require.NoError(t, err, "expected no error when creating token source")

	_, err = NewClient(nil)
	assert.Error(t, err, "expected error when token source is not provided")

	ctx := context.Background()
	client, err := NewClient(ts, WithContext(ctx))
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
}
