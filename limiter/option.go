package limiter

// Policy is an enumeration of the available rate limiting policies.
type Policy int

const (
	// PolicyWait is the policy that waits for the rate limiter to allow the
	// request to proceed.
	PolicyWait Policy = iota

	PolicyReject
)

type config struct {
	policy Policy
}

// Option is a functional option for configuring the limiter.
type Option func(*config)

// WithPolicy sets the policy option for the limiter.
func WithPolicy(policy Policy) Option {
	return func(c *config) {
		c.policy = policy
	}
}
