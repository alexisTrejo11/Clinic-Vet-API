// Package cqrs includes contains all interfaces to implements command and query buss
package cqrs

import (
	"reflect"
)

type Command any

type CommandHandler interface {
	Handle(command Command) CommandResult
}

type CommandBus interface {
	Register(commandType reflect.Type, handler CommandHandler) error
	Execute(command Command) CommandResult
}

type Query any

type QueryHandler[T any] interface {
	Handle(query Query) (T, error)
}

type QueryBus interface {
	Execute(query Query) (any, error)
	Register(queryType reflect.Type, handler any) error
}
