package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/vbauerster/mpb/v8"
	"golang.org/x/sync/errgroup"
)

func main() {
	outputDir := "Videos"
	logo := `
  ________     ___________   ___.           
 ╱  _____╱  ___╲__    ___╱_ _╲_ │__   ____  
╱   ╲  ___ ╱  _ ╲│    │ │  │  ╲ __ ╲_╱ __ ╲ 
╲    ╲_╲  (  <_> )    │ │  │  ╱ ╲_╲ ╲  ___╱ 
 ╲______  ╱╲____╱│____│ │____╱│___  ╱╲___  >
        ╲╱                        ╲╱     ╲╱ 
	`

	for {
		callClear()
		fmt.Println(logo)

		fmt.Print("Enter video URL (or 'q' to quit): ")
		url, err := getText()
		if url == "q" {
			fmt.Println("Goodbye!")
			return
		} else if url == "" {
			continue
		}
		if hasError(err) {
			continue
		}
		if err := ensureDir(outputDir); hasError(err) {
			fmt.Println("Press Enter to try again...")
			getText()
			continue
		}

		info, err := getInfo(url)
		if hasError(err) {
			continue
		}

		bestVideo, bestAudio, err := selectVideoAndAudio(info)
		if hasError(err) {
			continue
		}

		var videoPath, audioPath string
		var g errgroup.Group
		p := mpb.New(mpb.WithWidth(64), mpb.WithRefreshRate(180*time.Millisecond))

		// video
		g.Go(func() error {
			path, err := downloadOneFile(p, info, bestVideo, "Video only", outputDir)
			if err == nil {
				videoPath = path
			}
			return err
		})

		// audio
		g.Go(func() error {
			path, err := downloadOneFile(p, info, bestAudio, "Audio only", outputDir)
			if err == nil {
				audioPath = path
			}
			return err
		})

		if err := g.Wait(); err != nil {
			fmt.Printf("Download error: %v\n", err)
			fmt.Println("Please try again...")
			continue
		}

		p.Wait()
		fmt.Println()

		fileName := fmt.Sprintf("%v.mp4", sanitizeFileName(info.Title))
		outputPath := filepath.Join(outputDir, fileName)

		err = mergeFiles(videoPath, audioPath, outputPath)
		if hasError(err) {
			continue
		}

		fmt.Printf("Successfully downloaded: %v\n", sanitizeFileName(info.Title))
		answer, err := promptChoice("Download another video? (y/n): ", []string{"y", "n", "yes", "no"})
		if hasError(err) {
			continue
		}
		if answer == "n" || answer == "no" {
			fmt.Println("Goodbye!")
			return
		}
		continue
	}

}
