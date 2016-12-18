package location

import (
	"testing"
	"fmt"
)

func TestHaverSineDistance(t *testing.T) {

	home := Point{
		Longitude: -118.795078,
		Latitude: 34.271988,
	}

	office := Point{
		Longitude: -118.492862,
		Latitude: 34.158232,
	}

	dist := home.HaverSineDistance(office)

	fmt.Printf("distance in miles %f", dist)
}