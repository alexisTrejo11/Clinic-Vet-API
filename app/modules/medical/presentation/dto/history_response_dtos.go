package dto

import (
	"time"

	commondto "clinic-vet-api/app/shared/dto"
)

// MedHistResponse represents a medical history record summary
// swagger:model MedHistResponse
type MedHistoryResponse struct {
	// The unique identifier of the medical history record
	// Required: true
	// Example: 1
	ID uint `json:"id"`

	// The ID of the pet associated with this record
	// Required: true
	// Example: 5
	PetID uint `json:"pet_id"`

	// The date of the medical visit
	// Required: true
	// Format: date-time
	// Example: 2023-10-15T14:30:00Z
	Date time.Time `json:"date"`

	// The diagnosis made during the visit
	// Required: true
	// Max length: 500
	// Example: Otitis externa
	Diagnosis string `json:"diagnosis"`

	// The type of medical visit
	// Required: true
	// Enum: consultation, emergency, vaccination, surgery, checkup
	// Example: consultation
	VisitType string `json:"visit_type"`

	// The reason for the visit
	// Required: true
	// Max length: 200
	// Example: Routine checkup and vaccination
	VisitReason string `json:"visit_reason"`

	// Additional notes about the visit
	// Required: false
	// Max length: 1000
	// Example: Patient responded well to treatment
	Notes *string `json:"notes,omitempty"`

	// The medical condition observed
	// Required: true
	// Max length: 200
	// Example: Acute infection
	Condition string `json:"condition"`

	// The treatment prescribed
	// Required: true
	// Max length: 500
	// Example: Antibiotics for 7 days
	Treatment string `json:"treatment"`

	// The ID of the veterinarian who attended the visit
	// Required: true
	// Example: 3
	EmployeeID uint `json:"employee_id"`

	// The ID of the customer (pet customer)
	// Required: true
	// Example: 8
	CustomerID uint `json:"customer_id"`

	// The creation timestamp of the record
	// Required: true
	// Format: date-time
	// Example: 2023-10-15T16:45:00Z
	CreatedAt time.Time `json:"created_at"`

	// The last update timestamp of the record
	// Required: true
	// Format: date-time
	// Example: 2023-10-15T16:45:00Z
	UpdatedAt time.Time `json:"updated_at"`
}

// MedHistResponseDetail represents detailed medical history information
// swagger:model MedHistResponseDetail
type MedHistoryResponseDetail struct {
	// The unique identifier of the medical history record
	// Required: true
	// Example: 1
	ID uint `json:"id"`

	// Pet details associated with this record
	// Required: true
	Pet commondto.PetDetails `json:"pet"`

	// Customer details (pet customer)
	// Required: true
	Customer commondto.CustomerDetails `json:"customer"`

	// The date of the medical visit
	// Required: true
	// Format: date-time
	// Example: 2023-10-15T14:30:00Z
	Date time.Time `json:"date"`

	// The diagnosis made during the visit
	// Required: true
	// Max length: 500
	// Example: Otitis externa
	Diagnosis string `json:"diagnosis"`

	// Additional notes about the visit
	// Required: false
	// Max length: 1000
	// Example: Patient responded well to treatment
	Notes string `json:"notes,omitempty"`

	// The treatment prescribed
	// Required: true
	// Max length: 500
	// Example: Antibiotics for 7 days
	Treatment string `json:"treatment"`

	// Employee (veterinarian) details
	// Required: true
	Employee commondto.EmployeeDetails `json:"employee"`

	// The creation timestamp of the record
	// Required: true
	// Format: date-time
	// Example: 2023-10-15T16:45:00Z
	CreatedAt time.Time `json:"created_at"`

	// The last update timestamp of the record
	// Required: true
	// Format: date-time
	// Example: 2023-10-15T16:45:00Z
	UpdatedAt time.Time `json:"updated_at"`
}
