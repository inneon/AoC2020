package day13

import "testing"

func TestBusScheduler(t *testing.T) {
	t.Run("example 1a", func(t *testing.T) {
		got := EarliestBus(939, "7,13,x,x,59,x,31,19")
		want := 295
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1b", func(t *testing.T) {
		got := AlignedTimestamps("7,13,x,x,59,x,31,19")
		want := 1068781
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1b", func(t *testing.T) {
		got := alignedTimestamps([]positionAndID{{ id: 7, position: 0}, {id: 13, position: 1}})
		want := 77
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}