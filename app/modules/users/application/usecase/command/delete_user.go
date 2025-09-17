package command

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type DeleteUserHandler struct {
	userRepository repository.UserRepository
}

type DeleteUserCommand struct {
	userID     valueobject.UserID
	softDelete bool
}

func NewDeleteUserCommand(userID uint, softDelete bool) DeleteUserCommand {
	uid := valueobject.NewUserID(userID)
	cmd := &DeleteUserCommand{
		userID:     uid,
		softDelete: softDelete,
	}
	return *cmd
}

func NewDeleteUserHandler(userRepo repository.UserRepository) *DeleteUserHandler {
	return &DeleteUserHandler{
		userRepository: userRepo,
	}
}

func (d *DeleteUserHandler) Handle(ctx context.Context, command DeleteUserCommand) cqrs.CommandResult {
	if _, err := d.userRepository.FindByID(ctx, command.userID); err != nil {
		return *cqrs.FailureResult("failed to find user", err)
	}

	if command.softDelete {
		err := d.userRepository.SoftDelete(ctx, command.userID)
		if err != nil {
			return *cqrs.FailureResult("failed to delete user", err)
		}
	}

	err := d.userRepository.HardDelete(ctx, command.userID)
	if err != nil {
		return *cqrs.FailureResult("failed to delete user", err)
	}

	return *cqrs.SuccessResult(command.userID.String(), "user deleted successfully")
}
