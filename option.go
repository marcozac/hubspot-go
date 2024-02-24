package hubspot

import (
	"context"
	"net/url"
	"strings"
)

type clientConfig struct {
	ctx context.Context
}

type ClientOption func(*clientConfig)

// WithContext sets the context for the client.
//
// The context is used to initialize the oauth2 client, which is used to
// make requests to the HubSpot API. As stated in the [oauth2 documentation],
// the client is not valid beyond the lifetime of the context.
//
// [oauth2 documentation]: https://pkg.go.dev/golang.org/x/oauth2#NewClient
func WithContext(ctx context.Context) ClientOption {
	return func(c *clientConfig) {
		c.ctx = ctx
	}
}

// RequestConfig is a struct that contains options for making requests to the
// HubSpot API.
//
// All fields are exported only for documentation purposes. This struct is
// used exclusively as a target for the RequestOption type. Refer to the
// methods' documentation for a list of available options.
type RequestConfig struct {
	// Archived is a boolean that determines whether to include archived
	// objects in the response.
	Archived bool

	// Properties is a list of the properties to be returned in the response.
	//
	// According to the HubSpot API documentation, if any of the specified
	// properties are not present on the requested object(s), they will be
	// ignored.
	Properties []string
}

type RequestOption func(*RequestConfig)

// WithArchived sets the archived option for the request.
//
// See [RequestConfig.Archived] for more information.
func WithArchived(archived bool) RequestOption {
	return func(c *RequestConfig) {
		c.Archived = archived
	}
}

// WithProperties sets the properties option for the request.
//
// See [RequestConfig.Properties] for more information.
func WithProperties(properties ...string) RequestOption {
	return func(c *RequestConfig) {
		c.Properties = properties
	}
}

// applyRequestOptions applies the given options to the given request config,
// returning the same reference if the given config is not nil. Otherwise, a
// new config is created and returned with the options applied.
func applyRequestOptions(cfg *RequestConfig, opts ...RequestOption) *RequestConfig {
	if cfg == nil {
		cfg = &RequestConfig{}
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// applyArchived checks the archived field of the given request config and
// adds the appropriate query parameter to the given url.Values. It always
// adds the archived query parameter.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyArchivedOption(cfg *RequestConfig, query url.Values) {
	if cfg.Archived {
		query.Add("archived", "true")
	} else {
		query.Add("archived", "false")
	}
}

// applyPropertiesOption checks the properties field of the given request
// config and adds the appropriate query parameter to the given url.Values.
// If the properties field is empty, this function does nothing.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyPropertiesOption(cfg *RequestConfig, query url.Values) {
	if len(cfg.Properties) > 0 {
		query.Add("properties", strings.Join(cfg.Properties, ","))
	}
}
