package bus

import (
	"context"
	"fmt"
	"reflect"

	"clinic-vet-api/app/modules/medical/application/command"
	"clinic-vet-api/app/shared/cqrs"
)

// CommandHandler maneja comandos específicos
type CommandHandler[T cqrs.Command, R any] interface {
	Handle(cmd T) (R, error)
}

type CommandBus struct {
	handlers map[reflect.Type]interface{}
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[reflect.Type]interface{}),
	}
}

func (cb *CommandBus) Register(cmd cqrs.Command, handler any) {
	cmdType := reflect.TypeOf(cmd)
	cb.handlers[cmdType] = handler
}

func (cb *CommandBus) Execute(cmd cqrs.Command) (any, error) {
	cmdType := reflect.TypeOf(cmd)
	handler, exists := cb.handlers[cmdType]
	if !exists {
		return nil, fmt.Errorf("no handler registered for command type: %T", cmd)
	}

	handlerValue := reflect.ValueOf(handler)
	method := handlerValue.MethodByName("Handle")
	if !method.IsValid() {
		return nil, fmt.Errorf("handler for command type %T does not have Handle method", cmd)
	}

	args := []reflect.Value{
		reflect.ValueOf(cmd),
	}

	results := method.Call(args)
	if len(results) != 1 {
		return nil, fmt.Errorf("handler method must return exactly 1 value (result)")
	}

	if !results[1].IsNil() {
		return nil, results[1].Interface().(error)
	}

	return results[0].Interface(), nil
}

type MedicalSessionCommandBus struct {
	*CommandBus
	handlers command.MedicalSessionCommandHandlers
}

func NewMedicalSessionCommandBus(handlers command.MedicalSessionCommandHandlers) *MedicalSessionCommandBus {
	bus := NewCommandBus()
	return &MedicalSessionCommandBus{
		CommandBus: bus,
		handlers:   handlers,
	}
}

func (mh *MedicalSessionCommandBus) CreateMedicalSession(ctx context.Context, cmd command.CreateMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.CreateMedicalSession(ctx, cmd)
}

func (mh *MedicalSessionCommandBus) UpdateMedicalSession(ctx context.Context, cmd command.UpdateMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.UpdateMedicalSession(ctx, cmd)
}

func (mh *MedicalSessionCommandBus) SoftDeleteMedicalSession(ctx context.Context, cmd command.SoftDeleteMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.SoftDeleteMedicalSession(ctx, cmd)
}

func (mh *MedicalSessionCommandBus) HardDeleteMedicalSession(ctx context.Context, cmd command.HardDeleteMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.HardDeleteMedicalSession(ctx, cmd)
}
