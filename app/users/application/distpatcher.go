package userApplication

import (
	"fmt"
	"reflect"

	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
)

type CommandHandler interface {
	Handle(command any) userCommand.CommandResult
}

type CommandDispatcher struct {
	handlers map[reflect.Type]CommandHandler
}

func NewCommandDispatcher() *CommandDispatcher {
	return &CommandDispatcher{
		handlers: make(map[reflect.Type]CommandHandler),
	}
}

func (d *CommandDispatcher) Register(command any, handler CommandHandler) {
	commandType := reflect.TypeOf(command)
	d.handlers[commandType] = handler
}

func (d *CommandDispatcher) Dispatch(command any) userCommand.CommandResult {
	commandType := reflect.TypeOf(command)
	handler, ok := d.handlers[commandType]
	if !ok {
		return userCommand.FailureResult(
			"an ocurred dispatching command", fmt.Errorf("no handler registered for command type %s", commandType),
		)
	}
	return handler.Handle(command)
}
