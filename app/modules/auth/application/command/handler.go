// Package command contains command handlers for authentication operations
package command

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/core/service"
	"clinic-vet-api/app/modules/auth/application/jwt"
	"clinic-vet-api/app/shared/password"
)

type AuthCommandHandler interface {
	CustomerRegister(command CustomerRegisterCommand) AuthCommandResult
	EmployeeRegister(command EmployeeRegisterCommand) AuthCommandResult
	Login(command LoginCommand) AuthCommandResult
	RefreshSession(command RefreshSessionCommand) AuthCommandResult
	Logout(command LogoutCommand) AuthCommandResult
	LogoutAll(command LogoutAllCommand) AuthCommandResult
}

type authCommandHandler struct {
	userRepository  repository.UserRepository
	userAuthService service.UserSecurityService
	sessionRepo     repository.SessionRepository
	jwtService      jwt.JWTService
	passwordEncoder password.PasswordEncoder
}

func NewAuthCommandHandler(
	userRepository repository.UserRepository,
	userAuthService service.UserSecurityService,
	sessionRepo repository.SessionRepository,
	jwtService jwt.JWTService,
	passwordEncoder password.PasswordEncoder,
) AuthCommandHandler {
	return &authCommandHandler{
		userRepository:  userRepository,
		userAuthService: userAuthService,
		sessionRepo:     sessionRepo,
		jwtService:      jwtService,
		passwordEncoder: passwordEncoder,
	}
}
