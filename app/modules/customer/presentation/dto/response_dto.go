package dto

import (
	"time"

	"clinic-vet-api/app/core/domain/enum"
)

// CustomerResponse represents the customer response
// @Description Customer information response
type CustomerResponse struct {
	// Customer ID
	ID uint `json:"id" example:"cust_123456"`

	// Customer's first name
	FirstName string `json:"first_name" example:"John"`

	// Customer's last name
	LastName string `json:"last_name" example:"Doe"`

	// Customer's gender
	Gender enum.PersonGender `json:"gender" example:"male"`

	// Customer's date of birth
	DateOfBirth time.Time `json:"date_of_birth" example:"1990-01-15T00:00:00Z"`

	// URL to customer's photo
	PhotoURL *string `json:"photo_url,omitempty" example:"https://example.com/photo.jpg"`

	// Customer's phone number
	PhoneNumber *string `json:"phone_number,omitempty" example:"+1234567890"`

	// Customer's address
	Address *string `json:"address,omitempty" example:"123 Main St, City, State"`

	// Additional notes
	Notes *string `json:"notes,omitempty" example:"Allergic to penicillin"`

	// Whether the customer is active
	IsActive bool `json:"is_active" example:"true"`

	// Creation timestamp
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T12:00:00Z"`

	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-15T14:30:00Z"`
}
