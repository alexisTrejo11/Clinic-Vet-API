package command

import (
	"context"
	"fmt"
	"time"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type CustomerRegisterCommand struct {
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

func (h *authCommandHandler) CustomerRegister(command CustomerRegisterCommand) AuthCommandResult {
	if err := h.userAuthService.ValidateUserCredentials(command.CTX, command.Email, command.PhoneNumber, command.Password); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	user, err := command.toDomain()
	if err != nil {
		return FailureAuthResult(ErrDataParsingFailed, err)
	}

	if err := h.userAuthService.ProcessUserCreation(command.CTX, user); err != nil {
		return FailureAuthResult(ErrUserCreationFailed, err)
	}

	return SuccessAuthResult(nil, user.ID().String(), MsgUserCreatedSuccess)
}

func (command *CustomerRegisterCommand) toDomain() (*user.User, error) {
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
