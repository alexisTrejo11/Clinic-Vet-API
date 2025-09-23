package bus

import (
	appointcommand "clinic-vet-api/app/modules/appointment/application/command"
	icqrs "clinic-vet-api/app/shared/cqrs"
	"context"
)

type AppointmentCommandBus struct {
	handler appointcommand.AppointementCommandHandler
}

func NewAppointmentCommandBus(handler appointcommand.AppointementCommandHandler) *AppointmentCommandBus {
	return &AppointmentCommandBus{handler: handler}
}

func (bus *AppointmentCommandBus) RequestAppointmentByCustomer(ctx context.Context, cmd appointcommand.RequestApptByCustomerCommand) icqrs.CommandResult {
	return bus.handler.RequestAppointmentByCustomer(ctx, cmd)
}

func (bus *AppointmentCommandBus) CreateAppointment(ctx context.Context, cmd appointcommand.CreateApptCommand) icqrs.CommandResult {
	return bus.handler.CreateAppointment(ctx, cmd)
}

func (bus *AppointmentCommandBus) DeleteAppointment(ctx context.Context, cmd appointcommand.DeleteApptCommand) icqrs.CommandResult {
	return bus.handler.DeleteAppointment(ctx, cmd)
}

func (bus *AppointmentCommandBus) RescheduleAppointment(ctx context.Context, cmd appointcommand.RescheduleApptCommand) icqrs.CommandResult {
	return bus.handler.RescheduleAppointment(ctx, cmd)
}

func (bus *AppointmentCommandBus) CancelAppointment(ctx context.Context, cmd appointcommand.CancelApptCommand) icqrs.CommandResult {
	return bus.handler.CancelAppointment(ctx, cmd)
}

func (bus *AppointmentCommandBus) ConfirmAppointment(ctx context.Context, cmd appointcommand.ConfirmApptCommand) icqrs.CommandResult {
	return bus.handler.ConfirmAppointment(ctx, cmd)
}

func (bus *AppointmentCommandBus) MarkAppointmentAsNotAttend(ctx context.Context, cmd appointcommand.NotAttendApptCommand) icqrs.CommandResult {
	return bus.handler.MarkAppointmentAsNotAttend(ctx, cmd)
}

func (bus *AppointmentCommandBus) CompleteAppointment(ctx context.Context, cmd appointcommand.CompleteApptCommand) icqrs.CommandResult {
	return bus.handler.CompleteAppointment(ctx, cmd)
}

func (bus *AppointmentCommandBus) UpdateAppointment(ctx context.Context, cmd appointcommand.UpdateApptCommand) icqrs.CommandResult {
	return bus.handler.UpdateAppointment(ctx, cmd)
}
