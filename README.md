# 🎥 GoTube

[![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![FFmpeg](https://img.shields.io/badge/FFmpeg-Required-00599C?logo=ffmpeg&logoColor=white)](https://ffmpeg.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

<div align="center">
  <h3>⚡ High-performance YouTube Downloader CLI</h3>
  <img src="demonstration.gif" width="100%" alt="video example of work">
  <p>
    <i>Download highest quality video (1080p/4K) + best audio in parallel<br>
    Automatically merged into MP4 via FFmpeg</i>
  </p>
</div>

***

## 📥 Downloads (Releases)

Don't want to build from source? Download the latest ready-to-use binary for your OS:

👉 **[Go to Releases Page](https://github.com/zkqw3r/Gotube/releases/latest)**

| OS | File | Instructions |
|----|------|--------------|
| **Windows** | `gotube-windows-amd64.zip` | Just run `gotube.exe` |
| **Linux** | `gotube-linux-amd64.tar.gz` | `chmod +x gotube` then `./gotube` |
| **macOS** | `gotube-macos-*.zip` | Right-click app → Open (to bypass gatekeeper) |

***

## ✨ Features

- **🚀 Parallel Downloads** — Uses Goroutines to download video & audio streams simultaneously
- **🎯 Max Quality** — Automatically fetches the best video (1080p, 2K, 4K) and audio tracks
- **🔊 Smart Merge** — Uses FFmpeg to combine streams instantly without quality loss
- **📊 Visual Progress** — Clean multi-bar progress interface

## 🛠️ Tech Stack

- **Go (Golang)** — Concurrency & Core logic
- **FFmpeg** — Media merging engine
- **vbauerster/mpb** — Complex CLI progress bars
- **kkdai/youtube** — YouTube internal API wrapper

## 🚀 Build from Source

If you prefer to compile it yourself:

### Prerequisites
- **Go** 1.20+
- **FFmpeg** installed and in system `PATH`

### 1. Clone & Init
```bash
git clone https://github.com/zkqw3r/Gotube.git
cd Gotube
go mod tidy
```

### 2. Build

```bash
go build -o gotube .
```

## 🎮 Usage

Simply run the binary and follow the instructions:

```bash
./gotube
```
***

<div align="center">
  <sub>Built with ❤️ by <b>zkqw3r</b></sub>
</div>
