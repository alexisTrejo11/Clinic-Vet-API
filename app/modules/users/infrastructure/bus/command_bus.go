package bus

import (
	"context"

	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/modules/users/application/usecase/command"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/password"
)

type UserCommandBus interface {
	ChangePassword(ctx context.Context, c command.ChangePasswordCommand) cqrs.CommandResult
	ChangeEmail(ctx context.Context, c command.ChangeEmailCommand) cqrs.CommandResult
	ChangeUserStatus(ctx context.Context, c command.ChangeUserStatusCommand) cqrs.CommandResult
	DeleteUser(ctx context.Context, c command.DeleteUserCommand) cqrs.CommandResult
	CreateUser(ctx context.Context, c command.CreateUserCommand) cqrs.CommandResult
}

type userCommandBus struct {
	changePasswordHandler   command.ChangePasswordHandler
	changeEmailHandler      command.ChangeEmailHandler
	changeUserStatusHandler command.ChangeUserStatusHandler
	deleteUserHandler       command.DeleteUserHandler
	createUserHandler       command.CreateUserHandler
}

func NewUserCommandBus(userRepo repository.UserRepository, service *service.UserSecurityService) UserCommandBus {
	passwordEncoder := password.NewPasswordEncoder()

	changeUserStatusHandler := command.NewChangeUserStatusHandler(userRepo)
	changePasswordHandler := command.NewChangePasswordHandler(userRepo, passwordEncoder)
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

func (b *userCommandBus) ChangePassword(ctx context.Context, c command.ChangePasswordCommand) cqrs.CommandResult {
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
