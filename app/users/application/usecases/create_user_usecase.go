package userUsecase

import (
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/repositories"
)

type CreateUserUseCase struct {
	repo userRepository.UserRepository
}

func NewCreateUserUseCase(repo userRepository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo: repo,
	}
}

func (c *CreateUserUseCase) Execute() error {
	// Implementation for creating a user goes here
	return nil
}
