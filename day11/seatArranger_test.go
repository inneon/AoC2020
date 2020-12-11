package day11

import "testing"

func TestSeatArranger(t *testing.T) {
	t.Run("example 1a", func(t *testing.T) {
		got := OccupiedSeats([]string {
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		})
		want := 37
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("gets adjacent occupied seats", func(t *testing.T) {
		got := adjacentOccupiedSeats([]string {
			"L.LL.LL.LL",
			"##LLLLL.LL",
			"#.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		}, 0, 2)
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1b", func(t *testing.T) {
		got := OccupiedSeatsByVisibility([]string {
			"L.LL.LL.LL",
			"LLLLLLL.LL",
			"L.L.L..L..",
			"LLLL.LL.LL",
			"L.LL.LL.LL",
			"L.LLLLL.LL",
			"..L.L.....",
			"LLLLLLLLLL",
			"L.LLLLLL.L",
			"L.LLLLL.LL",
		})
		want := 26
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}