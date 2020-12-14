package day14

import "testing"

func TestMemoryMasker(t *testing.T) {
	// t.Run("writing to memory", func(t *testing.T) {
	// 	got := RunMaskedProgram([]string{"mem[49397] = 468472"})
	// 	want := 468472
	// 	if got != want {
	// 		t.Errorf("got %d want %d", got, want)
	// 	}
	// })

	t.Run("applying a mask", func(t *testing.T) {
		got := applyMask(11, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
		want := 73
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("applying a mask", func(t *testing.T) {
		got := applyMask(101, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
		want := 101
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1a", func(t *testing.T) {
		got := RunMaskedProgram([]string{
			"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			"mem[8] = 11",
			"mem[7] = 101",
			"mem[8] = 0",
		})
		want := 165
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}