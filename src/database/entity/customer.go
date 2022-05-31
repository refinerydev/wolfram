package entity

type Customer struct {
	ID           uint
	Name         string
	Email        string
	ImageUrl     string
	PhoneNumber  string
	Address      string
	Latitude     float64
	Longitude    float64
	CustomerType string
}
