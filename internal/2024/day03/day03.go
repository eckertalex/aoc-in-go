package day03

import (
	"fmt"
	"regexp"
	"strconv"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

type Solution struct{}

func New() *Solution {
	return &Solution{}
}

func mulitply(text string) int {
	r := regexp.MustCompile(`\d{1,3}`)
	matches := r.FindAllString(text, 2)

	x, err := strconv.Atoi(matches[0])
	if err != nil {
		fmt.Println(err)
	}
	y, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println(err)
	}

	return x * y
}

func (s *Solution) Part1(input *input.Input) string {
	text, err := input.Text()
	if err != nil {
		fmt.Println(err)
	}

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := r.FindAllString(text, -1)

	sum := 0
	for _, match := range matches {
		product := mulitply(match)
		sum += product
	}

	return strconv.Itoa(sum)
}

func (s *Solution) Part2(input *input.Input) string {
	text, err := input.Text()
	if err != nil {
		fmt.Println(err)
	}

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	matches := r.FindAllString(text, -1)

	sum := 0
	do := true
	for _, match := range matches {
		if match[0:3] == "mul" && do {
			product := mulitply(match)
			sum += product
		} else {
			do = match == "do()"
		}
	}

	return strconv.Itoa(sum)
}
