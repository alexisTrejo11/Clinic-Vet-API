package authCmd

import (
	"context"
	"errors"
	"time"

	jwtService "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
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

type loginHandler struct {
	userRepo    userRepository.UserRepository
	sessionRepo session.SessionRepository
	jwtService  jwtService.JWTService
}

func NewLoginHandler(
	userRepo userRepository.UserRepository,
	sessionRepo session.SessionRepository,
	jwtService jwtService.JWTService,
) *loginHandler {
	return &loginHandler{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
	}
}

func (h *loginHandler) Handle(cmd LoginCommand) AuthCommandResult {
	if cmd.Identifier == "" || cmd.Password == "" {
		return AuthCommandResult{CommandResult: shared.FailureResult("Identifier and password are required", errors.New("missing identifier or password"))}
	}

	user, err := h.Authenticate(&cmd)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("authentication failed", err)}

	}

	if user.Is2FAEnabled() {
		return AuthCommandResult{CommandResult: shared.FailureResult("2FA is enabled for this user, please complete the 2FA process", errors.New("2FA is enabled"))}
	}

	session, err := h.createSession(user.Id().String(), cmd)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("failed to create session", err)}
	}

	accesToken, err := h.jwtService.GenerateAccessToken(user.Id().String())
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("failed to generate access token", err)}
	}

	sessionResponse := SessionResponse{
		AccessToken:  accesToken,
		RefreshToken: session.RefreshToken,
		UserId:       session.UserId,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}

	return AuthCommandResult{
		Session:       sessionResponse,
		CommandResult: shared.SuccessResult(session.Id, "Login successful"),
	}
}

func (h *loginHandler) Authenticate(cmd *LoginCommand) (*user.User, error) {
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

func (h *loginHandler) createSession(userId string, cmd LoginCommand) (session.Session, error) {
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
