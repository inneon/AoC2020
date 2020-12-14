package day13

import "testing"

func TestBusScheduler(t *testing.T) {
	t.Run("example 1a", func(t *testing.T) {
		got := EarliestBus(939, "7,13,x,x,59,x,31,19")
		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}