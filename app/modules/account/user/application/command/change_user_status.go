package command

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type ChangeUserStatusCommand struct {
	userID valueobject.UserID
	status enum.UserStatus
}

func NewChangeUserStatusCommand(userID uint, status string) (ChangeUserStatusCommand, error) {
	uid := valueobject.NewUserID(userID)
	userStatus, err := enum.ParseUserStatus(status)
	if err != nil {
		return ChangeUserStatusCommand{}, err
	}

	cmd := ChangeUserStatusCommand{
		userID: uid,
		status: userStatus,
	}
	return cmd, nil
}

type ChangeUserStatusHandler struct {
	userRepository repository.UserRepository
}

func NewChangeUserStatusHandler(userRepository repository.UserRepository) *ChangeUserStatusHandler {
	return &ChangeUserStatusHandler{
		userRepository: userRepository,
	}
}

func (h *ChangeUserStatusHandler) Handle(ctx context.Context, command ChangeUserStatusCommand) cqrs.CommandResult {
	user, err := h.userRepository.FindByID(ctx, command.userID)
	if err != nil {
		return cqrs.FailureResult("error finding user", err)
	}

	switch command.status {
	case enum.UserStatusActive:
		if err := user.Activate(); err != nil {
			return cqrs.FailureResult("error activating user", err)
		}
	case enum.UserStatusInactive:
		if err := user.Deactivate(); err != nil {
			return cqrs.FailureResult("error deactivating user", err)
		}
	case enum.UserStatusBanned:
		if err := user.Ban(); err != nil {
			return cqrs.FailureResult("error banning user", err)
		}
	default:
		return cqrs.FailureResult("invalid user status", nil)
	}

	err = h.userRepository.Save(ctx, &user)
	if err != nil {
		return cqrs.FailureResult("error saving user", err)
	}

	return cqrs.SuccessResult("user status changed successfully")
}
