package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type ChangeUserStatusCommand struct {
	userID valueobject.UserID `json:"user_id"`
	status enum.UserStatus    `json:"status"`
	ctx    context.Context    `json:"ctx"`
}

func NewChangeUserStatusCommand(ctx context.Context, userIDInt int, status string) (*ChangeUserStatusCommand, error) {
	errorMessage := make([]string, 0) 

	userID, err := valueobject.NewUserID(userIDInt)



	cmd := &ChangeUserStatusCommand{
		us
	}
}

type ChangeUserStatusHandler struct {
	userRepository repository.UserRepository
}

func NewChangeUserStatusHandler(userRepository repository.UserRepository) *ChangeUserStatusHandler {
	return &ChangeUserStatusHandler{
		userRepository: userRepository,
	}
}

func (h *ChangeUserStatusHandler) Handle(command ChangeUserStatusCommand) error {
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

	if err := user.UpdateStatus(command.Status); err != nil {
		return err
	}

	err = h.userRepository.Save(command.CTX, &user)
	if err != nil {
		return err
	}

	return nil
}
