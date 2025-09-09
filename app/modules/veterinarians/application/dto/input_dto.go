// Package dto contains data transfer objects for the veterinarians module.
package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type CreateVetData struct {
	FirstName       string             `json:"first_name" validate:"required"`
	LastName        string             `json:"last_name" validate:"required"`
	Photo           string             `json:"photo"`
	LicenseNumber   string             `json:"license_number" validate:"required"`
	YearsExperience int                `json:"years_experience"`
	IsActive        bool               `json:"is_active"`
	Specialty       string             `json:"specialty"`
	ConsultationFee *valueobject.Money `json:"consultation_fee"`
	LaboralSchedule []ScheduleData     `json:"laboral_schedule"`
}

type ScheduleData struct {
	Day           string `json:"day"`
	EntryTime     int    `json:"entry_time"`
	DepartureTime int    `json:"departure_time"`
	StartBreak    int    `json:"start_break"`
	EndBreak      int    `json:"end_break"`
}

type UpdateVetData struct {
	VetID           valueobject.VetID  `json:"vet_id" validate:"required"`
	FirstName       *string            `json:"first_name"`
	LastName        *string            `json:"last_name"`
	Photo           *string            `json:"photo"`
	LicenseNumber   *string            `json:"license_number"`
	YearsExperience *int               `json:"years_experience"`
	Specialty       *string            `json:"specialty"`
	IsActive        *bool              `json:"is_active"`
	ConsultationFee *valueobject.Money `json:"consultation_fee"`
	LaboralSchedule *[]ScheduleData    `json:"laboral_schedule"`
}

type VetSearchParams struct {
	page.PageInput
	Filters VetFilters `json:"filters"`
	OrderBy VetOrderBy
}

type VetFilters struct {
	Name            *string            `json:"name"`
	LicenseNumber   *string            `json:"license_number"`
	Specialty       *enum.VetSpecialty `json:"specialty"`
	YearsExperience *struct {
		Min *int `json:"min"`
		Max *int `json:"max"`
	} `json:"years_experience"`
	IsActive *bool `json:"is_active"`
}

type VetOrderBy string

const (
	OrderByName            VetOrderBy = "name"
	OrderBySpecialty       VetOrderBy = "specialty"
	OrderByYearsExperience VetOrderBy = "years_experience"
	OrderByCreatedAt       VetOrderBy = "created_at"
)
