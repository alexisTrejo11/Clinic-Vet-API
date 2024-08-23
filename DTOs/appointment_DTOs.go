package DTOs

import "time"

type AppointmentInsertDTO struct {
	PetID   int32     `json:"pet_id" validate:"required"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service" validate:"required"`
	Date    time.Time `json:"date"`
}

type AppointmentUpdateDTO struct {
	Id      int32     `json:"id" validate:"required"`
	PetID   int32     `json:"pet_id"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service"`
	Date    time.Time `json:"date"`
}

type AppointmentDTO struct {
	Id      int32     `json:"id" validate:"required"`
	PetID   int32     `json:"pet_id" validate:"required"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service" validate:"required"`
	Date    time.Time `json:"date"`
}
