package dto

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/medical/session/application/query"
	commondto "clinic-vet-api/app/shared/dto"
)

// MedSessionResponse represents a complete medical session record
// swagger:model MedSessionResponse
type MedSessionResponse struct {
	// The unique identifier of the medical session
	// Required: true
	// Example: 1
	ID uint `json:"id"`

	// The ID of the pet associated with this session
	// Required: true
	// Example: 5
	PetID uint `json:"pet_id"`

	// The ID of the veterinarian who attended the session
	// Required: true
	// Example: 3
	EmployeeID uint `json:"employee_id"`

	// The date and time of the medical visit
	// Required: true
	// Format: date-time
	// Example: 2023-10-15T14:30:00Z
	Date time.Time `json:"date"`

	// The diagnosis made during the visit
	// Required: true
	// Max length: 1000
	// Example: Otitis externa
	Diagnosis string `json:"diagnosis"`

	// The type of medical visit
	// Required: true
	// Enum: consultation, emergency, vaccination, surgery, checkup, follow-up
	// Example: consultation
	VisitType string `json:"visit_type"`

	// The service provided during the visit
	// Required: true
	// Max length: 1000
	// Example: General Checkup
	ServiceProvided string `json:"service_provided"`

	// Additional notes about the visit
	// Required: false
	// Max length: 2000
	// Example: Patient responded well to treatment
	Notes *string `json:"notes,omitempty"`

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

	PetSessionSummaryResponse *PetSessionSummaryResponse `json:"pet_session_summary,omitempty"`
}

type PetSessionSummaryResponse struct {
	// The unique identifier of pet
	// Required: true
	// Example: 5
	PetID uint `json:"pet_id"`

	// The medical condition observed
	// Required: true
	// Max length: 200
	// Example: Stable
	Condition string `json:"condition"`

	// The treatment prescribed
	// Required: true
	// Max length: 1000
	// Example: Antibiotics for 7 days
	Treatment string `json:"treatment"`

	// The weight of the pet in kilograms
	// Required: false
	// Minimum: 0
	// Maximum: 200
	// Example: 12.5
	Weight *float64 `json:"weight,omitempty"`

	// The temperature of the pet in Celsius
	// Required: false
	// Minimum: 30
	// Maximum: 45
	// Example: 38.2
	Temperature *float64 `json:"temperature,omitempty"`

	// The heart rate of the pet in beats per minute
	// Required: false
	// Minimum: 20
	// Maximum: 300
	// Example: 120
	HeartRate *int `json:"heart_rate,omitempty"`

	// The respiratory rate of the pet in breaths per minute
	// Required: false
	// Minimum: 10
	// Maximum: 200
	// Example: 30
	RespiratoryRate *int `json:"respiratory_rate,omitempty"`

	// List of symptoms observed
	// Required: false
	// Example: ["fever", "coughing", "lethargy"]
	Symptoms []string `json:"symptoms,omitempty"`

	// List of medications prescribed
	// Required: false
	// Example: ["Amoxicillin 250mg", "Pain reliever"]
	Medications []string `json:"medications,omitempty"`

	// The scheduled follow-up date
	// Required: false
	// Format: date-time
	// Example: 2023-10-22T14:30:00Z
	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`
}

func FromResult(res *query.MedSessionResult) *MedSessionResponse {
	response := &MedSessionResponse{
		ID:              res.ID.Value(),
		EmployeeID:      res.EmployeeID.Value(),
		Date:            res.VisitDate,
		Diagnosis:       res.Diagnosis,
		VisitType:       res.VisitType.DisplayName(),
		ServiceProvided: res.ClinicService.DisplayName(),
		Notes:           res.Notes,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       res.UpdatedAt,
		PetSessionSummaryResponse: &PetSessionSummaryResponse{
			PetID:           res.PetDetailsResult.PetID.Value(),
			Condition:       res.PetDetailsResult.Condition.DisplayName(),
			Treatment:       res.PetDetailsResult.Treatment,
			Weight:          decimalToFloat64Ptr(res.PetDetailsResult.Weight),
			Temperature:     decimalToFloat64Ptr(res.PetDetailsResult.Temperature),
			HeartRate:       res.PetDetailsResult.HeartRate,
			RespiratoryRate: res.PetDetailsResult.RespiratoryRate,
			Symptoms:        res.PetDetailsResult.Symptoms,
			Medications:     res.PetDetailsResult.Medications,
			FollowUpDate:    res.PetDetailsResult.FollowUpDate,
		},
	}
	return response
}

func decimalToFloat64Ptr(d *valueobject.Decimal) *float64 {
	if d == nil {
		return nil
	}
	f := d.Float64()
	return &f
}

func FromResultList(results []query.MedSessionResult) []MedSessionResponse {
	dtos := make([]MedSessionResponse, len(results))
	for i, res := range results {
		dtos[i] = *FromResult(&res)
	}
	return dtos
}

// MedSessionResponseDetail represents detailed medical history information
// swagger:model MedSessionResponseDetail
type MedSessionResponseDetail struct {
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
