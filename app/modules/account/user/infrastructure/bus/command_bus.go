package bus

import (
	"context"

	authCmd "clinic-vet-api/app/modules/account/auth/local/application/command"
	"clinic-vet-api/app/modules/account/user/application/command"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/password"
)

type UserCommandBus interface {
	ResetPassword(ctx context.Context, c command.ResetPasswordCommand) cqrs.CommandResult
	RequestResetPassword(ctx context.Context, c command.RequestResetPasswordCommand) cqrs.CommandResult
	UpdatePassword(ctx context.Context, c command.UpdatePasswordCommand) cqrs.CommandResult
	ChangeEmail(ctx context.Context, c command.ChangeEmailCommand) cqrs.CommandResult
	ChangeUserStatus(ctx context.Context, c command.ChangeUserStatusCommand) cqrs.CommandResult

	ActivateAccount(ctx context.Context, c authCmd.ActivateAccountCommand) cqrs.CommandResult
	DeleteUser(ctx context.Context, c command.DeleteUserCommand) cqrs.CommandResult
	CreateUser(ctx context.Context, c command.CreateUserCommand) cqrs.CommandResult
}

type userCommandBus struct {
	changePasswordHandler   command.UpdatePasswordHandler
	changeEmailHandler      command.ChangeEmailHandler
	changeUserStatusHandler command.ChangeUserStatusHandler
	deleteUserHandler       command.DeleteUserHandler
	createUserHandler       command.CreateUserHandler
}

func NewUserCommandBus(userRepo repository.UserRepository, service *service.UserSecurityService) UserCommandBus {
	passwordEncoder := password.NewPasswordEncoder()

	changeUserStatusHandler := command.NewChangeUserStatusHandler(userRepo)
	changePasswordHandler := command.NewUpdatePasswordHandler(userRepo, passwordEncoder)
	changeEmailHandler := command.NewChangeEmailHandler(userRepo)
	deleteUserHandler := command.NewDeleteUserHandler(userRepo)
	createUserHandler := command.NewCreateUserHandler(userRepo, *service)

	bus := &userCommandBus{
		changePasswordHandler:   *changePasswordHandler,
		changeUserStatusHandler: *changeUserStatusHandler,
		changeEmailHandler:      changeEmailHandler,
		deleteUserHandler:       *deleteUserHandler,
		createUserHandler:       *createUserHandler,
	}

	return bus
}

func (b *userCommandBus) ResetPassword(ctx context.Context, c command.ResetPasswordCommand) cqrs.CommandResult {
	// Not implemented yet
	return cqrs.CommandResult{}
}

func (b *userCommandBus) RequestResetPassword(ctx context.Context, c command.RequestResetPasswordCommand) cqrs.CommandResult {
	// Not implemented yet
	return cqrs.CommandResult{}
}

func (b *userCommandBus) UpdatePassword(ctx context.Context, c command.UpdatePasswordCommand) cqrs.CommandResult {
	return b.changePasswordHandler.Handle(ctx, c)
}

func (b *userCommandBus) ChangeEmail(ctx context.Context, c command.ChangeEmailCommand) cqrs.CommandResult {
	return b.changeEmailHandler.Handle(ctx, c)
}

func (b *userCommandBus) DeleteUser(ctx context.Context, c command.DeleteUserCommand) cqrs.CommandResult {
	return b.deleteUserHandler.Handle(ctx, c)
}

func (b *userCommandBus) CreateUser(ctx context.Context, c command.CreateUserCommand) cqrs.CommandResult {
	return b.createUserHandler.Handle(ctx, c)
}

func (b *userCommandBus) ChangeUserStatus(ctx context.Context, c command.ChangeUserStatusCommand) cqrs.CommandResult {
	return b.changeUserStatusHandler.Handle(ctx, c)
}

func (b *userCommandBus) ActivateAccount(ctx context.Context, c authCmd.ActivateAccountCommand) cqrs.CommandResult {
	// Not implemented yet
	return cqrs.CommandResult{}
}
