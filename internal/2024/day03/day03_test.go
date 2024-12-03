package day03

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

func TestDay1(t *testing.T) {
	data := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	solution := New()

	t.Run("Part 1", func(t *testing.T) {
		want := "161"
		got := solution.Part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	data = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	t.Run("Part 2", func(t *testing.T) {
		want := "48"
		got := solution.Part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
