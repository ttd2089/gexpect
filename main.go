package gexpect

import (
	"errors"
	"reflect"
	"testing"
)

// DefaultErrorIsFmt is the default log message format for unmet expectations about error values.
const DefaultErrorIsFmt = "actual <%s>; expected <%s>"

// DefaultErrorIsPrefixFmt is the default log message format for unmet expectations about error
// values when one additional argument is provided to the the expect function. This format treats
// the additional argument as a prefix to the default format.
const DefaultErrorIsPrefixFmt = "%[3]v: actual <%[1]s>; expected <%[2]s>"

// DefaultExprIsFmt is the default log message format for unmet expectations about arbitrary value
// expressions.
const DefaultExprIsFmt = "actual <%+v>; expected <%+v>"

// DefaultExprIsPrefixFmt is the default log message format for unmet expectations about arbitrary
// value expressions when one additional argument is provided to the the expect function. This
// format treats the additional argument as a prefix to the default format.
const DefaultExprIsPrefixFmt = "%[3]v: actual <%[1]+v>; expected <%[2]+v>"

// DefaultConfig generates a ready-to-use Config with sane defaults.
func DefaultConfig() Config {
	return Config{
		Log: Fatalf,
		ErrContext: Context[error]{
			Eval: errors.Is,
			Fmt:  DefaultErrorIsFmt,
			FmtMap: map[int]string{
				1: DefaultErrorIsPrefixFmt,
			},
		},
		ExprContext: Context[any]{
			Eval: reflect.DeepEqual,
			Fmt:  DefaultExprIsFmt,
			FmtMap: map[int]string{
				1: DefaultExprIsPrefixFmt,
			},
		},
	}
}

var defaultConfig Config = DefaultConfig()

// ErrorIs forwards a call to the ErrorIs method on a default instance of Config.
func ErrorIs(t *testing.T, err, expected error, argv ...any) {
	defaultConfig.ErrorIs(t, err, expected, argv...)
}

// ErrorIs forwards a call to the ExprIs method on a default instance of Config.
func ExprIs(t *testing.T, expr, expected any, argv ...any) {
	defaultConfig.ExprIs(t, expr, expected, argv...)
}
