package command

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/core/service"

	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
)

type CreateUserCommand struct {
	email       valueobject.Email
	phoneNumber *valueobject.PhoneNumber
	password    string
	role        enum.UserRole
	status      enum.UserStatus
	ctx         context.Context
}

func NewCreateUserCommand(ctx context.Context, email string, phoneNumber *string, password, role string, status string) (CreateUserCommand, error) {
	var errorsMessages []string

	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	var phoneNumberVO *valueobject.PhoneNumber
	if phoneNumber != nil {
		phoneNumber, err := valueobject.NewPhoneNumber(*phoneNumber)
		if err != nil {
			errorsMessages = append(errorsMessages, err.Error())
		}

		phoneNumberVO = &phoneNumber
	}

	userRole, err := enum.ParseUserRole(role)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	userStatus, err := enum.ParseUserStatus(status)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	if password == "" {
		errorsMessages = append(errorsMessages, "password cannot be empty")
	}

	if len(errorsMessages) > 0 {
		return CreateUserCommand{}, apperror.MappingError(errorsMessages, "constructor", "Create User Command", "user")
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
		return *cqrs.FailureResult(ErrFailedMappingUser, errors.New("invalid command type"))
	}

	user, err := fromCreateCommand(command)
	if err != nil {
		return *cqrs.FailureResult(ErrFailedMappingUser, err)
	}

	err = uc.securityService.ValidateUserCredentials(command.ctx, command.email, command.phoneNumber, command.password)
	if err != nil {
		return *cqrs.FailureResult(ErrFailedValidatingUser, err)
	}

	if err := uc.securityService.ProcessUserCreation(command.ctx, user); err != nil {
		return *cqrs.FailureResult(ErrFailedProcessingUser, err)
	}
	return *cqrs.SuccessResult(user.ID().String(), ErrUserCreationSuccess)
}
