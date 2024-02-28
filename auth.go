package hubspot

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"sync"

	"golang.org/x/oauth2"

	"github.com/marcozac/hubspot-go/endpoint"
)

func NewEnvTokenSource(key string) (*EnvTokenSource, error) {
	if key == "" {
		return nil, ErrEnvKeyRequired
	}
	e := &EnvTokenSource{k: key}
	if err := e.Refresh(); err != nil {
		return nil, err
	}
	return e, nil
}

// EnvTokenSource implements the oauth2.TokenSource interface and retrieves
// the token from an environment variable.
//
// The token is cached in memory and, although it is not implemented a typical
// OAuth2 refreshing mechanism, the token can be refreshed (for example, if the
// the environment variable is updated) by calling the Refresh method.
//
// It is safe to use this token source from multiple goroutines.
type EnvTokenSource struct {
	k  string
	mu sync.RWMutex
	t  *oauth2.Token
}

var _ oauth2.TokenSource = (*EnvTokenSource)(nil)

// Token returns the cached token or an error if the token is invalid.
func (e *EnvTokenSource) Token() (*oauth2.Token, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if !e.t.Valid() {
		return nil, ErrTokenInvalid
	}
	return e.t, nil
}

// Refresh updates the token from the environment variable.
// It returns an error if the environment variable is not set.
func (e *EnvTokenSource) Refresh() error {
	// Lock immediately to prevent retrieving the token while it's
	// being refreshed
	e.mu.Lock()
	defer e.mu.Unlock()
	v := os.Getenv(e.k)
	if v == "" {
		return ErrEnvVarNotSet
	}
	e.t = &oauth2.Token{
		AccessToken: v,
		TokenType:   "bearer",
	}
	return nil
}

func NewEnvTokenSourceEncrypted(envKey string, encryptionKey []byte) (*EnvTokenSourceEncrypted, error) {
	if envKey == "" {
		return nil, ErrEnvKeyRequired
	}
	cb, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}
	e := &EnvTokenSourceEncrypted{k: envKey, cb: cb}
	if err := e.Refresh(); err != nil {
		return nil, err
	}
	return e, nil
}

// EnvTokenSourceEncrypted implements the oauth2.TokenSource interface and
// retrieves the token from an environment variable, as [EnvTokenSource],
// caching an encrypted version of the token in memory.
//
// The token is encrypted using AES in CFB mode with a random IV and the
// provided encryption key.
//
// It is safe to use this token source from multiple goroutines.
type EnvTokenSourceEncrypted struct {
	k   string
	mu  sync.RWMutex
	eat []byte // encrypted access token
	cb  cipher.Block
}

var _ oauth2.TokenSource = (*EnvTokenSourceEncrypted)(nil)

// Token returns the decrypted token or an error on decryption failure.
func (e *EnvTokenSourceEncrypted) Token() (*oauth2.Token, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	token, err := e.decrypt()
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: string(token),
		TokenType:   "bearer",
	}, nil
}

// Refresh updates the encrypted access token from the environment variable.
// It returns an error if the environment variable is not set or in case of
// encryption failure.
func (e *EnvTokenSourceEncrypted) Refresh() error {
	// Lock immediately to prevent retrieving the token while it's
	// being refreshed
	e.mu.Lock()
	defer e.mu.Unlock()
	v := os.Getenv(e.k)
	if v == "" {
		return ErrEnvVarNotSet
	}
	eat, err := e.encrypt([]byte(v))
	if err != nil {
		return err
	}
	e.eat = eat
	return nil
}

func (e *EnvTokenSourceEncrypted) encrypt(token []byte) ([]byte, error) {
	ciphertext := make([]byte, aes.BlockSize+len(token))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(e.cb, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], token)
	return ciphertext, nil
}

func (e *EnvTokenSourceEncrypted) decrypt() ([]byte, error) {
	if len(e.eat) < aes.BlockSize {
		return nil, ErrEncryptedAccessTokenTooShort
	}
	iv := e.eat[:aes.BlockSize]
	ciphertext := e.eat[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(e.cb, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

// OAuth2Endpoint is the OAuth 2.0 endpoint for HubSpot.
// It can be used to set the Endpoint field of an oauth2.Config.
var OAuth2Endpoint = oauth2.Endpoint{
	AuthURL:   endpoint.AuthorizeURL,
	TokenURL:  endpoint.OAuthToken,
	AuthStyle: oauth2.AuthStyleInParams,
}

// Token is an OAuth 2.0 as retrieved from [endpoint.OAuthToken] after a successful
// code exchange or refresh.
type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

// AccessToken is an OAuth 2.0 access token as retrieved from
// [endpoint.OAuthAccessTokens].
type AccessToken struct {
	HubID     int      `json:"hub_id,omitempty"`
	UserID    int      `json:"user_id,omitempty"`
	Scopes    []string `json:"scopes,omitempty"`
	TokenType string   `json:"token_type,omitempty"`
	User      string   `json:"user,omitempty"`
	HubDomain string   `json:"hub_domain,omitempty"`
	AppID     int      `json:"app_id,omitempty"`
	ExpiresIn int      `json:"expires_in,omitempty"`
	Token     string   `json:"token,omitempty"`
}

// RefreshToken is an OAuth 2.0 refresh token as retrieved from
// [endpoint.OAuthRefreshTokens].
type RefreshToken struct {
	HubID     int      `json:"hub_id,omitempty"`
	UserID    int      `json:"user_id,omitempty"`
	Scopes    []string `json:"scopes,omitempty"`
	TokenType string   `json:"token_type,omitempty"`
	User      string   `json:"user,omitempty"`
	HubDomain string   `json:"hub_domain,omitempty"`
	ClientID  string   `json:"client_id,omitempty"`
	Token     string   `json:"token,omitempty"`
}
