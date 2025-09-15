package command

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
)

type LogoutAllCommand struct {
	userID valueobject.UserID
	ctx    context.Context
}

func NewLogoutAllCommand(ctx context.Context, useridnUInt uint) *LogoutAllCommand {
	return &LogoutAllCommand{
		ctx:    ctx,
		userID: valueobject.NewUserID(useridnUInt),
	}
}

func (h *authCommandHandler) LogoutAll(command LogoutAllCommand) AuthCommandResult {
	user, err := h.userRepository.FindByID(command.ctx, command.userID)
	if err != nil {
		return FailureAuthResult("an error occurred finding user", err)
	}

	err = h.sessionRepo.DeleteAllUserSessions(command.ctx, user.ID())
	if err != nil {
		return FailureAuthResult("an error occurred delete user sessions", err)
	}

	return SuccessAuthResult(nil, user.ID().String(), "user sessions successfully deleted on all devices")
}
