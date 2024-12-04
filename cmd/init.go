package cmd

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//go:embed "templates"
var templateFS embed.FS

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Advent of Code challenge for a given year and day",
	RunE:  initChallenge,
}

func initChallenge(cmd *cobra.Command, args []string) error {
	initFiles(day, year)
	return nil
}

func initFiles(day, year int) {
	dir := fmt.Sprintf("internal/%d/day%02d", year, day)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	err = initDayFile(dir, fmt.Sprintf("%02d", day))
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	err = initTestFile(dir, fmt.Sprintf("%02d", day))
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	err = initInputFile(dir)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	// modifyExistingFile("main.go", year, day)
}

func initInputFile(dir string) error {
	inputFilePath := filepath.Join(dir, "input.txt")

	if _, err := os.Stat(inputFilePath); !os.IsNotExist(err) {
		return fmt.Errorf("file %s already exists", inputFilePath)
	}

	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer inputFile.Close()

	return nil
}

func initDayFile(dir, day string) error {
	dayFilePath := filepath.Join(dir, fmt.Sprintf("day%s.go", day))

	if _, err := os.Stat(dayFilePath); !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("file %s already exists", dayFilePath)
	}

	dayFile, err := os.Create(dayFilePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer dayFile.Close()

	tmpl, err := template.New("day").ParseFS(templateFS, "templates/day.tmpl")
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	data := map[string]any{
		"day": day,
	}

	err = tmpl.ExecuteTemplate(dayFile, "day", data)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}

func initTestFile(dir, day string) error {
	testFilePath := filepath.Join(dir, fmt.Sprintf("day%s_test.go", day))

	if _, err := os.Stat(testFilePath); !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("file %s already exists", testFilePath)
	}

	testFile, err := os.Create(testFilePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer testFile.Close()

	tmpl, err := template.New("day").ParseFS(templateFS, "templates/day.tmpl")
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	data := map[string]any{
		"day": day,
	}

	err = tmpl.ExecuteTemplate(testFile, "test", data)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}

// func modifyExistingFile(filePath string, year, day int) {
// 	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
//
// 	_, err = fmt.Fprintf(file, "\n// Initialized solution for Year %d, Day %d\n", year, day)
// 	if err != nil {
// 		fmt.Println("Error writing to file:", err)
// 	}
// }

func init() {
	rootCmd.AddCommand(initCmd)
}
