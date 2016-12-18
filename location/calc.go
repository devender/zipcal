package location

import (
	"math"
)

const (
	EARTH_RADIUS = 3959

	// WGS-84 ellipsoid params
	WGS_84_a = 6378137
	WGS_84_b = 6356752.314245
	WGS_84_f = 1 / 298.257223563
)

func toRadians(f float64) float64 {
	return f / 180 * math.Pi
}

func HaverSineDistance(p1 Point, p2 Point) float64 {

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

func FilterByDistanceFromPoint(p Point, dist float64, locations []Location) []Location {

	filteredLocs := make([]Location, 0)

	for _, loc := range locations {
		if HaverSineDistance(p, loc.Point) <= dist {
			filteredLocs = append(filteredLocs, loc)
		}
	}
	return filteredLocs
}