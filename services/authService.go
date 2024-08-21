package services

import (
	"errors"

	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/services/domainServices"
	"example.com/at/backend/api-vet/utils"
)

type AuthService interface {
	ValidateUniqueFields(userSignUpDTO DTOs.UserSignUpDTO) error
	ProcessSignUp(userSignUpDTO DTOs.UserSignUpDTO) (string, error)
	ProcessLogin(userDTO DTOs.UserDTO) (string, error)
	FindUser(userLoginDTO DTOs.UserLoginDTO) (DTOs.UserDTO, error)
	CheckPassword(hashPassword, givenPassword string) error
}

type AuthServiceImpl struct {
	authDomainService domainServices.AuthDomainService
	userRepository    repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository, ownerRepository repository.OwnerRepository) AuthServiceImpl {
	// Initializing the domain service internally
	authDomainService := domainServices.NewAuthDomainService(userRepository, ownerRepository)
	return AuthServiceImpl{
		authDomainService: authDomainService,
		userRepository:    userRepository,
	}
}

func (as *AuthServiceImpl) ValidateUniqueFields(userSignUpDTO DTOs.UserSignUpDTO) error {
	if userSignUpDTO.Email == "" && userSignUpDTO.Phone == "" {
		return errors.New("no credentials provided to create user")
	}

	if userSignUpDTO.Email != "" {
		isEmailTaken := as.userRepository.CheckEmailExists(userSignUpDTO.Email)
		if isEmailTaken {
			return errors.New("email is already taken")
		}
	}

	if userSignUpDTO.Phone != "" {
		isPhoneTaken := as.userRepository.CheckPhoneNumberExists(userSignUpDTO.Phone)
		if isPhoneTaken {
			return errors.New("phone number is already taken")
		}
	}

	return nil
}

func (as *AuthServiceImpl) CompleteSignUp(userSignUpDTO DTOs.UserSignUpDTO) (string, error) {
	newUser, err := as.authDomainService.ProcessUserCreation(userSignUpDTO)
	if err != nil {
		return "", errors.New("can't create user")
	}

	JWT, err := as.authDomainService.CreateJWT(newUser.ID, newUser.Role)
	if err != nil {
		return "", errors.New("can't create jwt token")
	}

	return JWT, nil
}

func (as *AuthServiceImpl) CheckPassword(hashPassword, givenPassword string) error {
	if err := utils.CheckPassword(hashPassword, givenPassword); err != nil {
		return err
	}

	return nil
}

func (as *AuthServiceImpl) CompleteLogin(userDTO DTOs.UserDTO) (string, error) {
	if err := as.authDomainService.ProcessUserLogin(userDTO); err != nil {
		return "", err
	}

	JWT, err := as.authDomainService.CreateJWT(userDTO.Id, userDTO.Role)
	if err != nil {
		return "", errors.New("can't create jwt token")
	}

	return JWT, nil
}
