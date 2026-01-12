package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
	"github.com/schollz/progressbar/v3"
)

func downloadOneFile(info *youtube.Video, format *youtube.Format, filenameSuffix string, outputDir string) (string, error) {
	stream, size, err := client.GetStream(info, format)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	fmt.Printf("%s Size: %.2f MB\n", filenameSuffix, float64(size)/1024/1024)

	safeTitle := sanitizeFileName(info.Title)
	fileName := fmt.Sprintf("%s_%s.mp4", safeTitle, filenameSuffix)
	fullPath := filepath.Join(outputDir, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bar := progressbar.DefaultBytes(
		size,
		"Downloading",
	)

	_, err = io.Copy(io.MultiWriter(file, bar), stream)
	if err != nil {
		return "", err
	}
	fmt.Println()

	return fullPath, nil
}
