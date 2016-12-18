package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"github.com/devender/zipcal/loader"
	"github.com/devender/zipcal/location"
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
	for _, loc := range sbLocations {
		log.Debug(loc)
	}

	return 0
}

