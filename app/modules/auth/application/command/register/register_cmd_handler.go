package register

import (
	"clinic-vet-api/app/modules/auth/application/command/result"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/shared/password"
	"context"
)

var (
	ErrValidationFailed   = "validation failed for unique credentials"
	ErrDataParsingFailed  = "failed to parse user data"
	ErrUserSaveFailed     = "failed to save user"
	ErrMissingCredentials = "either email or phone number is required"
	ErrUserCreationFailed = "user creation failed"
	MsgUserCreatedSuccess = "user successfully created"
)

type RegisterCommandHandler interface {
	CustomerRegister(ctx context.Context, command CustomerRegisterCommand) result.AuthCommandResult
	//AdminRegister(command AdminRegisterCommand) result.AuthCommandResult
	StaffRegister(ctx context.Context, command StaffRegisterCommand) result.AuthCommandResult

	//CustomerOAuthRegister(command CustomerOAuthRegisterCommand) result.AuthCommandResult
}

func NewRegisterCommandHandler(
	userRepo repository.UserRepository,
	passwordEncoder password.PasswordEncoder,
	userEvent event.UserEventProducer,
	serviceAuthService service.UserSecurityService,
) RegisterCommandHandler {
	return &registerCommandHandler{
		userRepo:        userRepo,
		userAuthService: serviceAuthService,
		userEvent:       userEvent,
	}
}

type registerCommandHandler struct {
	userAuthService service.UserSecurityService
	userEvent       event.UserEventProducer
	userRepo        repository.UserRepository
}
