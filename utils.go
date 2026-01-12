package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
