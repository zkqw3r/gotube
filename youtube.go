package main

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/kkdai/youtube/v2"
)

var client = youtube.Client{}

func selectVideoAndAudio(video *youtube.Video) (*youtube.Format, *youtube.Format, error) {
	video.Formats.Sort()
	var bestVideo *youtube.Format
	var bestAudio *youtube.Format

	for i := range video.Formats {
		f := &video.Formats[i]
		if f.AudioChannels == 0 && f.QualityLabel == "" {
			bestVideo = f
			break
		}
	}

	if bestVideo == nil {
		for i := range video.Formats {
			f := &video.Formats[i]
			if f.QualityLabel != "" {
				bestVideo = f
				break
			}
		}
	}

	if bestVideo == nil {
		return nil, nil, fmt.Errorf("No video streams found")
	}

	for i := range video.Formats {
		f := &video.Formats[i]
		if f.AudioChannels > 0 && f.QualityLabel == "" {
			bestAudio = f
			break
		}
	}

	if bestAudio == nil {
		for i := range video.Formats {
			f := &video.Formats[i]
			if f.AudioChannels > 0 {
				bestAudio = f
				break
			}
		}
	}

	if bestAudio == nil {
		return nil, nil, fmt.Errorf("No audio streams found")
	}

	return bestVideo, bestAudio, nil
}

func getInfo(url string) (*youtube.Video, error) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Get info..."
	s.Color("white", "bold")
	s.Start()

	video, err := client.GetVideo(url)
	if err != nil {
		return nil, fmt.Errorf("Video retrieval error:%w", err)
	}
	s.Stop()

	fmt.Printf("\nAuthor: %v\nVideo:  %v\nLength: %v\nViews:  %v\t\n\n", video.Author, video.Title, video.Duration, formatNumber(video.Views))
	return video, nil
}
