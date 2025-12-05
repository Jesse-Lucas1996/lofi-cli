package main

import (
	"fmt"
	"os"
)

var stations = map[string]string{
	"lofi-girl":    "https://play.streamafrica.net/lofiradio",
	"chilled-cow":  "https://stream.zeno.fm/0r0xa792kwzuv",
	"boxlofi":      "https://boxradio-edge-00.streamafrica.net/lofi",
	"the-chillout": "https://streams.radiomast.io/chill-out",
	"smooth-jazz":  "https://mediagw.e-tiger.net/stream/zc07",
	"deep-focus":   "https://stream-uk1.radioparadise.com/mellow-320",
	"groove-salad": "https://ice1.somafm.com/groovesalad-128-mp3",
	"drone-zone":   "https://ice1.somafm.com/dronezone-128-mp3",
	"deep-space":   "https://ice1.somafm.com/deepspaceone-128-mp3",
	"defcon":       "https://ice1.somafm.com/defcon-128-mp3",
}

func ListStations() {
	fmt.Println("\nAvailable Lofi Stations:")
	for name := range stations {
		fmt.Printf("• %s\n", name)
	}
}

func GetStationURL(stationName string) string {
	url, exists := stations[stationName]
	if !exists {
		fmt.Printf("Error: Station '%s' not found\n", stationName)
		os.Exit(1)
	}
	return url
}
