package main

import (
	"fmt"

	"github.com/kkdai/youtube/v2"
)

var client = youtube.Client{}

func findFormat(formats []youtube.Format, itag int) *youtube.Format {
	for i := range formats {
		if formats[i].ItagNo == itag {
			return &formats[i]
		}
	}
	return nil
}

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
	video, err := client.GetVideo(url)
	if err != nil {
		return nil, fmt.Errorf("Video retrieval error:%w", err)
	}
	return video, nil
}
