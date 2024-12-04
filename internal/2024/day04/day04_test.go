package day04

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

func TestDay04(t *testing.T) {
	data := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	solution := New()

	t.Run("Part 1", func(t *testing.T) {
		want := "18"
		got := solution.Part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := "9"
		got := solution.Part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}
