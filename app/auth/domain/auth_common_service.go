package services

/*
import (
	"errors"
	"fmt"
	"strconv"

	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/sqlc"
	"example.com/at/backend/api-vet/utils"
)

type AuthCommonService interface {
	ValidateUniqueFields(email, phoneNumber string) error
	CheckPassword(hashPassword, givenPassword string) error
	FindUserByEmailOrPhone(email, phoneNumber string) (*DTOs.UserDTO, error)
	CompleteLogin(userDTO DTOs.UserDTO) (string, error)
	CreateJWT(userId int32, role string) (string, error)
}

type authCommonServiceImpl struct {
	userRepository repository.UserRepository
	userMappers    mappers.UserMappers
	vetRepository  repository.VeterinarianRepository
}

func NewCommonAuthService(userRepository repository.UserRepository, ownerRepository repository.OwnerRepository, vetRepository repository.VeterinarianRepository) AuthCommonService {
	return &authCommonServiceImpl{
		userRepository: userRepository,
		vetRepository:  vetRepository,
	}
}

func (as *authCommonServiceImpl) ValidateUniqueFields(email, phoneNumber string) error {
	if email == "" && phoneNumber == "" {
		return errors.New("no credentials provided to create user")
	}

	if email != "" {
		isEmailTaken := as.userRepository.CheckEmailExists(email)
		if isEmailTaken {
			return errors.New("email is already taken")
		}
	}

	if phoneNumber != "" {
		isPhoneNumberTaken := as.userRepository.CheckPhoneNumberExists(phoneNumber)
		if isPhoneNumberTaken {
			return errors.New("phone number is already taken")
		}
	}

	return nil
}

func (as *authCommonServiceImpl) CheckPassword(hashPassword, givenPassword string) error {
	if err := utils.CheckPassword(hashPassword, givenPassword); err != nil {
		return err
	}

	return nil
}

func (as *authCommonServiceImpl) FindUserByEmailOrPhone(email, phoneNumber string) (*DTOs.UserDTO, error) {
	var user *sqlc.User
	var err error

	if email != "" {
		// Find user by email
		user, err = as.userRepository.GetUserByEmail(email)
		if err != nil {
			return nil, fmt.Errorf("error finding user by email: %w", err)
		}
	} else if phoneNumber != "" {
		// Find user by phone number
		user, err = as.userRepository.GetUserByPhoneNumber(phoneNumber)
		if err != nil {
			return nil, fmt.Errorf("error finding user by phone number: %w", err)
		}
	} else {
		return nil, fmt.Errorf("email or phone number must be provided")
	}

	// Map the found user to DTO
	userDTO := as.userMappers.MapSqlcEntityToDTO(user)
	return &userDTO, nil
}

func (as *authCommonServiceImpl) CompleteLogin(userDTO DTOs.UserDTO) (string, error) {
	if err := as.userRepository.UpdateUserLastLogin(userDTO.Id); err != nil {
		return "", err
	}

	JWT, err := as.CreateJWT(userDTO.Id, userDTO.Role)
	if err != nil {
		return "", errors.New("can't create jwt token")
	}

	return JWT, nil
}

func (as *authCommonServiceImpl) CreateJWT(userId int32, role string) (string, error) {
	var roles []string
	roles = append(roles, role)
	userIdStr := Int32ToString(userId)

	JWT, err := utils.GenerateJWT(userIdStr, roles)
	if err != nil {
		return "", err
	}

	return JWT, nil
}

func Int32ToString(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}
*/
