package bus

import (
	"fmt"
	"reflect"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
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

type MedicalHistoryCommandBus struct {
	*CommandBus
	handlers command.MedicalHistoryCommandHandlers
}

func NewMedicalHistoryCommandBus(handlers command.MedicalHistoryCommandHandlers) *MedicalHistoryCommandBus {
	bus := NewCommandBus()
	return &MedicalHistoryCommandBus{
		CommandBus: bus,
		handlers:   handlers,
	}
}

func (mh *MedicalHistoryCommandBus) CreateMedicalHistory(cmd command.CreateMedHistCommand) command.CreateMedHistResult {
	return mh.handlers.CreateMedicalHistory(cmd)
}

func (mh *MedicalHistoryCommandBus) UpdateMedicalHistory(cmd command.UpdateMedHistCommand) command.UpdateMedHistResult {
	return mh.handlers.UpdateMedicalHistory(cmd)
}

func (mh *MedicalHistoryCommandBus) SoftDeleteMedicalHistory(cmd command.SoftDeleteMedHistCommand) command.DeleteMedHistResult {
	return mh.handlers.SoftDeleteMedicalHistory(cmd)
}

func (mh *MedicalHistoryCommandBus) HardDeleteMedicalHistory(cmd command.HardDeleteMedHistCommand) command.DeleteMedHistResult {
	return mh.handlers.HardDeleteMedicalHistory(cmd)
}
