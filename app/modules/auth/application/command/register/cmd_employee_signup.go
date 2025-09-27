package register

import (
	"context"
	"errors"

	"clinic-vet-api/app/modules/auth/application/command/result"
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
)

type StaffRegisterCommand struct {
	email       vo.Email
	password    string
	role        enum.UserRole
	phoneNumber *vo.PhoneNumber
	employeeID  vo.EmployeeID
}

func NewStaffRegisterCommand(
	email vo.Email, password string, phoneNumber *vo.PhoneNumber, employeeID vo.EmployeeID,
) *StaffRegisterCommand {
	return &StaffRegisterCommand{
		email:       email,
		password:    password,
		phoneNumber: phoneNumber,
		employeeID:  employeeID,
	}
}

func (h *registerCommandHandler) StaffRegister(ctx context.Context, cmd StaffRegisterCommand) result.AuthCommandResult {
	if err := cmd.Validate(); err != nil {
		return result.AuthFailure(ErrValidationFailed, err)
	}

	if err := h.userAuthService.ValidateUserCredentials(ctx, cmd.email, cmd.phoneNumber, cmd.password); err != nil {
		return result.AuthFailure(ErrValidationFailed, err)
	}

	employee, err := h.userAuthService.ValidateEmployeeAccount(ctx, cmd.employeeID)
	if err != nil {
		return result.AuthFailure(ErrValidationFailed, err)
	}

	user := cmd.toEntity()
	if err := h.userAuthService.ProcessUserPersistence(ctx, user); err != nil {
		return result.AuthFailure(ErrUserCreationFailed, err)
	}

	go h.produceRegisterEvent(user, employee)

	return result.AuthSuccess(nil, user.ID().String(), "user successfully created")
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

func (cmd *StaffRegisterCommand) toEntity() *u.User {
	user := u.NewUserBuilder().
		WithRole(cmd.role).
		WithEmail(cmd.email).
		WithPassword(cmd.password).
		WithPhoneNumber(cmd.phoneNumber).
		WithEmployeeID(&cmd.employeeID).
		Build()
	return user
}

func (cmd *StaffRegisterCommand) Validate() error {
	if !cmd.role.IsValid() {
		return errors.New("invalid role")
	}
	if !cmd.role.IsStaff() {
		return errors.New("role must be a staff role")
	}

	if cmd.employeeID.IsZero() {
		return errors.New("employee ID cannot be empty")
	}

	if !cmd.email.IsValid() {
		return errors.New("invalid email format")
	}

	if cmd.password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}
