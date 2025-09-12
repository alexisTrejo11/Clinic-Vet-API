package dto

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/customer/application/command"
)

// UpdateCustomerRequest represents the request to update an existing customer
// @Description Request body for updating a customer
type UpdateCustomerRequest struct {
	// Customer's first name
	// Required: false
	// Minimum length: 2, Maximum length: 50
	FirstName *string `json:"first_name,omitempty" binding:"omitempty,min=2,max=50" example:"John"`

	// Customer's last name
	// Required: false
	// Minimum length: 2, Maximum length: 50
	LastName *string `json:"last_name,omitempty" binding:"omitempty,min=2,max=50" example:"Doe"`

	// Customer's gender
	// Required: false
	// Enum: male, female, not_specified
	Gender *enum.PersonGender `json:"gender,omitempty" binding:"omitempty,oneof=male female not_specified" example:"male"`

	// Customer's date of birth
	// Required: false
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" example:"1990-01-15T00:00:00Z"`

	// URL to customer's photo
	// Required: false
	Photo *string `json:"photo,omitempty" example:"https://example.com/photo.jpg"`

	// Customer's phone number in E.164 format
	// Required: false
	PhoneNumber *string `json:"phone_number,omitempty" binding:"omitempty,e164" example:"+1234567890"`

	// Customer's address
	// Required: false
	Address *string `json:"address,omitempty" example:"123 Main St, City, State"`

	// Additional notes about the customer
	// Required: false
	Notes *string `json:"notes,omitempty" example:"Allergic to penicillin"`
}

// ToCommand converts UpdateCustomerRequest to UpdateCustomerCommand
func (r *UpdateCustomerRequest) ToCommand(ctx context.Context, customerID uint) (*command.UpdateCustomerCommand, error) {
	return &command.UpdateCustomerCommand{
		ID:          valueobject.NewCustomerID(customerID),
		Photo:       r.Photo,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Gender:      r.Gender,
		DateOfBirth: r.DateOfBirth,
		PhoneNumber: r.PhoneNumber,
		Address:     r.Address,
		Notes:       r.Notes,
		CTX:         ctx,
	}, nil
}
