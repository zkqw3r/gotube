package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func hasError(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return true
	}
	return false
}

func getText() (string, error) {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return text, fmt.Errorf("Input error")
	}
	text = strings.TrimSpace(scanner.Text())
	return text, nil
}

func sanitizeFileName(name string) string {
	symbols := `\/:*?"<>|`
	for _, char := range symbols {
		name = strings.ReplaceAll(name, string(char), "_")
	}
	if len(name) > 200 {
		name = name[:200]
	}
	return name
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory %s: %w", dirName, err)
	}
	return nil
}

func callClear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func promptChoice(question string, validOptions []string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(question)
		if !scanner.Scan() {
			return "", fmt.Errorf("Input error")
		}
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		for _, opt := range validOptions {
			if input == opt {
				return input, nil
			}
		}

		fmt.Printf("Invalid input. Please enter 'y' or 'no'\n\n")
	}
}
