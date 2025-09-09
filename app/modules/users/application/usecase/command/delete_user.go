package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type DeleteUserHandler struct {
	userRepository repository.UserRepository
}

type DeleteUserCommand struct {
	userID     valueobject.UserID
	softDelete bool
	ctx        context.Context
}

func NewDeleteUserCommand(ctx context.Context, userID int, softDelete bool) (DeleteUserCommand, error) {
	uid, err := valueobject.NewUserID(userID)
	if err != nil {
		return DeleteUserCommand{}, err
	}
	cmd := &DeleteUserCommand{
		userID:     uid,
		softDelete: softDelete,
		ctx:        ctx,
	}
	return *cmd, nil
}

func NewDeleteUserHandler(userRepo repository.UserRepository) cqrs.CommandHandler {
	return &DeleteUserHandler{
		userRepository: userRepo,
	}
}

func (d *DeleteUserHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(DeleteUserCommand)

	if _, err := d.userRepository.GetByID(command.ctx, command.userID); err != nil {
		return cqrs.FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(command.ctx, command.userID, command.softDelete)
	if err != nil {
		return cqrs.FailureResult("failed to delete user", err)
	}

	return cqrs.SuccessResult(command.userID.String(), "user deleted successfully")
}
