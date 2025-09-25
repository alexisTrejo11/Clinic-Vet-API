package register

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type StaffRegisterCommand struct {
	email       valueobject.Email
	password    string
	phoneNumber valueobject.PhoneNumber
	employeeID  valueobject.EmployeeID
}

func NewStaffRegisterCommand(
	email valueobject.Email,
	password string,
	phoneNumber valueobject.PhoneNumber,
	employeeID valueobject.EmployeeID,
) *StaffRegisterCommand {
	return &StaffRegisterCommand{
		email:       email,
		password:    password,
		phoneNumber: phoneNumber,
		employeeID:  employeeID,
	}
}

func (h *registerCommandHandler) StaffRegister(ctx context.Context, command StaffRegisterCommand) result.AuthCommandResult {
	if err := h.userAuthService.ValidateUserCredentials(
		ctx, command.email, command.phoneNumber, command.password); err != nil {
		return result.AuthFailure(ErrValidationFailed, err)
	}

	employee, err := h.userAuthService.ValidateEmployeeAccount(ctx, command.employeeID)
	if err != nil {
		return result.AuthFailure(ErrValidationFailed, err)
	}

	user, err := command.toEntity()
	if err != nil {
		return result.AuthFailure(ErrUserCreationFailed, err)
	}

	if err := h.userAuthService.ProcessUserPersistence(ctx, user); err != nil {
		return result.AuthFailure(ErrUserCreationFailed, err)
	}

	go h.produceRegisterEvent(user, employee)

	return result.AuthSuccess(nil, user.ID().String(), "user successfully created")
}

func (command *StaffRegisterCommand) toEntity() (*u.User, error) {
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

func (h *registerCommandHandler) produceRegisterEvent(user *u.User, employee *employee.Employee) {
	event := &event.UserRegisteredEvent{
		UserID:   user.ID(),
		Email:    user.Email(),
		Role:     user.Role(),
		Employee: employee,
	}
	h.userEvent.Registered(*event)
}
