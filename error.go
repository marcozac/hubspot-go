package hubspot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var (
	// ErrTokenInvalid is returned when the access token is invalid.
	ErrTokenInvalid = errors.New("token is invalid")

	// ErrEnvKeyRequired is returned when an environment variable key is
	// required but not provided.
	ErrEnvKeyRequired = errors.New("environment variable key is required")

	// ErrEnvVarNotSet is returned when an environment variable is not set.
	ErrEnvVarNotSet = errors.New("environment variable not set")

	// ErrEncryptionKeyTooShort is returned when the encryption key is too
	// short.
	ErrEncryptionKeyTooShort = errors.New("encryption key too short")

	// ErrEncryptedAccessTokenTooShort is returned when the encrypted
	// access token is too short.
	ErrEncryptedAccessTokenTooShort = errors.New("encrypted access token too short")

	// ErrTokenSourceRequired is returned when a token source is required but
	// not provided.
	ErrTokenSourceRequired = errors.New("token source is required")

	// ErrNilParam is an error returned when a required parameter is nil.
	ErrNilParam = errors.New("parameter cannot be nil")
)

var _ error = (*HubSpotError)(nil)

// HubSpotError represents an error returned by the HubSpot API.
//
// It may include a list of sub-errors that can be unwrapped to get more
// detailed information about the error.
type HubSpotError struct {
	SubCategory   string              `json:"subCategory,omitempty"`
	Context       HubSpotErrorContext `json:"context,omitempty"`
	CorrelationID uuid.UUID           `json:"correlationId,omitempty"`
	Links         map[string]string   `json:"links,omitempty"`
	Message       string              `json:"message,omitempty"`
	Category      string              `json:"category,omitempty"`
	Errors        HubSpotSubErrors    `json:"errors,omitempty"`
}

func (e *HubSpotError) Error() string {
	return e.Message
}

func (e *HubSpotError) Unwrap() error {
	if len(e.Errors) > 0 {
		return e.Errors
	}
	return nil
}

// HubSpotSubErrors represents a list of sub-errors returned by the HubSpot API.
// The error message is a concatenation of all the sub-errors' messages.
//
// It may be unwrapped to get the list of sub-errors excluding the first one.
type HubSpotSubErrors []*HubSpotSubError

func (e HubSpotSubErrors) Error() string {
	var message strings.Builder
	for _, err := range e {
		message.WriteString(err.Error())
		message.WriteString(";")
	}
	return message.String()
}

func (e HubSpotSubErrors) Unwrap() error {
	if len(e) > 1 {
		return e[1:]
	}
	return nil
}

// HubSpotSubError represents a sub-error returned by the HubSpot API.
type HubSpotSubError struct {
	SubCategory string              `json:"subCategory,omitempty"`
	Code        string              `json:"code,omitempty"`
	In          string              `json:"in,omitempty"`
	Context     HubSpotErrorContext `json:"context,omitempty"`
	Message     string              `json:"message,omitempty"`
}

func (e *HubSpotSubError) Error() string {
	return e.Message
}

type HubSpotErrorContext struct {
	MissingScopes       []string `json:"missingScopes,omitempty"`
	InvalidPropertyName []string `json:"invalidPropertyName,omitempty"`
}

// HubSpotResponseError checks the response for errors and returns an error if
// any are found.
//
// It checks the response status code and, if it is less than 400, it decodes
// the response body into a [HubSpotError].
//
// When reporting an error, it always closes the response body.
func HubSpotResponseError(resp *http.Response) error {
	if resp.StatusCode < 400 {
		return nil
	}
	defer resp.Body.Close()
	hsErr := new(HubSpotError)
	if err := json.NewDecoder(resp.Body).Decode(hsErr); err != nil {
		return fmt.Errorf("decode hubspot error: %w", err)
	}
	return hsErr
}

func IsHubSpotError(err error) bool {
	_, ok := err.(*HubSpotError)
	return ok
}

func IsHubSpotSubErrors(err error) bool {
	_, ok := err.(HubSpotSubErrors)
	if ok {
		return true
	}
	ep := &HubSpotSubErrors{}
	return errors.As(err, &ep)
}

func IsHubSpotSubError(err error) bool {
	_, ok := err.(*HubSpotSubError)
	return ok
}

func IsAnyHubSpotError(err error) bool {
	switch err.(type) {
	case *HubSpotError, *HubSpotSubErrors, HubSpotSubErrors, *HubSpotSubError:
		return true
	}
	return false
}
