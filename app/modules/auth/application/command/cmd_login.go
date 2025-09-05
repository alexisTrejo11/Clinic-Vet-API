package command

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
)

const (
	ErrAuthenticationFailed  = "authentication failed"
	ErrTwoFactorRequired     = "2FA is enabled for this user, please complete the 2FA process"
	ErrSessionCreationFailed = "failed to create session"
	ErrAccessTokenGenFailed  = "failed to generate access token"
	ErrInvalidCredentials    = "user not found with provided credentials, please check your email/phone-number and password"
	ErrTwoFactorAuthConflict = "user has TwoFactorAuth auth login method"

	MsgLoginSuccess = "login successfully processed"

	DefaultSessionDuration = 7 * 24 * time.Hour
)

// LoginCommand represents the login request data
type LoginCommand struct {
	Identifier string          `json:"identifier" validate:"required"`
	Password   string          `json:"password" validate:"required"`
	RememberMe bool            `json:"remember_me"`
	IP         string          `json:"ip" validate:"required"`
	UserAgent  string          `json:"user_agent"`
	DeviceInfo string          `json:"source"`
	CTX        context.Context `json:"-"`
}

// loginHandler handles user authentication and session creation
type loginHandler struct {
	userRepo        repository.UserRepository
	sessionRepo     repository.SessionRepository
	jwtService      jwt.JWTService
	passwordEncoder password.PasswordEncoder
}

// NewLoginHandler creates a new instance of loginHandler
func NewLoginHandler(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	jwtService jwt.JWTService,
	passwordEncoder password.PasswordEncoder,
) AuthCommandHandler {
	return &loginHandler{
		userRepo:        userRepo,
		sessionRepo:     sessionRepo,
		jwtService:      jwtService,
		passwordEncoder: passwordEncoder,
	}
}

// Handle processes the login command and returns authentication result
func (h *loginHandler) Handle(cmd any) AuthCommandResult {
	command, ok := cmd.(LoginCommand)
	if !ok {
		return FailureAuthResult(ErrAuthenticationFailed, errors.New("invalid command type"))
	}

	user, err := h.authenticate(&command)
	if err != nil {
		return FailureAuthResult(ErrAuthenticationFailed, err)
	}

	if user.Is2FAEnabled() {
		return FailureAuthResult(
			ErrTwoFactorRequired,
			apperror.ConflictError("TwoFactorAuth", ErrTwoFactorAuthConflict),
		)
	}

	session, err := h.createSession(user.ID().String(), command)
	if err != nil {
		return FailureAuthResult(ErrSessionCreationFailed, err)
	}

	accessToken, err := h.jwtService.GenerateAccessToken(user.ID().String())
	if err != nil {
		return FailureAuthResult(ErrAccessTokenGenFailed, err)
	}

	response := getSessionResponse(session, accessToken)
	return SuccessAuthResult(&response, session.ID, MsgLoginSuccess)
}

func (h *loginHandler) authenticate(command *LoginCommand) (entity.User, error) {
	user, err := h.userRepo.GetByEmail(command.CTX, command.Identifier)
	if err != nil {
		user, err = h.userRepo.GetByPhone(command.CTX, command.Identifier)
		if err != nil {
			return entity.User{}, errors.New(ErrInvalidCredentials)
		}
	}

	if err := h.passwordEncoder.CheckPassword(command.Password, user.Password()); err != nil {
		return entity.User{}, errors.New(ErrInvalidCredentials)
	}

	return user, nil
}

// createSession creates and persists a new user session
func (h *loginHandler) createSession(userID string, command LoginCommand) (entity.Session, error) {
	refreshToken, err := h.jwtService.GenerateRefreshToken(userID)
	if err != nil {
		return entity.Session{}, err
	}

	now := time.Now()
	sessionDuration := DefaultSessionDuration

	if command.RememberMe {
		sessionDuration = 30 * 24 * time.Hour // 30 days
	}

	newSession := entity.Session{
		UserID:       userID,
		IPAddress:    command.IP,
		RefreshToken: refreshToken,
		CreatedAt:    now,
		ExpiresAt:    now.Add(sessionDuration),
		DeviceInfo:   command.DeviceInfo,
		UserAgent:    command.UserAgent,
	}

	if err := h.sessionRepo.Create(command.CTX, &newSession); err != nil {
		return entity.Session{}, err
	}

	return newSession, nil
}
