package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type ChangeUserStatusCommand struct {
	UserID valueobject.UserID `json:"user_id"`
	Status enum.UserStatus    `json:"status"`
	CTX    context.Context    `json:"ctx"`
}

type ChangeUserStatusHandler struct {
	userRepository repository.UserRepository
}

func NewChangeUserStatusHandler(userRepository repository.UserRepository) *ChangeUserStatusHandler {
	return &ChangeUserStatusHandler{
		userRepository: userRepository,
	}
}

func (h *ChangeUserStatusHandler) Handle(cmd cqrs.Command) error {
	command, ok := cmd.(ChangeUserStatusCommand)
	if !ok {
		return errors.New("invalid command type")
	}

	if command.UserID.IsZero() {
		return errors.New("user ID is required")
	}

	if command.Status == "" {
		return errors.New("status is required")
	}

	user, err := h.userRepository.GetByID(command.CTX, command.UserID)
	if err != nil {
		return err
	}

	// TODO: Implement

	err = h.userRepository.Save(command.CTX, &user)
	if err != nil {
		return err
	}

	return nil
}
