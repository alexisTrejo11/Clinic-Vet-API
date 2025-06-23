package authDto

import "time"

type Address struct {
	Street       string    `json:"street" validate:"required"`
	PhoneNumber  string    `json:"phone_number" validate:"required,e164"`
	Neighborhood string    `json:"neighborhood" validate:"required"`
	City         string    `json:"city" validate:"required"`
	Country      string    `json:"country" validate:"required"`
	ZipCode      time.Time `json:"zip_code" validate:"required"`
}
