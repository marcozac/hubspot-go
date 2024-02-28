package hubspot

import (
	"crypto/aes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/marcozac/hubspot-go/internal/testutil"
)

func TestEnvTokenSource(t *testing.T) {
	_, err := NewEnvTokenSource("")
	require.Error(t, err, "expected error when environment variable key is not provided")

	const ev = "TEST_HUBSPOT_ACCESS_TOKEN"
	_, err = NewEnvTokenSource(ev)
	require.Error(t, err, "expected error when environment variable is not set")

	at := uuid.NewString()
	t.Setenv(ev, at)

	ets, err := NewEnvTokenSource(ev)
	require.NoError(t, err, "expected no error when environment variable is set")
	token, err := ets.Token()
	require.NoError(t, err, "expected no error when token is valid")
	assert.Equal(t, at, token.AccessToken, "expected access token to match environment variable")

	ets.t.AccessToken = ""
	_, err = ets.Token()
	assert.Error(t, err, "expected error when token is invalid")
}

func TestEnvTokenSourceEncrypted(t *testing.T) {
	_, err := NewEnvTokenSourceEncrypted("", nil)
	require.Error(t, err, "expected error when environment variable key is not provided")

	const ev = "TEST_HUBSPOT_ACCESS_TOKEN"
	_, err = NewEnvTokenSourceEncrypted(ev, make([]byte, aes.BlockSize-1))
	require.Error(t, err, "expected error when encryption key is too short")

	encryptionKey := make([]byte, aes.BlockSize) // 16 bytes for AES-128
	_, err = rand.Read(encryptionKey)
	require.NoError(t, err, "expected no error when generating encryption key")

	_, err = NewEnvTokenSourceEncrypted(ev, encryptionKey)
	require.Error(t, err, "expected error when environment variable is not set")

	at := uuid.NewString()
	t.Setenv(ev, at)

	ets, err := NewEnvTokenSourceEncrypted(ev, encryptionKey)
	require.NoError(t, err, "expected no error when environment variable is set and encryption key is valid")
	token, err := ets.Token()
	require.NoError(t, err, "expected no error when token is valid")
	assert.Equal(t, at, token.AccessToken, "expected access token to match environment variable")

	ets.eat = make([]byte, aes.BlockSize-1)
	_, err = ets.Token()
	assert.Error(t, err, "expected error when encrypted access token is too short")

	// WARNING
	// This test is flaky because it relies on the global rand.Reader.
	// NEVER run this test in parallel with other tests that use rand.Reader.
	rr := rand.Reader
	defer func() {
		rand.Reader = rr
	}()

	rand.Reader = testutil.ReaderEOF{}
	err = ets.Refresh()
	assert.ErrorIs(t, err, io.EOF, "expected error from reader")
}
