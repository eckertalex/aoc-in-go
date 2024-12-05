package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Input struct {
	scanner *bufio.Scanner
	lines   chan string
}

func FromFile(year, day int) (*Input, error) {
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("failed to determine input path")
	}

	baseDir := filepath.Dir(filepath.Dir(callerFile))
	path := filepath.Join(baseDir, "internal", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day), "input.txt")

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file %s: %w", path, err)
	}

	return newInputFromReader(file, file), nil
}

func FromLiteral(input string) *Input {
	return newInputFromReader(strings.NewReader(input), nil)
}

func newInputFromReader(reader io.Reader, closer io.Closer) *Input {
	input := &Input{
		scanner: bufio.NewScanner(reader),
		lines:   make(chan string),
	}

	go func() {
		defer func() {
			if closer != nil {
				if err := closer.Close(); err != nil {
					fmt.Fprintf(os.Stderr, "error closing reader: %v\n", err)
				}
			}
		}()
		for input.scanner.Scan() {
			input.lines <- input.scanner.Text()
		}
		if err := input.scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "scanner error: %v\n", err)
		}
		close(input.lines)
	}()
	return input
}

func (i *Input) Lines() <-chan string {
	return i.lines
}

func (i *Input) LinesSlice() []string {
	var lines []string
	for line := range i.Lines() {
		lines = append(lines, line)
	}
	return lines
}

func (i *Input) Text() string {
	lines := i.LinesSlice()
	return strings.Join(lines, "\n")
}

func (i *Input) Ints() <-chan int {
	ints := make(chan int)
	go func() {
		defer close(ints)
		for line := range i.lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			value, err := strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error converting line to int: %v\n", err)
				continue
			}
			ints <- value
		}
	}()
	return ints
}

func (i *Input) IntsSlice() []int {
	intsChan := i.Ints()
	var ints []int
	for value := range intsChan {
		ints = append(ints, value)
	}
	return ints
}

func (i *Input) RuneMatrix() [][]rune {
	var runeMatrix [][]rune

	for line := range i.Lines() {
		runeMatrix = append(runeMatrix, []rune(line))
	}

	return runeMatrix
}
