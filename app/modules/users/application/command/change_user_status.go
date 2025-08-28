package userDomainCommand

import (
	"context"
	"errors"

	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type ChangeUserStatusCommand struct {
	UserId int                   `json:"user_id"`
	Status userDomain.UserStatus `json:"status"`
	CTX    context.Context       `json:"ctx"`
}

type ChangeUserStatusHandler struct {
	userRepository userDomain.UserRepository
}

func NewChangeUserStatusHandler(userRepository userDomain.UserRepository) *ChangeUserStatusHandler {
	return &ChangeUserStatusHandler{
		userRepository: userRepository,
	}
}

func (h *ChangeUserStatusHandler) Handle(command ChangeUserStatusCommand) error {
	if command.UserId == 0 {
		return errors.New("user ID is required")
	}

	if command.Status == "" {
		return errors.New("status is required")
	}

	user, err := h.userRepository.GetByID(command.CTX, command.UserId)
	if err != nil {
		return err
	}

	if err := user.UpdateStatus(command.Status); err != nil {
		return err
	}

	err = h.userRepository.Save(command.CTX, &user)
	if err != nil {
		return err
	}

	return nil
}
