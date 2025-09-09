package bus

import (
	"errors"
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
)

type authCommandBus struct {
	handlers map[reflect.Type]command.AuthCommandHandler
}

func NewAuthCommandBus(
	sessionRepo repository.SessionRepository,
	userRepo repository.UserRepository,
	jwtService jwt.JWTService,
) command.AuthCommandBus {
	bus := &authCommandBus{
		handlers: make(map[reflect.Type]command.AuthCommandHandler),
	}

	bus.Register(reflect.TypeOf(command.RefreshSessionCommand{}), command.NewRefreshSessionHandler(userRepo, sessionRepo, jwtService))
	bus.Register(reflect.TypeOf(command.SignupCommand{}), command.NewSignupCommandHandler(userRepo, *service.NewUserSecurityService(userRepo)))
	bus.Register(reflect.TypeOf(command.LogoutCommand{}), command.NewLogoutHandler(userRepo, sessionRepo))
	bus.Register(reflect.TypeOf(command.LoginCommand{}), command.NewLoginHandler(userRepo, sessionRepo, jwtService, &password.PasswordEncoderImpl{}))
	bus.Register(reflect.TypeOf(command.LogoutAllCommand{}), command.NewLogoutAllHandler(userRepo, sessionRepo))

	return bus
}

func (bus *authCommandBus) Register(commandType reflect.Type, handler command.AuthCommandHandler) error {
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}

	bus.handlers[commandType] = handler
	return nil
}

func (bus *authCommandBus) Dispatch(cmd any) command.AuthCommandResult {
	commandType := reflect.TypeOf(cmd)
	handler, exists := bus.handlers[commandType]
	if !exists {
		return command.FailureAuthResult(
			fmt.Sprintf("No handler for command type: %s", commandType.Name()),
			errors.New("unhandled command type"),
		)
	}

	result := handler.Handle(cmd)

	return result
}
