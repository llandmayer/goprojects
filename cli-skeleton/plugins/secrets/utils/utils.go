package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func TerragruntRootPath() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %w", err)
	}

	for {
		terragruntDir := filepath.Join(root, "terragrunt")

		if s, err := os.Stat(terragruntDir); err == nil && s.IsDir() {
			return terragruntDir, nil
		}

		parent := filepath.Dir(root)
		if parent == root {
			break
		}
		root = parent
	}

	return "", fmt.Errorf("terragrunt directory not found, please run this command from the right directory")
}

func LS(path string) ([]string, error) {
	var dirs []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, nil
}

// Does a find file based on the rootPath
// Return the absolute path of the file
func FindFilePath(rootPath string, filename string) ([]string, error) {
	var foundPath []string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && info.Name() == filename {
			path, _ := filepath.Abs(path)

			foundPath = append(foundPath, path)

			return filepath.SkipDir
		}
		return nil
	})
	if len(foundPath) == 0 {
		return nil, fmt.Errorf("file not found: %s, %v", filename, err)
	}
	return foundPath, nil
}

func ExistInSlice(slice []string, elementString string) bool {
	for _, element := range slice {
		if element == elementString {
			return true
		}
	}
	return false
}

type Table struct {
	Content [][]string
}

// AddRow appends a new row of data to the table content
func (t *Table) AddRow(row []string) {
	t.Content = append(t.Content, row)
}

// Initiabtes a new table
// Expects headers on initialisation
func NewTable(headers []string) *Table {
	return &Table{Content: [][]string{headers}}
}

// Print prints the table in a formatted way
func (t *Table) Print() {
	columnLengths := make([]int, len(t.Content[0]))
	for _, line := range t.Content {
		for i, val := range line {
			if len(val) > columnLengths[i] {
				columnLengths[i] = len(val)
			}
		}
	}

	var lineLength int
	for _, c := range columnLengths {
		lineLength += c + 3
	}
	lineLength += 1
	for i, line := range t.Content {
		if i == 0 { // table header
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2))
		}
		for j, val := range line {
			fmt.Printf("| %-*s ", columnLengths[j], val)
			if j == len(line)-1 {
				fmt.Printf("|\n")
			}
		}
		if i == 0 || i == len(t.Content)-1 {
			fmt.Printf("+%s+\n", strings.Repeat("-", lineLength-2))
		}
	}
}

func OpenWithEditor(filePath string) error {
	cmd := exec.Command("vim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	return nil
}

func CreateFile(filePath string, content []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

func ReadFromCli(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", question)
	text, err := reader.ReadString('\n')
	fmt.Println(text)

	if err != nil {
		return "", fmt.Errorf("failed to get string from buffer: %v", err)
	}

	return text, err
}
