package dto

import (
	"time"

	"clinic-vet-api/app/modules/appointment/application/command"
	"clinic-vet-api/app/modules/core/domain/enum"
)

// CreateApptRequest represents the request to create an appointment
// @Description Request body for creating a new appointment
type CreateApptRequest struct {
	// ID of the customer making the appointment
	// Required: true
	CustomerID uint `json:"customer_id" binding:"required" example:"123"`

	// ID of the pet for the appointment
	// Required: true
	PetID uint `json:"pet_id" binding:"required" example:"456"`

	// Date and time of the appointment
	// Required: true
	// Format: RFC3339
	Datetime time.Time `json:"datetime" binding:"required" example:"2024-03-15T10:30:00Z"`

	// Reason for the appointment
	// Required: true
	Reason string `json:"reason" binding:"required" example:"Annual checkup"`

	// Service requested
	// Required: true
	Service string `json:"service" binding:"required" example:"Vaccination"`

	// Additional notes for the appointment (optional)
	Notes *string `json:"notes" example:"Patient has allergy to penicillin"`

	// Status of the appointment
	// Required: true
	Status string `json:"status" binding:"required" example:"scheduled"`
}

// RequestApptResponse represents the response for appointment request
// @Description Response body for appointment request confirmation
type RequestApptResponse struct {
	// ID of the pet for the appointment
	// Required: true
	PetID uint `json:"pet_id" binding:"required" example:"456"`

	// Date and time of the appointment
	// Required: true
	// Format: RFC3339
	Datetime time.Time `json:"datetime" binding:"required" example:"2024-03-15T10:30:00Z"`

	// Reason for the appointment
	// Required: true
	Reason string `json:"reason" binding:"required" example:"Annual checkup"`

	// Service requested
	// Required: true
	Service string `json:"service" binding:"required" example:"Vaccination"`

	// Additional notes for the appointment (optional)
	Notes *string `json:"notes" example:"Patient has allergy to penicillin"`
}

func (r *CreateApptRequest) ToCommand(employeeID *uint) (*command.CreateApptCommand, error) {
	clinicService, err := enum.ParseClinicService(r.Service)
	if err != nil {
		return nil, err
	}
	apptStatus, err := enum.ParseAppointmentStatus(r.Status)
	if err != nil {
		return nil, err
	}

	return command.NewCreateApptCommand(
		r.CustomerID,
		r.PetID,
		employeeID,
		clinicService,
		r.Datetime,
		apptStatus,
		r.Notes,
	), nil
}
