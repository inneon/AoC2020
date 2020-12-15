package day14

import (
	"testing"
	"reflect"
)

func TestMemoryMasker(t *testing.T) {
	t.Run("writing to memory", func(t *testing.T) {
		got := RunMaskedProgram([]string{"mem[49397] = 468472"})
		want := 468472
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

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
	t.Run("applying a v2 mask", func(t *testing.T) {
		got := applyV2Mask(42, "X1101X", 0)
		want := []int{26, 27, 58, 59}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 2a", func(t *testing.T) {
		got := RunV2Program([]string{
			"mask = 000000000000000000000000000000X1001X",
			"mem[42] = 100",
			"mask = 00000000000000000000000000000000X0XX",
			"mem[26] = 1",
		})
		want := 208
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}