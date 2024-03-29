package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	input "github.com/eckertalex/aoc-in-go/internal"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitLines(input *input.Input) ([]int, []int, error) {
	var ls, rs []int
	for line := range input.Lines() {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			log.Println("Invalid line format:", line)
			continue
		}

		l, errl := strconv.Atoi(fields[0])
		r, errr := strconv.Atoi(fields[1])
		if errl != nil || errr != nil {
			log.Println("Error parsing numbers on line:", line)
			continue
		}

		ls = append(ls, l)
		rs = append(rs, r)
	}

	return ls, rs, nil
}

func part1(input *input.Input) int {
	ls, rs, err := splitLines(input)
	if err != nil {
		log.Println("error splitting lines")
		os.Exit(1)
	}

	sort.Ints(ls)
	sort.Ints(rs)

	length := min(len(ls), len(rs))
	xs := make([]int, length)
	for i := 0; i < length; i++ {
		xs[i] = abs(ls[i] - rs[i])
	}

	return sum(xs)
}

func part2(input *input.Input) int {
	ls, rs, err := splitLines(input)
	if err != nil {
		log.Println("error splitting lines")
		os.Exit(1)
	}

	var xs []int
	for _, l := range ls {
		count := 0
		for _, r := range rs {
			if l == r {
				count++
			}
		}
		xs = append(xs, l*count)
	}

	return sum(xs)
}

func main() {
	fmt.Println(part1(input.FromFile()))
	fmt.Println(part2(input.FromFile()))
}
