package day12

import "testing"

func TestJoltAdaptor(t *testing.T) {
	t.Run("moving north", func(t *testing.T) {
		x, y := FollowDirections([]string{
			"N1",
		})
		wantX := 0
		wantY := 1
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving sourth", func(t *testing.T) {
		x, y := FollowDirections([]string{
			"S2",
		})
		wantX := 0
		wantY := -2
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})
	t.Run("moving east", func(t *testing.T) {
		x, y := FollowDirections([]string{
			"E1",
		})
		wantX := 1
		wantY := 0
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving west", func(t *testing.T) {
		x, y := FollowDirections([]string{
			"W2",
		})
		wantX := -2
		wantY := 0
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving forwards without turning", func(t *testing.T) {
		x, y := FollowDirections([]string{
			"F2",
		})
		wantX := 2
		wantY := 0
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving forwards after turning left", func(t *testing.T) {
		x, y := FollowDirections([]string{
			"L180",
			"F2",
		})
		wantX := -2
		wantY :=0
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving forwards after turning right", func(t *testing.T) {
		x, y := FollowDirections([]string{
			"R90",
			"F2",
		})
		wantX := 0
		wantY := -2
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})
}