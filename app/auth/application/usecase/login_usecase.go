package authUsecase

import (
	"errors"
	"strconv"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application"
	authDto "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/dtos"
	authRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/repositories"
	authDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/repositories"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Execute(dto authDto.RequestLogin, deviceInfo, ipAddress string) (*authDto.TokenResponse, error)
}

type loginUseCase struct {
	userRepo    userRepository.UserRepository
	sessionRepo authRepository.SessionRepository
	jwtService  application.JWTService
}

func NewLoginUseCase(userRepo userRepository.UserRepository, sessionRepo authRepository.SessionRepository, jwtService application.JWTService) LoginUseCase {
	return &loginUseCase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		jwtService:  jwtService,
	}
}

func (uc *loginUseCase) Execute(dto authDto.RequestLogin, deviceInfo, ipAddress string) (*authDto.TokenResponse, error) {
	user, err := uc.userRepo.FindByEmail(dto.IdentifierField)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("account is not activated")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	userIdStr := strconv.Itoa(int(user.ID))
	accessToken, err := uc.jwtService.GenerateAccessToken(userIdStr)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := uc.jwtService.GenerateRefreshToken(userIdStr)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	session := &authDomain.Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		DeviceInfo:   deviceInfo,
		IPAddress:    ipAddress,
		IsActive:     true,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour), // 7 days
	}

	if err := uc.sessionRepo.Create(session); err != nil {
		return nil, errors.New("failed to create session")
	}

	return &authDto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    3600, // 1 hour
		TokenType:    "Bearer",
	}, nil
}
