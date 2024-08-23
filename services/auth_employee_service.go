package services

import (
	"errors"

	DTOs "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/services/domainServices"
)

type AuthEmployeeService interface {
	CompleteSignUp(userSignUpDTO DTOs.UserEmployeeSignUpDTO, vetDTO DTOs.VetDTO) (string, error)
}

type authEmployeeServiceImpl struct {
	authDomainService domainServices.AuthDomainService
	userRepository    repository.UserRepository
	vetRepository     repository.VeterinarianRepository
}

func NewAuthEmployeeService(userRepository repository.UserRepository, ownerRepository repository.OwnerRepository, vetRepository repository.VeterinarianRepository) AuthEmployeeService {
	// Initializing the domain service internally
	authDomainService := domainServices.NewAuthDomainService(userRepository, ownerRepository, vetRepository)
	return &authEmployeeServiceImpl{
		userRepository:    userRepository,
		authDomainService: authDomainService,
		vetRepository:     vetRepository,
	}
}

func (as *authEmployeeServiceImpl) CompleteSignUp(userSignUpDTO DTOs.UserEmployeeSignUpDTO, vetDTO DTOs.VetDTO) (string, error) {
	newUser, err := as.authDomainService.ProcessEmployeeUserCreation(userSignUpDTO, vetDTO)
	if err != nil {
		return "", errors.New("can't create user")
	}

	JWT, err := as.authDomainService.CreateJWT(newUser.ID, newUser.Role)
	if err != nil {
		return "", errors.New("can't create jwt token")
	}

	return JWT, nil
}
