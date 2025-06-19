package services

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/sqlc"
	"example.com/at/backend/api-vet/utils"
)

type AuthClientService interface {
	CompleteSignUp(userSignUpDTO DTOs.UserSignUpDTO) (string, error)
}

type authClientServiceImpl struct {
	authCommonService AuthCommonService
	userRepository    repository.UserRepository
	vetRepository     repository.VeterinarianRepository
	ownerMappers      mappers.OwnerMapper
	userMappers       mappers.UserMappers
	ownerRepository   repository.OwnerRepository
}

func NewClientAuthService(authCommonService AuthCommonService, userRepository repository.UserRepository, ownerRepository repository.OwnerRepository, vetRepository repository.VeterinarianRepository) AuthClientService {
	return &authClientServiceImpl{
		vetRepository:     vetRepository,
		userRepository:    userRepository,
		ownerRepository:   ownerRepository,
		authCommonService: authCommonService,
	}
}

func (as *authClientServiceImpl) CompleteSignUp(userSignUpDTO DTOs.UserSignUpDTO) (string, error) {
	newUser, err := as.ProcessClientUserCreation(userSignUpDTO)
	if err != nil {
		return "", err
	}

	if err := as.CreateOwner(*newUser, userSignUpDTO); err != nil {
		return "", err
	}

	JWT, err := as.authCommonService.CreateJWT(newUser.ID, newUser.Role)
	if err != nil {
		return "", err
	}

	return JWT, nil
}

func (as *authClientServiceImpl) CreateOwner(userData sqlc.CreateUserRow, userSignUpDTO DTOs.UserSignUpDTO) error {
	ownerCreateParams, err := as.ownerMappers.MapSignUpDataToCreateParams(userData.ID, userSignUpDTO)
	if err != nil {
		return err
	}

	if _, err := as.ownerRepository.CreateOwner(*ownerCreateParams); err != nil {
		return err
	}

	return nil
}

func (as *authClientServiceImpl) ProcessClientUserCreation(userSignUpDTO DTOs.UserSignUpDTO) (*sqlc.CreateUserRow, error) {
	passwordHashed, err := utils.HashPassword(userSignUpDTO.Password)
	if err != nil {
		return nil, err
	}
	params := as.userMappers.MapSignUpDTOToCreateParams(userSignUpDTO.Name, userSignUpDTO.LastName, userSignUpDTO.Email, userSignUpDTO.PhoneNumber, "Common-Owner")
	params.Password = passwordHashed

	newUser, err := as.userRepository.CreateUser(params)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
