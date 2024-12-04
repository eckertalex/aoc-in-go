package day04

import (
	"strconv"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

type Solution struct{}

func New() *Solution {
	return &Solution{}
}

func isWordMatch(matrix [][]rune, word []rune, row, col int, direction [2]int) bool {
	rows, cols := len(matrix), len(matrix[0])

	for i, char := range word {
		newRow := row + i*direction[0]
		newCol := col + i*direction[1]

		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || matrix[newRow][newCol] != char {
			return false
		}
	}

	return true
}

func countWordOccurrences(matrix [][]rune, word []rune) int {
	count := 0

	directions := [8][2]int{
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // down-left
		{1, 0},   // down
		{1, 1},   // down-right
	}

	for row, line := range matrix {
		for col := range line {
			for _, direction := range directions {
				if isWordMatch(matrix, word, row, col, direction) {
					count++
				}
			}
		}
	}

	return count
}

func (s *Solution) Part1(input *input.Input) string {
	matrix := input.RuneMatrix()
	wordRune := []rune("XMAS")
	count := countWordOccurrences(matrix, wordRune)
	return strconv.Itoa(count)
}

func isXMas(matrix [][]rune, centerRow, centerCol int) bool {
	rows, cols := len(matrix), len(matrix[0])

	directions := [4][2]int{
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}
	if matrix[centerRow][centerCol] != 'A' {
		return false
	}

	for _, direction := range directions {
		top := centerRow + direction[0]
		topCol := centerCol + direction[1]
		bottom := centerRow - direction[0]
		bottomCol := centerCol - direction[1]

		if top < 0 || top >= rows || topCol < 0 || topCol >= cols ||
			bottom < 0 || bottom >= rows || bottomCol < 0 || bottomCol >= cols {
			return false
		}

		if (matrix[top][topCol] != 'M' || matrix[bottom][bottomCol] != 'S') &&
			(matrix[top][topCol] != 'S' || matrix[bottom][bottomCol] != 'M') {
			return false
		}
	}

	return true
}

func countXMas(matrix [][]rune) int {
	count := 0

	for row, line := range matrix {
		for col := range line {
			if isXMas(matrix, row, col) {
				count++
			}
		}
	}

	return count
}

func (s *Solution) Part2(input *input.Input) string {
	matrix := input.RuneMatrix()
	count := countXMas(matrix)
	return strconv.Itoa(count)
}
