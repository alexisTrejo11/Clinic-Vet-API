package authCmd

import (
	"context"
	"errors"
	"time"

	jwtService "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
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
	userRepo    userDomain.UserRepository
	sessionRepo session.SessionRepository
	jwtService  jwtService.JWTService
}

func NewLoginHandler(
	userRepo userDomain.UserRepository,
	sessionRepo session.SessionRepository,
	jwtService jwtService.JWTService,
) AuthCommandHandler {
	return &loginHandler{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
	}
}

func (h *loginHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(LoginCommand)

	if command.Identifier == "" || command.Password == "" {
		return FailureAuthResult("Identifier and password are required", errors.New("missing identifier or password"))
	}

	user, err := h.Authenticate(&command)
	if err != nil {
		return FailureAuthResult("authentication failed", err)
	}

	if user.Is2FAEnabled() {
		return FailureAuthResult("2FA is enabled for this user, please complete the 2FA process", errors.New("2FA is enabled"))
	}

	session, err := h.createSession(user.Id().String(), command)
	if err != nil {
		return FailureAuthResult("failed to create session", err)
	}

	accesToken, err := h.jwtService.GenerateAccessToken(user.Id().String())
	if err != nil {
		return FailureAuthResult("failed to generate access token", err)
	}

	response := getSessionResponse(session, accesToken)
	return SuccessAuthResult(&response, session.Id, "login successfully processed")
}

func (h *loginHandler) Authenticate(command *LoginCommand) (userDomain.User, error) {
	user, err := h.userRepo.GetByEmail(command.CTX, command.Identifier)
	if err == nil {
		return user, nil
	}

	user, err = h.userRepo.GetByPhone(command.CTX, command.Identifier)
	if err == nil {
		return user, nil
	}

	if err := shared.CheckPassword(command.Password, user.Password()); err == nil {
		return user, nil
	}

	return userDomain.User{}, errors.New("user not found with provided credentials, please check your email/phone-number and password")
}

func (h *loginHandler) createSession(userId string, command LoginCommand) (session.Session, error) {
	refresh, err := h.jwtService.GenerateRefreshToken(userId)
	if err != nil {
		return session.Session{}, err
	}

	newSession := session.Session{
		UserId:       userId,
		IpAddress:    command.IP,
		RefreshToken: refresh,
		CreatedAt:    time.Now(),
		DeviceInfo:   command.DeviceInfo,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		UserAgent:    command.UserAgent,
	}

	if err := h.sessionRepo.Create(command.CTX, &newSession); err != nil {
		return session.Session{}, err
	}

	return newSession, nil
}
