package usecase

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/address"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/profile"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type ProfileUpdate struct {
	UserID      valueobject.UserID      `json:"user_id"`
	Name        *valueobject.PersonName `json:"name"`
	Gender      *string                 `json:"gender"`
	ProfilePic  *string                 `json:"profile_pic"`
	Bio         *string                 `json:"bio"`
	DateOfBirth *time.Time              `json:"date_of_birth"`
	Address     *address.Address        `json:"address"`
}

type ProfileUseCases interface {
	GetUserProfile(ctx context.Context, userID valueobject.UserID) (profile.Profile, error)
	UpdateProfileUseCase(ctx context.Context, request ProfileUpdate) error
}

type GetProfileByIDUseCase struct {
	repo repository.ProfileRepository
}

type UpdateUserProfileUseCase struct {
	repo repository.ProfileRepository
}

type profileUseCasesImpl struct {
	getProfile    *GetProfileByIDUseCase
	updateProfile *UpdateUserProfileUseCase
}

func NewProfileUseCases(
	repo repository.ProfileRepository,
) ProfileUseCases {
	return &profileUseCasesImpl{
		getProfile:    &GetProfileByIDUseCase{repo: repo},
		updateProfile: &UpdateUserProfileUseCase{repo: repo},
	}
}

func (p *profileUseCasesImpl) UpdateProfileUseCase(ctx context.Context, request ProfileUpdate) error {
	return p.updateProfile.Execute(ctx, request)
}

func (p *profileUseCasesImpl) GetUserProfile(ctx context.Context, userID valueobject.UserID) (profile.Profile, error) {
	return p.getProfile.Execute(ctx, userID)
}

func (uc *UpdateUserProfileUseCase) Execute(ctx context.Context, request ProfileUpdate) error {
	profile, err := uc.repo.GetByUserID(ctx, request.UserID.Value())
	if err != nil {
		return err
	}

	applyProfileUpdates(&profile, request)
	uc.repo.Update(ctx, &profile)
	return nil
}

func (uc *GetProfileByIDUseCase) Execute(ctx context.Context, userID valueobject.UserID) (profile.Profile, error) {
	profileEntity, err := uc.repo.GetByUserID(ctx, userID.Value())
	if err != nil {
		return profile.Profile{}, err
	}
	return profileEntity, nil
}

func applyProfileUpdates(profile *profile.Profile, request ProfileUpdate) {
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
