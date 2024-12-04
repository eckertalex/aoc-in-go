package main

import (
	"os"

	"github.com/eckertalex/aoc-in-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
