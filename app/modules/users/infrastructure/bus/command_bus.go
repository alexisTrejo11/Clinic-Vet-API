package usecase

import (
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
)

type UserCommandBus struct {
	handlers map[reflect.Type]cqrs.CommandHandler
}

func NewUserCommandBus() cqrs.CommandBus {
	return &UserCommandBus{
		handlers: make(map[reflect.Type]cqrs.CommandHandler),
	}
}

func (d *UserCommandBus) Register(commandType reflect.Type, handler cqrs.CommandHandler) error {
	cmd := reflect.TypeOf(commandType)
	d.handlers[cmd] = handler

	return nil
}

func (d *UserCommandBus) Execute(command cqrs.Command) cqrs.CommandResult {
	commandType := reflect.TypeOf(command)
	handler, ok := d.handlers[commandType]
	if !ok {
		return cqrs.FailureResult(
			"an ocurred dispatching command", fmt.Errorf("no handler registered for command type %s", commandType),
		)
	}
	return handler.Handle(command)
}

func (d *UserCommandBus) RegisterCurrentCommands(userRepo repository.UserRepository) {
	d.Register(command.ChangePasswordCommand{}, command.NewChangePasswordHandler(userRepo, &password.PasswordEncoderImpl{}))
	d.Register(command.ChangeEmailCommand{}, command.NewChangeEmailHandler(userRepo))
	d.Register(command.ChangePasswordCommand{}, command.NewChangePasswordHandler(userRepo))
	d.Register(command.DeleteUserCommand{}, command.NewDeleteUserHandler(userRepo))
	d.Register(command.CreateUserCommand{}, command.NewCreateUserHandler(userRepo))
}
