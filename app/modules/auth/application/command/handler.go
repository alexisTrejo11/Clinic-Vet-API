package command

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
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
