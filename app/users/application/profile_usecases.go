package userApplication

import (
	"context"
	"time"

	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type UpdateProfileRequest struct {
	UserId      user.UserId
	FirstName   *string
	LastName    *string
	Gender      *string
	ProfilePic  *string
	Bio         *string
	DateOfBirth *time.Time
	Address     *string
}

type ProfileUseCases interface {
	GetUserProfile(ctx context.Context, userId int) (user.Profile, error)
	UpdateProfileUseCase(ctx context.Context, request UpdateProfileRequest) error
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
	repo          userRepo.UserRepository
}

func NewProfileUseCases(
	repo userRepo.UserRepository,
	getProfile *GetProfileByIdUseCase,
	updateProfile *UpdateUserProfileUseCase,
) ProfileUseCases {
	return &profileUseCasesImpl{
		repo:          repo,
		getProfile:    getProfile,
		updateProfile: updateProfile,
	}
}

func (p *profileUseCasesImpl) UpdateProfileUseCase(ctx context.Context, request UpdateProfileRequest) error {
	return p.updateProfile.Execute(ctx, request)
}

func (p *profileUseCasesImpl) GetUserProfile(ctx context.Context, userId int) (user.Profile, error) {
	return p.getProfile.Execute(ctx, userId)
}

func (uc *UpdateUserProfileUseCase) Execute(ctx context.Context, request UpdateProfileRequest) error {
	user, err := uc.repo.GetByIdWithProfile(ctx, request.UserId.GetValue())
	if err != nil {
		return err
	}

	applyProfileUpdates(user, request)
	uc.repo.UpdateProfile(ctx, request.UserId.GetValue(), user.Profile())
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

func applyProfileUpdates(userEntity *user.User, request UpdateProfileRequest) {
	profile := userEntity.Profile()

	if request.FirstName != nil {
		profile.Name.FirstName = *request.FirstName
	}
	if request.LastName != nil {
		profile.Name.LastName = *request.LastName
	}
	if request.Address != nil {
		profile.Location = *request.Address
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
		profile.Gender = *request.Gender
	}

	userEntity.SetProfile(&profile)
}
