package authCmd

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	sessionRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type LoginCommand struct {
	Identifier string          `json:"identifier"`
	Password   string          `json:"password"`
	RememberMe bool            `json:"remember_me"`
	IP         string          `json:"ip"`
	UserAgent  string          `json:"user_agent"`
	DeviceInfo string          `json:"source"`
	CTX        context.Context `json:"-"`
}

type LoginHandler interface {
	Handle(cmd *LoginCommand) (SessionResponse, error)
}

type loginHander struct {
	userRepo    userRepository.UserRepository
	sessionRepo sessionRepo.SessionRepository
	jwtService  jwt.JWTService
}

type SessionResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserId       string    `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewLoginHandler(
	userRepo userRepository.UserRepository,
	sessionRepo sessionRepo.SessionRepository,
	jwtService jwt.JWTService,
) LoginHandler {
	return &loginHander{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
	}
}

func (h *loginHander) Handle(cmd *LoginCommand) (SessionResponse, error) {
	// Validate the command
	if cmd.Identifier == "" || cmd.Password == "" {
		return SessionResponse{}, errors.New("email/phoneNumber and password are required")
	}

	// Authenticate the user
	user, err := h.Authenticate(cmd)
	if err != nil {
		return SessionResponse{}, err
	}

	if user.Is2FAEnabled() {
		return SessionResponse{}, errors.New("2FA is enabled for this user, please complete the 2FA process")
	}

	session, err := h.createSession(user.Id().String(), *cmd)
	if err != nil {
		return SessionResponse{}, err
	}

	accesToken, err := h.jwtService.GenerateAccessToken(user.Id().String())
	if err != nil {
		return SessionResponse{}, err
	}

	return SessionResponse{
		AccessToken:  accesToken,
		RefreshToken: session.RefreshToken,
		UserId:       session.UserId,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}, nil
}

func (h *loginHander) Authenticate(cmd *LoginCommand) (*user.User, error) {
	user, err := h.userRepo.GetByEmail(cmd.CTX, cmd.Identifier)
	if err == nil {
		return user, nil
	}

	user, err = h.userRepo.GetByPhone(cmd.CTX, cmd.Identifier)
	if err == nil {
		return user, nil
	}

	if err := shared.CheckPassword(cmd.Password, user.Password()); err == nil {
		return user, nil
	}

	return nil, errors.New("user not found with provided credentials, please check your email/phone-number and password")
}

func (h *loginHander) createSession(userId string, cmd LoginCommand) (session.Session, error) {
	refresh, err := h.jwtService.GenerateRefreshToken(userId)
	if err != nil {
		return session.Session{}, err
	}

	newSession := session.Session{
		UserId:       userId,
		IpAddress:    cmd.IP,
		RefreshToken: refresh,
		CreatedAt:    time.Now(),
		DeviceInfo:   cmd.DeviceInfo,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		UserAgent:    cmd.UserAgent,
	}

	if err := h.sessionRepo.Create(cmd.CTX, &newSession); err != nil {
		return session.Session{}, err
	}

	return newSession, nil
}
