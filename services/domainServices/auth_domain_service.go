package domainServices

import (
	"strconv"

	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/sqlc"
	"example.com/at/backend/api-vet/utils"
)

type AuthDomainService interface {
	ProcessClientUserCreation(userSignUpDTO DTOs.UserSignUpDTO) (*sqlc.CreateUserRow, error)
	ProcessEmployeeUserCreation(userSignUpDTO DTOs.UserEmployeeSignUpDTO, vetDTO DTOs.VetDTO) (*sqlc.CreateUserRow, error)
	CreateJWT(userId int32, role string) (string, error)
	ProcessUserLogin(UserDTO DTOs.UserDTO) error
	CreateOwner(userData sqlc.CreateUserRow, userSignUpDTO DTOs.UserSignUpDTO) error
}

type AuthDomainServiceImpl struct {
	userMappers     mappers.UserMappers
	ownerMappers    mappers.OwnerMapper
	userRepository  repository.UserRepository
	ownerRepository repository.OwnerRepository
	vetRepository   repository.VeterinarianRepository
}

func NewAuthDomainService(userRepository repository.UserRepository, ownerRepository repository.OwnerRepository, vetRepository repository.VeterinarianRepository) AuthDomainService {
	return &AuthDomainServiceImpl{
		userRepository:  userRepository,
		vetRepository:   vetRepository,
		ownerRepository: ownerRepository,
	}
}

func (ads AuthDomainServiceImpl) ProcessClientUserCreation(userSignUpDTO DTOs.UserSignUpDTO) (*sqlc.CreateUserRow, error) {
	passwordHashed, err := utils.HashPassword(userSignUpDTO.Password)
	if err != nil {
		return nil, err
	}
	params := ads.userMappers.MapSignUpDTOToCreateParams(userSignUpDTO.Name, userSignUpDTO.LastName, userSignUpDTO.Email, userSignUpDTO.PhoneNumber, "Common-Owner")
	params.Password = passwordHashed

	newUser, err := ads.userRepository.CreateUser(params)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (ads AuthDomainServiceImpl) ProcessUserLogin(UserDTO DTOs.UserDTO) error {
	if err := ads.userRepository.UpdateUserLastLogin(UserDTO.Id); err != nil {
		return err
	}

	return nil
}

func (ads AuthDomainServiceImpl) CreateJWT(userId int32, role string) (string, error) {
	var roles []string
	roles = append(roles, role)
	userIdStr := Int32ToString(userId)

	JWT, err := utils.GenerateJWT(userIdStr, roles)
	if err != nil {
		return "", err
	}

	return JWT, nil
}

func (ads AuthDomainServiceImpl) CreateOwner(userData sqlc.CreateUserRow, userSignUpDTO DTOs.UserSignUpDTO) error {
	ownerCreateParams, err := ads.ownerMappers.MapSignUpDataToCreateParams(userData.ID, userSignUpDTO)
	if err != nil {
		return err
	}

	if _, err := ads.ownerRepository.CreateOwner(*ownerCreateParams); err != nil {
		return err
	}

	return nil
}

func Int32ToString(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

func (ads AuthDomainServiceImpl) ProcessEmployeeUserCreation(userSignUpDTO DTOs.UserEmployeeSignUpDTO, vetDTO DTOs.VetDTO) (*sqlc.CreateUserRow, error) {
	passwordHashed, err := utils.HashPassword(userSignUpDTO.Password)
	if err != nil {
		return nil, err
	}
	params := ads.userMappers.MapSignUpDTOToCreateParams(vetDTO.Name, "", userSignUpDTO.Email, userSignUpDTO.PhoneNumber, "Veterinarian")
	params.Password = passwordHashed

	newUser, err := ads.userRepository.CreateUser(params)
	if err != nil {
		return nil, err
	}

	ads.vetRepository.AddUserIdToExisitngVet(vetDTO.Id, newUser.ID)

	return newUser, nil
}
