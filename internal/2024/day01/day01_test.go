package day01

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

func TestDay1(t *testing.T) {
	data := `3   4
4   3
2   5
1   3
3   9
3   3`

	solution := New()

	t.Run("Part 1", func(t *testing.T) {
		want := "11"
		got := solution.Part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := "31"
		got := solution.Part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
