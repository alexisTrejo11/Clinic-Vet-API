package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
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
	OwnerID        *int
	VeterinarianID *int
	Status         enum.UserStatus
	DateOfBirth    time.Time
	Profile        CreateProfileCommand
	Ctx            context.Context
}

type CreateUserHandler struct {
	repo            repository.UserRepository
	securityService *service.UseSecurityService
}

func NewCreateUserHandler(repo repository.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repo:            repo,
		securityService: service.NewUseSecurityService(repo),
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

	return shared.SuccessResult(user.ID().String(), "user created successfully")
}

func FromCreateCommand(command CreateUserCommand) (*entity.User, error) {
	return &entity.User{}, nil
}
