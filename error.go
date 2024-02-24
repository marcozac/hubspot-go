package hubspot

import "errors"

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
