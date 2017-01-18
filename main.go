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

	//test locations
	sbLocations = append(sbLocations, location.Location{
		Address1:    "2197 E. Los Angeles Avenue",
		City:        "Simi Valley",
		Name:        "Los Angeles & 1st",
		PostalCode:  "93065",
		Country:     "US",
		SubDivision: "CA",
		Point:       location.Point{Latitude: location.Angle(34.272064), Longitude:  location.Angle(-118.777687) },
		Metadata:    "[{\"key\":\"locationType\", \"value\":[\"payment\", \"payout\"]},{\"key\":\"services\",\"value\":[\"moneygram_payout\", \"moneygram_billpay\"]},{\"key\": \"canPayoutHighValue\" ,\"value\": \"true\"}]",
	})
	sbLocations = append(sbLocations, location.Location{
		Address1:    "3197 E. Los Angeles Avenue",
		City:        "Simi Valley",
		Name:        "Los Angeles & 1st",
		PostalCode:  "93065",
		Country:     "US",
		SubDivision: "CA",
		Point:       location.Point{Latitude: location.Angle(34.272064), Longitude:  location.Angle(-118.777687) },
		Metadata:    "[{\"key\":\"locationType\", \"value\":[\"partner\"]},{\"key\":\"provider\",\"value\":[\"dolex\"]},{\"key\": \"canPayoutHighValue\" ,\"value\": \"true\"}]",
	})

	var b bytes.Buffer
	fmt.Fprintf(&b, "INSERT INTO insikt.geolocation_location (name,address1,city,state,postal_code,country,latitude,longitude,metadata,enabled) VALUES \n");
	for _, loc := range sbLocations {
		if loc.Country == "US" {
			fmt.Fprintf(&b, "%s\n,", loc.Sql())
		}
	}

	b.Truncate(b.Len() - 1)
	fmt.Fprintf(&b, ";\n")
	b.WriteTo(os.Stdout)
	return 0
}
