package day06

import "testing"

func TestFormFiller(t *testing.T) {
	t.Run("example1", func (t *testing.T) {
		got := FillForm([]string {
			"abc",
			"",
			"a",
			"b",
			"c",
			"",
			"ab",
			"ac",
			"",
			"a",
			"a",
			"a",
			"a",
			"",
			"b",
		})
		want := 11
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}