package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
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

	entity, err := h.sessionRepo.GetByUserAndID(command.CTX, command.UserID.String(), command.RefreshToken)
	if err != nil {
		return FailureAuthResult("Session not found", err)
	}

	access, err := h.jwtService.GenerateAccessToken(command.UserID.String())
	if err != nil {
		return FailureAuthResult("Failed to generate access token", err)
	}

	response := getSessionResponse(entity, access)
	return SuccessAuthResult(&response, entity.ID, "session successfully refreshed")
}

func (h *RefreshSessionHandler) validateExisitngUser(command RefreshSessionCommand) error {
	exists, err := h.userRepo.ExistsByID(command.CTX, command.UserID.GetValue())
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("user", "id", command.UserID.String())
	}

	return nil
}

func getSessionResponse(entity entity.Session, access string) SessionResponse {
	sessionResponse := &SessionResponse{
		UserID:       entity.UserID,
		AccessToken:  access,
		RefreshToken: entity.RefreshToken,
		ExpiresAt:    entity.ExpiresAt,
		CreatedAt:    entity.CreatedAt,
	}

	return *sessionResponse
}
