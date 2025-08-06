package userUsecase

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type UserCommandUsecase interface {
	CreateUserUseCase() *CreateUserUseCase
	DeleteUserUseCase() *DeleteUserUseCase
}

type CreateUserUseCase struct {
	repo userRepository.UserRepository
}

type DeleteUserUseCase struct {
	userRepository userRepository.UserRepository
}

type ChangeUserPasswordUseCase struct {
	userRepository userRepository.UserRepository
}

type BanUserUseCase struct {
	userRepository userRepository.UserRepository
}

type UnbanUserUseCase struct {
	userRepository userRepository.UserRepository
}

type UpdateUserStatusUseCase struct {
	userRepository userRepository.UserRepository
}

type UpdateUserCredentialsUseCase struct {
	userRepository userRepository.UserRepository
}

type userCommandUsecaseImpl struct {
	createUserUseCase *CreateUserUseCase
	deleteUserUseCase *DeleteUserUseCase
}

func NewUserCommandUsecase(repo userRepository.UserRepository) UserCommandUsecase {
	return &userCommandUsecaseImpl{
		createUserUseCase: &CreateUserUseCase{repo: repo},
		deleteUserUseCase: &DeleteUserUseCase{userRepository: repo},
	}
}

func (u *userCommandUsecaseImpl) CreateUserUseCase() *CreateUserUseCase {
	return u.createUserUseCase
}
func (u *userCommandUsecaseImpl) DeleteUserUseCase() *DeleteUserUseCase {
	return u.deleteUserUseCase
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, command userCommand.CreateUserCommand) userCommand.CommandResult {
	user, err := FromCreateCommand(command)
	if err != nil {
		userCommand.FailureResult("an error ocurrred mappping user", err)
	}

	if err := uc.validateBuissnessRules(ctx, *user); err != nil {
		return userCommand.FailureResult("an error ocurrred validating user", err)
	}

	if err := uc.proceessCreation(ctx, user); err != nil {
		return userCommand.FailureResult("an error ocurrred creating user", err)
	}

	return userCommand.SuccesResult(user.Id().String(), "user created successfully")
}

func (d *DeleteUserUseCase) Execute(ctx context.Context, userId int, softDelete bool) userCommand.CommandResult {
	if _, err := d.userRepository.GetById(ctx, userId); err != nil {
		return userCommand.FailureResult("failed to find user", err)
	}

	err := d.userRepository.Delete(ctx, userId, softDelete)
	if err != nil {
		return userCommand.FailureResult("failed to delete user", err)
	}

	return userCommand.FailureResult("User deleted successfully", nil)
}

func (uc *CreateUserUseCase) validateBuissnessRules(ctx context.Context, user userDomain.User) error {
	if err := userDomain.ValidatePassword(user.Password()); err != nil {
		return err
	}

	if err := uc.validateUniqueConstraints(ctx, user); err != nil {
		return err
	}

	return nil
}

// TO AUTH?
func (uc *CreateUserUseCase) validateUniqueConstraints(ctx context.Context, user userDomain.User) error {
	if exists, err := uc.repo.ExistsByEmail(ctx, user.Email().String()); err == nil {
		if err != nil {
			return err
		}
		if exists {
			return errors.New("email already exists")
		}
	}

	if exists, err := uc.repo.ExistsByPhone(ctx, user.PhoneNumber().String()); err == nil {
		if err != nil {
			return err
		}
		if exists {
			return errors.New("phone number already exists")
		}
	}

	return nil
}

func (uc *CreateUserUseCase) proceessCreation(ctx context.Context, user *userDomain.User) error {
	if err := uc.hashPassword(user); err != nil {
		return err
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	return nil
	// Event --> userCreatedEvent := userDomain.NewUserCreatedEvent(user)
}

func (u *CreateUserUseCase) hashPassword(user *userDomain.User) error {
	hashedPw, err := shared.HashPassword(user.Password())
	if err != nil {
		return err
	}

	user.SetPassword(hashedPw)
	return nil
}

func FromCreateCommand(command userCommand.CreateUserCommand) (*userDomain.User, error) {
	return &userDomain.User{}, nil
}
