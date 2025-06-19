package dtos

import "time"

type AppointmentResponse struct {
	Id      int32     `json:"id"`
	PetID   int32     `json:"pet_id"`
	VetID   int32     `json:"vet_id"`
	Service string    `json:"service"`
	Status  string    `json:"status"`
	OwnerID int32     `json:"owner_id"`
	Date    time.Time `json:"date"`
}

type AppointmentNamedResponse struct {
	Pet     string    `json:"pet"`
	Owner   string    `json:"owner_id"`
	Vet     string    `json:"vet_id"`
	Service string    `json:"service"`
	Status  string    `json:"status"`
	Date    time.Time `json:"date"`
}
