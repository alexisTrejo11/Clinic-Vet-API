package authCmd

import (
	"context"

	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type LogoutAllCommand struct {
	UserId int             `json:"user_id"`
	CTX    context.Context `json:"-"`
}

type logoutAllHandler struct {
	userRepository    userDomain.UserRepository
	sessionRepository session.SessionRepository
}

func NewLogoutAllHandler(
	userRepository userDomain.UserRepository,
	sessionRepository session.SessionRepository,
) *logoutAllHandler {
	return &logoutAllHandler{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (h *logoutAllHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(LogoutAllCommand)

	user, err := h.userRepository.GetByID(command.CTX, command.UserId)
	if err != nil {
		return FailureAuthResult("an error ocurred finding user", err)
	}

	err = h.sessionRepository.DeleteAllUserSessions(command.CTX, user.Id().String())
	if err != nil {
		return FailureAuthResult("an error ocurred delete user sessions", err)
	}

	return SuccessAuthResult(nil, user.Id().String(), "user sessions sucessfully deleted on all devices")
}
