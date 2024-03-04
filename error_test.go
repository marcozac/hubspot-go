package hubspot

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHubSpotError(t *testing.T) {
	err := &HubSpotError{
		Message: "error message",
		Errors: HubSpotSubErrors{
			{
				Message: "sub error message 1",
			},
			{
				Message: "sub error message 2",
			},
		},
	}
	assert.Error(t, err)
	assert.ErrorAs(t, err, &HubSpotSubErrors{})

	subErr := errors.Unwrap(err)
	assert.NotNil(t, subErr)

	subErr2 := errors.Unwrap(subErr)
	assert.NotNil(t, subErr2)

	assert.Nil(t, errors.Unwrap(subErr2))
}

func TestIsHubSpotError(t *testing.T) {
	assert.True(t, IsHubSpotError(&HubSpotError{}))
	assert.False(t, IsHubSpotError(errors.New("not a HubSpotError")))
	assert.False(t, IsHubSpotError(nil))

	assert.True(t, IsHubSpotSubErrors(&HubSpotSubErrors{}))
	assert.True(t, IsHubSpotSubErrors(HubSpotSubErrors{}))
	assert.False(t, IsHubSpotSubErrors(errors.New("not a HubSpotSubErrors")))
	assert.False(t, IsHubSpotSubErrors(nil))

	assert.True(t, IsHubSpotSubError(&HubSpotSubError{}))
	assert.False(t, IsHubSpotSubError(errors.New("not a HubSpotSubError")))
	assert.False(t, IsHubSpotSubError(nil))

	assert.True(t, IsAnyHubSpotError(&HubSpotError{}))
	assert.True(t, IsAnyHubSpotError(&HubSpotSubErrors{}))
	assert.True(t, IsAnyHubSpotError(HubSpotSubErrors{}))
	assert.True(t, IsAnyHubSpotError(&HubSpotSubError{}))
	assert.False(t, IsAnyHubSpotError(errors.New("not a HubSpotError")))
	assert.False(t, IsAnyHubSpotError(nil))
}
