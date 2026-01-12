package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func downloadOneFile(info *youtube.Video, format *youtube.Format, filenameSuffix string) (string, error) {
	stream, size, err := client.GetStream(info, format)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	fmt.Printf("Downloading %s... Size: %.2f MB\n", filenameSuffix, float64(size)/1024/1024)
	safeTitle := strings.ReplaceAll(info.Title, "/", "_")
	safeTitle = strings.ReplaceAll(safeTitle, "\\", "_")
	fileName := fmt.Sprintf("%s_%s.mp4", safeTitle, filenameSuffix)

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
