package command

import (
	appt "clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/enum"
	apperror "clinic-vet-api/app/shared/error/application"
	"context"
)

func createCommandToDomain(ctx context.Context, command CreateApptCommand) (*appt.Appointment, error) {
	status := enum.AppointmentStatusPending
	if command.status != nil {
		status = *command.status
	}

	appointmentEntity, err := appt.CreateAppointment(
		ctx,
		command.petID,
		command.customerID,
		appt.WithEmployeeID(command.vetID),
		appt.WithService(command.service),
		appt.WithScheduledDate(command.datetime),
		appt.WithNotes(command.notes),
		appt.WithStatus(status),
	)
	if err != nil {
		return nil, apperror.MappingError([]string{err.Error()}, "constructor", "domain", "Appointment")
	}
	return appointmentEntity, nil
}

func requestByCustomerCommandToDomain(ctx context.Context, command RequestApptByCustomerCommand) (*appt.Appointment, error) {
	appointmentEntity, err := appt.CreateAppointment(
		ctx,
		command.petID,
		command.customerID,
		appt.WithService(command.service),
		appt.WithScheduledDate(command.requestedDate),
		appt.WithNotes(command.notes),
		appt.WithStatus(enum.AppointmentStatusPending),
	)
	if err != nil {
		return nil, apperror.MappingError([]string{err.Error()}, "constructor", "domain", "Appointment")
	}

	return appointmentEntity, nil
}
