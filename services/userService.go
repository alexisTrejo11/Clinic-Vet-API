package services

import (
	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/repository"
)

type UserService interface {
	CreateUser(userSignupDTO dtos.UserSignupDTO) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return UserServiceImpl{
		userRepository: userRepository,
	}
}

func (us UserServiceImpl) CreateUser(userSignupDTO dtos.UserSignupDTO) error {

}
