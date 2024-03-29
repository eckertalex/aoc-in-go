package main

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal"
)

func TestDay1(t *testing.T) {
	data := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	t.Run("Part 1", func(t *testing.T) {
		want := 2
		got := part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := 4
		got := part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}
