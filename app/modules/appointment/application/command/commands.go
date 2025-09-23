// Package command contains the command definitions for appointment-related operations.
package command

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type CreateApptCommand struct {
	customerID valueobject.CustomerID
	petID      valueobject.PetID
	vetID      *valueobject.EmployeeID
	service    enum.ClinicService
	datetime   time.Time
	status     *enum.AppointmentStatus
	notes      *string
}

func NewCreateApptCommand(
	customerIDInt,
	petIDInt uint,
	vetIDInt *uint,
	service enum.ClinicService,
	dateTime time.Time,
	status enum.AppointmentStatus,
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
		status:     &status,
		service:    service,
	}
}

type CancelApptCommand struct {
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
	reason        string
}

func NewCancelApptCommand(id uint, vetID *uint, reason string) *CancelApptCommand {
	var vetIDObj *valueobject.EmployeeID
	if vetID != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetID)
		vetIDObj = &vetIDVal
	}

	return &CancelApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
		vetID:         vetIDObj,
		reason:        reason,
	}
}

type CompleteApptCommand struct {
	id         valueobject.AppointmentID
	employeeID *valueobject.EmployeeID
	notes      *string
}

func NewCompleteApptCommand(id uint, vetIDInt *uint, notes string) *CompleteApptCommand {
	cmd := &CompleteApptCommand{
		id:    valueobject.NewAppointmentID(id),
		notes: &notes,
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
}

func NewConfirmAppointmentCommand(appointIDInt, vetIDInt uint) *ConfirmApptCommand {
	return &ConfirmApptCommand{
		id:         valueobject.NewAppointmentID(appointIDInt),
		employeeID: valueobject.NewEmployeeID(vetIDInt),
	}
}

type DeleteApptCommand struct {
	appointmentID valueobject.AppointmentID
}

func NewDeleteApptCommand(id uint) *DeleteApptCommand {
	return &DeleteApptCommand{
		appointmentID: valueobject.NewAppointmentID(id),
	}
}

type NotAttendApptCommand struct {
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
}

func NewNotAttendApptCommand(id uint, vetIDUint *uint) *NotAttendApptCommand {
	var vetID *valueobject.EmployeeID
	if vetIDUint != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetIDUint)
		vetID = &vetIDVal
	}

	return &NotAttendApptCommand{

		appointmentID: valueobject.NewAppointmentID(id),
		vetID:         vetID,
	}
}

type RescheduleApptCommand struct {
	appointmentID  valueobject.AppointmentID
	veterinarianID *valueobject.EmployeeID
	datetime       time.Time
}

func NewRescheduleApptCommand(appointIDInt uint, vetID *uint, dateTime time.Time, reason *string) *RescheduleApptCommand {
	var veterinarianID *valueobject.EmployeeID
	if vetID != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetID)
		veterinarianID = &vetIDVal
	}
	return &RescheduleApptCommand{
		appointmentID:  valueobject.NewAppointmentID(appointIDInt),
		veterinarianID: veterinarianID,
		datetime:       dateTime,
	}
}

type UpdateApptCommand struct {
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
	status        *enum.AppointmentStatus
	notes         *string
	service       *enum.ClinicService
}

func NewUpdateApptCommand(
	appointIDInt uint,
	vetIDInt *uint,
	status string,
	notes *string,
	service *enum.ClinicService,
) *UpdateApptCommand {
	var vetID *valueobject.EmployeeID
	if vetIDInt != nil {
		vetIDObj := valueobject.NewEmployeeID(*vetIDInt)
		vetID = &vetIDObj
	}

	return &UpdateApptCommand{
		appointmentID: valueobject.NewAppointmentID(appointIDInt),
		vetID:         vetID,
		service:       service,
		notes:         notes,
	}
}

type RequestApptByCustomerCommand struct {
	petID         valueobject.PetID
	customerID    valueobject.CustomerID
	requestedDate time.Time
	service       enum.ClinicService
	notes         *string
}

func NewRequestApptByCustomerCommand(
	petID uint,
	customerID uint,
	requestedDate time.Time,
	service string,
	notes *string,
) (*RequestApptByCustomerCommand, error) {
	parsedService, err := enum.ParseClinicService(service)
	if err != nil {
		return nil, err
	}

	return &RequestApptByCustomerCommand{
		petID:         valueobject.NewPetID(petID),
		requestedDate: requestedDate,
		customerID:    valueobject.NewCustomerID(customerID),
		service:       parsedService,
		notes:         notes,
	}, nil
}
