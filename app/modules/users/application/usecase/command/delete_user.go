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
	ctx        context.Context
}

func NewDeleteUserCommand(ctx context.Context, userID uint, softDelete bool) DeleteUserCommand {
	uid := valueobject.NewUserID(userID)
	cmd := &DeleteUserCommand{
		userID:     uid,
		softDelete: softDelete,
		ctx:        ctx,
	}
	return *cmd
}

func NewDeleteUserHandler(userRepo repository.UserRepository) cqrs.CommandHandler {
	return &DeleteUserHandler{
		userRepository: userRepo,
	}
}

func (d *DeleteUserHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(DeleteUserCommand)
	if !ok {
		return *cqrs.FailureResult("invalid command type", nil)
	}

	if _, err := d.userRepository.FindByID(command.ctx, command.userID); err != nil {
		return *cqrs.FailureResult("failed to find user", err)
	}

	if command.softDelete {
		err := d.userRepository.SoftDelete(command.ctx, command.userID)
		if err != nil {
			return *cqrs.FailureResult("failed to delete user", err)
		}
	}

	err := d.userRepository.HardDelete(command.ctx, command.userID)
	if err != nil {
		return *cqrs.FailureResult("failed to delete user", err)
	}

	return *cqrs.SuccessResult(command.userID.String(), "user deleted successfully")
}
