package appointmentCmd

import (
	"context"
	"fmt"
	"reflect"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type Command interface{}

type CommandHandler interface{}

type CommandBus interface {
	Register(commandType reflect.Type, handler CommandHandler) error
	Execute(ctx context.Context, command Command) shared.CommandResult
}

type appointmentCommandBus struct {
	handlers map[reflect.Type]CommandHandler
}

func NewAppointmentCommandBus(appointmentRepo appointmentDomain.AppointmentRepository) CommandBus {
	bus := &appointmentCommandBus{
		handlers: make(map[reflect.Type]CommandHandler),
	}

	bus.registerHandlers(appointmentRepo)
	return bus
}

func (bus *appointmentCommandBus) registerHandlers(appointmentRepo appointmentDomain.AppointmentRepository) {
	bus.Register(reflect.TypeOf(CreateAppointmentCommand{}), NewCreateAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(UpdateAppointmentCommand{}), NewUpdateAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(DeleteAppointmentCommand{}), NewDeleteAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(RescheduleAppointmentCommand{}), NewRescheduleAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(ConfirmAppointmentCommand{}), NewConfirmAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(CancelAppointmentCommand{}), NewCancelAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(CompleteAppointmentCommand{}), NewCompleteAppointmentHandler(appointmentRepo))
	bus.Register(reflect.TypeOf(MarkAsNotPresentedCommand{}), NewMarkAsNotPresentedHandler(appointmentRepo))
}

func (bus *appointmentCommandBus) Register(commandType reflect.Type, handler CommandHandler) error {
	if handler == nil {
		return fmt.Errorf("handler cannot be nil for command type %s", commandType.Name())
	}

	bus.handlers[commandType] = handler
	return nil
}

func (bus *appointmentCommandBus) Execute(ctx context.Context, command Command) shared.CommandResult {
	commandType := reflect.TypeOf(command)
	handler, exists := bus.handlers[commandType]

	if !exists {
		return shared.FailureResult(
			"unhandled command type",
			fmt.Errorf("no handler registered for command type %s", commandType.Name()),
		)
	}

	switch cmd := command.(type) {
	case CreateAppointmentCommand:
		h := handler.(CreateAppointmentHandler)
		return h.Handle(ctx, cmd)

	case UpdateAppointmentCommand:
		h := handler.(UpdateAppointmentHandler)
		return h.Handle(ctx, cmd)

	case DeleteAppointmentCommand:
		h := handler.(DeleteAppointmentHandler)
		return h.Handle(ctx, cmd)

	case RescheduleAppointmentCommand:
		h := handler.(RescheduleAppointmentHandler)
		return h.Handle(ctx, cmd)

	case ConfirmAppointmentCommand:
		h := handler.(ConfirmAppointmentHandler)
		return h.Handle(ctx, cmd)

	case CancelAppointmentCommand:
		h := handler.(CancelAppointmentHandler)
		return h.Handle(ctx, cmd)

	case CompleteAppointmentCommand:
		h := handler.(CompleteAppointmentHandler)
		return h.Handle(ctx, cmd)

	case MarkAsNotPresentedCommand:
		h := handler.(MarkAsNotPresentedHandler)
		return h.Handle(ctx, cmd)

	default:
		return shared.FailureResult("not registered command", fmt.Errorf("unknown command type: %s", commandType.Name()))
	}
}

type AppointmentCommandService struct {
	commandBus      CommandBus
	appointmentRepo appointmentDomain.AppointmentRepository
}

func NewAppointmentCommandService(commandBus CommandBus, appointmentRepo appointmentDomain.AppointmentRepository) *AppointmentCommandService {
	return &AppointmentCommandService{
		commandBus:      commandBus,
		appointmentRepo: appointmentRepo,
	}
}

// Convenience methods for common operations
func (s *AppointmentCommandService) CreateAppointment(command CreateAppointmentCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}

func (s *AppointmentCommandService) UpdateAppointment(command UpdateAppointmentCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}

func (s *AppointmentCommandService) DeleteAppointment(command DeleteAppointmentCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}

func (s *AppointmentCommandService) RescheduleAppointment(command RescheduleAppointmentCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}

func (s *AppointmentCommandService) ConfirmAppointment(command ConfirmAppointmentCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}

func (s *AppointmentCommandService) CancelAppointment(command CancelAppointmentCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}

func (s *AppointmentCommandService) CompleteAppointment(command CompleteAppointmentCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}

func (s *AppointmentCommandService) MarkAsNotPresented(command MarkAsNotPresentedCommand) shared.CommandResult {
	return s.commandBus.Execute(context.Background(), command)
}
