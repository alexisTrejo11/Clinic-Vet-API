package authUsecase

import (
	"errors"

	authRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/repositories"
)

type LogoutAllUseCase struct {
	sessionRepo authRepository.SessionRepository
}

func NewLogoutAllUseCase(sessionRepo authRepository.SessionRepository) *LogoutAllUseCase {
	return &LogoutAllUseCase{
		sessionRepo: sessionRepo,
	}
}

func (uc *LogoutAllUseCase) Execute(userID uint) error {
	if err := uc.sessionRepo.DeleteAllUserSessions(userID); err != nil {
		return errors.New("failed to logout from all devices")
	}

	return nil
}
