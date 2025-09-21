package bus

import (
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

func (mh *MedicalSessionCommandBus) CreateMedicalSession(cmd command.CreateMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.CreateMedicalSession(cmd)
}

func (mh *MedicalSessionCommandBus) UpdateMedicalSession(cmd command.UpdateMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.UpdateMedicalSession(cmd)
}

func (mh *MedicalSessionCommandBus) SoftDeleteMedicalSession(cmd command.SoftDeleteMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.SoftDeleteMedicalSession(cmd)
}

func (mh *MedicalSessionCommandBus) HardDeleteMedicalSession(cmd command.HardDeleteMedSessionCommand) cqrs.CommandResult {
	return mh.handlers.HardDeleteMedicalSession(cmd)
}
