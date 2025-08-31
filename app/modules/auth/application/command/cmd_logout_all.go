package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type LogoutAllCommand struct {
	UserID valueobject.UserID `json:"user_id"`
	CTX    context.Context    `json:"-"`
}

type logoutAllHandler struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func NewLogoutAllHandler(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
) *logoutAllHandler {
	return &logoutAllHandler{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (h *logoutAllHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(LogoutAllCommand)

	user, err := h.userRepository.GetByID(command.CTX, command.UserID)
	if err != nil {
		return FailureAuthResult("an error ocurred finding user", err)
	}

	err = h.sessionRepository.DeleteAllUserSessions(command.CTX, user.ID().String())
	if err != nil {
		return FailureAuthResult("an error ocurred delete user sessions", err)
	}

	return SuccessAuthResult(nil, user.ID().String(), "user sessions sucessfully deleted on all devices")
}
