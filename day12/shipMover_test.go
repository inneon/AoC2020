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

	t.Run("moving south", func(t *testing.T) {
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

	t.Run("following the waypoint", func(t *testing.T) {
		x, y := FollowWaypoint([]string{
			"F2",
		})
		wantX := 20
		wantY := 2
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving the waypoint north", func(t *testing.T) {
		x, y := FollowWaypoint([]string{
			"N1",
			"F2",
		})
		wantX := 20
		wantY := 4
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving the waypoint south", func(t *testing.T) {
		x, y := FollowWaypoint([]string{
			"S2",
			"F2",
		})
		wantX := 20
		wantY := -2
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving the waypoint east", func(t *testing.T) {
		x, y := FollowWaypoint([]string{
			"E1",
			"F2",
		})
		wantX := 22
		wantY := 2
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("moving the waypoint west", func(t *testing.T) {
		x, y := FollowWaypoint([]string{
			"W2",
			"F2",
		})
		wantX := 16
		wantY := 2
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("rotating the waypoint left", func(t *testing.T) {
		x, y := FollowWaypoint([]string{
			"L90",
			"F2",
		})
		wantX := -2
		wantY := 20
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})

	t.Run("rotating the waypoint right", func(t *testing.T) {
		x, y := FollowWaypoint([]string{
			"R180",
			"F2",
		})
		wantX := -20
		wantY := -2
		if x != wantX || y != wantY {
			t.Errorf("got (%d, %d) want (%d, %d) x", x, y, wantX, wantY)
		}
	})
}