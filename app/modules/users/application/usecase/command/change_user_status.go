package command

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type ChangeUserStatusCommand struct {
	userID valueobject.UserID
	status enum.UserStatus
	ctx    context.Context
}

func NewChangeUserStatusCommand(ctx context.Context, userID uint, status string) (ChangeUserStatusCommand, error) {
	uid := valueobject.NewUserID(userID)
	userStatus, err := enum.ParseUserStatus(status)
	if err != nil {
		return ChangeUserStatusCommand{}, err
	}

	cmd := &ChangeUserStatusCommand{
		userID: uid,
		status: userStatus,
		ctx:    ctx,
	}
	return *cmd, nil
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

	if command.userID.IsZero() {
		return errors.New("user ID is required")
	}

	if command.status == "" {
		return errors.New("status is required")
	}

	user, err := h.userRepository.FindByID(command.ctx, command.userID)
	if err != nil {
		return err
	}

	// TODO: Implement

	err = h.userRepository.Save(command.ctx, &user)
	if err != nil {
		return err
	}

	return nil
}
