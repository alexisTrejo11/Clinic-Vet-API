// Package usecase implements the business logic for user profile management.
package usecase

import (
	"context"

	p "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/profile"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
)

type CreateProfileData struct {
	UserID     valueobject.UserID `json:"user_id"`
	ProfilePic string             `json:"profile_pic"`
	Bio        string             `json:"bio"`
}

type UpdateProfileData struct {
	UserID     valueobject.UserID `json:"user_id"`
	ProfilePic *string            `json:"profile_pic"`
	Bio        *string            `json:"bio"`
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
	if request.Bio != nil {
		profile.Bio = *request.Bio
	}
	if request.ProfilePic != nil {
		profile.PhotoURL = *request.ProfilePic
	}
}
