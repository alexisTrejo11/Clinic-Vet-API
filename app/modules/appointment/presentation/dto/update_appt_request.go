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
	AppointmentID uint
	VetID         *uint
	Status        string
	Reason        *string
	Notes         *string
	Service       *string
}

type RequestApptmentData struct {
	PetID    int
	Service  string
	Datetime string
	Reason   string
	Notes    *string
}

type RescheduleApptRequest struct {
	AppointmentID uint
	Datetime      CustomDate
	Reason        *string
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

type RescheduleAppointData struct {
	// @Required Required: Date and time for the appointment (RFC3339 format)
	NewDateTime time.Time `json:"datetime" binding:"required" example:"2024-01-15T10:30:00Z"`

	// Optional: Reason for the appointment
	Reason string `json:"reason,omitempty" binding:"omitempty,max=500" example:"Routine checkup"`
}

func (a *RescheduleAppointData) ToCommandWithVetID(ctx context.Context, apptID, customernID uint) command.RescheduleApptCommand {
	return *command.NewRescheduleApptCommand(ctx, apptID, &customernID, a.NewDateTime, &a.Reason)
}

// Clean sanitizes and validates the appointment data
func (a *RescheduleAppointData) Clean() {
	a.Reason = strings.TrimSpace(a.Reason)
}

func (a *RescheduleAppointData) ToCommand(ctx context.Context, appointmentID uint) command.RescheduleApptCommand {
	return *command.NewRescheduleApptCommand(ctx, appointmentID, nil, a.NewDateTime, &a.Reason)
}

func (r *RequestApptmentData) ToCommand(ctx context.Context, customerID uint) (*command.RequestApptByCustomerCommand, error) {
	return command.NewRequestApptByCustomerCommand(ctx, uint(r.PetID), customerID, r.Service, r.Datetime, r.Reason, r.Notes)
}

func (r *RescheduleApptRequest) ToCommand(ctx context.Context, employeID *uint) (command.RescheduleApptCommand, error) {
	if r.Datetime.IsZero() {
		return command.RescheduleApptCommand{}, fmt.Errorf("date time cannot be zero")
	}

	if r.Datetime.Before(time.Now()) {
		return command.RescheduleApptCommand{}, fmt.Errorf("date time cannot be in the past")
	}

	rescheduleCommand := command.NewRescheduleApptCommand(ctx, r.AppointmentID, employeID, r.Datetime.Time, r.Reason)
	return *rescheduleCommand, nil
}
