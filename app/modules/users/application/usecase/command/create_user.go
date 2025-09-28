package command

import (
	"context"

	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"

	"clinic-vet-api/app/shared/cqrs"
)

type CreateUserCommand struct {
	email       valueobject.Email
	phoneNumber *valueobject.PhoneNumber
	password    string
	role        enum.UserRole
	status      enum.UserStatus
}

func NewCreateUserCommand(email string, phoneNumber *string, password, role string, status string) (CreateUserCommand, error) {
	var phoneNumberVO *valueobject.PhoneNumber
	if phoneNumber != nil {
		phone := valueobject.NewPhoneNumberNoErr(*phoneNumber)
		phoneNumberVO = &phone
	}
	return CreateUserCommand{
		email:       valueobject.NewEmailNoErr(email),
		phoneNumber: phoneNumberVO,
		password:    password,
		role:        enum.UserRole(role),
		status:      enum.UserStatus(status),
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

func (uc *CreateUserHandler) Handle(ctx context.Context, command CreateUserCommand) cqrs.CommandResult {
	if err := command.Validate(); err != nil {
		return *cqrs.FailureResult(ErrInvalidUserData, err)
	}

	user := command.ToEntity()
	err := uc.securityService.ValidateUserCredentials(ctx, command.email, command.phoneNumber, command.password)
	if err != nil {
		return *cqrs.FailureResult(ErrFailedValidatingUser, err)
	}

	if err := uc.securityService.ProcessUserPersistence(ctx, user); err != nil {
		return *cqrs.FailureResult(ErrFailedProcessingUser, err)
	}

	return *cqrs.SuccessCreateResult(user.ID().String(), ErrUserCreationSuccess)
}

func (cmd CreateUserCommand) ToEntity() *u.User {
	return u.NewUserBuilder().
		WithEmail(cmd.email).
		WithPassword(cmd.password).
		WithPhoneNumber(cmd.phoneNumber).
		WithStatus(cmd.status).
		WithRole(cmd.role).
		Build()
}

func (cmd CreateUserCommand) Validate() error {
	if err := cmd.email.Validate(); err != nil {
		return err
	}

	if cmd.phoneNumber != nil {
		if err := cmd.phoneNumber.Validate(); err != nil {
			return err
		}
	}

	if err := cmd.role.Validate(); err != nil {
		return err
	}

	if err := cmd.status.Validate(); err != nil {
		return err
	}

	return nil
}
