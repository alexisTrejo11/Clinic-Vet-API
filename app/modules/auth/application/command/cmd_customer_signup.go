package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
)

// SignupCommand represents the user registration request data
type SignupCommand struct {
	// User credentials - at least one is required
	Email       valueobject.Email        `json:"email" validate:"omitempty,email"`
	PhoneNumber *valueobject.PhoneNumber `json:"phone_number" validate:"omitempty,phone"`
	Password    string                   `json:"password" validate:"required,min=8"`
	Role        enum.UserRole            `json:"role" validate:"required"`

	CTX context.Context `json:"-" validate:"required"`
	// Personal details
	FirstName   string            `json:"first_name" validate:"required,min=2,max=50"`
	LastName    string            `json:"last_name" validate:"required,min=2,max=50"`
	Gender      enum.PersonGender `json:"gender" validate:"required"`
	DateOfBirth time.Time         `json:"date_of_birth" validate:"required"`
	Location    string            `json:"location" validate:"required"`
}

type SignupHandler struct {
	userRepo        repository.UserRepository
	securityService service.UserSecurityService
}

// NewSignupCommandHandler creates a new instance of signupHandler
func NewSignupCommandHandler(userRepo repository.UserRepository, securityService service.UserSecurityService) AuthCommandHandler {
	return &SignupHandler{
		userRepo:        userRepo,
		securityService: securityService,
	}
}

// Handle processes the signup command and returns authentication result
func (h *SignupHandler) Handle(cmd any) AuthCommandResult {
	command, ok := cmd.(SignupCommand)
	if !ok {
		return FailureAuthResult(ErrInvalidCommandType, errors.New("expected SignupCommand"))
	}

	if err := h.securityService.ValidateUserCredentials(command.CTX, command.Email, command.PhoneNumber, command.Password); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	user, err := h.toDomain(command)
	if err != nil {
		return FailureAuthResult(ErrDataParsingFailed, err)
	}

	if err := h.securityService.ProcessUserCreation(command.CTX, user); err != nil {
		return FailureAuthResult(ErrUserCreationFailed, err)
	}

	return SuccessAuthResult(nil, user.ID().String(), MsgUserCreatedSuccess)
}

// toDomain converts SignupCommand to domain entity with comprehensive validation
func (h *SignupHandler) toDomain(command SignupCommand) (*user.User, error) {
	userEntity, err := user.NewUser(command.Role, enum.UserStatusPending,
		user.WithEmail(command.Email),
		user.WithPhoneNumber(command.PhoneNumber),
		user.WithPassword(command.Password),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build user entity: %w", err)
	}

	return userEntity, nil
}
