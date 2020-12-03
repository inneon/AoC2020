package day02

import "testing"

func TestPassword(t *testing.T) {
	t.Run("Can parse a line", func(t *testing.T) {
		got, _ := parse("1-13 a: abcde")
		want := passwordLine{
			lowerBound: 1,
			upperBound: 13,
			character: "a",
			password: "abcde",
		}
		if *got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("When a password matches it rule", func(t *testing.T){
		got := checkRule(passwordLine{
			lowerBound: 1,
			upperBound: 3,
			character: "a",
			password: "abcda",
		})
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("When a password doesn't match its rule", func(t *testing.T){
		got := checkRule(passwordLine{
			lowerBound: 1,
			upperBound: 3,
			character: "b",
			password: "cdefg",
		})
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Example 1a", func(t *testing.T) {
		got := NumberOfMatches([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})
		want := 2
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("When a password matches the second rule", func(t *testing.T){
		got := checkSecondRule(passwordLine{
			lowerBound: 1,
			upperBound: 3,
			character: "a",
			password: "abcda",
		})
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("When a password doesn't match its second rule", func(t *testing.T){
		got := checkSecondRule(passwordLine{
			lowerBound: 2,
			upperBound: 9,
			character: "c",
			password: "ccccccccc",
		})
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Example 1a", func(t *testing.T) {
		got := NumberOfSecondaryMatches([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})
		want := 1
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}