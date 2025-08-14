package authCmd

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	sessionRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain/repositories"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type RefreshSessionCommand struct {
	UserId       int             `json:"user_id"`
	RefreshToken string          `json:"session_id"`
	CTX          context.Context `json:"-"`
}

type RefreshSessionHandler interface {
	Handle(command RefreshSessionCommand) (SessionResponse, error)
}

type refresSessionHandler struct {
	userRepo    userRepo.UserRepository
	sessionRepo sessionRepo.SessionRepository
	jwtService  jwt.JWTService
}

func NewRefreshSessionHandler(
	userRepo userRepo.UserRepository,
	sessionRepo sessionRepo.SessionRepository,
	jwtService jwt.JWTService,
) RefreshSessionHandler {
	return &refresSessionHandler{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (h *refresSessionHandler) Handle(command RefreshSessionCommand) (SessionResponse, error) {
	_, err := h.userRepo.GetById(command.CTX, command.UserId)
	if err != nil {
		return SessionResponse{}, err
	}

	session, err := h.sessionRepo.FindByUserAndRefreshToken(command.CTX, strconv.Itoa(command.UserId), command.RefreshToken)
	if err != nil {
		return SessionResponse{}, err
	}

	access, err := h.jwtService.GenerateAccessToken(strconv.Itoa(command.UserId))
	if err != nil {
		return SessionResponse{}, err
	}

	return SessionResponse{
		UserId:       strconv.Itoa(command.UserId),
		AccessToken:  access,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}, nil
}
