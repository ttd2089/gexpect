package gexpect

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type call struct {
	t      *testing.T
	format string
	args   []any
}

func (c call) String() string {
	var sb strings.Builder
	sb.WriteString("(%p, %q, [")
	argSpecs := make([]string, 0, len(c.args))
	for range c.args {
		argSpecs = append(argSpecs, "<%v>")
	}
	sb.WriteString(strings.Join(argSpecs, ", "))
	sb.WriteString("])")
	return fmt.Sprintf(sb.String(), append([]any{c.t, c.format}, c.args...)...)
}

type logSpy struct {
	expected []call
	actual   []call
}

func (spy *logSpy) log(t *testing.T, format string, args ...any) {
	spy.actual = append(spy.actual, call{t, format, args})
}

func (spy *logSpy) expect(t *testing.T, format string, args ...any) {
	spy.expected = append(spy.expected, call{t, format, args})
}

func (spy *logSpy) verify(t *testing.T) {
	nCalls := len(spy.actual)
	nExpected := len(spy.expected)
	if nCalls != nExpected {
		t.Errorf("expected calls: %d; actual %d\n", nExpected, nCalls)
	}
	n := nCalls
	if nExpected < n {
		n = nExpected
	}
	for i := 0; i < n; i++ {
		actual := spy.actual[i]
		expected := spy.expected[i]
		if reflect.DeepEqual(actual, expected) {
			continue
		}
		t.Errorf("calls[%d]: expected %s; actual %s\n", i, expected, actual)
	}
	for i := n; i < nExpected; i++ {
		t.Errorf("calls[%d]: missing expected call %s\n", n, spy.expected[n])
	}
	for i := n; i < nCalls; i++ {
		t.Errorf("calls[%d]: unexpected call %s\n", n, spy.actual[n])
	}
}
