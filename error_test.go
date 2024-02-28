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
