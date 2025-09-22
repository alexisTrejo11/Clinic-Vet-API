// Package usecase implements the business logic for user profile management.
package usecase

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
)

type ProfileUseCases interface {
	GetUserProfile(ctx context.Context, userID uint, customerID uint, employeID uint) (map[string]any, error)
}

type profileUseCasesImpl struct {
	profileRepo repository.ProfileRepository
}

func NewProfileUseCases(repo repository.ProfileRepository) ProfileUseCases {
	return &profileUseCasesImpl{profileRepo: repo}
}

// Implement GetUserProfile to fetch user profile details by user ID.
func (uc *profileUseCasesImpl) GetUserProfile(ctx context.Context, userID uint, customerID uint, employeID uint) (map[string]any, error) {
	var role enum.UserRole
	if customerID != 0 {
		role = enum.UserRoleCustomer
	} else if employeID != 0 {
		role = enum.UserRoleVeterinarian
	} else {
		role = enum.UserRoleAdmin
	}
	profile, err := uc.profileRepo.GetMapByUserID(ctx, valueobject.NewUserID(userID), role)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
