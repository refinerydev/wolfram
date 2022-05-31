package entity

type Customer struct {
	ID           uint
	Name         string
	ImageUrl     string
	Email        string
	PhoneNumber  string
	Address      string
	Latitude     float64
	Longitude    float64
	CustomerType string
}
