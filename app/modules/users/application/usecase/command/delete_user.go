package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type deleteUserHandler struct {
	userRepository repository.UserRepository
}

type DeleteUserCommand struct {
	UserID     valueobject.UserID `json:"user_id" validate:"required"`
	SoftDelete bool               `json:"soft_delete"`
	CTX        context.Context    `json:"-"`
}

func NewDeleteUserHandler(userRepo repository.UserRepository) DeleteUserHandler {
	return &deleteUserHandler{
		userRepository: userRepo,
	}
}

type DeleteUserHandler interface {
	Handle(cmd any) cqrs.CommandResult
}

func (d *deleteUserHandler) Handle(cmd any) cqrs.CommandResult {
	command := cmd.(DeleteUserCommand)

	if _, err := d.userRepository.GetByID(command.CTX, command.UserID); err != nil {
		return cqrs.FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(command.CTX, command.UserID, command.SoftDelete)
	if err != nil {
		return cqrs.FailureResult("failed to delete user", err)
	}

	return cqrs.SuccessResult(command.UserID.String(), "user deleted successfully")
}
