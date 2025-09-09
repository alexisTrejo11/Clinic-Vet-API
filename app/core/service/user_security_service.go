package service

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type UserSecurityService struct {
	repository repository.UserRepository
}

func NewUserSecurityService(repository repository.UserRepository) *UserSecurityService {
	return &UserSecurityService{
		repository: repository,
	}
}

func (s *UserSecurityService) ProccesUserCreation(ctx context.Context, user user.User) error {
	return s.repository.Save(ctx, &user)
}

func (s *UserSecurityService) AuthenticateUser(ctx context.Context, email string, password string) (user.User, error) {
	return user.User{}, nil
}

func (s *UserSecurityService) ValidateUserCreation(ctx context.Context, email valueobject.Email, phone valueobject.PhoneNumber, rawPassword string) error {
	return nil
}

func (s *UserSecurityService) IsEmailUnique(ctx context.Context, email string) (bool, error) {
	exists, err := s.repository.ExistsByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return !exists, nil
}
