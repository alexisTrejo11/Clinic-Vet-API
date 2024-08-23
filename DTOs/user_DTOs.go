package DTOs

import (
	"time"
)

type UserSignUpDTO struct {
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=8"`
	Birthday    string `json:"birthday"`
	Photo       string `json:"photo"`
	Genre       Genre  `json:"genre" validate:"required,oneof=male female other"`
}

type UserLoginDTO struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" validate:"required"`
}

type UserAddressInsertDTO struct {
	Street       string    `json:"street" validate:"required"`
	PhoneNumber  string    `json:"phone_number" validate:"required,e164"`
	Neighborhood string    `json:"neighborhood" validate:"required"`
	City         string    `json:"city" validate:"required"`
	Country      string    `json:"country" validate:"required"`
	ZipCode      time.Time `json:"zip_code" validate:"required"`
}

type UserDTO struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type Genre string

const (
	Male   Genre = "male"
	Female Genre = "female"
	Other  Genre = "other"
)

type UserEmployeeSignUpDTO struct {
	Email          string `json:"email" validate:"required,email"`
	PhoneNumber    string `json:"phone_number" validate:"required,e164"`
	Password       string `json:"password" validate:"required,min=8"`
	VeterinarianId int32  `json:"veterinarian_id" validate:"required"`
}

type UserEmployeeLoginDTO struct {
	VeterinarianId *int32 `json:"veterinarian_id"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password" validate:"required"`
}
