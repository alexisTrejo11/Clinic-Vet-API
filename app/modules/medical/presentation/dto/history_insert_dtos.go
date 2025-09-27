// Package dto contains data transfer objects for medical http history operations
package dto

import (
	"time"
)

type AdminCreateMedSessionRequest struct {
	PetID      uint `json:"pet_id" validate:"required"`
	CustomerID uint `json:"customer_id" validate:"required"`
	EmployeeID uint `json:"employee_id"`
	CreateMedSessionRequest
}

// CreateMedSessionRequest  represents the request to create a new medical history entry
// swagger:model CreateMedSessionRequest
type CreateMedSessionRequest struct {
	// The summary of the pet's medical condition during the visit
	// Required: true
	PetDetails PetSummaryRequest `json:"pet_details" validate:"required,dive"`

	// The date of the medical visit
	// Required: true
	// Example: 2023-10-15T14:30:00Z
	VisitDate time.Time `json:"date" validate:"required"`

	// The diagnosis made during the visit
	// Required: true
	// Example: Otitis externa
	Diagnosis string `json:"diagnosis" validate:"required,max=500"`

	// The type of medical visit
	// Required: true
	// Example: consultation
	VisitType string `json:"visit_type" validate:"required,min=3,max=50"`

	// The clinic service provided during the visit
	// Required: true
	// Example: Routine checkup
	ClinicService string `json:"clinic_service" validate:"required,max=200"`

	// Additional notes about the visit
	// Required: false
	// Example: Patient responded well to treatment
	Notes *string `json:"notes,omitempty" validate:"omitempty,max=1000"`
}

type PetSummaryRequest struct {
	// The ID of the pet
	// Required: true
	// Example: 1
	PetID uint `json:"pet_id" validate:"required"`

	// The medical condition observed
	// Required: true
	// Example: Acute infection
	Condition string `json:"condition" validate:"required,max=200"`

	// The treatment prescribed
	// Required: true
	// Example: Antibiotics for 7 days
	Treatment string `json:"treatment" validate:"required,max=500"`

	// The symptoms observed
	// Required: false
	// Example: Ear scratching, head shaking
	Symptoms []string `json:"symptoms,omitempty" validate:"omitempty,dive,max=100"`

	// The medications prescribed
	// Required: false
	// Example: Amoxicillin, ear drops
	Medications []string `json:"medications,omitempty" validate:"omitempty,dive,max=100"`

	// The follow-up date for the next visit
	// Required: false
	// Example: 2023-10-22T14:30:00Z
	FollowUpDate *time.Time `json:"follow_up_date,omitempty" validate:"omitempty"`

	// The weight of the pet during the visit (in kg)
	// Required: false
	// Example: 12.5
	Weight *float64 `json:"weight,omitempty" validate:"omitempty,gt=0"`

	// The heart rate of the pet during the visit (in bpm)
	// Required: false
	// Example: 90
	HeartRate *int `json:"heart_rate,omitempty" validate:"omitempty,gt=0"`

	// The respiratory rate of the pet during the visit (in breaths per minute)
	// Required: false
	// Example: 20
	RespiratoryRate *int `json:"respiratory_rate,omitempty" validate:"omitempty,gt=0"`

	// The body temperature of the pet during the visit (in °C)
	// Required: false
	// Example: 38.5
	Temperature *float64 `json:"temperature,omitempty" validate:"omitempty,gt=0"`
}

// UpdateMedSessionRequest represents the request to update a medical history entry
// swagger:model UpdateMedSessionRequest
type UpdateMedSessionRequest struct {
	// The date of the medical visit
	// Required: false
	// Example: 2023-10-15T14:30:00Z
	Date *time.Time `json:"date,omitempty" validate:"omitempty"`

	// The diagnosis made during the visit
	// Required: false
	// Example: Otitis externa
	Diagnosis *string `json:"diagnosis,omitempty" validate:"omitempty,max=500"`

	// The type of medical visit
	// Required: false
	// Example: consultation
	VisitType *string `json:"visit_type,omitempty" validate:"omitempty,oneof=consultation emergency vaccination surgery checkup"`

	// The clinic service provided during the visit
	// Required: false
	// Example: Routine checkup
	ClinicService *string `json:"clinic_service,omitempty" validate:"omitempty,max=200"`

	// Additional notes about the visit
	// Required: false
	// Example: Patient responded well to treatment
	Notes *string `json:"notes,omitempty" validate:"omitempty,max=1000"`

	// The medical condition observed
	// Required: false
	// Example: Acute infection
	Condition *string `json:"condition,omitempty" validate:"omitempty,max=200"`

	// The treatment prescribed
	// Required: false
	// Example: Antibiotics for 7 days
	Treatment *string `json:"treatment,omitempty" validate:"omitempty,max=500"`
}
