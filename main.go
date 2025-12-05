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
	station string
	rootCmd = &cobra.Command{
		Use:   "lofi",
		Short: "A CLI tool to play lofi music in your terminal",
		Long:  `A simple command-line interface to play lofi music streams in your terminal.`,
	}

	playCmd = &cobra.Command{
		Use:   "play",
		Short: "Play lofi music",
		Run: func(cmd *cobra.Command, args []string) {
			playLofi(station)
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
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(listCmd)
}

func playLofi(stationName string) {
	url := GetStationURL(stationName)
	
	fmt.Printf("🎵 Now playing: %s\n", stationName)
	fmt.Println("Press Ctrl+C to stop")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	cmd := exec.Command("mpv", "--no-video", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting playback: %v\n", err)
		os.Exit(1)
	}

	<-sigChan
	fmt.Println("\nStopping playback...")
	cmd.Process.Signal(syscall.SIGTERM)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
