package bus

import (
	"fmt"
	"reflect"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	appointcommand "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	icqrs "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	infraerr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure"
)

type appointmentCommandBus struct {
	handlers map[reflect.Type]icqrs.CommandHandler
}

func NewApptCommandBus(appointmentRepo repository.AppointmentRepository) icqrs.CommandBus {
	bus := &appointmentCommandBus{
		handlers: make(map[reflect.Type]icqrs.CommandHandler),
	}

	bus.registerHandlers(appointmentRepo)
	return bus
}

func (bus *appointmentCommandBus) registerHandlers(appointmentRepo repository.AppointmentRepository) {
	bus.Register(reflect.TypeOf(appointcommand.CreateApptCommand{}), appointcommand.NewCreateApptHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.UpdateApptCommand{}), appointcommand.NewUpdateApptHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.DeleteApptCommand{}), appointcommand.NewDeleteApptHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.RescheduleApptCommand{}), appointcommand.NewRescheduleApptHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.ConfirmApptCommand{}), appointcommand.NewConfirmApptHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.CancelApptCommand{}), appointcommand.NewCancelApptHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.CompleteApptCommand{}), appointcommand.NewCompleteApptHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(appointcommand.NotAttendApptCommand{}), appointcommand.NewNotAttendApptHandler(appointmentRepo))
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
	case appointcommand.CreateApptCommand:
		h := handler.(*appointcommand.CreateApptHandler)
		return h.Handle(cmd)

	case appointcommand.UpdateApptCommand:
		h := handler.(*appointcommand.UpdateApptHandler)
		return h.Handle(cmd)

	case appointcommand.DeleteApptCommand:
		h := handler.(*appointcommand.DeleteApptHandler)
		return h.Handle(cmd)

	case appointcommand.RescheduleApptCommand:
		h := handler.(*appointcommand.RescheduleApptHandler)
		return h.Handle(cmd)

	case appointcommand.ConfirmApptCommand:
		h := handler.(*appointcommand.ConfirmApptHandler)
		return h.Handle(cmd)

	case appointcommand.CancelApptCommand:
		h := handler.(*appointcommand.CancelApptHandler)
		return h.Handle(cmd)

	case appointcommand.CompleteApptCommand:
		h := handler.(*appointcommand.CompleteApptHandler)
		return h.Handle(cmd)

	case appointcommand.NotAttendApptCommand:
		h := handler.(*appointcommand.NotAttendApptHandler)
		return h.Handle(cmd)

	default:
		return icqrs.FailureResult(
			"not supported command operation",
			infraerr.InvalidCommandErr("unhandled command operation", commandType.Name(), "appointment"),
		)
	}
}

type ApptCommandService struct {
	commandBus      icqrs.CommandBus
	appointmentRepo repository.AppointmentRepository
}

func NewApptCommandService(commandBus icqrs.CommandBus, appointmentRepo repository.AppointmentRepository) *ApptCommandService {
	return &ApptCommandService{
		commandBus:      commandBus,
		appointmentRepo: appointmentRepo,
	}
}

// Convenience methods for common operations
func (s *ApptCommandService) CreateAppointment(cmd appointcommand.CreateApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *ApptCommandService) UpdateAppointment(cmd appointcommand.UpdateApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *ApptCommandService) DeleteAppointment(cmd appointcommand.DeleteApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *ApptCommandService) RescheduleAppointment(cmd appointcommand.RescheduleApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *ApptCommandService) ConfirmAppointment(cmd appointcommand.ConfirmApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *ApptCommandService) CancelAppointment(cmd appointcommand.CancelApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *ApptCommandService) CompleteAppointment(cmd appointcommand.CompleteApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}

func (s *ApptCommandService) MarkAsNotPresented(cmd appointcommand.NotAttendApptCommand) icqrs.CommandResult {
	return s.commandBus.Execute(cmd)
}
