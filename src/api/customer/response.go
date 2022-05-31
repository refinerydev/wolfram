package customer

import "github.com/refinerydev/wolfram/src/database/entity"

type response struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	ImageUrl     string  `json:"image_url"`
	Email        string  `json:"email"`
	PhoneNumber  string  `json:"phone_number"`
	Address      string  `json:"address"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"long"`
	CustomerType string  `json:"customer_type"`
}

func customerFormatter(customer entity.Customer) response {
	formatter := response{
		ID:           customer.ID,
		Name:         customer.Name,
		ImageUrl:     customer.ImageUrl,
		Email:        customer.Email,
		PhoneNumber:  customer.PhoneNumber,
		Address:      customer.Address,
		Latitude:     customer.Latitude,
		Longitude:    customer.Longitude,
		CustomerType: customer.CustomerType,
	}

	return formatter
}

func customerListFormatter(customers []entity.Customer) []response {
	formatter := []response{}

	for _, customer := range customers {
		c := customerFormatter(customer)
		formatter = append(formatter, c)
	}
	return formatter
}
