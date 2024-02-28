package limiter

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// New10 returns a new rate limiter that allows a burst of `b` requests every
// 10 seconds.
//
// It can be used to create a rate limiter with a custom burst size, for
// example, to match the rate limits of Professional and Enterprise accounts.
func New10(b int) *rate.Limiter {
	return rate.NewLimiter(rate.Every(10*time.Second), b)
}

// NewDefault10 returns a new rate limiter that allows a burst of 100 requests
// every 10 seconds.
//
// According to [HubSpot's API documentation], this is the default rate limit
// for OAuth apps and Free and Starter accounts.
func NewDefault10() *rate.Limiter {
	return New10(100)
}

// NewInf returns a new rate limiter that allows an infinite burst of requests.
func NewInf() *rate.Limiter {
	return rate.NewLimiter(rate.Inf, 0)
}

// WrapHTTPClient returns a new HTTP client that wraps the given client transport
// with the given rate limiter. If the client is nil, it is replaced with
// http.DefaultClient. If the rate limiter is nil, it is replaced with a new
// rate limiter that allows an infinite burst of requests.
func WrapHTTPClient(c *http.Client, l *rate.Limiter, opts ...Option) *http.Client {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	if c == nil {
		c = http.DefaultClient
	}
	if l == nil {
		l = NewInf()
	}
	rt := c.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &http.Client{
		CheckRedirect: c.CheckRedirect,
		Jar:           c.Jar,
		Timeout:       c.Timeout,
		Transport: &Transport{
			rt:     rt,
			lim:    l,
			policy: cfg.policy,
		},
	}
}

var _ http.RoundTripper = (*Transport)(nil)

type Transport struct {
	rt     http.RoundTripper
	lim    *rate.Limiter
	policy Policy
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	switch t.policy {
	case PolicyWait:
		if err := t.lim.Wait(req.Context()); err != nil {
			return nil, err
		}
	case PolicyReject:
		if !t.lim.Allow() {
			return nil, ErrRateLimitExceeded
		}
	default:
		return nil, ErrUnknownPolicy
	}
	return t.rt.RoundTrip(req)
}

var (
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	ErrUnknownPolicy     = errors.New("unknown policy")
)
