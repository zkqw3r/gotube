# 🎥 Gotube

[![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![FFmpeg](https://img.shields.io/badge/FFmpeg-Required-00599C?logo=ffmpeg&logoColor=white)](https://ffmpeg.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

<div align="center">
  <h3>⚡ High-quality YouTube Downloader CLI</h3>
  <img src="demo.gif" width="100%" alt="video example of work">
  <i>Download 1080p+ video and best audio, automatically merged via FFmpeg.</i>
</div>

## ✨ Features

- **🎯 Best Quality** - Automatically selects the best video stream (1080p, 2K, 4K)
- **🔊 Crystal Audio** - Prioritizes Opus/AAC high-bitrate audio
- **🔗 Smart Merge** - Uses FFmpeg to combine streams without re-encoding (copy codec)

## 🛠️ Tech Stack

- **Go (Golang)** - Core logic
- **kkdai/youtube** - YouTube internal API wrapper
- **FFmpeg** - Media processing engine
- **os/exec** - External process management

## 📂 Project Structure

| File | Description |
|------|-------------|
| `main.go` | Entry point, error handling, and high-level logic |
| `youtube.go` | Wrapper for YouTube API |
| `downloader.go` | Stream downloading logic |
| `ffmpeg.go` | Media merging commands |
| `utils.go` | Helper functions |

## 🚀 Installation

### Prerequisites
- **Go** 1.20+ installed
- **FFmpeg** installed and added to system PATH

### 1. Clone the repository
```bash
git clone https://github.com/zkqw3r/Gotube
cd Gotube
```

### 2. Initialize modules
```bash
go mod tidy
```

### 3. Build the binary
```bash
go build -o gotube .
```

## 🎮 Usage

Simply run the executable:

```bash
./gotube
# or run directly with Go:
go run .
```

## 📝 To-Do List

- [x] Add progress bar (TUI)
- [ ] Command line arguments (flags)

---

<div align="center">
  <sub>Built with ❤️ by zkqw3r</sub>
</div>

