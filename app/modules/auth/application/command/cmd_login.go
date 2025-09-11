package command

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/jwt"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
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

	if user.CanLogin() {
		return FailureAuthResult(
			ErrAuthenticationFailed,
			apperror.ConflictError("UserStatus", "user is not active, please contact support"),
		)
	}

	if user.IsTwoFactorEnabled() {
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

	return SuccessAuthResult(
		getSessionResponse(session, accessToken),
		session.ID,
		MsgLoginSuccess,
	)
}

func (h *loginHandler) authenticate(command *LoginCommand) (user.User, error) {
	userEntity, err := h.userRepo.FindByEmail(command.CTX, command.Identifier)
	if err != nil {
		userEntity, err = h.userRepo.FindByPhone(command.CTX, command.Identifier)
		if err != nil {
			return user.User{}, errors.New(ErrInvalidCredentials)
		}
	}

	if err := h.passwordEncoder.CheckPassword(command.Password, userEntity.Password()); err != nil {
		return user.User{}, errors.New(ErrInvalidCredentials)
	}

	return userEntity, nil
}

// createSession creates and persists a new user session
func (h *loginHandler) createSession(userID string, command LoginCommand) (auth.Session, error) {
	refreshToken, err := h.jwtService.GenerateRefreshToken(userID)
	if err != nil {
		return auth.Session{}, err
	}

	now := time.Now()
	sessionDuration := DefaultSessionDuration

	if command.RememberMe {
		sessionDuration = 30 * 24 * time.Hour // 30 days
	}

	newSession := auth.Session{
		UserID:       userID,
		IPAddress:    command.IP,
		RefreshToken: refreshToken,
		CreatedAt:    now,
		ExpiresAt:    now.Add(sessionDuration),
		DeviceInfo:   command.DeviceInfo,
		UserAgent:    command.UserAgent,
	}

	if err := h.sessionRepo.Create(command.CTX, &newSession); err != nil {
		return auth.Session{}, err
	}

	return newSession, nil
}
