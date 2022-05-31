package customer

type request struct {
	Name         string  `json:"name" validate:"required"`
	Email        string  `json:"email" validate:"required,email"`
	PhoneNumber  string  `json:"phone_number" validate:"required"`
	Address      string  `json:"address" validate:"required"`
	Latitude     float64 `json:"lat" validate:"required,numeric"`
	Longitude    float64 `json:"long" validate:"required,numeric"`
	CustomerType string  `json:"customer_type" validate:"required"`
}
