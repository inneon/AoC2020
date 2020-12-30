package day18

import (
	"testing"
)

func TestBadMathser(t *testing.T) {
	t.Run("when the input is a single number", func(t *testing.T) {
		got := interpretRightToLeft("123")
		want := 123
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("when the input is two numbers added", func(t *testing.T) {
		got := interpretRightToLeft("123 + 456")
		want := 579
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("when the input is a single bracketed number", func(t *testing.T) {
		got := interpretRightToLeft("(1 + 2)")
		want := 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("when the input is a bracketed", func(t *testing.T) {
		got := interpretRightToLeft("(1 + (2 * 3))")
		want := 7
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}