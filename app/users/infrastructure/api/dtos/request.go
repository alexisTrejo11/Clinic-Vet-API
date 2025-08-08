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

type UpdateProfileRequest struct {
	UserId   int             `json:"user_id" validate:"required"`
	Bio      *string         `json:"bio" validate:"min=0,max=500"`
	PhotoURL *string         `json:"photo_url" validate:"omitempty,url"`
	Name     *string         `json:"name" validate:"omitempty"`
	Address  *AddressRequest `json:"address" validate:"omitempty"`
}

type AddressRequest struct {
	Street              string  `json:"street" validate:"required"`
	City                string  `json:"city" validate:"required"`
	State               string  `json:"state" validate:"required"`
	ZipCode             string  `json:"zip_code" validate:"required"`
	Country             string  `json:"country" validate:"required"`
	BuildingType        string  `json:"building_type" validate:"omitempty,oneof=house apartment office other"`
	BuildingOuterNumber string  `json:"building_outer_number" validate:"required"`
	BuildingInnerNumber *string `json:"building_inner_number" validate:"omitempty"`
}
