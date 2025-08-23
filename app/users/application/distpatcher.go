package userApplication

import (
	"fmt"
	"reflect"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type CommandHandler interface {
	Handle(command any) shared.CommandResult
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

func (d *CommandDispatcher) Dispatch(command any) shared.CommandResult {
	commandType := reflect.TypeOf(command)
	handler, ok := d.handlers[commandType]
	if !ok {
		return shared.FailureResult(
			"an ocurred dispatching command", fmt.Errorf("no handler registered for command type %s", commandType),
		)
	}
	return handler.Handle(command)
}

func (d *CommandDispatcher) RegisterCurrentCommands(userRepo userDomain.UserRepository) {
	d.Register(userCommand.ChangePasswordCommand{}, userCommand.NewChangePasswordHandler(userRepo))
	d.Register(userCommand.ChangeEmailCommand{}, userCommand.NewChangeEmailHandler(userRepo))
	d.Register(userCommand.ChangePasswordCommand{}, userCommand.NewChangePasswordHandler(userRepo))
	d.Register(userCommand.DeleteUserCommand{}, userCommand.NewDeleteUserHandler(userRepo))
	d.Register(userCommand.CreateUserCommand{}, userCommand.NewCreateUserHandler(userRepo))
}
