package day17

import (
	"testing"
)

func TestConwayHyperCuber(t *testing.T) {
	t.Run("example2", func(t *testing.T) {
		got := RunHyperCube([]string{
			".#.",
			"..#",
			"###",
		}, 6)
		want := 848
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}