package command

import (
	"context"
	"time"

	u "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type EmployeeSignupCommand struct {
	email       valueobject.Email
	password    string
	phoneNumber *valueobject.PhoneNumber
	employeeID  valueobject.VetID
	ctx         context.Context
}

func NewEmployeeSignupCommand(
	ctx context.Context,
	email valueobject.Email,
	password string,
	phoneNumber *valueobject.PhoneNumber,
	employeeID valueobject.VetID,
) *EmployeeSignupCommand {
	return &EmployeeSignupCommand{
		email:       email,
		password:    password,
		phoneNumber: phoneNumber,
		employeeID:  employeeID,
		ctx:         ctx,
	}
}

type EmployeeSignupHandler struct {
	userRepository      repository.UserRepository
	employeeRepository  repository.VetRepository
	userSecurityService service.UserSecurityService
}

func NewEmployeeSignupHandler(
	userRepository repository.UserRepository,
	employeeRepository repository.VetRepository,
	userSecurityService service.UserSecurityService,
) *EmployeeSignupHandler {
	return &EmployeeSignupHandler{
		userRepository:      userRepository,
		employeeRepository:  employeeRepository,
		userSecurityService: userSecurityService,
	}
}

func (h *EmployeeSignupHandler) Handle(cmd cqrs.Command) AuthCommandResult {
	command, ok := cmd.(*EmployeeSignupCommand)
	if !ok {
		return FailureAuthResult(ErrAuthenticationFailed, ErrInvalidCommand)
	}

	if err := h.userSecurityService.ValidateUserCredentials(
		command.ctx, command.email, command.phoneNumber, command.password); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	if err := h.userSecurityService.ValidateEmployeeAccount(command.ctx, command.employeeID); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	user, err := h.createUser(command)
	if err != nil {
		return FailureAuthResult(ErrUserCreationFailed, err)
	}

	if err := h.userSecurityService.ProcessUserCreation(command.ctx, user); err != nil {
		return FailureAuthResult(ErrUserCreationFailed, err)
	}

	return SuccessAuthResult(nil, user.ID().String(), "user successfully created")
}

func (h *EmployeeSignupHandler) createUser(command *EmployeeSignupCommand) (*u.User, error) {
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
