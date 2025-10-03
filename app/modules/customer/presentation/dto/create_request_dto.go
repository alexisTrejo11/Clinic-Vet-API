// Package dto contains Data Transfer Objects for the customer module
package dto

import (
	"clinic-vet-api/app/modules/customer/application/command"
	"time"
)

// CreateCustomerRequest represents the request to create a new customer
// @Description Request body for creating a new customer
type CreateCustomerRequest struct {
	// Customer's first name
	// Required: true
	// Minimum length: 2, Maximum length: 50
	FirstName string `json:"first_name" binding:"required,min=2,max=50" example:"John"`

	// Customer's last name
	// Required: true
	// Minimum length: 2, Maximum length: 50
	LastName string `json:"last_name" binding:"required,min=2,max=50" example:"Doe"`

	// Customer's gender
	// Required: true
	// Enum: male, female, not_specified
	Gender string `json:"gender" binding:"required" example:"male"`

	// Customer's date of birth
	// Required: true
	DateOfBirth time.Time `json:"date_of_birth" binding:"required" example:"1990-01-15T00:00:00Z"`

	// URL to customer's photo
	// Required: true
	// Example: https://example.com/photo.jpg
	Photo string `json:"photo,omitempty" binding:"required" example:"https://example.com/photo.jpg"`
}

// ToCommand converts CreateCustomerRequest to CreateCustomerCommand
func (r *CreateCustomerRequest) ToCommand() (command.CreateCustomerCommand, error) {
	return command.NewCreateCustomerCommand(
		r.Photo,
		r.FirstName,
		r.LastName,
		r.Gender,
		r.DateOfBirth,
	)
}
