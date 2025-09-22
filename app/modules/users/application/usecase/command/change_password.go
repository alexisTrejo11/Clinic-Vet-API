package command

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/password"
)

type ChangePasswordCommand struct {
	UserID          valueobject.UserID
	CurrentPassword string
	NewPassword     string
}

type ChangePasswordHandler struct {
	userRepository  repository.UserRepository
	passwordEncoder password.PasswordEncoder
}

func NewChangePasswordHandler(userRepo repository.UserRepository, passwordEncoder password.PasswordEncoder) *ChangePasswordHandler {
	return &ChangePasswordHandler{
		userRepository:  userRepo,
		passwordEncoder: passwordEncoder,
	}
}

func (c *ChangePasswordHandler) Handle(ctx context.Context, command ChangePasswordCommand) cqrs.CommandResult {
	user, err := c.userRepository.FindByID(ctx, command.UserID)
	if err != nil {
		return *cqrs.FailureResult("failed to find user", err)
	}

	if err := c.changePassword(ctx, &user, command); err != nil {
		return *cqrs.FailureResult("failed to change password", err)
	}

	if err := c.userRepository.Save(ctx, &user); err != nil {
		return *cqrs.FailureResult("failed to update user", err)
	}

	return *cqrs.SuccessResult(user.ID().String(), "password changed successfully")
}

func (c *ChangePasswordHandler) changePassword(ctx context.Context, user *user.User, command ChangePasswordCommand) error {
	valid := c.passwordEncoder.CheckPassword(user.HashedPassword(), command.CurrentPassword)
	if !valid {
		return apperror.ValidationError("Current password is incorrect")
	}

	if err := user.UpdatePassword(command.NewPassword); err != nil {
		return err
	}

	hashedPassword, err := c.passwordEncoder.HashPassword(command.NewPassword)
	if err != nil {
		return err
	}

	user.SetHashedPassword(hashedPassword)

	return c.userRepository.Save(ctx, user)
}
