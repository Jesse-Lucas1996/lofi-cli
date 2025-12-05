# Lofi CLI

A simple command-line interface to play lofi music in your terminal, written in Go.

## Prerequisites

- Go 1.21 or higher
- MPV media player installed on your system

## Installation

1. Install MPV media player if you haven't already:
   ```bash
   # On Ubuntu/Debian
   sudo apt-get install mpv
   
   # On macOS
   brew install mpv
   
   # On Arch Linux
   sudo pacman -S mpv
   ```

2. Install the CLI tool globally:
   ```bash
   go install github.com/Jesse-Lucas1996/lofi-cli@latest
   ```

3. **Important**: Add Go's bin directory to your PATH (one-time setup):
   
   For **bash** users:
   ```bash
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
   source ~/.bashrc
   ```
   
   For **zsh** users:
   ```bash
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
   source ~/.zshrc
   ```
   
   For **fish** users:
   ```bash
   fish_add_path (go env GOPATH)/bin
   ```

4. Verify installation:
   ```bash
   lofi-cli list
   # or use the alias
   lofi list
   ```

### Updating to Latest Version

If you've already installed lofi-cli and want to update to the latest version:

```bash
go clean -modcache
go install github.com/Jesse-Lucas1996/lofi-cli@latest

git clone https://github.com/Jesse-Lucas1996/lofi-cli.git
cd lofi-cli
go build -o $(go env GOPATH)/bin/lofi-cli
```

## Usage

To play lofi music:
```bash
lofi play
```

To play a specific station:
```bash
lofi play --station lofi-girl
```

To list all available stations:
```bash
lofi list
```

## Available Stations

- **lofi-girl**: Classic lofi beats radio
- **chilled-cow**: Relaxing lofi music stream
- **boxlofi**: Box Lofi radio stream
- **the-chillout**: The Chillout station
- **smooth-jazz**: Smooth jazz radio
- **deep-focus**: Radio Paradise Mellow (perfect for focus)
- **groove-salad**: Ambient/downtempo grooves (SomaFM)
- **drone-zone**: Atmospheric ambient music (SomaFM)
- **deep-space**: Deep space ambient sounds (SomaFM)
- **defcon**: Hacker/cyber ambient music (SomaFM)

## Controls

- Press `Ctrl+C` to stop playback

## Features

- Simple and intuitive CLI interface
- Multiple lofi stations to choose from
- Clean and efficient Go implementation
- Easy installation and usage

## License

MIT License 