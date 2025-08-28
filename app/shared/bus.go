package shared

import (
	"context"
	"reflect"
)

type Command interface{}

type CommandHandler interface {
	Handle(command Command) CommandResult
}

type CommandBus interface {
	Register(commandType reflect.Type, handler CommandHandler) error
	Execute(ctx context.Context, command Command) CommandResult
}
