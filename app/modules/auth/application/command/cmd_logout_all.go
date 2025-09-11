package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
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

type logoutAllHandler struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func NewLogoutAllHandler(userRepository repository.UserRepository, sessionRepository repository.SessionRepository) *logoutAllHandler {
	return &logoutAllHandler{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (h *logoutAllHandler) Handle(cmd any) AuthCommandResult {
	command, ok := cmd.(LogoutAllCommand)
	if !ok {
		return FailureAuthResult(ErrAuthenticationFailed, errors.New("invalid command type"))
	}

	user, err := h.userRepository.FindByID(command.ctx, command.userID)
	if err != nil {
		return FailureAuthResult("an error ocurred finding user", err)
	}

	err = h.sessionRepository.DeleteAllUserSessions(command.ctx, user.ID())
	if err != nil {
		return FailureAuthResult("an error ocurred delete user sessions", err)
	}

	return SuccessAuthResult(nil, user.ID().String(), "user sessions sucessfully deleted on all devices")
}
