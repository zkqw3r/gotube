package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	url, err := getUrl()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Get info...")

	info, err := getInfo(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}

	safeTitle := strings.ReplaceAll(info.Title, "/", "_")
	safeTitle = strings.ReplaceAll(safeTitle, "\\", "_")
	fmt.Printf("Info received\n\n")

	bestVideo, bestAudio, err := selectVideoAndAudio(info)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}
	videoPath, err := downloadOneFile(info, bestVideo, "video_only")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}
	audioPath, err := downloadOneFile(info, bestAudio, "audio_only")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}
	outputPath := fmt.Sprintf("%v.mp4", safeTitle)

	err = mergeFiles(videoPath, audioPath, outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully downloaded %v\n", outputPath)
}
