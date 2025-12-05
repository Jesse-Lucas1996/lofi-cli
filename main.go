package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	station   string
	visualize bool
	rootCmd   = &cobra.Command{
		Use:   "lofi",
		Short: "A CLI tool to play lofi music in your terminal",
		Long:  `A simple command-line interface to play lofi music streams in your terminal.`,
	}

	playCmd = &cobra.Command{
		Use:   "play",
		Short: "Play lofi music",
		Run: func(cmd *cobra.Command, args []string) {
			playLofi(station, visualize)
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List available stations",
		Run: func(cmd *cobra.Command, args []string) {
			ListStations()
		},
	}
)

func init() {
	playCmd.Flags().StringVarP(&station, "station", "s", "lofi-girl", "Station to play (default: lofi-girl)")
	playCmd.Flags().BoolVarP(&visualize, "visualize", "v", false, "Show audio visualizer (requires cava)")
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(listCmd)
}

func playLofi(stationName string, showVisualizer bool) {
	url := GetStationURL(stationName)
	
	fmt.Printf("🎵 Now playing: %s\n", stationName)
	if showVisualizer {
		fmt.Println("🎨 Visualizer enabled")
	}
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var mpvCmd, cavaCmd *exec.Cmd

	if showVisualizer {
		if _, err := exec.LookPath("cava"); err != nil {
			fmt.Println("⚠️  Warning: cava not found. Install it with: sudo apt-get install cava")
			fmt.Println("Continuing without visualizer...")
			fmt.Println()
			showVisualizer = false
		}
	}

	if showVisualizer {
		mpvCmd = exec.Command("mpv", 
			"--no-video",
			"--really-quiet",
			url)
		
		cavaCmd = exec.Command("cava")
		cavaCmd.Stdin = os.Stdin
		cavaCmd.Stdout = os.Stdout
		cavaCmd.Stderr = os.Stderr
		
		if err := cavaCmd.Start(); err != nil {
			fmt.Printf("⚠️  Could not start cava: %v\n", err)
			fmt.Println("Continuing without visualizer...")
			fmt.Println()
			showVisualizer = false
		}
	}

	if !showVisualizer {
		mpvCmd = exec.Command("mpv", "--no-video", url)
		mpvCmd.Stdout = os.Stdout
		mpvCmd.Stderr = os.Stderr
	}

	if err := mpvCmd.Start(); err != nil {
		fmt.Printf("Error starting playback: %v\n", err)
		os.Exit(1)
	}

	<-sigChan
	fmt.Println("\n\nStopping playback...")
	mpvCmd.Process.Signal(syscall.SIGTERM)
	if cavaCmd != nil && cavaCmd.Process != nil {
		cavaCmd.Process.Signal(syscall.SIGTERM)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
