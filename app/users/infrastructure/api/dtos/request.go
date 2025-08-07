package userDtos

import "time"

type CreateUserRequest struct {
	Email          string     `json:"email" validate:"required,email"`
	Password       string     `json:"password" validate:"required,min=8"`
	PhoneNumber    string     `json:"phone_number" validate:"required"`
	Role           string     `json:"role" validate:"required,oneof=customer veterinarian admin"`
	Address        string     `json:"address" validate:"required"`
	OwnerId        *int       `json:"owner_id" validate:"required"`
	VeterinarianId *int       `json:"veterinarian_id" validate:"omitempty"`
	Gender         *string    `json:"gender" validate:"required,oneof=male, female, other"`
	Location       *string    `json:"location" validate:"required"`
	DateOfBirth    *time.Time `json:"date_of_birth" validate:"required"`
}
