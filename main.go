package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	outputDir := "Videos"

	url, err := getUrl()
	check(err)

	err = ensureDir(outputDir)
	check(err)

	fmt.Println("Get info...")
	info, err := getInfo(url)
	check(err)
	fmt.Printf("Info received\n\n")

	bestVideo, bestAudio, err := selectVideoAndAudio(info)
	check(err)

	videoPath, err := downloadOneFile(info, bestVideo, "video_only", outputDir)
	check(err)

	audioPath, err := downloadOneFile(info, bestAudio, "audio_only", outputDir)
	check(err)

	fileName := fmt.Sprintf("%v.mp4", sanitizeFileName(info.Title))
	outputPath := filepath.Join(outputDir, fileName)

	err = mergeFiles(videoPath, audioPath, outputPath)
	check(err)

	fmt.Printf("Successfully downloaded %v\n", outputPath)
}
