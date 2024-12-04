package cmd

import (
	"fmt"
	"time"

	"github.com/eckertalex/aoc-in-go/internal/2024/day01"
	"github.com/eckertalex/aoc-in-go/internal/2024/day02"
	"github.com/eckertalex/aoc-in-go/internal/2024/day03"
	input "github.com/eckertalex/aoc-in-go/internal/input"
	"github.com/eckertalex/aoc-in-go/internal/util"
	"github.com/spf13/cobra"
)

type Solution interface {
	Part1(input *input.Input) string
	Part2(input *input.Input) string
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specific Advent of Code challenge",
	RunE:  runChallenge,
}

func loadSolution(year, day int) (Solution, error) {
	switch year {
	case 2024:
		switch day {
		case 1:
			return day01.New(), nil
		case 2:
			return day02.New(), nil
		case 3:
			return day03.New(), nil
		default:
			return nil, fmt.Errorf("no solution found for year %d, day %d", year, day)
		}
	default:
		return nil, fmt.Errorf("no solutions found for year %d", year)
	}
}

func runChallenge(cmd *cobra.Command, args []string) error {
	solution, err := loadSolution(year, day)
	if err != nil {
		return fmt.Errorf("failed to load solution: %v", err)
	}

	inputData, err := input.FromFile(year, day)
	if err != nil {
		return fmt.Errorf("failed to load input: %v", err)
	}

	var result string
	var startTime time.Time
	switch part {
	case 1:
		startTime = time.Now()
		result = solution.Part1(inputData)
	case 2:
		startTime = time.Now()
		result = solution.Part2(inputData)
	default:
		return fmt.Errorf("invalid part number: %d (must be 1 or 2)", part)
	}
	elapsedTime := time.Since(startTime)

	fmt.Printf("%s in %v\n", result, elapsedTime)

	err = util.CopyToClipboard(result)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}
