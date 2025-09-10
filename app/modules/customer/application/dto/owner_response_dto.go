package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
)

type CustomerDetail struct {
	ID               uint              `json:"id"`
	Photo            string            `json:"photo"`
	Name             string            `json:"name"`
	PhoneNumber      string            `json:"phone"`
	DateOfBirth      time.Time         `json:"date_of_birth"`
	EmergencyContact *string           `json:"emergency_contact,omitempty" db:"emergency_contact"`
	Notes            *string           `json:"notes,omitempty" db:"notes"`
	IsActive         bool              `json:"is_active" db:"is_active"`
	Pets             []dto.PetResponse `json:"pets"`
}

type CustomerSummary struct {
	ID          int    `json:"id"`
	Photo       string `json:"photo"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone"`
	IsActive    bool   `json:"is_active"`
	PetCount    int    `json:"pet_count"`
}

type CustomerListDetail struct {
	Customers []CustomerDetail `json:"owners"`
}
