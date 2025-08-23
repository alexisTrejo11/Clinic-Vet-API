package userApplication

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type ProfileUpdate struct {
	UserId      int                      `json:"user_id"`
	Name        *valueObjects.PersonName `json:"name"`
	Gender      *string                  `json:"gender"`
	ProfilePic  *string                  `json:"profile_pic"`
	Bio         *string                  `json:"bio"`
	DateOfBirth *time.Time               `json:"date_of_birth"`
	Address     *userDomain.Address      `json:"address"`
}

type ProfileUseCases interface {
	GetUserProfile(ctx context.Context, userId int) (userDomain.Profile, error)
	UpdateProfileUseCase(ctx context.Context, request ProfileUpdate) error
}

type GetProfileByIdUseCase struct {
	repo userDomain.ProfileRepository
}

type UpdateUserProfileUseCase struct {
	repo userDomain.ProfileRepository
}

type profileUseCasesImpl struct {
	getProfile    *GetProfileByIdUseCase
	updateProfile *UpdateUserProfileUseCase
}

func NewProfileUseCases(
	repo userDomain.ProfileRepository,

) ProfileUseCases {
	return &profileUseCasesImpl{
		getProfile:    &GetProfileByIdUseCase{repo: repo},
		updateProfile: &UpdateUserProfileUseCase{repo: repo},
	}
}

func (p *profileUseCasesImpl) UpdateProfileUseCase(ctx context.Context, request ProfileUpdate) error {
	return p.updateProfile.Execute(ctx, request)
}

func (p *profileUseCasesImpl) GetUserProfile(ctx context.Context, userId int) (userDomain.Profile, error) {
	return p.getProfile.Execute(ctx, userId)
}

func (uc *UpdateUserProfileUseCase) Execute(ctx context.Context, request ProfileUpdate) error {
	profile, err := uc.repo.GetByUserId(ctx, request.UserId)
	if err != nil {
		return err
	}

	applyProfileUpdates(&profile, request)
	uc.repo.Update(ctx, &profile)
	return nil
}

func (uc *GetProfileByIdUseCase) Execute(ctx context.Context, userId int) (userDomain.Profile, error) {
	profile, err := uc.repo.GetByUserId(ctx, userId)
	if err != nil {
		return userDomain.Profile{}, err
	}
	return profile, nil
}

func applyProfileUpdates(profile *userDomain.Profile, request ProfileUpdate) {
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
