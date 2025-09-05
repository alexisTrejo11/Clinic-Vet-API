package dto

import "time"

// @Description Represents the request body for creating a new userDomain.
type CreateUserRequest struct {
	// The user's email address. (required, must be a valid email format)
	Email string `json:"email" validate:"required,email"`
	// The user's password. (required, minimum 8 characters)
	Password string `json:"password" validate:"required,min=8"`
	// The user's phone number. (required)
	PhoneNumber string `json:"phone_number" validate:"required"`
	// The role of the user (e.g., "customer", "veterinarian", "admin"). (required)
	Role string `json:"role" validate:"required,oneof=customer veterinarian admin"`
	// The user's address. (required)
	Address string `json:"address" validate:"required"`
	// The unique ID of the owner. (required)
	OwnerID *int `json:"owner_id" validate:"required`
	// The unique ID of the veterinarian. (optional)
	VeterinarianID *int `json:"veterinarian_id" validate:"omitempty"`
	// The user's gender. (required, must be "male", "female", or "other")
	Gender *string `json:"gender" validate:"required,oneof=male, female, other"`
	// The user's location. (required)
	Location *string `json:"location" validate:"required"`
	// The user's date of birth. (required)
	DateOfBirth *time.Time `json:"date_of_birth" validate:"required"`
}
