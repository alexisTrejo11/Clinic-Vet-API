package dto

import (
	"context"
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
)

// RequestApptByOwnerData represents the structure for appointment booking requests
// @Description Appointment booking request data with validation and sanitization
type RequestApptByOwnerData struct {
	// @Required Required: The ID of the pet for the appointment
	PetID uint `json:"petId" binding:"required,min=1" example:"123"`

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
func (a *RequestApptByOwnerData) Clean() {
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

func (a *RequestApptByOwnerData) ToCommand(ctx context.Context, userOwnerID uint) (command.CreateApptCommand, error) {
	clinicService, err := enum.ParseClinicService(a.Service)
	if err != nil {
		return command.CreateApptCommand{}, err
	}

	visitsReason, err := enum.ParseVisitReason(a.Reason)
	if err != nil && a.Reason != "" {
		return command.CreateApptCommand{}, err
	}

	ownerAppointCMD := command.NewCreateApptCommand(
		ctx,
		userOwnerID,
		a.PetID,
		nil, // No vet assigned at booking
		clinicService,
		a.DateTime,
		enum.AppointmentStatusPending, // Default status is pending
		visitsReason,
		a.Notes,
	)
	return *ownerAppointCMD, nil
}

type RescheduleAppointData struct {
	// @Required Required: Date and time for the appointment (RFC3339 format)
	NewDateTime time.Time `json:"datetime" binding:"required" example:"2024-01-15T10:30:00Z"`

	// Optional: Reason for the appointment
	Reason string `json:"reason,omitempty" binding:"omitempty,max=500" example:"Routine checkup"`
}

func (a *RescheduleAppointData) ToCommandWithVetID(ctx context.Context, apptID, ownernID uint) command.RescheduleApptCommand {
	return *command.NewRescheduleApptCommand(ctx, apptID, &ownernID, a.NewDateTime, &a.Reason)
}

// Clean sanitizes and validates the appointment data
func (a *RescheduleAppointData) Clean() {
	a.Reason = strings.TrimSpace(a.Reason)
}

func (a *RescheduleAppointData) ToCommand(ctx context.Context, appointmentID uint) command.RescheduleApptCommand {
	return *command.NewRescheduleApptCommand(ctx, appointmentID, nil, a.NewDateTime, &a.Reason)
}
