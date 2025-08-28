package authCmd

import (
	"context"

	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type LogoutCommand struct {
	UserId       int
	RefreshToken string
	CTX          context.Context
}

type logoutHandler struct {
	userRepository userDomain.UserRepository
	sessionRepo    session.SessionRepository
}

func NewLogoutHandler(
	userRepository userDomain.UserRepository,
	sessionRepo session.SessionRepository,
) AuthCommandHandler {
	return &logoutHandler{
		userRepository: userRepository,
		sessionRepo:    sessionRepo,
	}
}

func (h *logoutHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(LogoutCommand)

	user, err := h.userRepository.GetByID(command.CTX, command.UserId)
	if err != nil {
		return FailureAuthResult("an error ocurred finding user", err)
	}

	err = h.sessionRepo.DeleteUserSession(command.CTX, string(command.UserId), command.RefreshToken)
	if err != nil {
		return FailureAuthResult("an error ocurred deleting session", err)
	}

	return SuccessAuthResult(nil, user.Id().String(), "session successfully deleted")
}
