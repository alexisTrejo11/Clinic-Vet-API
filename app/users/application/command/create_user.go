package userCommand

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type CreateUserHandler struct {
	repo userRepository.UserRepository
}

func NewCreateUserHandler(repo userRepository.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repo: repo,
	}
}

type CreateProfileCommand struct {
	FirstName   string
	LastName    string
	Gender      string
	ProfilePic  string
	Bio         string
	DateOfBirth time.Time
	Address     string
}

type CreateUserCommand struct {
	Email          string
	PhoneNumber    string
	Password       string
	Gender         string
	Phone          string
	Address        string
	Role           string
	OwnerId        *int
	VeterinarianId *int
	Status         string
	DateOfBirth    time.Time
	Profile        CreateProfileCommand
	Ctx            context.Context
}

func (uc *CreateUserHandler) Handle(cmd any) CommandResult {
	command := cmd.(CreateUserCommand)

	user, err := FromCreateCommand(command)
	if err != nil {
		FailureResult("an error ocurrred mappping user", err)
	}

	if err := uc.validateBuissnessRules(command.Ctx, *user); err != nil {
		return FailureResult("an error ocurrred validating user", err)
	}

	if err := uc.proceessCreation(command.Ctx, user); err != nil {
		return FailureResult("an error ocurrred creating user", err)
	}

	return SuccesResult(user.Id().String(), "user created successfully")
}

func (uc *CreateUserHandler) validateBuissnessRules(ctx context.Context, user userDomain.User) error {
	if err := userDomain.ValidatePassword(user.Password()); err != nil {
		return err
	}

	if err := uc.validateUniqueConstraints(ctx, user); err != nil {
		return err
	}

	return nil
}

func (uc *CreateUserHandler) validateUniqueConstraints(ctx context.Context, user userDomain.User) error {
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

func (uc *CreateUserHandler) proceessCreation(ctx context.Context, user *userDomain.User) error {
	if err := uc.hashPassword(user); err != nil {
		return err
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	return nil
	// Event --> userCreatedEvent := userDomain.NewUserCreatedEvent(user)
}

func (u *CreateUserHandler) hashPassword(user *userDomain.User) error {
	hashedPw, err := shared.HashPassword(user.Password())
	if err != nil {
		return err
	}

	user.SetPassword(hashedPw)
	return nil
}

func FromCreateCommand(command CreateUserCommand) (*userDomain.User, error) {
	return &userDomain.User{}, nil
}
