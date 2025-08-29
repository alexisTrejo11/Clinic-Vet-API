package command

import (
	"errors"
	"fmt"
	"reflect"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
)

type AuthCommandHandler interface {
	Handle(cmd any) AuthCommandResult
}

type AuthCommandBus interface {
	Register(commandType reflect.Type, handler AuthCommandHandler) error
	Dispatch(command any) AuthCommandResult
}

type authCommandBus struct {
	handlers map[reflect.Type]AuthCommandHandler
}

func NewAuthCommandBus(
	sessionRepo repository.SessionRepository,
	userRepo repository.UserRepository,
	jwtService jwt.JWTService,
) AuthCommandBus {
	bus := &authCommandBus{
		handlers: make(map[reflect.Type]AuthCommandHandler),
	}

	// Registra los handlers correctamente
	bus.Register(reflect.TypeOf(RefreshSessionCommand{}), NewRefreshSessionHandler(userRepo, sessionRepo, jwtService))
	bus.Register(reflect.TypeOf(SignupCommand{}), NewSignupCommandHandler(userRepo))
	bus.Register(reflect.TypeOf(LogoutCommand{}), NewLogoutHandler(userRepo, sessionRepo))
	bus.Register(reflect.TypeOf(LoginCommand{}), NewLoginHandler(userRepo, sessionRepo, jwtService))
	bus.Register(reflect.TypeOf(LogoutAllCommand{}), NewLogoutAllHandler(userRepo, sessionRepo))

	return bus
}

func (bus *authCommandBus) Register(commandType reflect.Type, handler AuthCommandHandler) error {
	if handler == nil {
		return fmt.Errorf("handler is nil")
	}

	bus.handlers[commandType] = handler
	return nil
}

func (bus *authCommandBus) Dispatch(command any) AuthCommandResult {
	commandType := reflect.TypeOf(command)
	handler, exists := bus.handlers[commandType]
	if !exists {
		return FailureAuthResult(
			fmt.Sprintf("No handler for command type: %s", commandType.Name()),
			errors.New("unhandled command type"),
		)
	}

	// Usa una interfaz común en lugar de type assertions específicos
	result := handler.Handle(command)

	return result
}
