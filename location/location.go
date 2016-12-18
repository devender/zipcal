package location

import (
	"math"
)

const (
	EARTH_RADIUS = 3959 //miles
)

type Angle float64

type Point struct {
	Latitude  Angle
	Longitude Angle
}

// Radians returns the angle in radians.
func (a Angle) Radians() float64 {
	return float64(a / 180 * math.Pi)
}

type Location struct {
	Name        string
	PhoneNumber string
	Address1    string
	Address2    string
	Address3    string
	City        string
	Country     string
	PostalCode  string
	Point
}

func (p1 Point) HaverSineDistance(p2 Point) float64 {

	lat1 := p1.Latitude.Radians()
	lat2 := p2.Latitude.Radians()

	lon1 := p1.Longitude.Radians()
	lon2 := p2.Longitude.Radians()

	dLng := lon2 - lon1
	dLat := lat2 - lat1

	sindLat := math.Sin(dLat / 2);
	sindLng := math.Sin(dLng / 2);

	a := math.Pow(sindLat, 2) + math.Pow(sindLng, 2) * math.Cos(lat1) * math.Cos(lat2);

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a));

	return EARTH_RADIUS * c
}
