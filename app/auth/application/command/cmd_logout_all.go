package authCmd

import (
	"context"
	"errors"

	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
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

func (h *logoutAllHandler) Handle(command LogoutAllCommand) AuthCommandResult {
	user, err := h.userRepository.GetById(command.CTX, command.UserId)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("user not found", err)}
	}

	if user == nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("user not found", errors.New("user not found"))}
	}

	err = h.sessionRepository.DeleteAllUserSessions(command.CTX, user.Id().String())
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("failed to delete user sessions", err)}
	}

	return AuthCommandResult{CommandResult: shared.SuccessResult("", "user logged out from all devices")}
}
