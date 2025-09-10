// Package dto contains data transfer objects for the veterinarians module.
package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type CreateEmployeeData struct {
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

type UpdateEmployeeData struct {
	EmployeeID      valueobject.EmployeeID `json:"vet_id" validate:"required"`
	FirstName       *string                `json:"first_name"`
	LastName        *string                `json:"last_name"`
	Photo           *string                `json:"photo"`
	LicenseNumber   *string                `json:"license_number"`
	YearsExperience *int                   `json:"years_experience"`
	Specialty       *string                `json:"specialty"`
	IsActive        *bool                  `json:"is_active"`
	ConsultationFee *valueobject.Money     `json:"consultation_fee"`
	LaboralSchedule *[]ScheduleData        `json:"laboral_schedule"`
}
