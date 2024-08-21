package dtos

import "time"

type UserSignupDTO struct {
	Email    string    `json:"email" validate:"required"`
	Phone    string    `json:"phone" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	LastName string    `json:"last_name" validate:"required"`
	Password string    `json:"pasword" validate:"required"`
	BirthDay time.Time `json:"birthday" validate:"required"`
	Genre    string    `json:"genre"`
}

type UserLoginDTO struct {
	Email    string
	Name     string
	LastName string
	Password string
	BirthDay time.Time
}

type UserAdressInsertDTO struct {
	Street       string    `json:"email" validate:"required"`
	Number       string    `json:"phone" validate:"required"`
	Neighborhood string    `json:"neighborhood" validate:"required"`
	City         string    `json:"last_name" validate:"required"`
	Country      string    `json:"pasword" validate:"required"`
	ZipCode      time.Time `json:"zip_code" validate:"required"`
}
