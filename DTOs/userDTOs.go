package DTOs

import "time"

type UserSignUpDTO struct {
	Email    string    `json:"email" validate:"required"`
	Phone    string    `json:"phone" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	LastName string    `json:"last_name" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Photo    string    `json:"photo"`
	BirthDay time.Time `json:"birthday" validate:"required"`
	Genre    string    `json:"genre"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone"`
}

type UserAddressInsertDTO struct {
	Street       string    `json:"street" validate:"required"`
	Number       string    `json:"number" validate:"required"`
	Neighborhood string    `json:"neighborhood" validate:"required"`
	City         string    `json:"city" validate:"required"`
	Country      string    `json:"country" validate:"required"`
	ZipCode      time.Time `json:"zip_code" validate:"required"`
}

type UserDTO struct {
	Id             int32  `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	Phone          string `json:"phone"`
	Role           string `json:"role"`
}
