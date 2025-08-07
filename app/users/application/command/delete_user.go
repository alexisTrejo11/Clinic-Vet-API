package userCommand

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type DeleteUserHandler struct {
	userRepository userRepository.UserRepository
}

func (d *DeleteUserHandler) Handle(ctx context.Context, userId int, softDelete bool) shared.CommandResult {
	if _, err := d.userRepository.GetByIdWithProfile(ctx, userId); err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(ctx, userId, softDelete)
	if err != nil {
		return shared.FailureResult("failed to delete user", err)
	}

	return shared.SuccesResult(strconv.Itoa(userId), "user deleted successfully")
}
