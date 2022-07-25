package gexpect

import "testing"

// A LogMethod is a function that logs a message using an instance of *testing.T
type LogMethod func(t *testing.T, format string, args ...any)

// Fatalf logs a message using the (*testing.T).Fatalf method.
func Fatalf(t *testing.T, format string, args ...any) {
	t.Fatalf(format, args...)
}

// Errorf logs a message using the (*testing.T).Errorf method.
func Errorf(t *testing.T, format string, args ...any) {
	t.Errorf(format, args...)
}

// Logf logs a message usin the (*testing.T).Logf method.
func Logf(t *testing.T, format string, args ...any) {
	t.Logf(format, args...)
}
