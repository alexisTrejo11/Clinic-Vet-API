package services

import (
	"errors"
	"fmt"

	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/services/domainServices"
	"example.com/at/backend/api-vet/sqlc"
	"example.com/at/backend/api-vet/utils"
)

type AuthService interface {
	ValidateUniqueFields(userSignUpDTO DTOs.UserSignUpDTO) error
	CompleteSignUp(userSignUpDTO DTOs.UserSignUpDTO) (string, error)
	CompleteLogin(userDTO DTOs.UserDTO) (string, error)
	FindUser(userLoginDTO DTOs.UserLoginDTO) (*DTOs.UserDTO, error)
	CheckPassword(hashPassword, givenPassword string) error
}

type AuthServiceImpl struct {
	authDomainService domainServices.AuthDomainService
	userRepository    repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository, ownerRepository repository.OwnerRepository) AuthService {
	// Initializing the domain service internally
	authDomainService := domainServices.NewAuthDomainService(userRepository, ownerRepository)
	return &AuthServiceImpl{
		authDomainService: authDomainService,
		userRepository:    userRepository,
	}
}

func (as *AuthServiceImpl) ValidateUniqueFields(userSignUpDTO DTOs.UserSignUpDTO) error {
	if userSignUpDTO.Email == "" && userSignUpDTO.PhoneNumber == "" {
		return errors.New("no credentials provided to create user")
	}

	if userSignUpDTO.Email != "" {
		isEmailTaken := as.userRepository.CheckEmailExists(userSignUpDTO.Email)
		if isEmailTaken {
			return errors.New("email is already taken")
		}
	}

	if userSignUpDTO.PhoneNumber != "" {
		isPhoneNumberTaken := as.userRepository.CheckPhoneNumberExists(userSignUpDTO.PhoneNumber)
		if isPhoneNumberTaken {
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

	if err := as.authDomainService.CreateOwner(*newUser, userSignUpDTO); err != nil {
		return "", errors.New("can't create owner")
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

func (as *AuthServiceImpl) FindUser(userLoginDTO DTOs.UserLoginDTO) (*DTOs.UserDTO, error) {
	var user *sqlc.User
	var err error

	if userLoginDTO.Email != "" {
		// Find user by email
		user, err = as.userRepository.GetUserByEmail(userLoginDTO.Email)
		if err != nil {
			return nil, fmt.Errorf("error finding user by email: %w", err)
		}
	} else if userLoginDTO.PhoneNumber != "" {
		// Find user by phone number
		user, err = as.userRepository.GetUserByPhoneNumber(userLoginDTO.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("error finding user by phone number: %w", err)
		}
	} else {
		return nil, fmt.Errorf("email or phone number must be provided")
	}

	// Map the found user to DTO
	userDTO := mappers.MapUserSqlcToDTO(user)
	return &userDTO, nil
}
