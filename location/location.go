package location

import (
	"math"
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

