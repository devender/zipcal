package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"github.com/devender/zipcal/loader"
	"github.com/devender/zipcal/location"
	"fmt"
	"bytes"
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

	/*
	home := location.Point{
		Longitude: -118.795078,
		Latitude: 34.271988,
	}

	filt := location.FilterByDistanceFromPoint(home, 10, sbLocations)

	for _, a := range filt {
		fmt.Printf("%v\n", a)
	}
	*/
	var b bytes.Buffer
	fmt.Fprintf(&b, "INSERT INTO insikt.geolocation_location (name,address1,city,state,postal_code,country,latitude,longitude,metadata,enabled) VALUES \n");
	for _, loc := range sbLocations {
		if loc.Country == "US" {
			fmt.Fprintf(&b, "%s\n,", loc.Sql())
		}
	}
	b.Truncate(b.Len()-1)
	fmt.Fprintf(&b, ";\n")
	b.WriteTo(os.Stdout)
	return 0
}

