package day10

import "testing"

func TestJoltAdaptor(t *testing.T) {
	t.Run("example 1a", func(t *testing.T) {
		small, large := BuildAdaptorChain([]int{
			16,
			10,
			15,
			5,
			1,
			11,
			7,
			19,
			6,
			12,
			4,
		})
		got := small*large
		want := 35
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 2a", func(t *testing.T) {
		small, large := BuildAdaptorChain([]int{
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		})
		got := small*large
		want := 220
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}