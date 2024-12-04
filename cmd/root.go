package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var (
	year int
	day  int
	part int

	rootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "Advent of Code CLI tool",
		Long:  `A CLI tool to run Advent of Code solutions for specific years, days, and parts`,
	}
)

func getCurrentYear() int {
	return time.Now().Year()
}

func getCurrentDay() int {
	now := time.Now()
	if now.Month() == time.December {
		return now.Day()
	}
	return 1
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", getCurrentYear(), "Advent of Code year (defaults to current year)")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", getCurrentDay(), "Day of the challenge (defaults to current day in December, or 1)")
	rootCmd.PersistentFlags().IntVarP(&part, "part", "p", 1, "Part of the challenge (1 or 2)")

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(initCmd)
}
