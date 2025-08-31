package controller

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

// @Description Represents the request body for updating a user's profile.
type UpdateProfileRequest struct {
	// The unique ID of the userDomain. (required)
	UserID int `json:"user_id" validate:"required"`
	// A brief biography of the userDomain. (optional, max 500 characters)
	Bio *string `json:"bio" validate:"min=0,max=500"`
	// The URL of the user's profile photo. (optional, must be a valid URL)
	PhotoURL *string `json:"photo_url" validate:"omitempty,url"`
	// The name of the userDomain. (optional)
	Name *string `json:"name" validate:"omitempty"`
	// The user's address. (optional)
	Address *AddressRequest `json:"address" validate:"omitempty"`
}

// @Description Represents the address details for a userDomain.
type AddressRequest struct {
	// The street name. (required)
	Street string `json:"street" validate:"required"`
	// The city name. (required)
	City string `json:"city" validate:"required"`
	// The state or province name. (required)
	State string `json:"state" validate:"required"`
	// The postal code. (required)
	ZipCode string `json:"zip_code" validate:"required"`
	// The country name. (required)
	Country string `json:"country" validate:"required"`
	// The type of building (e.g., "house", "apartment", "office", "other"). (optional)
	BuildingType string `json:"building_type" validate:"omitempty,oneof=house apartment office other"`
	// The outer number of the building. (required)
	BuildingOuterNumber string `json:"building_outer_number" validate:"required"`
	// The inner number of the building. (optional)
	BuildingInnerNumber *string `json:"building_inner_number" validate:"omitempty"`
}
