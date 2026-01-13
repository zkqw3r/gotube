package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func downloadOneFile(p *mpb.Progress, info *youtube.Video, format *youtube.Format, filenameSuffix string, outputDir string) (string, error) {
	stream, size, err := client.GetStream(info, format)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	safeTitle := sanitizeFileName(info.Title)
	fileName := fmt.Sprintf("%s_%s.mp4", safeTitle, filenameSuffix)
	fullPath := filepath.Join(outputDir, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bar := p.AddBar(size,
		mpb.PrependDecorators(
			decor.Name(filenameSuffix, decor.WC{W: len(filenameSuffix) + 1, C: decor.DindentRight}),
			decor.CountersKibiByte("% .2f / % .2f"),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_GO, 90),
			decor.Name(" ] "),
			decor.Percentage(decor.WCSyncSpace),
		),
	)

	proxyReader := bar.ProxyReader(stream)
	defer proxyReader.Close()

	_, err = io.Copy(file, proxyReader)
	if err != nil {
		return "", err
	}
	return fullPath, nil
}
