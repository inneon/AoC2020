package day15

import (
	"testing"
)

func TestRambunctiousRecitor(t *testing.T) {
	t.Run("when the turn is in the initial phase", func(t *testing.T) {
		got := MemoryGame([]int{0,3,6}, 2)
		want := 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("when the previous number has not been said", func(t *testing.T) {
		got := MemoryGame([]int{0,3,6}, 4)
		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("when the previous number has been said", func(t *testing.T) {
		got := MemoryGame([]int{0,3,6}, 5)
		want := 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1a", func(t *testing.T) {
		got := MemoryGame([]int{0,3,6}, 2020)
		want := 436
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}