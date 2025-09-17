package command

import (
	"context"
	"time"

	u "clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/enum"
	event "clinic-vet-api/app/core/domain/event/user_event"
	"clinic-vet-api/app/core/domain/valueobject"
)

type EmployeeRegisterCommand struct {
	email       valueobject.Email
	password    string
	phoneNumber *valueobject.PhoneNumber
	employeeID  valueobject.EmployeeID
	ctx         context.Context
}

func NewEmployeeRegisterCommand(
	ctx context.Context,
	email valueobject.Email,
	password string,
	phoneNumber *valueobject.PhoneNumber,
	employeeID valueobject.EmployeeID,
) *EmployeeRegisterCommand {
	return &EmployeeRegisterCommand{
		email:       email,
		password:    password,
		phoneNumber: phoneNumber,
		employeeID:  employeeID,
		ctx:         ctx,
	}
}

func (h *authCommandHandler) EmployeeRegister(command EmployeeRegisterCommand) AuthCommandResult {
	if err := h.userAuthService.ValidateUserCredentials(
		command.ctx, command.email, command.phoneNumber, command.password); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	if err := h.userAuthService.ValidateEmployeeAccount(command.ctx, command.employeeID); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	user, err := command.toEntity()
	if err != nil {
		return FailureAuthResult(ErrUserCreationFailed, err)
	}

	if err := h.userAuthService.ProcessUserPersistence(command.ctx, user); err != nil {
		return FailureAuthResult(ErrUserCreationFailed, err)
	}

	go h.produceRegisterEvent(user)

	return SuccessAuthResult(nil, user.ID().String(), "user successfully created")
}

func (command *EmployeeRegisterCommand) toEntity() (*u.User, error) {
	user, err := u.CreateUser(
		enum.UserRoleVeterinarian,
		enum.UserStatusPending,
		u.WithEmail(command.email),
		u.WithPassword(command.password),
		u.WithPhoneNumber(command.phoneNumber),
		u.WithEmployeeID(command.employeeID),
		u.WithJoinedAt(time.Now()),
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *authCommandHandler) produceRegisterEvent(user *u.User) {
	event := &event.UserRegisteredEvent{
		UserID: user.ID(),
		Email:  user.Email(),
		Role:   user.Role(),
	}
	h.userEvent.Registered(*event)
}
