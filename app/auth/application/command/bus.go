package authCmd

import (
	"fmt"
	"reflect"

	"errors"

	jwtService "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type AuthCommandBus interface {
	Register(commandType reflect.Type, handler shared.CommandHandler) error
	Dispatch(command any) AuthCommandResult
}

type authCommandBus struct {
	handlers map[reflect.Type]shared.CommandHandler
}

func NewAuthCommandBus(
	sessionRepo session.SessionRepository,
	userRepo userDomain.UserRepository,
	jwtService jwtService.JWTService,
) AuthCommandBus {
	bus := &authCommandBus{
		handlers: make(map[reflect.Type]shared.CommandHandler),
	}

	bus.Register(reflect.TypeOf(RefreshSessionCommand{}), NewRefreshSessionHandler(userRepo, sessionRepo, jwtService))
	bus.Register(reflect.TypeOf(SignupCommand{}), NewSignupCommandHandler(userRepo))

	return bus
}

func (bus *authCommandBus) Register(commandType reflect.Type, handler shared.CommandHandler) error {
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
		return AuthCommandResult{CommandResult: shared.FailureResult("Invalid command", errors.New("unhandled command type"))}
	}

	switch cmd := command.(type) {
	case RefreshSessionCommand:
		h := handler.(refreshSessionHandler)
		return h.Handle(cmd)
	case SignupCommand:
		h := handler.(signupHandler)
		return h.Handle(cmd)
	case LogoutCommand:
		h := handler.(logoutHandler)
		return h.Handle(cmd)
	case LoginCommand:
		h := handler.(loginHandler)
		return h.Handle(cmd)
	case LogoutAllCommand:
		h := handler.(logoutAllHandler)
		return h.Handle(cmd)
	default:
		return AuthCommandResult{CommandResult: shared.FailureResult("Invalid command type", fmt.Errorf("no handler registered for command type %s", commandType.Name()))}
	}

}
