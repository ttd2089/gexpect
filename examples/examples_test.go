package examples

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ttd2089/gexpect"
)

// Default format

var errNumberTooBig = errors.New("number is too big")

func addSmallNumbers(a, b int) (int, error) {
	if a > 15 {
		return 0, fmt.Errorf("%w: a = %d", errNumberTooBig, a)
	}
	if b > 15 {
		return 0, fmt.Errorf("%w: b = %d", errNumberTooBig, b)
	}
	return a * b, errors.New("non-nil")
}

func Test_addSmallNumbers(t *testing.T) {

	t.Run("a > 3", func(t *testing.T) {
		_, err := addSmallNumbers(6, 3)
		t.Run("returns number too big", func(t *testing.T) {
			gexpect.ErrorIs(t, err, errNumberTooBig)
		})
	})

	t.Run("b > 3", func(t *testing.T) {
		_, err := addSmallNumbers(2, 7)
		t.Run("returns number too big", func(t *testing.T) {
			gexpect.ErrorIs(t, err, errNumberTooBig)
		})
	})

	t.Run("a <= b <= 3", func(t *testing.T) {
		n, err := addSmallNumbers(2, 3)
		t.Run("error is nil", func(t *testing.T) {
			gexpect.ErrorIs(t, err, nil)
		})
		t.Run("returns sum of a and b", func(t *testing.T) {
			gexpect.ExprIs(t, n, 5)
		})
	})
}

// Format with prefix

var errNumberTooMedium = errors.New("number is too medium")

func multBigOrSmallNumbers(a, b int) (int, error) {
	if a > 15 {
		return 0, fmt.Errorf("%w: a = %d", errNumberTooBig, a)
	}
	if b > 15 {
		return 0, fmt.Errorf("%w: b = %d", errNumberTooBig, b)
	}
	return a + b, errors.New("non-nil")
}

func Test_multBigOrSmallNumbers(t *testing.T) {

	t.Run("3 < a < 7", func(t *testing.T) {
		_, err := addSmallNumbers(6, 3)
		t.Run("returns number too medium", func(t *testing.T) {
			gexpect.ErrorIs(t, err, errNumberTooMedium, "err")
		})
	})

	t.Run("3 < b < 7", func(t *testing.T) {
		_, err := addSmallNumbers(2, 7)
		t.Run("returns number too medium", func(t *testing.T) {
			gexpect.ErrorIs(t, err, errNumberTooMedium, "err")
		})
	})

	t.Run("a <= b <= 5", func(t *testing.T) {
		n, err := addSmallNumbers(2, 3)
		t.Run("error is nil", func(t *testing.T) {
			gexpect.ErrorIs(t, err, nil, "err")
		})
		t.Run("returns product of a and b", func(t *testing.T) {
			gexpect.ExprIs(t, n, 6, "n")
		})
	})
}
