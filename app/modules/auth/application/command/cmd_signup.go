package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/service"
)

// SignupCommand represents the user registration request data
type SignupCommand struct {
	// User credentials - at least one is required
	Email       valueobject.Email       `json:"email" validate:"omitempty,email"`
	PhoneNumber valueobject.PhoneNumber `json:"phone_number" validate:"omitempty,phone"`
	Password    string                  `json:"password" validate:"required,min=8"`

	// Personal details
	FirstName   string            `json:"first_name" validate:"required,min=2,max=50"`
	LastName    string            `json:"last_name" validate:"required,min=2,max=50"`
	Gender      enum.PersonGender `json:"gender" validate:"required"`
	DateOfBirth time.Time         `json:"date_of_birth" validate:"required"`
	Location    string            `json:"location" validate:"required"`
	Address     string            `json:"address" validate:"required"`
	Role        enum.UserRole     `json:"role" validate:"required"`

	// Optional profile information
	ProfilePicture string `json:"profile_picture" validate:"omitempty,url"`
	Bio            string `json:"bio" validate:"omitempty,max=500"`

	// Veterinarian-specific details (optional)
	LicenseNumber   *string            `json:"license_number,omitempty"`
	Specialty       *enum.VetSpecialty `json:"specialty,omitempty"`
	YearsExperience *int               `json:"years_experience,omitempty" validate:"omitempty,min=0,max=50"`

	// Request metadata
	CTX       context.Context `json:"-"`
	IP        string          `json:"ip" validate:"required,ip"`
	UserAgent string          `json:"user_agent" validate:"required"`
	Source    string          `json:"source" validate:"required"`
}

// signupHandler handles user registration operations
type SignupHandler struct {
	userRepo        repository.UserRepository
	securityService service.UserSecurityService
}

// NewSignupCommandHandler creates a new instance of signupHandler
func NewSignupCommandHandler(userRepo repository.UserRepository, securityService service.UserSecurityService) AuthCommandHandler {
	return &SignupHandler{
		userRepo:        userRepo,
		securityService: securityService,
	}
}

// Handle processes the signup command and returns authentication result
func (h *SignupHandler) Handle(cmd any) AuthCommandResult {
	command, ok := cmd.(SignupCommand)
	if !ok {
		return FailureAuthResult(ErrInvalidCommandType, errors.New("expected SignupCommand"))
	}
	if err := h.securityService.ValidateUserCreation(command.CTX, command.Email, command.PhoneNumber, command.Password); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	user, err := h.toDomain(command)
	if err != nil {
		return FailureAuthResult(ErrDataParsingFailed, err)
	}

	if err := h.securityService.ProccesUserCreation(command.CTX, *user); err != nil {
		return FailureAuthResult(ErrUserCreationFailed, err)
	}

	return SuccessAuthResult(nil, user.ID().String(), MsgUserCreatedSuccess)
}

// toDomain converts SignupCommand to domain entity with comprehensive validation
func (h *SignupHandler) toDomain(command SignupCommand) (*user.User, error) {
	userEntity, err := user.NewUser(valueobject.UserID{}, command.Role, enum.UserStatusPending,
		user.WithEmail(command.Email),
		user.WithPhoneNumber(command.PhoneNumber),
		user.WithPassword(command.Password),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build user entity: %w", err)
	}

	return userEntity, nil
}
