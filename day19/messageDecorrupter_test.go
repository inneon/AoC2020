package day19

import (
	"testing"
)

func TestMessageDecorrupter(t *testing.T) {
	t.Run("literal rule", func(t *testing.T) {
		got := doesMatch("a", map[int]string{ 0: "a" }, 0)
		want := []string{""}
		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("nested rule", func(t *testing.T) {
		got := doesMatch("a", map[int]string{ 0: "1", 1: "a" }, 0)
		want := []string{""}
		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("compound rule", func(t *testing.T) {
		got := doesMatch("ab", map[int]string{ 0: "1 2", 1: "a", 2: "b"}, 0)
		want := []string{""}
		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("choice rule", func(t *testing.T) {
		got := doesMatch("ba", map[int]string{ 0: "1 2 | 2 1", 1: "a", 2: "b"}, 0)
		want := []string{""}
		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Both types of rules - true", func(t *testing.T) {
		got := doesMatch("ababbb", map[int]string{ 
				0: "4 1 5", 
				1: "2 3 | 3 2", 
				2: "4 4 | 5 5", 
				3: "4 5 | 5 4", 
				4: "a", 
				5: "b",
			},
			0,
		)
		want := []string{""}
		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Both types of rules - false", func(t *testing.T) {
		got := doesMatch("aaaabbb", map[int]string{ 
				0: "4 1 5", 
				1: "2 3 | 3 2", 
				2: "4 4 | 5 5", 
				3: "4 5 | 5 4", 
				4: "a", 
				5: "b",
			},
			0,
		)
		want := []string{"b"}
		if len(got) != len(want) || got[0] != want[0] {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("GetMatchingRules", func(t *testing.T) {
		got := GetMatchingRules([]string {
			"0: 4 1 5",
			"1: 2 3 | 3 2",
			"2: 4 4 | 5 5",
			"3: 4 5 | 5 4",
			"4: \"a\"",
			"5: \"b\"",
			"",
			"ababbb",
			"bababa",
			"abbbab",
			"aaabbb",
			"aaaabbb",
		})
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}