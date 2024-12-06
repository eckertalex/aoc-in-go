package day06

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

func TestDay06(t *testing.T) {
	data := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	solution := New()

	t.Run("Part 1", func(t *testing.T) {
		want := "41"
		got := solution.Part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := "6"
		got := solution.Part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
