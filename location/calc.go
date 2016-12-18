package location

import (
	"math"
)

func toRadians(f float64) float64 {
	return f / 180 * math.Pi
}

func FilterByDistanceFromPoint(p Point, dist float64, locations []Location) []Location {

	filteredLocs := make([]Location, 0)

	for _, loc := range locations {
		if p.HaverSineDistance(loc.Point) <= dist {
			filteredLocs = append(filteredLocs, loc)
		}
	}
	return filteredLocs
}