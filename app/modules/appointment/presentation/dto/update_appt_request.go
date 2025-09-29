package dto

import (
	"strings"
	"time"

	"clinic-vet-api/app/modules/appointment/application/command"
)

type UpdateApptRequest struct {
	// Optional : New status of the appointment (e.g., "confirmed", "cancelled", "completed")
	Status *string `json:"status" binding:"required,oneof=scheduled confirmed cancelled completed no_show" example:"confirmed"`

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

func (a *RescheduleAppointData) ToCommand(appointmentID uint) (command.RescheduleApptCommand, error) {
	return command.NewRescheduleApptCommand(appointmentID, nil, a.NewDateTime, &a.Reason)
}

func (r *RequestApptmentData) ToCommand(customerID uint) (command.RequestApptByCustomerCommand, error) {
	return command.NewRequestApptByCustomerCommand(uint(r.PetID), customerID, r.Datetime, r.Service, r.Notes)
}

func (r *UpdateApptRequest) ToCommand(appointmentID uint) (command.UpdateApptCommand, error) {
	return command.NewUpdateApptCommand(
		appointmentID,
		r.Status,
		r.Notes,
		r.Service,
	)
}
