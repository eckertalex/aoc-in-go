package main

import (
	"testing"

	input "github.com/eckertalex/aoc-in-go/internal"
)

func TestDay1(t *testing.T) {
	data := `3   4
4   3
2   5
1   3
3   9
3   3`

	t.Run("Part 1", func(t *testing.T) {
		want := 11
		got := part1(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		want := 31
		got := part2(input.FromLiteral(data))

		if want != got {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}
