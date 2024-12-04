package day02

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

func TestDay02(t *testing.T) {
	data := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	solution := New()

	t.Run("Part 1", func(t *testing.T) {
		want := "2"
		got := solution.Part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := "4"
		got := solution.Part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
