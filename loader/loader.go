package loader

type Location struct {
	Name        string
	PhoneNumber string
	Address1    string
	Address2    string
	Address3    string
	City        string
	Country     string
	PostalCode  string
	Latitude    string
	Longitude   string
}

// Loader is something that loads data from somewhere and returns a slice of locations
type Loader interface {
	Load() []Location
}
