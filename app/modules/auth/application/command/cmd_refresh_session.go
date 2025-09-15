package command

import (
	"context"

	"clinic-vet-api/app/core/domain/entity/auth"
	"clinic-vet-api/app/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type RefreshSessionCommand struct {
	UserID       valueobject.UserID `json:"user_id"`
	RefreshToken string             `json:"session_id"`
	CTX          context.Context    `json:"-"`
}

func (h *authCommandHandler) RefreshSession(command RefreshSessionCommand) AuthCommandResult {
	if err := h.validateExisitngUser(command); err != nil {
		return FailureAuthResult("Error occurred in user validation", err)
	}

	session, err := h.sessionRepo.GetByUserAndID(command.CTX, command.UserID, command.RefreshToken)
	if err != nil {
		return FailureAuthResult("Session not found", err)
	}

	access, err := h.jwtService.GenerateAccessToken(command.UserID.String())
	if err != nil {
		return FailureAuthResult("Failed to generate access token", err)
	}

	response := getSessionResponse(session, access)
	return SuccessAuthResult(response, session.ID, "session successfully refreshed")
}

func (h *authCommandHandler) validateExisitngUser(command RefreshSessionCommand) error {
	exists, err := h.userRepository.ExistsByID(command.CTX, command.UserID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("user", "id", command.UserID.String())
	}

	return nil
}

func getSessionResponse(entity auth.Session, access string) *SessionResponse {
	return &SessionResponse{
		UserID:       entity.UserID,
		AccessToken:  access,
		RefreshToken: entity.RefreshToken,
		ExpiresAt:    entity.ExpiresAt,
		CreatedAt:    entity.CreatedAt,
	}
}
