package location

import (
	"math"
	"fmt"
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
	SubDivision string
	Country     string
	PostalCode  string
	Brand       string
	Metadata    string
	Point
}

func (l *Location) String() string {
	return fmt.Sprintf("{ name=%s, brand=%s, address1=%s, address2=%s, "+"city=%s, state=%s, postal=%s, country=%s, phone=%s, lat=%v, log=%v}",
		l.Name, l.Brand, l.Address1, l.Address2, l.City, l.SubDivision, l.PostalCode, l.Country,
		l.Point.Latitude, l.Point.Longitude)
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
	a := math.Pow(sindLat, 2) + math.Pow(sindLng, 2)*math.Cos(lat1)*math.Cos(lat2);
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a));
	return EARTH_RADIUS * c
}
