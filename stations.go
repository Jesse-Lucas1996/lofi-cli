package main

import "fmt"

var stations = map[string]string{
	"lofi-girl":      "https://play.streamafrica.net/lofiradio",
	"chilled-cow":    "https://stream.zeno.fm/0r0xa792kwzuv",
	"lofi-synthwave": "https://stream.zeno.fm/f3wvbbqmdg8uv",
}

func ListStations() {
	fmt.Println("\nAvailable Lofi Stations:")
	for name := range stations {
		fmt.Printf("• %s\n", name)
	}
}

func GetStationURL(stationName string) (string, bool) {
	url, exists := stations[stationName]
	return url, exists
}
