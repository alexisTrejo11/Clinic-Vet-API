package register

import (
	"context"
	"fmt"
	"time"

	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	commondto "clinic-vet-api/app/shared/dto"
)

type CustomerRegisterCommand struct {
	Email       valueobject.Email
	PhoneNumber *valueobject.PhoneNumber
	Password    string
	Role        enum.UserRole

	FirstName   string
	LastName    string
	Gender      enum.PersonGender
	DateOfBirth time.Time
	Location    string
}

func (h *registerCommandHandler) CustomerRegister(ctx context.Context, command CustomerRegisterCommand) result.AuthCommandResult {
	if err := h.userAuthService.ValidateUserCredentials(ctx, command.Email, command.PhoneNumber, command.Password); err != nil {
		return result.AuthFailure(ErrValidationFailed, err)
	}

	user, err := command.toDomain()
	if err != nil {
		return result.AuthFailure(ErrDataParsingFailed, err)
	}

	if err := h.userAuthService.ProcessUserPersistence(ctx, user); err != nil {
		return result.AuthFailure(ErrUserCreationFailed, err)
	}

	go h.ProduceEvent(user, command)

	return result.AuthSuccess(nil, user.ID().String(), MsgUserCreatedSuccess)
}

func (command *CustomerRegisterCommand) toDomain() (*user.User, error) {
	userEntity, err := user.CreateUser(command.Role, enum.UserStatusPending,
		user.WithEmail(command.Email),
		user.WithPhoneNumber(command.PhoneNumber),
		user.WithPassword(command.Password),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build user entity: %w", err)
	}

	return userEntity, nil
}

func (h *registerCommandHandler) ProduceEvent(user *user.User, cmd CustomerRegisterCommand) {
	name, _ := valueobject.NewPersonName(cmd.FirstName, cmd.LastName)
	event := event.UserRegisteredEvent{
		UserID: user.ID(),
		Role:   user.Role(),
		Email:  user.Email(),
		PersonalData: commondto.PersonalData{
			Name:        name,
			Gender:      cmd.Gender,
			DateOfBirth: cmd.DateOfBirth, Location: cmd.Location,
		},
	}

	h.userEvent.Registered(event)
}
