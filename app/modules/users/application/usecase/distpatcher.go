package usecase

import (
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
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

func (d *CommandDispatcher) RegisterCurrentCommands(userRepo repository.UserRepository) {
	d.Register(command.ChangePasswordCommand{}, command.NewChangePasswordHandler(userRepo))
	d.Register(command.ChangeEmailCommand{}, command.NewChangeEmailHandler(userRepo))
	d.Register(command.ChangePasswordCommand{}, command.NewChangePasswordHandler(userRepo))
	d.Register(command.DeleteUserCommand{}, command.NewDeleteUserHandler(userRepo))
	d.Register(command.CreateUserCommand{}, command.NewCreateUserHandler(userRepo))
}
