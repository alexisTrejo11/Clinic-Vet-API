package dto

import (
	"context"
	"fmt"
	"strings"
	"time"

	"clinic-vet-api/app/modules/appointment/application/command"
	"clinic-vet-api/app/modules/core/domain/enum"
)

type UpdateApptRequest struct {
	// @Required Required: Unique identifier of the appointment to update
	AppointmentID uint `json:"appointmentId" binding:"required" example:"123"`

	// Optional: ID of the veterinarian assigned to the appointment
	VetID *uint `json:"vetId,omitempty" binding:"omitempty,min=1" example:"456"`

	// @Required Required: New status of the appointment (e.g., "confirmed", "cancelled", "completed")
	Status string `json:"status" binding:"required,oneof=scheduled confirmed cancelled completed no_show" example:"confirmed"`

	// Optional: Reason for the appointment update (e.g., cancellation reason)
	Reason *string `json:"reason,omitempty" binding:"omitempty,max=500" example:"Patient rescheduled"`

	// Optional: Additional notes or observations about the appointment
	Notes *string `json:"notes,omitempty" binding:"omitempty,max=1000" example:"Patient requires special handling"`

	// Optional: Service type for the appointment (e.g., "checkup", "vaccination", "surgery")
	Service *string `json:"service,omitempty" binding:"omitempty,max=100" example:"Annual checkup"`
}

type RequestApptmentData struct {
	// @Required Required: ID of the pet for the appointment
	PetID int `json:"petId" binding:"required,min=1" example:"789"`

	// @Required Required: Type of service requested
	Service string `json:"service" binding:"required,max=100" example:"Vaccination"`

	// @Required Required: Date and time for the appointment (ISO 8601 format)
	Datetime time.Time `json:"datetime" binding:"required" example:"2024-01-15T14:30:00Z"`

	// @Required Required: Reason for the appointment
	Reason string `json:"reason" binding:"required,max=500" example:"Annual vaccination"`

	// Optional: Additional notes or special instructions
	Notes *string `json:"notes,omitempty" binding:"omitempty,max=1000" example:"Pet is anxious around other animals"`
}

type RescheduleAppointData struct {
	// @Required Required: New date and time for the appointment (RFC3339 format)
	NewDateTime time.Time `json:"datetime" binding:"required" example:"2024-01-15T10:30:00Z"`

	// Optional: Reason for rescheduling the appointment
	Reason string `json:"reason,omitempty" binding:"omitempty,max=500" example:"Routine checkup"`
}

// Clean sanitizes and validates the appointment data
func (a *RescheduleAppointData) Clean() {
	a.Reason = strings.TrimSpace(a.Reason)
}

func (a *RescheduleAppointData) ToCommand(ctx context.Context, appointmentID uint) command.RescheduleApptCommand {
	return *command.NewRescheduleApptCommand(ctx, appointmentID, nil, a.NewDateTime, &a.Reason)
}

func (r *RequestApptmentData) ToCommand(ctx context.Context, customerID uint) (*command.RequestApptByCustomerCommand, error) {
	return command.NewRequestApptByCustomerCommand(ctx, uint(r.PetID), customerID, r.Datetime, r.Reason, r.Service, r.Notes)
}

func (r *UpdateApptRequest) ToCommand(ctx context.Context) (command.UpdateApptCommand, error) {
	var clinicService *enum.ClinicService
	if r.Service != nil {
		serviceEnum, err := enum.ParseClinicService(*r.Service)
		if err != nil {
			return command.UpdateApptCommand{}, fmt.Errorf("invalid service: %w", err)
		}
		clinicService = &serviceEnum
	}

	updateCommand := command.NewUpdateApptCommand(
		ctx,
		r.AppointmentID,
		r.VetID,
		r.Status,
		r.Reason,
		r.Notes,
		clinicService)

	return *updateCommand, nil
}
