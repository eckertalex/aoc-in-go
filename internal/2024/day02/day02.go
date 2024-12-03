package day02

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

type Solution struct{}

func New() *Solution {
	return &Solution{}
}

func splitLinesToInt2D(input *input.Input) [][]int {
	var int2D [][]int

	for line := range input.Lines() {
		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error parsing numbers on line:", line)
				continue
			}
			row[i] = n
		}
		int2D = append(int2D, row)
	}

	return int2D
}

func isSafeReport(report []int) bool {
	isInc, isDec := true, true
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff > 0 {
			isDec = false
		} else if diff < 0 {
			isInc = false
		} else {
			isInc, isDec = false, false
		}
	}

	return isInc || isDec
}

func remove(slice []int, index int) []int {
	cs := make([]int, len(slice))
	copy(cs, slice)

	return append(cs[:index], cs[index+1:]...)
}

func (s *Solution) Part1(input *input.Input) string {
	reports := splitLinesToInt2D(input)

	safeCount := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

func (s *Solution) Part2(input *input.Input) string {
	reports := splitLinesToInt2D(input)

	safeCount := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			rs := remove(report, i)
			if isSafeReport(rs) {
				safeCount++
				break
			}
		}
	}

	return strconv.Itoa(safeCount)
}
