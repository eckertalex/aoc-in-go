package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	input "github.com/eckertalex/aoc-in-go/internal"
)

func mulitply(text string) int {
	r := regexp.MustCompile(`\d{1,3}`)
	matches := r.FindAllString(text, 2)

	x, err := strconv.Atoi(matches[0])
	if err != nil {
		log.Println(err)
	}
	y, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Println(err)
	}

	return x * y
}

func part1(input *input.Input) int {
	text, err := input.Text()
	if err != nil {
		log.Println(err)
	}

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := r.FindAllString(text, -1)

	sum := 0
	for _, match := range matches {
		product := mulitply(match)
		sum += product
	}

	return sum
}

func part2(input *input.Input) int {
	text, err := input.Text()
	if err != nil {
		log.Println(err)
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

	return sum
}

func main() {
	fmt.Println(part1(input.FromFile()))
	fmt.Println(part2(input.FromFile()))
}
