package day01

import "testing"

func TestExpenseReport(t *testing.T) {
	t.Run("Finding the pair of numbers", func(t *testing.T){
		got := FindPair([]int{2019,1}, 2020)
		want := 2019
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Example 1a", func(t *testing.T){
		got := FindPair([]int{1721, 979, 366, 299, 675, 1456}, 2020)
		want := 514579
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
