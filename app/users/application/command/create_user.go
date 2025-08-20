package userCommand

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

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
	Status         user.UserStatus
	DateOfBirth    time.Time
	Profile        CreateProfileCommand
	Ctx            context.Context
}

type CreateUserHandler struct {
	repo userRepository.UserRepository
}

func NewCreateUserHandler(repo userRepository.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repo: repo,
	}
}

func (uc *CreateUserHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(CreateUserCommand)

	user, err := FromCreateCommand(command)
	if err != nil {
		shared.FailureResult("an error ocurrred mappping user", err)
	}

	if err := uc.validateBuissnessRules(command.Ctx, *user); err != nil {
		return shared.FailureResult("an error ocurrred validating user", err)
	}

	if err := uc.proceessCreation(command.Ctx, user); err != nil {
		return shared.FailureResult("an error ocurrred creating user", err)
	}

	return shared.SuccessResult(user.Id().String(), "user created successfully")
}

func (uc *CreateUserHandler) validateBuissnessRules(ctx context.Context, userEntity user.User) error {
	if err := user.ValidatePassword(userEntity.Password()); err != nil {
		return err
	}

	if err := uc.validateUniqueConstraints(ctx, userEntity); err != nil {
		return err
	}

	return nil
}

func (uc *CreateUserHandler) validateUniqueConstraints(ctx context.Context, user user.User) error {
	exists, err := uc.repo.ExistsByEmail(ctx, user.Email().String())
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	exists, err = uc.repo.ExistsByPhone(ctx, user.PhoneNumber().String())
	if err != nil {
		return err
	}

	if exists {
		return errors.New("phone number already exists")
	}

	return nil
}

func (uc *CreateUserHandler) proceessCreation(ctx context.Context, user *user.User) error {
	if err := uc.hashPassword(user); err != nil {
		return err
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	return nil
	// Event --> userCreatedEvent := user.NewUserCreatedEvent(user)
}

func (u *CreateUserHandler) hashPassword(user *user.User) error {
	hashedPw, err := shared.HashPassword(user.Password())
	if err != nil {
		return err
	}

	user.SetPassword(hashedPw)
	return nil
}

func FromCreateCommand(command CreateUserCommand) (*user.User, error) {
	return &user.User{}, nil
}
