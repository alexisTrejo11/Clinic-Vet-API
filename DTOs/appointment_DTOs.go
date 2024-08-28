package DTOs

import "time"

type AppointmentRequestInsertDTO struct {
	PetID   int32     `json:"pet_id" validate:"required"`
	Service string    `json:"service" validate:"required"`
	Date    time.Time `json:"date"`
}

type AppointmentInsertDTO struct {
	PetID   int32     `json:"pet_id" validate:"required"`
	OwnerID int32     `json:"owner_id" validate:"required"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service" validate:"required"`
	Status  string    `json:"status"`
	Date    time.Time `json:"date"`
}

type AppointmentOwnerUpdateDTO struct {
	Id      int32     `json:"id" validate:"required"`
	PetID   int32     `json:"pet_id"`
	Service string    `json:"service"`
	Date    time.Time `json:"date"`
}

type AppointmentUpdateDTO struct {
	Id      int32     `json:"id" validate:"required"`
	PetID   int32     `json:"pet_id"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service"`
	Status  string    `json:"status"`
	OwnerID int32     `json:"owner_id"`
	Date    time.Time `json:"date"`
}

type AppointmentDTO struct {
	Id      int32     `json:"id"`
	PetID   int32     `json:"pet_id"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service"`
	Status  string    `json:"status"`
	OwnerID int32     `json:"owner_id"`
	Date    time.Time `json:"date"`
}

type AppointmentNamedDTO struct {
	Pet     string    `json:"pet"`
	Owner   string    `json:"owner_id"`
	Vet     string    `json:"vet_id"`
	Service string    `json:"service"`
	Status  string    `json:"status"`
	Date    time.Time `json:"date"`
}
