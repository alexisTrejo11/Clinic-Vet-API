package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

// CreateUserCommand represents the command to create a new user.
type CreateUserCommand struct {
	email       valueobject.Email
	phoneNumber valueobject.PhoneNumber
	password    string
	role        enum.UserRole
	status      enum.UserStatus
	profile     CreateProfileCommand
	ctx         context.Context
}

// NewCreateUserCommand creates a new instance of CreateUserCommand mapping primitive data types.
func NewCreateUserCommand(ctx context.Context, email, phoneNumber, password, role string, status string) (CreateUserCommand, error) {
	var errorsMessages []string

	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	phoneNumberVO, err := valueobject.NewPhoneNumber(phoneNumber)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	userRole, err := enum.ParseUserRole(role)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	userStatus, err := enum.ParseUserStatus(status)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	if len(errorsMessages) > 0 {
		return CreateUserCommand{}, apperror.MappingError(errorsMessages, "constructor", "CreateUserCommand", "user")
	}

	if password == "" {
		errorsMessages = append(errorsMessages, "password cannot be empty")
	}

	return CreateUserCommand{
		email:       emailVO,
		phoneNumber: phoneNumberVO,
		password:    password,
		role:        userRole,
		status:      userStatus,
		ctx:         ctx,
	}, nil
}

// CreateUserHandler handles the creation of a new user.
type CreateUserHandler struct {
	repo            repository.UserRepository
	securityService service.UserSecurityService
}

func NewCreateUserHandler(repo repository.UserRepository, securityService service.UserSecurityService) *CreateUserHandler {
	return &CreateUserHandler{
		repo:            repo,
		securityService: securityService,
	}
}

func (uc *CreateUserHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(CreateUserCommand)
	if !ok {
		return cqrs.FailureResult(ErrFailedMappingUser, errors.New("invalid command type"))
	}

	user, err := fromCreateCommand(command)
	if err != nil {
		return cqrs.FailureResult(ErrFailedMappingUser, err)
	}

	if err := uc.securityService.ValidateUserCreation(
		command.ctx,
		command.email,
		command.phoneNumber,
		command.password,
	); err != nil {
		return cqrs.FailureResult(ErrFailedValidatingUser, err)
	}

	if err := uc.securityService.ProccesUserCreation(command.ctx, *user); err != nil {
		return cqrs.FailureResult(ErrFailedProcessingUser, err)
	}
	return cqrs.SuccessResult(user.ID().String(), ErrUserCreationSuccess)
}
