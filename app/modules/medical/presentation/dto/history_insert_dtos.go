// Package dto contains data transfer objects for medical http history operations
package dto

import (
	"time"
)

type AdminCreateMedHistoryRequest struct {
	PetID      uint `json:"pet_id" validate:"required"`
	CustomerID uint `json:"customer_id" validate:"required"`
	EmployeeID uint `json:"employee_id"`
	CreateMedHistoryRequest
}

// CreateMedHistoryRequest  represents the request to create a new medical history entry
// swagger:model CreateMedHistoryRequest
type CreateMedHistoryRequest struct {
	// The date of the medical visit
	// Required: true
	// Example: 2023-10-15T14:30:00Z
	Date time.Time `json:"date" validate:"required"`

	// The diagnosis made during the visit
	// Required: true
	// Example: Otitis externa
	Diagnosis string `json:"diagnosis" validate:"required,max=500"`

	// The type of medical visit
	// Required: true
	// Example: consultation
	VisitType string `json:"visit_type" validate:"required,oneof=consultation emergency vaccination surgery checkup"`

	// The reason for the visit
	// Required: true
	// Example: Routine checkup
	VisitReason string `json:"visit_reason" validate:"required,max=200"`

	// Additional notes about the visit
	// Required: false
	// Example: Patient responded well to treatment
	Notes *string `json:"notes,omitempty" validate:"omitempty,max=1000"`

	// The medical condition observed
	// Required: true
	// Example: Acute infection
	Condition string `json:"condition" validate:"required,max=200"`

	// The treatment prescribed
	// Required: true
	// Example: Antibiotics for 7 days
	Treatment string `json:"treatment" validate:"required,max=500"`
}

// UpdateMedHistoryRequest represents the request to update a medical history entry
// swagger:model UpdateMedHistoryRequest
type UpdateMedHistoryRequest struct {
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

	// The reason for the visit
	// Required: false
	// Example: Routine checkup
	VisitReason *string `json:"visit_reason,omitempty" validate:"omitempty,max=200"`

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
