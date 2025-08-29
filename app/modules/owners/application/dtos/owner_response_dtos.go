package dto

import (
	"time"

	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/dtos"
)

type OwnerResponse struct {
	Id               int                   `json:"id"`
	Photo            string                `json:"photo"`
	Name             string                `json:"name"`
	PhoneNumber      string                `json:"phone"`
	DateOfBirth      time.Time             `json:"date_of_birth"`
	Address          *string               `json:"address,omitempty" db:"address"`
	EmergencyContact *string               `json:"emergency_contact,omitempty" db:"emergency_contact"`
	Notes            *string               `json:"notes,omitempty" db:"notes"`
	IsActive         bool                  `json:"is_active" db:"is_active"`
	Pets             []petDTOs.PetResponse `json:"pets"`
}

type OwnerSummaryResponse struct {
	Id          int    `json:"id"`
	Photo       string `json:"photo"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	IsActive    bool   `json:"is_active"`
	PetCount    int    `json:"pet_count"`
}

type OwnerListResponse struct {
	Owners []OwnerResponse `json:"owners"`
}
