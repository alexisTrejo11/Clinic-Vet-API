package bus

import (
	appointcommand "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	icqrs "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type AppointmentCommandBus struct {
	handler appointcommand.AppointementCommandHandler
}

func (bus *AppointmentCommandBus) RequestAppointmentByCustomer(cmd appointcommand.RequestApptByCustomerCommand) icqrs.CommandResult {
	return bus.handler.RequestAppointmentByCustomer(cmd)
}

func (bus *AppointmentCommandBus) CreateAppointment(cmd appointcommand.CreateApptCommand) icqrs.CommandResult {
	return bus.handler.CreateAppointment(cmd)
}

func (bus *AppointmentCommandBus) DeleteAppointment(cmd appointcommand.DeleteApptCommand) icqrs.CommandResult {
	return bus.handler.DeleteAppointment(cmd)
}

func (bus *AppointmentCommandBus) RescheduleAppointment(cmd appointcommand.RescheduleApptCommand) icqrs.CommandResult {
	return bus.handler.RescheduleAppointment(cmd)
}

func (bus *AppointmentCommandBus) CancelAppointment(cmd appointcommand.CancelApptCommand) icqrs.CommandResult {
	return bus.handler.CancelAppointment(cmd)
}

func (bus *AppointmentCommandBus) ConfirmAppointment(cmd appointcommand.ConfirmApptCommand) icqrs.CommandResult {
	return bus.handler.ConfirmAppointment(cmd)
}

func (bus *AppointmentCommandBus) MarkAppointmentAsNotAttend(cmd appointcommand.NotAttendApptCommand) icqrs.CommandResult {
	return bus.handler.MarkAppointmentAsNotAttend(cmd)
}

func (bus *AppointmentCommandBus) CompleteAppointment(cmd appointcommand.CompleteApptCommand) icqrs.CommandResult {
	return bus.handler.CompleteAppointment(cmd)
}

func (bus *AppointmentCommandBus) UpdateAppointment(cmd appointcommand.UpdateApptCommand) icqrs.CommandResult {
	return bus.handler.UpdateAppointment(cmd)
}
