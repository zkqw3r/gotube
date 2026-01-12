package main

import (
	"fmt"
	"os"
	"os/exec"
)

func mergeFiles(videoPath, audioPath, outputPath string) error {
	fmt.Printf("\nCombining using ffmpeg...\n")

	cmd := exec.Command("ffmpeg",
		"-i", videoPath,
		"-i", audioPath,
		"-c:v", "copy",
		"-c:a", "aac",
		"-strict", "experimental",
		"-y", outputPath,
	)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error ffmpeg: %v. Ensure that ffmpeg is installed and added to PATH.", err)
	}
	os.Remove(videoPath)
	os.Remove(audioPath)

	return nil
}
