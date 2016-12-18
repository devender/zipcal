package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"github.com/devender/zipcal/loader"
	"github.com/devender/zipcal/location"
	"fmt"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	os.Exit(realMain())
}

func loadStarbucksLocations() []location.Location {
	sb := loader.Starbucks{Source: "/Users/dgollapally/Downloads/data/starbucks.csv"}
	return sb.Load()
}

func realMain() int {

	sbLocations := loadStarbucksLocations()

	home := location.Point{
		Longitude: -118.795078,
		Latitude: 34.271988,
	}

	filt := location.FilterByDistanceFromPoint(home, 10, sbLocations)

	for _, a := range filt {
		fmt.Printf("%v\n", a)
	}

	return 0
}

