package loader

import (
	"github.com/devender/zipcal/location"
)

// Loader is something that loads data from somewhere and returns a slice of locations
type Loader interface {
	Load() []location.Location
}
