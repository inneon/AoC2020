package day17

import (
	"testing"
)

func TestConwayCuber(t *testing.T) {
	t.Run("parsing input", func(t *testing.T) {
		got, _ := parseInput([]string{
			".#.",
			"..#",
			"###",
		})
		if !got[0][0][1] ||
			!got[0][1][2] ||
			!got[0][2][0] ||
			!got[0][2][1] ||
			!got[0][2][2] {
			t.Errorf("got %v", got)
		}
	})

	t.Run("getting bounds", func(t *testing.T) {
		_, got := parseInput([]string{
			".#.",
			"..#",
			"###",
		})
		if got.minX != 0 || got.maxX != 2 || 
			got.minY != 0 || got.maxY != 2 || 
			got.minZ != 0 || got.maxZ != 0 {
			t.Errorf("got %v", got)
		}
	})

	t.Run("getting neighbours", func(t *testing.T) {
		dimensions, _ := parseInput([]string{
			".#.",
			"..#",
			"###",
		})
		got := neigbours(dimensions, 1, 2, 0)
		want := 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("counting", func(t *testing.T) {
		dimensions, bounds := parseInput([]string{
			".#.",
			"..#",
			"###",
		})
		got := count(dimensions, bounds)
		want := 5
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example1", func(t *testing.T) {
		got := RunCube([]string{
			".#.",
			"..#",
			"###",
		}, 6)
		want := 112
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}