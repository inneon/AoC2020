package day07

import "testing"

func TestBagPacker(t *testing.T) {
	t.Run("one instruction", func(t *testing.T) {
		got := CountOutermost([]string{"bright white bags contain 1 shiny gold bag."})
		want := 1
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("nesting bags", func(t *testing.T) {
		got := CountOutermost([]string{
			"bright white bags contain 1 shiny gold bag.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		})
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1a", func(t *testing.T) {
		got := CountOutermost([]string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		})
		want := 4
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("one instruction", func(t *testing.T) {
		got := CountInside([]string{"shiny gold bags contain no other bags."})
		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("simple nesting", func(t *testing.T) {
		got := CountInside([]string{
			"shiny gold bags contain 1 faded blue bag, 2 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		})
		want := 3
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 1b", func(t *testing.T) {
		got := CountInside([]string{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			"bright white bags contain 1 shiny gold bag.",
			"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
			"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
			"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
			"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			"faded blue bags contain no other bags.",
			"dotted black bags contain no other bags.",
		})
		want := 32
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		got := CountInside([]string{
			"shiny gold bags contain 2 dark red bags.",
			"dark red bags contain 2 dark orange bags.",
			"dark orange bags contain 2 dark yellow bags.",
			"dark yellow bags contain 2 dark green bags.",
			"dark green bags contain 2 dark blue bags.",
			"dark blue bags contain 2 dark violet bags.",
			"dark violet bags contain no other bags.",
		})
		want := 126
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}