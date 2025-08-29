package command

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
)

type RefreshSessionCommand struct {
	UserId       int             `json:"user_id"`
	RefreshToken string          `json:"session_id"`
	CTX          context.Context `json:"-"`
}

type refreshSessionHandler struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	jwtService  jwt.JWTService
}

func NewRefreshSessionHandler(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtService jwt.JWTService,
) AuthCommandHandler {
	return &refreshSessionHandler{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
	}
}

func (h *refreshSessionHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(RefreshSessionCommand)

	_, err := h.userRepo.GetByID(command.CTX, command.UserId)
	if err != nil {
	}

	entity, err := h.sessionRepo.GetByUserAndId(command.CTX, strconv.Itoa(command.UserId), command.RefreshToken)
	if err != nil {
		return FailureAuthResult("Session not found", err)
	}

	access, err := h.jwtService.GenerateAccessToken(strconv.Itoa(command.UserId))
	if err != nil {
		return FailureAuthResult("Failed to generate access token", err)
	}

	response := getSessionResponse(entity, access)
	return SuccessAuthResult(&response, entity.Id, "session successfully refreshed")
}

func getSessionResponse(entity entity.Session, access string) SessionResponse {
	sessionResponse := &SessionResponse{
		UserID:       entity.UserId,
		AccessToken:  access,
		RefreshToken: entity.RefreshToken,
		ExpiresAt:    entity.ExpiresAt,
		CreatedAt:    entity.CreatedAt,
	}

	return *sessionResponse
}
