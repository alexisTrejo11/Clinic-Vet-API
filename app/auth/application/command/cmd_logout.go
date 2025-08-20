package authCmd

import (
	"context"
	"errors"
	"strconv"

	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type LogoutCommand struct {
	UserId       int
	RefreshToken string
	CTX          context.Context
}

type logoutHandler struct {
	userRepository userRepo.UserRepository
	sessionRepo    session.SessionRepository
}

func NewlogoutHandler(userRepository userRepo.UserRepository, sessionRepo session.SessionRepository) *logoutHandler {
	return &logoutHandler{
		userRepository: userRepository,
		sessionRepo:    sessionRepo,
	}
}

func (h *logoutHandler) Handle(command LogoutCommand) AuthCommandResult {
	user, err := h.userRepository.GetById(command.CTX, command.UserId)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("an error occurred while fetching user", err)}
	}

	if user == nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("user not found", errors.New("user not found"))}
	}

	err = h.sessionRepo.DeleteUserSession(command.CTX, string(command.UserId), command.RefreshToken)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("an error occurred while deleting user session", err)}
	}

	return AuthCommandResult{CommandResult: shared.SuccessResult(strconv.Itoa(command.UserId), "User logged out successfully")}
}
