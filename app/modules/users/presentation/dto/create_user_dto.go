package dto

import (
	"time"

	"clinic-vet-api/app/modules/users/application/usecase/command"
)

// AdminCreateUserRequest represents the payload for creating a new user by an administrator
// This endpoint allows administrators to create user accounts with specific roles and permissions
// swagger:model AdminCreateUserRequest
type AdminCreateUserRequest struct {
	// The user's email address
	// Required: true
	// Format: email
	// Example: user@example.com
	// Max length: 255
	Email string `json:"email" validate:"required,email"`

	// The user's password
	// Required: true
	// Min length: 8
	// Max length: 72
	// Example: SecurePassword123!
	// Pattern: ^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^\da-zA-Z]).{8,}$
	Password string `json:"password" validate:"required,min=8"`

	// The user's phone number in E.164 format
	// Required: false
	// Format: e164
	// Example: +1234567890
	// Max length: 20
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`

	// The role assigned to the user
	// Required: true
	// Enum: customer, veterinarian, admin, receptionist, groomer
	// Example: veterinarian
	Role string `json:"role" validate:"required"`

	// The user's gender identity
	// Required: true
	// Enum: male, female, other
	// Example: female
	Gender *string `json:"gender" validate:"required,oneof=male female other"`

	// The user's physical location or address
	// Required: true
	// Min length: 5
	// Max length: 255
	// Example: 123 Main St, City, State 12345
	Location *string `json:"location" validate:"required"`

	// The status of the user account
	// Required: false
	// Enum: active, inactive, pending
	// Default: pending
	// Example: active
	Status string `json:"status" validate:"omitempty,oneof=active inactive pending"`

	// The user's date of birth
	// Required: true
	// Format: date
	// Example: 1990-01-15
	// Minimum: 1900-01-01
	// Maximum: current date - 13 years (must be at least 13 years old)
	DateOfBirth *time.Time `json:"date_of_birth" validate:"required"`
}

func (c *AdminCreateUserRequest) ToCommand() (command.CreateUserCommand, error) {
	return command.NewCreateUserCommand(c.Email, c.PhoneNumber, c.Password, c.Role, c.Status)
}
