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

   # On Windows
   reconsider your life
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

### Basic Commands

Play lofi music (default station):
```bash
lofi play
```

Play a specific station:
```bash
lofi play --station groove-salad
# or use the short flag
lofi play -s drone-zone
```

List all available stations:
```bash
lofi list
```

### 🎨 Visualizer Mode

Play with a beautiful audio visualizer (requires cava):
```bash
lofi play --visualize
# or use the short flag
lofi play -v -s deep-space
```

**To install cava:**
```bash
# Ubuntu/Debian
sudo apt-get install cava

# macOS
brew install cava

# Arch Linux
sudo pacman -S cava
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

- 🎵 Simple and intuitive CLI interface
- 📻 10 verified lofi/chill music stations
- 🎨 Optional audio visualizer mode (cava integration)
- ⚡ Clean and efficient Go implementation
- 🚀 Easy installation and usage
- 🎧 Multiple music genres: lofi, jazz, ambient, synthwave

## License

MIT License 