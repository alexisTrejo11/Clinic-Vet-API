package dto

import (
	"time"

	petdto "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/usecase/dtos"
)

type OwnerDetail struct {
	ID               int                 `json:"id"`
	Photo            string              `json:"photo"`
	Name             string              `json:"name"`
	PhoneNumber      string              `json:"phone"`
	DateOfBirth      time.Time           `json:"date_of_birth"`
	Address          *string             `json:"address,omitempty" db:"address"`
	EmergencyContact *string             `json:"emergency_contact,omitempty" db:"emergency_contact"`
	Notes            *string             `json:"notes,omitempty" db:"notes"`
	IsActive         bool                `json:"is_active" db:"is_active"`
	Pets             []petdto.PetDetails `json:"pets"`
}

type OwnerSummary struct {
	ID          int    `json:"id"`
	Photo       string `json:"photo"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	IsActive    bool   `json:"is_active"`
	PetCount    int    `json:"pet_count"`
}

type OwnerListDetail struct {
	Owners []OwnerDetail `json:"owners"`
}
