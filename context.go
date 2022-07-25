package gexpect

import (
	"testing"
)

// A Context is a group of related comparison and message formatting properties to be used
// together when identifying and reporting errors.
type Context[T any] struct {

	// Eval is the function used determine whether a value is expected. The first parameter is the
	// actual value to evaluate and the second parameter is the expected value to compare against.
	Eval func(T, T) bool

	// Fmt is the default format string used to report unmet expectations.
	Fmt string

	// FmtMap contains alternate format strings for reporting unmet expectations. The key
	// corresponds to the length of the argv argument. If the map is nil or does not contain an
	// entry corresponding to the length of a given value of argv then ErrorIsFmt will be used.
	FmtMap map[int]string
}

func (c *Context[T]) expect(t *testing.T, log LogMethod, actual, expected T, argv ...any) {
	if c.Eval(actual, expected) {
		return
	}
	format := c.Fmt
	if c.FmtMap != nil {
		if f, ok := c.FmtMap[len(argv)]; ok {
			format = f
		}
	}
	log(t, format, append([]any{actual, expected}, argv...)...)
}
