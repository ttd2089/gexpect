package gexpect

import (
	"testing"
)

func Test_expect(t *testing.T) {

	t.Run("Eval(a, b) == true", func(t *testing.T) {

		t.Run("does not log", func(t *testing.T) {
			s := logSpy{}
			defer s.verify(t)
			c := Context[any]{Fmt: "format", Eval: func(_, _ any) bool { return true }}
			c.expect(t, s.log, 1, 2)
		})

	})

	t.Run("Eval(a, b) == false", func(t *testing.T) {

		t.Run("no matching entry in FmtMap/logs using Log and Fmt", func(t *testing.T) {
			s := logSpy{}
			defer s.verify(t)
			c := Context[any]{Fmt: "format", Eval: func(_, _ any) bool { return false }}
			s.expect(t, c.Fmt, 1, 2)
			c.expect(t, s.log, 1, 2)
		})

		t.Run("matching entry in ErrorIsFmtMap/logs using Log and FmtMap[n]", func(t *testing.T) {
			s := logSpy{}
			defer s.verify(t)
			c := Context[any]{
				Eval: func(_, _ any) bool { return false },
				FmtMap: map[int]string{
					2: "format from map",
				},
			}
			s.expect(t, c.FmtMap[2], 1, 2, 3, 4)
			c.expect(t, s.log, 1, 2, 3, 4)
		})
	})
}
