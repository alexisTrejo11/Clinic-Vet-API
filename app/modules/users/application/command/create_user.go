package userDomainCommand

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
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
	Status         userDomain.UserStatus
	DateOfBirth    time.Time
	Profile        CreateProfileCommand
	Ctx            context.Context
}

type CreateUserHandler struct {
	repo            userDomain.UserRepository
	securityService *userDomain.UseSecurityService
}

func NewCreateUserHandler(repo userDomain.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repo:            repo,
		securityService: userDomain.NewUseSecurityService(repo),
	}
}

func (uc *CreateUserHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(CreateUserCommand)

	user, err := FromCreateCommand(command)
	if err != nil {
		shared.FailureResult("an error ocurrred mappping user", err)
	}

	if err := uc.securityService.ValidateCreation(command.Ctx, *user); err != nil {
		return shared.FailureResult("an error occurred validating user", err)
	}

	if err := uc.securityService.HashPassword(user); err != nil {
		return shared.FailureResult("an error occurred hashing password", err)
	}

	if err := uc.repo.Save(command.Ctx, user); err != nil {
		return shared.FailureResult("an error occurred saving user", err)
	}

	return shared.SuccessResult(user.Id().String(), "user created successfully")
}

func FromCreateCommand(command CreateUserCommand) (*userDomain.User, error) {
	return &userDomain.User{}, nil
}
