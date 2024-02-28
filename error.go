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
)

var _ error = (*HubSpotError)(nil)

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

type HubSpotSubErrors []HubSpotSubError

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
