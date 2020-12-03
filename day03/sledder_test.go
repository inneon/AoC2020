package day03

import "testing"

func TestSledder(t *testing.T) {
	t.Run("When there are no trees", func(t *testing.T) {
		got := CountTrees([]string{"..."})
		want := 0
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("When there is one tree", func(t *testing.T) {
		got := CountTrees([]string{"....", "...#"})
		want := 1
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("When the pattern repeats", func(t *testing.T) {
		got := CountTrees([]string{"....", "...#", "..#."})
		want := 2
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Example 1a", func(t *testing.T) {
		got := CountTrees([]string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		})
		want := 7
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}