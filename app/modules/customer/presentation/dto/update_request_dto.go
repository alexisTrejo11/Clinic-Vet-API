package dto

import (
	"time"

	"clinic-vet-api/app/modules/customer/application/command"
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
	Gender *string `json:"gender,omitempty" binding:"omitempty,oneof=male female not_specified" example:"male"`

	// Customer's date of birth
	// Required: false
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" example:"1990-01-15T00:00:00Z"`

	// URL to customer's photo
	// Required: false
	Photo *string `json:"photo,omitempty" example:"https://example.com/photo.jpg"`
}

// ToCommand converts UpdateCustomerRequest to UpdateCustomerCommand
func (r *UpdateCustomerRequest) ToCommand(customerID uint) (command.UpdateCustomerCommand, error) {
	return command.NewUpdateCustomerCommand(
		customerID,
		r.Photo,
		r.FirstName,
		r.LastName,
		r.Gender,
		r.DateOfBirth,
	)

}
