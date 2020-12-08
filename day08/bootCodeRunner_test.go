package day08

import "testing"

func TestBootCode(t *testing.T) {
	t.Run("example 1a", func(t *testing.T) {
		got := Run([]string{
			"nop +0",
			"acc +1",
			"jmp +4",
			"acc +3",
			"jmp -3",
			"acc -99",
			"acc +1",
			"jmp -4",
			"acc +6",
		})
		want := 5
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}