// Package usecase implements the business logic for user profile management.
package usecase

import (
	"context"

	p "clinic-vet-api/app/core/domain/entity/user/profile"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
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
	GetUserProfile(ctx context.Context, userID uint) (map[string]any, error)
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

// Implement GetUserProfile to fetch user profile details by user ID.
func (uc *profileUseCasesImpl) GetUserProfile(ctx context.Context, userID uint) (map[string]any, error) {
	/*
		profile, err := uc.profileRepo.GetMapByUserID(ctx, valueobject.NewUserID(userID))
		if err != nil {
			return nil, err
		}
	*/

	dummyData := map[string]any{
		"user_id":     "123",
		"profile_pic": "https://example.com/profile.jpg",
		"bio":         "This is a sample bio.",
		"name":        "John Doe", // Placeholder for actual user name
		"gender":      "male",
		"birth_date":  "1990-01-01", // Placeholder for actual birth date
		"email":       "email@email",
		"phone":       "1234567890",
		"location":    "City, Country",
		"joined_at":   "2020-01-01", // Placeholder for actual join date
	}
	return dummyData, nil
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
