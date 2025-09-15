package bus

import (
	"fmt"
	"reflect"

	repository "clinic-vet-api/app/core/repositories"
	"clinic-vet-api/app/core/service"
	"clinic-vet-api/app/modules/users/application/usecase/command"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/password"
)

type UserCommandBus struct {
	handlers map[reflect.Type]cqrs.CommandHandler
}

func NewUserCommandBus(userRepo repository.UserRepository) *UserCommandBus {
	bus := &UserCommandBus{
		handlers: make(map[reflect.Type]cqrs.CommandHandler),
	}
	bus.RegisterCommands(userRepo)
	return bus
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
			"an occurred dispatching command", fmt.Errorf("no handler registered for command type %s", commandType),
		)
	}
	return handler.Handle(command)
}

func (d *UserCommandBus) RegisterCommands(userRepo repository.UserRepository) {
	d.Register(reflect.TypeOf(command.ChangePasswordCommand{}), command.NewChangePasswordHandler(userRepo, &password.PasswordEncoderImpl{}))
	d.Register(reflect.TypeOf(command.ChangeEmailCommand{}), command.NewChangeEmailHandler(userRepo))
	d.Register(reflect.TypeOf(command.ChangePasswordCommand{}), command.NewChangePasswordHandler(userRepo, &password.PasswordEncoderImpl{}))
	d.Register(reflect.TypeOf(command.DeleteUserCommand{}), command.NewDeleteUserHandler(userRepo))
	d.Register(reflect.TypeOf(command.CreateUserCommand{}), command.NewCreateUserHandler(userRepo, *service.NewUserSecurityService(userRepo)))
}
