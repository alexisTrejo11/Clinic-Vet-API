package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
)

type LogoutCommand struct {
	UserID       valueobject.UserID
	RefreshToken string
	CTX          context.Context
}

type logoutHandler struct {
	userRepository repository.UserRepository
	sessionRepo    repository.SessionRepository
}

func NewLogoutHandler(
	userRepository repository.UserRepository,
	sessionRepo repository.SessionRepository,
) AuthCommandHandler {
	return &logoutHandler{
		userRepository: userRepository,
		sessionRepo:    sessionRepo,
	}
}

func (h *logoutHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(LogoutCommand)

	user, err := h.userRepository.GetByID(command.CTX, command.UserID)
	if err != nil {
		return FailureAuthResult("an error ocurred finding user", err)
	}

	err = h.sessionRepo.DeleteUserSession(command.CTX, command.UserID, command.RefreshToken)
	if err != nil {
		return FailureAuthResult("an error ocurred deleting session", err)
	}

	return SuccessAuthResult(nil, user.ID().String(), "session successfully deleted")
}
