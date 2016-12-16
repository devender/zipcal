package loader

import (
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
	"encoding/csv"
)

type Starbucks struct {
	Source string
}

func (s *Starbucks) Load() []Location {
	log.Debugf("reading from %s", s.Source)

	f, err := os.Open(s.Source)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	r := csv.NewReader(f)

	locations := make([]Location, 1000)

	for {
		r, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		locations = append(locations, Location{
			Name: r[1], PhoneNumber: r[4], Address1: r[7], Address2: r[8], Address3: r[9], City: r[10],
			Country: r[12], PostalCode: r[13], Latitude: r[15], Longitude: r[16],
		})

	}

	return locations
}
