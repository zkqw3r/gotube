package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getUrl() (string, error) {
	var url string
	fmt.Print("Enter the video link: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return url, fmt.Errorf("Input error")
	}
	url = strings.TrimSpace(scanner.Text())
	return url, nil
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
