# Lofi CLI

A simple command-line interface to play lofi music in your terminal, written in Go.

## Prerequisites

- Go 1.21 or higher
- MPV media player installed on your system

## Installation

1. Install MPV media player if you haven't already:
   ```bash
   sudo apt-get install mpv
   ```

2. Install the CLI tool:
   ```bash
   go install github.com/Jesse-Lucas1996/lofi-cli@latest
   ```

   Or build from source:
   ```bash
   git clone https://github.com/Jesse-Lucas1996/lofi-cli.git
   cd lofi-cli
   go build
   sudo mv lofi-cli /usr/local/bin/lofi
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

- lofi-girl: Classic lofi beats
- chilled-cow: Relaxing lofi music

## Controls

- Press `Ctrl+C` to stop playback

## Features

- Simple and intuitive CLI interface
- Multiple lofi stations to choose from
- Clean and efficient Go implementation
- Easy installation and usage

## License

MIT License 