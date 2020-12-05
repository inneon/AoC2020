package day05


import "testing"

func TestTickerReader(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		got := getID("FBFBBFFRLR")
		want := 357
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("example2", func(t *testing.T) {
		got := getID("BFFFBBFRRR")
		want := 567
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("example3", func(t *testing.T) {
		got := getID("FFFBBBFRRR")
		want := 119
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("gets the biggest of the examples", func(t *testing.T) {
		got, _ := Range([]string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"})
		want := 820
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("gets the missing seat", func(t *testing.T) {
		got := Missing([]string{"1000", "0100", "0110", "0111"})
		want := 5
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}