// Package usecase implements the business logic for user profile management.
package usecase

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/address"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/profile"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type CreateProfileData struct {
	UserID      valueobject.UserID     `json:"user_id"`
	Name        valueobject.PersonName `json:"name"`
	Gender      enum.PersonGender      `json:"gender"`
	ProfilePic  string                 `json:"profile_pic"`
	Bio         string                 `json:"bio"`
	DateOfBirth time.Time              `json:"date_of_birth"`
	Address     *address.Address       `json:"address"`
}

type UpdateProfileData struct {
	UserID      valueobject.UserID      `json:"user_id"`
	Name        *valueobject.PersonName `json:"name"`
	Gender      *string                 `json:"gender"`
	ProfilePic  *string                 `json:"profile_pic"`
	Bio         *string                 `json:"bio"`
	DateOfBirth *time.Time              `json:"date_of_birth"`
	Address     *address.Address        `json:"address"`
}

type ProfileUseCases interface {
	GetUserProfile(ctx context.Context, userID valueobject.UserID) (p.Profile, error)
	CreateProfile(ctx context.Context, request CreateProfileData) error
	UpdateProfile(ctx context.Context, request UpdateProfileData) error
}

type profileUseCasesImpl struct {
	profileRepo repository.ProfileRepository
}

func NewProfileUseCases(repo repository.ProfileRepository) ProfileUseCases {
	return &profileUseCasesImpl{profileRepo: repo}
}

func (uc *profileUseCasesImpl) CreateProfile(ctx context.Context, request CreateProfileData) error {
	profile := &p.Profile{
		UserID:   request.UserID,
		Bio:      request.Bio,
		PhotoURL: request.ProfilePic,
	}
	return uc.profileRepo.Create(ctx, profile)
}

func (uc *profileUseCasesImpl) GetUserProfile(ctx context.Context, userID valueobject.UserID) (p.Profile, error) {
	profile, err := uc.profileRepo.GetByUserID(ctx, userID)
	if err != nil {
		return p.Profile{}, err
	}
	return profile, nil
}

func (uc *profileUseCasesImpl) UpdateProfile(ctx context.Context, request UpdateProfileData) error {
	profile, err := uc.profileRepo.GetByUserID(ctx, request.UserID)
	if err != nil {
		return err
	}

	applyProfileUpdates(&profile, request)
	uc.profileRepo.Update(ctx, &profile)
	return nil
}

func applyProfileUpdates(profile *p.Profile, request UpdateProfileData) {
	if request.Address != nil {
		profile.Address = request.Address
	}
	if request.Bio != nil {
		profile.Bio = *request.Bio
	}
	if request.DateOfBirth != nil {
		profile.DateOfBirth = request.DateOfBirth
	}
	if request.ProfilePic != nil {
		profile.PhotoURL = *request.ProfilePic
	}
}
