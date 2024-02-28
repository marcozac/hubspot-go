package hubspot

import (
	"context"
	"net/url"
	"strconv"
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

// applyClientOptions applies the given options to the given client config,
// returning the same reference if the given config is not nil. Otherwise, a
// new config is created and returned with the options applied.
func applyClientOptions(cfg *clientConfig, opts ...ClientOption) *clientConfig {
	if cfg == nil {
		cfg = &clientConfig{
			ctx: context.Background(),
		}
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// RequestConfig is a struct that contains options for making requests to the
// HubSpot API.
//
// All fields are exported only for documentation purposes. This struct is
// used exclusively as a target for the RequestOption type. Refer to the
// methods' documentation for a list of available options.
type RequestConfig struct {
	// Archived is a boolean that determines whether to return only archived
	// objects in the response.
	Archived bool

	// Properties is a list of the properties to be returned in the response.
	//
	// According to the HubSpot API documentation, if any of the specified
	// properties are not present on the requested object(s), they will be
	// ignored.
	Properties []string

	// PropertiesWithHistory is a list of the properties to be returned in the
	// response, along with their history.
	//
	// According to the HubSpot API documentation, if any of the specified
	// properties are not present on the requested object(s), they will be
	// ignored.
	PropertiesWithHistory []string

	// Associations is a list of the associations to be returned in the
	// response.
	//
	// According to the HubSpot API documentation, if any of the specified
	// associations are not present on the requested object(s), they will be
	// ignored.
	Associations []string

	// Limit is the maximum number of objects to return in the response.
	//
	// According to the HubSpot API documentation, the maximum value for this
	// field is 100.
	Limit int

	// After is the paging cursor token used to get the next page of results.
	After string
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

// WithPropertiesWithHistory sets the properties with history option for the
// request.
//
// See [RequestConfig.PropertiesWithHistory] for more information.
func WithPropertiesWithHistory(properties ...string) RequestOption {
	return func(c *RequestConfig) {
		c.PropertiesWithHistory = properties
	}
}

// WithAssociations sets the associations option for the request.
//
// See [RequestConfig.Associations] for more information.
func WithAssociations(associations ...string) RequestOption {
	return func(c *RequestConfig) {
		c.Associations = associations
	}
}

// WithLimit sets the limit option for the request.
//
// See [RequestConfig.Limit] for more information.
func WithLimit(limit int) RequestOption {
	return func(c *RequestConfig) {
		c.Limit = limit
	}
}

// WithAfter sets the after option for the request.
//
// See [RequestConfig.After] for more information.
func WithAfter(after string) RequestOption {
	return func(c *RequestConfig) {
		c.After = after
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

const (
	QueryKeyArchived              = "archived"
	QueryKeyProperties            = "properties"
	QueryKeyPropertiesWithHistory = "propertiesWithHistory"
	QueryKeyAssociations          = "associations"
	QueryKeyLimit                 = "limit"
	QueryKeyAfter                 = "after"
)

// queryOption is a function that applies a query parameter to the given
// url.Values.
type queryOption func(cfg *RequestConfig, query url.Values)

// applyQueryOptions applies the given options to the given request config,
// setting also (overwriting) the url.RawQuery field with the resulting
// query string.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyQueryOptions(cfg *RequestConfig, url *url.URL, opts ...queryOption) {
	q := url.Query()
	for _, opt := range opts {
		opt(cfg, q)
	}
	url.RawQuery = getRAWQuery(q)
}

func getRAWQuery(q url.Values) string {
	var sb strings.Builder
	for k, v := range q {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(strings.Join(v, ","))
		sb.WriteString("&")
	}
	return strings.TrimSuffix(sb.String(), "&")
}

// applyArchived checks the archived field of the given request config and
// adds the appropriate query parameter to the given url.Values. It always
// adds the "archived" query parameter.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyArchivedQuery(cfg *RequestConfig, query url.Values) {
	if cfg.Archived {
		query.Set(QueryKeyArchived, "true")
	} else {
		query.Set(QueryKeyArchived, "false")
	}
}

// applyPropertiesOption checks the properties field of the given request
// config and adds the appropriate query parameter to the given url.Values.
// If the cfg.Properties field is empty, this function does nothing.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyPropertiesQuery(cfg *RequestConfig, query url.Values) {
	for _, p := range cfg.Properties {
		query.Add(QueryKeyProperties, p)
	}
}

// applyPropertiesWithHistoryOption checks the properties with history field of the given request
// config and adds the appropriate query parameter to the given url.Values.
// If the cfg.PropertiesWithHistory field is empty, this function does nothing.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyPropertiesWithHistoryQuery(cfg *RequestConfig, query url.Values) {
	for _, p := range cfg.PropertiesWithHistory {
		query.Add(QueryKeyPropertiesWithHistory, p)
	}
}

// applyAssociationsOption checks the associations field of the given request
// config and adds the appropriate query parameter to the given url.Values.
// If the cfg.Associations field is empty, this function does nothing.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyAssociationsQuery(cfg *RequestConfig, query url.Values) {
	for _, a := range cfg.Associations {
		query.Add(QueryKeyAssociations, a)
	}
}

// applyLimitOption checks the limit field of the given request
// config and adds the appropriate query parameter to the given url.Values.
// If the cfg.Limit field is 0, this function does nothing.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyLimitQuery(cfg *RequestConfig, query url.Values) {
	if cfg.Limit > 0 {
		query.Set(QueryKeyLimit, strconv.Itoa(cfg.Limit))
	}
}

// applyAfterOption checks the after field of the given request
// config and adds the appropriate query parameter to the given url.Values.
// If the cfg.After field is empty, this function does nothing.
//
// NOTE: this function performs a write operation on the given url.Values
// (map), so it is not safe to call concurrently.
func applyAfterQuery(cfg *RequestConfig, query url.Values) {
	if cfg.After != "" {
		query.Set(QueryKeyAfter, cfg.After)
	}
}
