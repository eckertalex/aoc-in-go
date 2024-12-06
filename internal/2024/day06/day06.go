package day06

import (
	"strconv"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

type Solution struct{}

func New() *Solution {
	return &Solution{}
}

var directionOrder = []rune{'^', '>', 'v', '<'}

var directionMap = map[rune][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func rotateClockwise(index int) int {
	return (index + 1) % len(directionOrder)
}

func findInitialPosAndDir(matrix [][]rune) ([2]int, int) {
	var pos [2]int
	var dirIndex int

	for i, row := range matrix {
		for j, col := range row {
			if _, exists := directionMap[col]; exists {
				pos = [2]int{i, j}
				for index, sym := range directionOrder {
					if sym == col {
						dirIndex = index
						break
					}
				}
			}
		}
	}

	return pos, dirIndex
}

func walk(matrix [][]rune, visited map[[2]int]int, guardPos [2]int, dirIndex int) bool {
	currentPos := guardPos
	currentDirIndex := dirIndex

	for {
		dir := directionMap[directionOrder[currentDirIndex]]
		x := currentPos[0] + dir[0]
		y := currentPos[1] + dir[1]

		if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) {
			return false
		}

		if matrix[x][y] == '#' {
			currentDirIndex = rotateClockwise(currentDirIndex)
			continue
		}

		currentPos = [2]int{x, y}
		visited[currentPos]++

		if visited[currentPos] > 4 {
			return true
		}
	}
}

func (s *Solution) Part1(input *input.Input) string {
	matrix := input.RuneMatrix()

	guardPos, dirIndex := findInitialPosAndDir(matrix)

	visited := make(map[[2]int]int)
	visited[guardPos] = 1

	walk(matrix, visited, guardPos, dirIndex)

	return strconv.Itoa(len(visited))
}

func (s *Solution) Part2(input *input.Input) string {
	matrix := input.RuneMatrix()

	guardPos, dirIndex := findInitialPosAndDir(matrix)

	visited := make(map[[2]int]int)
	visited[guardPos] = 1

	walk(matrix, visited, guardPos, dirIndex)

	circularObstacles := 0
	for visitedPos := range visited {
		seen := make(map[[2]int]int)
		seen[guardPos] = 1

		modifiedMatrix := make([][]rune, len(matrix))
		for row := range matrix {
			modifiedMatrix[row] = make([]rune, len(matrix[row]))
			copy(modifiedMatrix[row], matrix[row])
		}
		modifiedMatrix[visitedPos[0]][visitedPos[1]] = '#'

		if walk(modifiedMatrix, seen, guardPos, dirIndex) {
			circularObstacles++
		}
	}

	return strconv.Itoa(circularObstacles)
}
