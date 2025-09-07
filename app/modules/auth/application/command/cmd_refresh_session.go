package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type RefreshSessionCommand struct {
	UserID       valueobject.UserID `json:"user_id"`
	RefreshToken string             `json:"session_id"`
	CTX          context.Context    `json:"-"`
}

type RefreshSessionHandler struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	jwtService  jwt.JWTService
}

func NewRefreshSessionHandler(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtService jwt.JWTService,
) AuthCommandHandler {
	return &RefreshSessionHandler{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
	}
}

func (h *RefreshSessionHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(RefreshSessionCommand)

	if err := h.validateExisitngUser(command); err != nil {
		return FailureAuthResult("Error ocurred validatin user", err)
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
	return SuccessAuthResult(&response, session.ID, "session successfully refreshed")
}

func (h *RefreshSessionHandler) validateExisitngUser(command RefreshSessionCommand) error {
	exists, err := h.userRepo.ExistsByID(command.CTX, command.UserID.Value())
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("user", "id", command.UserID.String())
	}

	return nil
}

func getSessionResponse(entity auth.Session, access string) SessionResponse {
	sessionResponse := &SessionResponse{
		UserID:       entity.UserID,
		AccessToken:  access,
		RefreshToken: entity.RefreshToken,
		ExpiresAt:    entity.ExpiresAt,
		CreatedAt:    entity.CreatedAt,
	}

	return *sessionResponse
}
