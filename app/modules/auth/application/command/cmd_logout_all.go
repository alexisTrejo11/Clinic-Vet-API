package command

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type LogoutAllCommand struct {
	userID valueobject.UserID
	ctx    context.Context
}

func NewLogoutAllCommand(idnInt int) (LogoutAllCommand, error) {
	userID, err := valueobject.NewUserID(idnInt)
	if err != nil {
		return LogoutAllCommand{}, apperror.FieldValidationError("userId", strconv.Itoa(idnInt), err.Error())
	}

	cmd := &LogoutAllCommand{
		ctx:    context.Background(),
		userID: userID,
	}

	return *cmd, nil
}

type logoutAllHandler struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func NewLogoutAllHandler(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
) *logoutAllHandler {
	return &logoutAllHandler{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (h *logoutAllHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(LogoutAllCommand)

	user, err := h.userRepository.GetByID(command.ctx, command.userID)
	if err != nil {
		return FailureAuthResult("an error ocurred finding user", err)
	}

	err = h.sessionRepository.DeleteAllUserSessions(command.ctx, user.ID())
	if err != nil {
		return FailureAuthResult("an error ocurred delete user sessions", err)
	}

	return SuccessAuthResult(nil, user.ID().String(), "user sessions sucessfully deleted on all devices")
}
