// Package command contains all the implementation related to auth operations
package command

import "reflect"

type AuthCommandHandler interface {
	Handle(cmd any) AuthCommandResult
}

type AuthCommandBus interface {
	Register(commandType reflect.Type, handler AuthCommandHandler) error
	Dispatch(command any) AuthCommandResult
}
