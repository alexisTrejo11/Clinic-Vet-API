package authCmd

import (
	"context"
	"strconv"

	jwtService "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"

	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type RefreshSessionCommand struct {
	UserId       int             `json:"user_id"`
	RefreshToken string          `json:"session_id"`
	CTX          context.Context `json:"-"`
}

type refreshSessionHandler struct {
	userRepo    userRepo.UserRepository
	sessionRepo session.SessionRepository
	jwtService  jwtService.JWTService
}

func NewRefreshSessionHandler(
	userRepo userRepo.UserRepository,
	sessionRepo session.SessionRepository,
	jwtService jwtService.JWTService,
) *refreshSessionHandler {
	return &refreshSessionHandler{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
	}
}

func (h *refreshSessionHandler) Handle(command RefreshSessionCommand) AuthCommandResult {
	_, err := h.userRepo.GetById(command.CTX, command.UserId)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("User not found", err)}
	}

	session, err := h.sessionRepo.GetByUserAndId(command.CTX, strconv.Itoa(command.UserId), command.RefreshToken)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("Session not found", err)}
	}

	access, err := h.jwtService.GenerateAccessToken(strconv.Itoa(command.UserId))
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("Failed to generate access token", err)}
	}

	sessionResponse := SessionResponse{
		UserId:       strconv.Itoa(command.UserId),
		AccessToken:  access,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}
	return AuthCommandResult{
		Session:       sessionResponse,
		CommandResult: shared.SuccessResult("", "Session refreshed successfully"),
	}
}
