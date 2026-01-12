package main

import (
	"fmt"

	"github.com/kkdai/youtube/v2"
)

var client = youtube.Client{}

func selectVideoAndAudio(video *youtube.Video) (*youtube.Format, *youtube.Format, error) {
	var bestVideo *youtube.Format
	var bestAudio *youtube.Format

	// Search tag for video
	for i := range video.Formats {
		if video.Formats[i].ItagNo == 399 {
			bestVideo = &video.Formats[i]
			break
		}
	}
	if bestVideo == nil {
		for i := range video.Formats {
			if video.Formats[i].ItagNo == 248 {
				bestVideo = &video.Formats[i]
				break
			}
		}
	}
	if bestVideo == nil {
		for i := range video.Formats {
			if video.Formats[i].ItagNo == 137 {
				bestVideo = &video.Formats[i]
				break
			}
		}
	}
	if bestVideo == nil {
		return nil, nil, fmt.Errorf("Couldn't find a suitable video in normal quality")
	}

	// Search tag for audio
	for i := range video.Formats {
		if video.Formats[i].ItagNo == 251 {
			bestAudio = &video.Formats[i]
			break
		}
	}
	if bestAudio == nil {
		for i := range video.Formats {
			if video.Formats[i].ItagNo == 140 {
				bestAudio = &video.Formats[i]
				break
			}
		}
	}
	if bestAudio == nil {
		return nil, nil, fmt.Errorf("Couldn't find the audio")
	}

	return bestVideo, bestAudio, nil
}

func getInfo(url string) (*youtube.Video, error) {
	video, err := client.GetVideo(url)
	if err != nil {
		return nil, fmt.Errorf("Video retrieval error:%w", err)
	}
	return video, nil
}
