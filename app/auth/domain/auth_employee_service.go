package services

import (
	"errors"

	DTOs "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/utils"
)

type AuthEmployeeService interface {
	CompleteSignUp(userSignUpDTO DTOs.UserEmployeeSignUpDTO, vetDTO DTOs.VetDTO) (string, error)
}

type authEmployeeServiceImpl struct {
	userRepository    repository.UserRepository
	vetRepository     repository.VeterinarianRepository
	userMappers       mappers.UserMappers
	authCommonService AuthCommonService
}

func NewAuthEmployeeService(userRepository repository.UserRepository, vetRepository repository.VeterinarianRepository, authCommonService AuthCommonService) AuthEmployeeService {
	return &authEmployeeServiceImpl{
		userRepository:    userRepository,
		vetRepository:     vetRepository,
		authCommonService: authCommonService,
	}
}

func (as *authEmployeeServiceImpl) CompleteSignUp(userSignUpDTO DTOs.UserEmployeeSignUpDTO, vetDTO DTOs.VetDTO) (string, error) {
	passwordHashed, err := utils.HashPassword(userSignUpDTO.Password)
	if err != nil {
		return "", err
	}

	params := as.userMappers.MapSignUpDTOToCreateParams(vetDTO.Name, "", userSignUpDTO.Email, userSignUpDTO.PhoneNumber, "Veterinarian")
	params.Password = passwordHashed

	newUser, err := as.userRepository.CreateUser(params)
	if err != nil {
		return "", err
	}

	as.vetRepository.AddUserIdToExisitngVet(vetDTO.Id, newUser.ID)

	JWT, err := as.authCommonService.CreateJWT(newUser.ID, "Veterinarian")
	if err != nil {
		return "", errors.New("can't create jwt token")
	}

	return JWT, nil
}
