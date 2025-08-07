package userCommand

import (
	"context"
	"errors"

	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type ChangeUserStatusCommand struct {
	UserId int             `json:"user_id"`
	Status user.UserStatus `json:"status"`
	CTX    context.Context `json:"ctx"`
}

type ChangeUserStatusHandler struct {
	userRepository userRepository.UserRepository
}

func NewChangeUserStatusHandler(userRepository userRepository.UserRepository) *ChangeUserStatusHandler {
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

	user, err := h.userRepository.GetByIdWithProfile(command.CTX, command.UserId)
	if err != nil {
		return err
	}

	if err := user.UpdateStatus(command.Status); err != nil {
		return err
	}

	err = h.userRepository.Save(command.CTX, user)
	if err != nil {
		return err
	}

	return nil
}
