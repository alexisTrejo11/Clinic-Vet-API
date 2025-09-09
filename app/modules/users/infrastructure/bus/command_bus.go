package bus

import (
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
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
