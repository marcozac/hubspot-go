package limiter

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

func TestLimiter(t *testing.T) {
	t.Run("New10", func(t *testing.T) {
		limiter := New10(100)
		require.NotNil(t, limiter)
	})

	t.Run("NewDefault10", func(t *testing.T) {
		limiter := NewDefault10()
		require.NotNil(t, limiter)
	})

	t.Run("NewInf", func(t *testing.T) {
		limiter := NewInf()
		require.NotNil(t, limiter)
	})

	t.Run("WrapHTTPClient", func(t *testing.T) {
		originalClient := &http.Client{}
		limiter := NewDefault10()
		client := WrapHTTPClient(originalClient, limiter)
		require.NotNil(t, client)
		require.IsType(t, &Transport{}, client.Transport)
	})

	t.Run("RoundTripPolicyWait", func(t *testing.T) {
		client := WrapHTTPClient(nil, nil, WithPolicy(PolicyWait)) // to cover the rest of the code
		transport := client.Transport.(*Transport)
		transport.lim.SetBurst(1)
		transport.lim.SetLimit(rate.Every(time.Second))

		for i := 0; i < 2; i++ {
			req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
			require.NoError(t, err)
			_, err = client.Do(req)
			assert.NoError(t, err)
		}
	})

	t.Run("RoundTripPolicyReject", func(t *testing.T) {
		client := WrapHTTPClient(http.DefaultClient, New10(1), WithPolicy(PolicyReject))
		for i := 0; i < 2; i++ {
			req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
			require.NoError(t, err)
			if _, err := client.Do(req); i == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		}
	})
}
