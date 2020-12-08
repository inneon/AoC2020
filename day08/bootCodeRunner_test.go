package day08

import "testing"

func TestBootCode(t *testing.T) {
	t.Run("example 1a", func(t *testing.T) {
		got, code := Run([]string{
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
		wantAccum := 5
		wantCode := Loop
		if got != wantAccum {
			t.Errorf("got %d want %d", got, wantAccum)
		}
		if code != wantCode {
			t.Errorf("got %d want %d", code, wantCode)
		}
	})

	t.Run("example 1b", func(t *testing.T) {
		got := Fix([]string{
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
		want := 8
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}