package authUsecase

import (
	"errors"

	authDto "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/dtos"
	authRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/repositories"
)

type LogoutUseCase struct {
	sessionRepo authRepository.SessionRepository
}

func NewLogoutUseCase(sessionRepo authRepository.SessionRepository) *LogoutUseCase {
	return &LogoutUseCase{
		sessionRepo: sessionRepo,
	}
}

func (uc *LogoutUseCase) Execute(dto authDto.RequestLogout) error {
	// Find session by refresh token
	session, err := uc.sessionRepo.FindByRefreshToken(dto.RefreshToken)
	if err != nil {
		return errors.New("invalid refresh token")
	}

	// Delete session
	if err := uc.sessionRepo.DeleteSession(session.ID); err != nil {
		return errors.New("failed to logout")
	}

	return nil
}
