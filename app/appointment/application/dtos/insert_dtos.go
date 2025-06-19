package dtos

import "time"

type AppointmentRequestCreate struct {
	PetID   int32     `json:"pet_id" validate:"required"`
	Service string    `json:"service" validate:"required"`
	Date    time.Time `json:"date"`
}

type AppointmentCreate struct {
	PetID   int32     `json:"pet_id" validate:"required"`
	OwnerID int32     `json:"owner_id" validate:"required"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service" validate:"required"`
	Status  string    `json:"status"`
	Date    time.Time `json:"date"`
}

type AppointmentOwnerUpdate struct {
	Id      int32     `json:"id" validate:"required"`
	PetID   int32     `json:"pet_id"`
	Service string    `json:"service"`
	Date    time.Time `json:"date"`
}

type AppointmentUpdate struct {
	Id      int32     `json:"id" validate:"required"`
	PetID   int32     `json:"pet_id"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service"`
	Status  string    `json:"status"`
	OwnerID int32     `json:"owner_id"`
	Date    time.Time `json:"date"`
}
