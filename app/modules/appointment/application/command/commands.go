// Package command contains the command definitions for appointment-related operations.
package command

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"context"
	"time"
)

type CreateApptCommand struct {
	ctx        context.Context
	customerID valueobject.CustomerID
	petID      valueobject.PetID
	vetID      *valueobject.EmployeeID
	service    enum.ClinicService
	datetime   time.Time
	status     *enum.AppointmentStatus
	reason     enum.VisitReason
	notes      *string
}

func NewCreateApptCommand(
	ctx context.Context,
	customerIDInt,
	petIDInt uint,
	vetIDInt *uint,
	service enum.ClinicService,
	dateTime time.Time,
	status enum.AppointmentStatus,
	reason enum.VisitReason,
	notes *string,
) *CreateApptCommand {
	var vetID *valueobject.EmployeeID
	if vetIDInt != nil {
		vetIDObj := valueobject.NewEmployeeID(*vetIDInt)
		vetID = &vetIDObj
	}

	return &CreateApptCommand{
		customerID: valueobject.NewCustomerID(customerIDInt),
		petID:      valueobject.NewPetID(petIDInt),
		vetID:      vetID,
		datetime:   dateTime,
		notes:      notes,
		reason:     reason,
		status:     &status,
		service:    service,
	}
}

type CancelApptCommand struct {
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
	reason        string
	ctx           context.Context
}

func NewCancelApptCommand(ctx context.Context, id uint, vetID *uint, reason string) *CancelApptCommand {
	var vetIDObj *valueobject.EmployeeID
	if vetID != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetID)
		vetIDObj = &vetIDVal
	}

	return &CancelApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		vetID:         vetIDObj,
		reason:        reason,
		ctx:           ctx,
	}
}

type CompleteApptCommand struct {
	id         valueobject.AppointmentID
	employeeID *valueobject.EmployeeID
	notes      *string
	ctx        context.Context
}

func NewCompleteApptCommand(ctx context.Context, id uint, vetIDInt *uint, notes string) *CompleteApptCommand {
	cmd := &CompleteApptCommand{
		id:    valueobject.NewAppointmentID(id),
		notes: &notes,
		ctx:   ctx,
	}

	if vetIDInt != nil {
		employeeID := valueobject.NewEmployeeID(*vetIDInt)
		cmd.employeeID = &employeeID
	}
	return cmd
}

type ConfirmApptCommand struct {
	id         valueobject.AppointmentID
	employeeID valueobject.EmployeeID
	ctx        context.Context
}

func NewConfirmAppointmentCommand(ctx context.Context, appointIDInt, vetIDInt uint) *ConfirmApptCommand {
	return &ConfirmApptCommand{
		ctx:        ctx,
		id:         valueobject.NewAppointmentID(appointIDInt),
		employeeID: valueobject.NewEmployeeID(vetIDInt),
	}
}

type DeleteApptCommand struct {
	appointmentID valueobject.AppointmentID
	ctx           context.Context
}

func NewDeleteApptCommand(id uint, ctx context.Context) *DeleteApptCommand {
	return &DeleteApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		ctx:           ctx,
	}
}

type NotAttendApptCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
}

func NewNotAttendApptCommand(ctx context.Context, id uint, vetIDUint *uint) *NotAttendApptCommand {
	var vetID *valueobject.EmployeeID
	if vetIDUint != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetIDUint)
		vetID = &vetIDVal
	}

	return &NotAttendApptCommand{
		ctx:           ctx,
		appointmentID: valueobject.NewAppointmentID(id),
		vetID:         vetID,
	}
}

type RescheduleApptCommand struct {
	ctx            context.Context
	appointmentID  valueobject.AppointmentID
	veterinarianID *valueobject.EmployeeID
	datetime       time.Time
	reason         *string
}

func NewRescheduleApptCommand(ctx context.Context, appointIDInt uint, vetID *uint, dateTime time.Time, reason *string) *RescheduleApptCommand {
	var veterinarianID *valueobject.EmployeeID
	if vetID != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetID)
		veterinarianID = &vetIDVal
	}

	return &RescheduleApptCommand{
		ctx:            ctx,
		appointmentID:  valueobject.NewAppointmentID(appointIDInt),
		veterinarianID: veterinarianID,
		reason:         reason,
		datetime:       dateTime,
	}
}

type UpdateApptCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
	status        *enum.AppointmentStatus
	reason        *string
	notes         *string
	service       *enum.ClinicService
}

func NewUpdateApptCommand(
	ctx context.Context,
	appointIDInt uint,
	vetIDInt *uint,
	status string,
	reason,
	notes *string,
	service *enum.ClinicService,
) *UpdateApptCommand {
	var vetID *valueobject.EmployeeID
	if vetIDInt != nil {
		vetIDObj := valueobject.NewEmployeeID(*vetIDInt)
		vetID = &vetIDObj
	}

	return &UpdateApptCommand{
		ctx:           ctx,
		appointmentID: valueobject.NewAppointmentID(appointIDInt),
		vetID:         vetID,
		service:       service,
		reason:        reason,
		notes:         notes,
	}
}

type RequestApptByCustomerCommand struct {
	ctx           context.Context
	petID         valueobject.PetID
	customerID    valueobject.CustomerID
	requestedDate time.Time
	reason        enum.VisitReason
	service       enum.ClinicService
	notes         *string
}

func NewRequestApptByCustomerCommand(
	ctx context.Context,
	petID uint,
	customerID uint,
	requestedDate string,
	reason string,
	service string,
	notes *string,
) (*RequestApptByCustomerCommand, error) {
	parsedDate, err := time.Parse(time.RFC822, requestedDate)
	if err != nil {
		return nil, err
	}

	parsedReason, err := enum.ParseVisitReason(reason)
	if err != nil {
		return nil, err
	}

	parsedService, err := enum.ParseClinicService(service)
	if err != nil {
		return nil, err
	}

	return &RequestApptByCustomerCommand{
		ctx:           ctx,
		petID:         valueobject.NewPetID(petID),
		requestedDate: parsedDate,
		customerID:    valueobject.NewCustomerID(customerID),
		reason:        parsedReason,
		service:       parsedService,
		notes:         notes,
	}, nil
}
