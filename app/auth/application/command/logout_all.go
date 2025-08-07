package authCmd

import (
	"context"
	"errors"

	sessionRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain/repositories"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type LogoutAllCommand struct {
	UserId user.UserId     `json:"user_id"`
	CTX    context.Context `json:"-"`
}

type LogoutAllHandler interface {
	Handle(command LogoutAllCommand) error
}

type logoutAllHandlerImpl struct {
	userRepository    userRepository.UserRepository
	sessionRepository sessionRepo.SessionRepository
}

func NewLogoutAllHandler(
	userRepository userRepository.UserRepository,
	sessionRepository sessionRepo.SessionRepository,
) LogoutAllHandler {
	return logoutAllHandlerImpl{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (h logoutAllHandlerImpl) Handle(command LogoutAllCommand) error {
	user, err := h.userRepository.GetById(command.CTX, command.UserId.GetValue())
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	err = h.sessionRepository.DeleteAllUserSessions(command.CTX, user.Id().String())
	if err != nil {
		return err
	}

	return nil
}
