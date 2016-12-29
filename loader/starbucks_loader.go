package loader

import (
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
	"encoding/csv"
	"strconv"
	"github.com/devender/zipcal/location"
)

type Starbucks struct {
	Source string
}

func (s *Starbucks) Load() []location.Location {
	log.Debugf("reading from %s", s.Source)

	f, err := os.Open(s.Source)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	r := csv.NewReader(f)

	locations := make([]location.Location, 1000)

	for {
		r, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		latitude, err1 := strconv.ParseFloat(r[15], 64)
		longitude, err2 := strconv.ParseFloat(r[16], 64)

		if err1 == nil && err2 == nil && latitude != 0 && longitude != 0 {
			locations = append(locations, location.Location{
				Name: r[1],
				Brand: r[2],
				PhoneNumber: r[4],
				Address1: r[7],
				Address2: r[8],
				Address3: r[9],
				City: r[10],
				SubDivision: r[11],
				Country: r[12],
				PostalCode: r[13],
				Point: location.Point{
					Longitude: location.Angle(longitude),
					Latitude: location.Angle(latitude),

				},
			})
		}
	}

	return locations
}
