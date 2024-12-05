package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	input "github.com/eckertalex/aoc-in-go/internal/input"
)

type Solution struct{}

func New() *Solution {
	return &Solution{}
}

func splitRulesUpdates(input *input.Input) ([][]string, [][]string) {
	var rules [][]string
	var updates [][]string

	foundLineBreak := false
	for line := range input.Lines() {
		switch {
		case line == "":
			foundLineBreak = true
		case !foundLineBreak:
			rule := strings.Split(line, "|")
			rules = append(rules, rule)
		case foundLineBreak:
			pages := strings.Split(line, ",")
			updates = append(updates, pages)
		}
	}

	return rules, updates
}

func sortByRules(rules [][]string) func(a, b string) int {
	return func(a, b string) int {
		for _, rule := range rules {
			if rule[0] == a && rule[1] == b {
				return -1
			}
		}
		return 0
	}
}

func (s *Solution) Part1(input *input.Input) string {
	rules, updates := splitRulesUpdates(input)

	midSum := 0
	for _, update := range updates {
		correctlySorted := slices.IsSortedFunc(update, sortByRules(rules))
		if correctlySorted {
			mid := update[len(update)/2]
			n, err := strconv.Atoi(mid)
			if err != nil {
				fmt.Println(err)
				continue
			}

			midSum += n
		}
	}

	return strconv.Itoa(midSum)
}

func (s *Solution) Part2(input *input.Input) string {
	rules, updates := splitRulesUpdates(input)

	midSum := 0
	for _, update := range updates {
		correctlySorted := slices.IsSortedFunc(update, sortByRules(rules))
		if !correctlySorted {
			slices.SortFunc(update, sortByRules(rules))

			mid := update[len(update)/2]
			n, err := strconv.Atoi(mid)
			if err != nil {
				fmt.Println(err)
				continue
			}

			midSum += n
		}
	}

	return strconv.Itoa(midSum)
}
