package authCmd

import (
	"context"
	"errors"

	sessionRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain/repositories"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type LogoutCommand struct {
	UserId       int
	RefreshToken string
	CTX          context.Context
}

type LogoutHandler interface {
	Handle(command LogoutCommand) error
}

type logoutCommandHandler struct {
	userRepository userRepo.UserRepository
	sessionRepo    sessionRepo.SessionRepository
}

func NewLogoutCommandHandler(userRepository userRepo.UserRepository) LogoutHandler {
	return &logoutCommandHandler{
		userRepository: userRepository,
	}
}

func (h *logoutCommandHandler) Handle(command LogoutCommand) error {
	user, err := h.userRepository.GetById(command.CTX, command.UserId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	err = h.sessionRepo.DeleteUserSession(command.CTX, string(command.UserId), command.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}
