package services

import (
	"errors"

	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/services/domainServices"
)

type AuthClientService interface {
	CompleteSignUp(userSignUpDTO DTOs.UserSignUpDTO) (string, error)
}

type authClientServiceImpl struct {
	authDomainService domainServices.AuthDomainService
	userRepository    repository.UserRepository
}

func NewClientAuthService(userRepository repository.UserRepository, ownerRepository repository.OwnerRepository) AuthClientService {
	// Initializing the domain service internally
	authDomainService := domainServices.NewAuthDomainService(userRepository, ownerRepository)
	return &authClientServiceImpl{
		authDomainService: authDomainService,
		userRepository:    userRepository,
	}
}

func (as *authClientServiceImpl) CompleteSignUp(userSignUpDTO DTOs.UserSignUpDTO) (string, error) {
	newUser, err := as.authDomainService.ProcessClientUserCreation(userSignUpDTO)
	if err != nil {
		return "", errors.New("can't create user")
	}

	if err := as.authDomainService.CreateOwner(*newUser, userSignUpDTO); err != nil {
		return "", errors.New("can't create owner")
	}

	JWT, err := as.authDomainService.CreateJWT(newUser.ID, newUser.Role)
	if err != nil {
		return "", errors.New("can't create jwt token")
	}

	return JWT, nil
}
