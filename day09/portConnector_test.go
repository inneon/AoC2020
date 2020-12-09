package day09

import "testing"

func TestPortConnector(t *testing.T) {
	t.Run("when a code is valid", func(t *testing.T) {
		got := isSum([]int{
			35,
			20,
			15,
			25,
			47,
		}, 40)
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("when a code is invalid", func(t *testing.T) {
		got := isSum([]int{
			35,
			20,
			15,
			25,
			47,
		}, 2)
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("example 1a", func(t *testing.T) {
		got, _ := FirstInvalid([]int{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}, 5)
		want := 127
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1b", func(t *testing.T) {
		got := EncryptionWeakness([]int{
			35,
			20,
			15,
			25,
			47,
			40,
			62,
			55,
			65,
			95,
			102,
			117,
			150,
			182,
			127,
			219,
			299,
			277,
			309,
			576,
		}, 5)
		want := 62
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}