package input

import (
	"bufio"
	"io"
	"log"
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

func FromFile() *Input {
	_, caller, _, ok := runtime.Caller(1)
	if !ok {
		log.Println("Failed to determine input path")
		os.Exit(1)
	}

	path := filepath.Join(filepath.Dir(caller), "input.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Println("Failed to open file:", err)
		os.Exit(1)
	}

	return newInputFromReader(file, file)
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
					log.Printf("error closing reader: %v", err)
				}
			}
		}()

		for input.scanner.Scan() {
			input.lines <- input.scanner.Text()
		}

		if err := input.scanner.Err(); err != nil {
			log.Printf("scanner error: %v", err)
		}

		close(input.lines)
	}()

	return input
}

func (i *Input) Lines() <-chan string {
	return i.lines
}

func (i *Input) LinesSlice() ([]string, error) {
	var lines []string
	for line := range i.Lines() {
		lines = append(lines, line)
	}

	return lines, nil
}

func (i *Input) Text() (string, error) {
	lines, err := i.LinesSlice()
	if err != nil {
		return "", err
	}

	return strings.Join(lines, ""), nil
}

func (i *Input) Ints() (<-chan int, error) {
	ints := make(chan int)

	go func() {
		defer close(ints)

		for line := range i.lines {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Printf("error converting line to int: %v", err)
				continue
			}
			ints <- value
		}
	}()

	return ints, nil
}

func (i *Input) IntsSlice() ([]int, error) {
	intsChan, err := i.Ints()
	if err != nil {
		return nil, err
	}

	var ints []int
	for value := range intsChan {
		ints = append(ints, value)
	}

	return ints, nil
}
