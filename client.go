package hubspot

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func NewClient(ts oauth2.TokenSource, opts ...Option) (*Client, error) {
	if ts == nil {
		return nil, ErrTokenSourceRequired
	}
	cfg := &config{
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(cfg)
	}
	var client Client
	client.hc = oauth2.NewClient(cfg.ctx, ts)
	return &client, nil
}

type Client struct {
	hc *http.Client
}

type config struct {
	ctx context.Context
}

type Option func(*config)

// WithContext sets the context for the client.
//
// The context is used to initialize the oauth2 client, which is used to
// make requests to the HubSpot API. As stated in the [oauth2 documentation],
// the client is not valid beyond the lifetime of the context.
//
// [oauth2 documentation]: https://pkg.go.dev/golang.org/x/oauth2#NewClient
func WithContext(ctx context.Context) Option {
	return func(c *config) {
		c.ctx = ctx
	}
}
