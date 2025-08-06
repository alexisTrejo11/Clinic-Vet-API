package userUsecase

import (
	"context"

	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type ProfileUseCases interface {
	GetUserProfile(ctx context.Context, userId int) (userDomain.Profile, error)
	UpdateProfileUseCase(ctx context.Context, command userCommand.UpdateProfileCommand) error
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

func (p *profileUseCasesImpl) UpdateProfileUseCase(ctx context.Context, command userCommand.UpdateProfileCommand) error {
	return p.updateProfile.Execute(ctx, command)
}

func (p *profileUseCasesImpl) GetUserProfile(ctx context.Context, userId int) (userDomain.Profile, error) {
	return p.getProfile.Execute(ctx, userId)
}

func (uc *UpdateUserProfileUseCase) Execute(ctx context.Context, command userCommand.UpdateProfileCommand) error {
	user, err := uc.repo.GetById(ctx, command.UserId.GetValue())
	if err != nil {
		return err
	}

	applyProfileUpdates(user, command)
	uc.repo.UpdateProfile(ctx, command.UserId.GetValue(), user.Profile())
	return nil
}

func (uc *GetProfileByIdUseCase) Execute(ctx context.Context, userId int) (userDomain.Profile, error) {
	user, err := uc.repo.GetById(ctx, userId)
	if err != nil {
		return userDomain.Profile{}, err
	}
	profile := user.Profile()
	return profile, nil
}

func applyProfileUpdates(user *userDomain.User, command userCommand.UpdateProfileCommand) {
	profile := user.Profile()

	if command.FirstName != nil {
		profile.Name.FirstName = *command.FirstName
	}
	if command.LastName != nil {
		profile.Name.LastName = *command.LastName
	}
	if command.Address != nil {
		profile.Location = *command.Address
	}
	if command.Bio != nil {
		profile.Bio = *command.Bio
	}
	if command.DateOfBirth != nil {
		profile.DateOfBirth = command.DateOfBirth
	}
	if command.ProfilePic != nil {
		profile.PhotoURL = *command.ProfilePic
	}
	if command.Gender != nil {
		profile.Gender = *command.Gender
	}

	user.SetProfile(&profile)
}
