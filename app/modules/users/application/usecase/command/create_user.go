package command

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

// Error messages as variables
const (
	ErrFailedMappingUser   = "failed to map user from command"
	ErrFailedValidation    = "failed to validate user creation"
	ErrFailedHashPassword  = "failed to hash password"
	ErrFailedSaveUser      = "failed to save user"
	ErrInvalidRole         = "invalid user role"
	ErrInvalidStatus       = "invalid user status"
	ErrInvalidGender       = "invalid gender"
	ErrInvalidDateOfBirth  = "invalid date of birth"
	ErrUserCreationSuccess = "user created successfully"
)

type createProfileCommand struct {
	firstName   string
	lastName    string
	gender      string
	profilePic  string
	bio         string
	dateOfBirth time.Time
	address     string
}

type createUserCommand struct {
	email          string
	phoneNumber    string
	password       string
	gender         string
	phone          string
	address        string
	role           string
	ownerID        *int
	veterinarianID *int
	status         string
	dateOfBirth    time.Time
	profile        createProfileCommand
	ctx            context.Context
}

type CreateUserHandler struct {
	repo            repository.UserRepository
	securityService *service.UseSecurityService
}

// NewCreateUserCommand creates a new user command from primitive types
func NewCreateUserCommand(
	ctx context.Context,
	email, phoneNumber, password, gender, phone, address, role string,
	ownerID, veterinarianID *int,
	status string,
	dateOfBirth time.Time,
	firstName, lastName, profilePic, bio string,
) (createUserCommand, error) {
	if email == "" {
		return createUserCommand{}, errors.New(ErrInvalidEmail)
	}
	if phoneNumber == "" {
		return createUserCommand{}, errors.New(ErrInvalidPhone)
	}
	if role == "" {
		return createUserCommand{}, errors.New(ErrInvalidRole)
	}
	if status == "" {
		return createUserCommand{}, errors.New(ErrInvalidStatus)
	}
	if gender == "" {
		return createUserCommand{}, errors.New(ErrInvalidGender)
	}
	if dateOfBirth.IsZero() {
		return createUserCommand{}, errors.New(ErrInvalidDateOfBirth)
	}

	profile := createProfileCommand{
		firstName:   firstName,
		lastName:    lastName,
		gender:      gender,
		profilePic:  profilePic,
		bio:         bio,
		dateOfBirth: dateOfBirth,
		address:     address,
	}

	return createUserCommand{
		email:          email,
		phoneNumber:    phoneNumber,
		password:       password,
		gender:         gender,
		phone:          phone,
		address:        address,
		role:           role,
		ownerID:        ownerID,
		veterinarianID: veterinarianID,
		status:         status,
		dateOfBirth:    dateOfBirth,
		profile:        profile,
		ctx:            ctx,
	}, nil
}

func NewCreateUserHandler(repo repository.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{
		repo:            repo,
		securityService: service.NewUseSecurityService(repo),
	}
}

func (uc *CreateUserHandler) Handle(cmd any) cqrs.CommandResult {
	command, ok := cmd.(createUserCommand)
	if !ok {
		return cqrs.FailureResult(ErrFailedMappingUser, errors.New("invalid command type"))
	}

	user, err := fromCreateCommand(command)
	if err != nil {
		return cqrs.FailureResult(ErrFailedMappingUser, err)
	}

	if err := uc.securityService.ValidateCreation(command.ctx, *user); err != nil {
		return cqrs.FailureResult(ErrFailedValidation, err)
	}

	if err := uc.securityService.HashPassword(user); err != nil {
		return cqrs.FailureResult(ErrFailedHashPassword, err)
	}

	if err := uc.repo.Save(command.ctx, user); err != nil {
		return cqrs.FailureResult(ErrFailedSaveUser, err)
	}

	return cqrs.SuccessResult(user.ID().String(), ErrUserCreationSuccess)
}

func fromCreateCommand(command createUserCommand) (*entity.User, error) {
	// Map email
	email, err := valueobject.NewEmail(command.email)
	if err != nil {
		return nil, errors.New(ErrInvalidEmail)
	}

	// Map phone number
	phoneNumber, err := valueobject.NewPhoneNumber(command.phoneNumber)
	if err != nil {
		return nil, errors.New(ErrInvalidPhone)
	}

	// Map role
	userRole, err := enum.NewUserRole(command.role)
	if err != nil {
		return nil, errors.New(ErrInvalidRole)
	}

	// Map status
	userStatus, err := enum.UserStatus(command.status)
	if err != nil {
		return nil, errors.New(ErrInvalidStatus)
	}

	// Map gender
	userGender := enum.NewGender(command.gender)
	if err != nil {
		return nil, errors.New(ErrInvalidGender)
	}

	// Create profile
	profile := entity.Profile{
		FirstName:   command.profile.firstName,
		LastName:    command.profile.lastName,
		Gender:      userGender,
		ProfilePic:  command.profile.profilePic,
		Bio:         command.profile.bio,
		DateOfBirth: command.profile.dateOfBirth,
		Address:     command.profile.address,
	}

	// Create user
	user := &entity.User{
		Email:          email,
		PhoneNumber:    phoneNumber,
		Password:       command.password, // Will be hashed later
		Gender:         userGender,
		Phone:          command.phone,
		Address:        command.address,
		Role:           userRole,
		OwnerID:        command.ownerID,
		VeterinarianID: command.veterinarianID,
		Status:         userStatus,
		DateOfBirth:    command.dateOfBirth,
		Profile:        profile,
	}

	return user, nil
}
