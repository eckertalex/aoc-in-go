package day07

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

func TestDay07(t *testing.T) {
	data := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	solution := New()

	t.Run("Part 1", func(t *testing.T) {
		want := "3749"
		got := solution.Part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := "11387"
		got := solution.Part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
