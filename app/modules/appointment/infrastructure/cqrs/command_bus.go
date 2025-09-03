package cqrs

import (
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	appointcommand "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	icqrs "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	infraerr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure"
)

type appointmentCommandBus struct {
	handlers map[reflect.Type]icqrs.CommandHandler
}

func NewAppointmentCommandBus(appointmentRepo repository.AppointmentRepository) icqrs.CommandBus {
	bus := &appointmentCommandBus{
		handlers: make(map[reflect.Type]icqrs.CommandHandler),
	}

	bus.registerHandlers(appointmentRepo)
	return bus
}

func (bus *appointmentCommandBus) registerHandlers(appointmentRepo repository.AppointmentRepository) {
	bus.Register(reflect.TypeOf(appointcommand.CreateAppointmentCommand{}), appointcommand.NewCreateAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.UpdateAppointmentCommand{}), appointcommand.NewUpdateAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.DeleteAppointmentCommand{}), appointcommand.NewDeleteAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.RescheduleAppointmentCommand{}), appointcommand.NewRescheduleAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.ConfirmAppointmentCommand{}), appointcommand.NewConfirmAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.CancelAppointmentCommand{}), appointcommand.NewCancelAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.CompleteAppointmentCommand{}), appointcommand.NewCompleteAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.NotAttendAppointmentCommand{}), appointcommand.NewNotAttendAppointmentHandler(appointmentRepo))
}

func (bus *appointmentCommandBus) Register(commandType reflect.Type, handler icqrs.CommandHandler) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for command type %s", commandType.Name())
	}

	bus.handlers[commandType] = handler
	return nil
}

func (bus *appointmentCommandBus) Execute(command icqrs.Command) icqrs.CommandResult {
	commandType := reflect.TypeOf(command)
	handler, exists := bus.handlers[commandType]

	if !exists {
		return icqrs.FailureResult(
			"unhandled command type",
			infraerr.InvalidCommandErr("no handler registered for command type %s", commandType.Name(), "appointment"),
		)
	}

	switch cmd := command.(type) {
	case appointcommand.CreateAppointmentCommand:
		h := handler.(*appointcommand.CreateAppointmentHandler)
		return h.Handle(cmd)

	case appointcommand.UpdateAppointmentCommand:
		h := handler.(*appointcommand.UpdateAppointmentHandler)
		return h.Handle(cmd)

	case appointcommand.DeleteAppointmentCommand:
		h := handler.(*appointcommand.DeleteAppointmentHandler)
		return h.Handle(cmd)

	case appointcommand.RescheduleAppointmentCommand:
		h := handler.(*appointcommand.RescheduleAppointmentHandler)
		return h.Handle(cmd)

	case appointcommand.ConfirmAppointmentCommand:
		h := handler.(*appointcommand.ConfirmAppointmentHandler)
		return h.Handle(cmd)

	case appointcommand.CancelAppointmentCommand:
		h := handler.(*appointcommand.CancelAppointmentHandler)
		return h.Handle(cmd)

	case appointcommand.CompleteAppointmentCommand:
		h := handler.(*appointcommand.CompleteAppointmentHandler)
		return h.Handle(cmd)

	case appointcommand.NotAttendAppointmentCommand:
		h := handler.(*appointcommand.NotAttendAppointmentHandler)
		return h.Handle(cmd)

	default:
		return icqrs.FailureResult(
			"not supported command operation",
			infraerr.InvalidCommandErr("unhandled command operation", commandType.Name(), "appointment"),
		)
	}
}

type AppointmentCommandService struct {
	commandBus      icqrs.CommandBus
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentCommandService(commandBus icqrs.CommandBus, appointmentRepo repository.AppointmentRepository) *AppointmentCommandService {
	return &AppointmentCommandService{
		commandBus:      commandBus,
		appointmentRepo: appointmentRepo,
	}
}

// Convenience methods for common operations
func (s *AppointmentCommandService) CreateAppointment(cmd appointcommand.CreateAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *AppointmentCommandService) UpdateAppointment(cmd appointcommand.UpdateAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *AppointmentCommandService) DeleteAppointment(cmd appointcommand.DeleteAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *AppointmentCommandService) RescheduleAppointment(cmd appointcommand.RescheduleAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *AppointmentCommandService) ConfirmAppointment(cmd appointcommand.ConfirmAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *AppointmentCommandService) CancelAppointment(cmd appointcommand.CancelAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *AppointmentCommandService) CompleteAppointment(cmd appointcommand.CompleteAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *AppointmentCommandService) MarkAsNotPresented(cmd appointcommand.NotAttendAppointmentCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}
