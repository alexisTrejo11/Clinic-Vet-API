package dto

import (
	"context"
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
)

// RequestOwnerAppointmentData represents the structure for appointment booking requests
// @Description Appointment booking request data with validation and sanitization
type RequestOwnerAppointmentData struct {
	// @Required Required: The ID of the pet for the appointment
	PetID int `json:"petId" binding:"required,min=1" example:"123"`

	// @Required Required: Type of service requested (e.g., "checkup", "vaccination", "surgery")
	Service string `json:"service" binding:"required,min=2,max=100" example:"Annual Checkup"`

	// @Required Required: Date and time for the appointment (RFC3339 format)
	DateTime time.Time `json:"datetime" binding:"required" example:"2024-01-15T10:30:00Z"`

	// Optional: Reason for the appointment
	Reason string `json:"reason,omitempty" binding:"omitempty,max=500" example:"Routine checkup"`

	// Optional: Additional notes for the appointment
	Notes *string `json:"notes,omitempty" binding:"omitempty,max=1000" example:"Patient has allergy to penicillin"`
}

// Clean sanitizes and validates the appointment data
func (a *RequestOwnerAppointmentData) Clean() {
	a.Service = strings.TrimSpace(a.Service)
	a.Reason = strings.TrimSpace(a.Reason)

	a.Service = strings.ToLower(a.Service)

	if a.Reason != "" {
		a.Reason = strings.ToUpper(a.Reason[:1]) + a.Reason[1:]
	}

	if a.Notes != nil {
		cleanedNotes := strings.TrimSpace(*a.Notes)
		a.Notes = &cleanedNotes
	}
}

func (a *RequestOwnerAppointmentData) ToCommand(userOwnerID int) (command.CreateAppointmentCommand, error) {
	ownerAppointCMD, err := command.NewCreateAppointCmd(
		context.Background(),
		userOwnerID,
		a.PetID,
		nil,
		a.Service,
		a.DateTime,
		"pending",
		a.Reason,
		a.Notes,
	)
	return *ownerAppointCMD, err
}

type OwnerRescheduleAppointData struct {
	// @Required Required: Date and time for the appointment (RFC3339 format)
	NewDateTime time.Time `json:"datetime" binding:"required" example:"2024-01-15T10:30:00Z"`

	// Optional: Reason for the appointment
	Reason string `json:"reason,omitempty" binding:"omitempty,max=500" example:"Routine checkup"`
}

// Clean sanitizes and validates the appointment data
func (a *OwnerRescheduleAppointData) Clean() {
	a.Reason = strings.TrimSpace(a.Reason)
}

func (a *OwnerRescheduleAppointData) ToCommand(appointmentID int) (command.RescheduleAppointmentCommand, error) {
	return command.NewRescheduleAppointmentCommand(context.Background(), appointmentID, nil, a.NewDateTime, &a.Reason)
}
