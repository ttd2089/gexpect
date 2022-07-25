package gexpect

import "testing"

// A Config is a configuration object that defines how unmet expectations will be reported.
type Config struct {

	// Log is the LogMethod that unmet expectations will be reported with.
	Log LogMethod

	// ErrContext is a Context for evaluating error values.
	ErrContext Context[error]

	// ExprContext is a Context for evaluating value expressions.
	ExprContext Context[any]
}

// ErrorIs logs a message when the actual error does not match the expected error.
func (c *Config) ErrorIs(t *testing.T, err, expected error, argv ...any) {
	c.ErrContext.expect(t, c.Log, err, expected, argv...)
}

// ExprIs logs a message when the actual expression does not match the expected value.
func (c *Config) ExprIs(t *testing.T, expr, expected any, argv ...any) {
	c.ExprContext.expect(t, c.Log, expr, expected, argv...)
}
