package dto

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/customer/application/command"
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
	Gender enum.PersonGender `json:"gender" binding:"required,oneof=male female not_specified" example:"male"`

	// Customer's date of birth
	// Required: true
	DateOfBirth time.Time `json:"date_of_birth" binding:"required" example:"1990-01-15T00:00:00Z"`

	// URL to customer's photo
	// Required: false
	Photo string `json:"photo,omitempty" example:"https://example.com/photo.jpg"`

	// Customer's phone number in E.164 format
	// Required: false
	PhoneNumber string `json:"phone_number,omitempty" binding:"omitempty,e164" example:"+1234567890"`
}

// ToCommand converts CreateCustomerRequest to CreateCustomerCommand
func (r *CreateCustomerRequest) ToCommand(ctx context.Context) (*command.CreateCustomerCommand, error) {
	return &command.CreateCustomerCommand{
		Photo:       r.Photo,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Gender:      r.Gender,
		DateOfBirth: r.DateOfBirth,
		CTX:         ctx,
	}, nil
}
