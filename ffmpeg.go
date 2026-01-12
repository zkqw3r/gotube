package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func mergeFiles(videoPath, audioPath, outputPath string) error {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Combining video and audio with FFmpeg..."
	s.Color("white", "bold")
	s.Start()

	cmd := exec.Command("ffmpeg",
		"-i", videoPath,
		"-i", audioPath,
		"-c:v", "copy",
		"-c:a", "aac",
		"-strict", "experimental",
		"-y", outputPath,
	)

	output, err := cmd.CombinedOutput()
	s.Stop()
	if err != nil {
		return fmt.Errorf("Error ffmpeg: %v\nLog:\n%s", err, string(output))
	}

	os.Remove(videoPath)
	os.Remove(audioPath)

	return nil
}
