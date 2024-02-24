package hubspot

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	// use a different key to avoid conflicts with auth tests
	const envTokenKey = "TEST_CLIENT_HUBSPOT_ACCESS_TOKEN"
	// TODO
	// Replace with a valid access token when real tests are implemented
	t.Setenv(envTokenKey, uuid.NewString())
	ts, err := NewEnvTokenSource(envTokenKey)
	require.NoError(t, err, "expected no error when creating token source")

	_, err = NewClient(nil)
	assert.Error(t, err, "expected error when token source is not provided")

	ctx := context.Background()
	client, err := NewClient(ts, WithContext(ctx))
	require.NoError(t, err, "expected no error when creating client")

	_ = client // TODO: test client
}
