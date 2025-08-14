package userApplication

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type ProfileUpdate struct {
	UserId      int                      `json:"user_id"`
	Name        *valueObjects.PersonName `json:"name"`
	Gender      *string                  `json:"gender"`
	ProfilePic  *string                  `json:"profile_pic"`
	Bio         *string                  `json:"bio"`
	DateOfBirth *time.Time               `json:"date_of_birth"`
	Address     *user.Address            `json:"address"`
}

type ProfileUseCases interface {
	GetUserProfile(ctx context.Context, userId int) (user.Profile, error)
	UpdateProfileUseCase(ctx context.Context, request ProfileUpdate) error
}

type GetProfileByIdUseCase struct {
	repo userRepo.UserRepository
}

type UpdateUserProfileUseCase struct {
	repo userRepo.UserRepository
}

type profileUseCasesImpl struct {
	getProfile    *GetProfileByIdUseCase
	updateProfile *UpdateUserProfileUseCase
}

func NewProfileUseCases(
	repo userRepo.UserRepository,

) ProfileUseCases {
	return &profileUseCasesImpl{
		getProfile:    &GetProfileByIdUseCase{repo: repo},
		updateProfile: &UpdateUserProfileUseCase{repo: repo},
	}
}

func (p *profileUseCasesImpl) UpdateProfileUseCase(ctx context.Context, request ProfileUpdate) error {
	return p.updateProfile.Execute(ctx, request)
}

func (p *profileUseCasesImpl) GetUserProfile(ctx context.Context, userId int) (user.Profile, error) {
	return p.getProfile.Execute(ctx, userId)
}

func (uc *UpdateUserProfileUseCase) Execute(ctx context.Context, request ProfileUpdate) error {
	user, err := uc.repo.GetByIdWithProfile(ctx, request.UserId)
	if err != nil {
		return err
	}

	applyProfileUpdates(user, request)
	uc.repo.UpdateProfile(ctx, request.UserId, user.Profile())
	return nil
}

func (uc *GetProfileByIdUseCase) Execute(ctx context.Context, userId int) (user.Profile, error) {
	userEntity, err := uc.repo.GetByIdWithProfile(ctx, userId)
	if err != nil {
		return user.Profile{}, err
	}
	profile := userEntity.Profile()
	return profile, nil
}

func applyProfileUpdates(userEntity *user.User, request ProfileUpdate) {
	profile := userEntity.Profile()

	if (request.Name != nil) && (*request.Name == valueObjects.PersonName{}) {
		profile.Name = *request.Name
	}

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
	if request.Gender != nil {
		profile.Gender = valueObjects.NewGender(*request.Gender)
	}

	userEntity.SetProfile(&profile)
}
