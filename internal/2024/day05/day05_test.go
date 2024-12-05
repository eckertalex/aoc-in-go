package day05

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

func TestDay05(t *testing.T) {
	data := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	solution := New()

	t.Run("Part 1", func(t *testing.T) {
		want := "143"
		got := solution.Part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := "123"
		got := solution.Part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
