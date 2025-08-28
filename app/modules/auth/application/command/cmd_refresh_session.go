package authCmd

import (
	"context"
	"strconv"

	jwtService "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type RefreshSessionCommand struct {
	UserId       int             `json:"user_id"`
	RefreshToken string          `json:"session_id"`
	CTX          context.Context `json:"-"`
}

type refreshSessionHandler struct {
	userRepo    userDomain.UserRepository
	sessionRepo session.SessionRepository
	jwtService  jwtService.JWTService
}

func NewRefreshSessionHandler(
	userRepo userDomain.UserRepository,
	sessionRepo session.SessionRepository,
	jwtService jwtService.JWTService,
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

func getSessionResponse(entity session.Session, access string) SessionResponse {
	sessionResponse := &SessionResponse{
		UserID:       entity.UserId,
		AccessToken:  access,
		RefreshToken: entity.RefreshToken,
		ExpiresAt:    entity.ExpiresAt,
		CreatedAt:    entity.CreatedAt,
	}

	return *sessionResponse
}
