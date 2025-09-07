package command

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
)

const (
	ErrValidationFailed   = "validation failed for unique credentials"
	ErrDataParsingFailed  = "failed to parse user data"
	ErrUserSaveFailed     = "failed to save user"
	ErrEmailAlreadyExists = "email already exists"
	ErrPhoneAlreadyExists = "phone number already exists"
	ErrInvalidCommandType = "invalid command type"
	ErrMissingCredentials = "either email or phone number is required"

	MsgUserCreatedSuccess = "user successfully created"

	MaxConcurrentValidations = 2
)

// SignupCommand represents the user registration request data
type SignupCommand struct {
	// User credentials - at least one is required
	Email       *string `json:"email" validate:"omitempty,email"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,phone"`
	Password    string  `json:"password" validate:"required,min=8"`
	UserID      int     `json:"user_id"`

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
type signupHandler struct {
	userRepo        repository.UserRepository
	passwordEndoder password.PasswordEncoder
}

// NewSignupCommandHandler creates a new instance of signupHandler
func NewSignupCommandHandler(userRepo repository.UserRepository, passwordEndoder password.PasswordEncoder) AuthCommandHandler {
	return &signupHandler{
		userRepo:        userRepo,
		passwordEndoder: passwordEndoder,
	}
}

// Handle processes the signup command and returns authentication result
func (h *signupHandler) Handle(cmd any) AuthCommandResult {
	command, ok := cmd.(SignupCommand)
	if !ok {
		return FailureAuthResult(ErrInvalidCommandType, errors.New("expected SignupCommand"))
	}

	if err := h.validateCredentialsPresent(&command); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	if err := h.validateUniqueCredentials(&command); err != nil {
		return FailureAuthResult(ErrValidationFailed, err)
	}

	if err := user.ValidatePassword(command.Password); err != nil {
		return FailureAuthResult(err.Error(), err)
	}

	user, err := h.toDomain(command)
	if err != nil {
		return FailureAuthResult(ErrDataParsingFailed, err)
	}

	if err := h.hashUserPassword(user); err != nil {
		return FailureAuthResult("Error Hashing password", err)
	}

	if err := h.userRepo.Save(command.CTX, user); err != nil {
		return FailureAuthResult(ErrUserSaveFailed, err)
	}

	return SuccessAuthResult(nil, user.ID().String(), MsgUserCreatedSuccess)
}

// validateCredentialsPresent ensures at least email or phone is provided
func (h *signupHandler) validateCredentialsPresent(command *SignupCommand) error {
	if command.Email == nil && command.PhoneNumber == nil {
		return errors.New(ErrMissingCredentials)
	}
	return nil
}

// validateUniqueCredentials checks if email and phone are unique using concurrent validation
func (h *signupHandler) validateUniqueCredentials(command *SignupCommand) error {
	var wg sync.WaitGroup
	errChan := make(chan error, MaxConcurrentValidations)

	if command.Email != nil && *command.Email != "" {
		wg.Add(1)
		go h.validateEmailUnique(command.CTX, *command.Email, &wg, errChan)
	}

	if command.PhoneNumber != nil && *command.PhoneNumber != "" {
		wg.Add(1)
		go h.validatePhoneUnique(command.CTX, *command.PhoneNumber, &wg, errChan)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// validateEmailUnique checks if email is already registered
func (h *signupHandler) validateEmailUnique(ctx context.Context, email string, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	exists, err := h.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		errChan <- fmt.Errorf("email validation error: %w", err)
		return
	}

	if exists {
		errChan <- apperror.ConflictError("Email", ErrEmailAlreadyExists)
		return
	}
}

// validatePhoneUnique checks if phone number is already registered
func (h *signupHandler) validatePhoneUnique(ctx context.Context, phone string, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	exists, err := h.userRepo.ExistsByPhone(ctx, phone)
	if err != nil {
		errChan <- fmt.Errorf("phone validation error: %w", err)
		return
	}

	if exists {
		errChan <- apperror.ConflictError("PhoneNumber", ErrPhoneAlreadyExists)
		return
	}
}

// toDomain converts SignupCommand to domain entity with comprehensive validation
func (h *signupHandler) toDomain(command SignupCommand) (*user.User, error) {
	var errorsMessages []string

	userID, err := valueobject.NewUserID(command.UserID)
	if err != nil {
		errorsMessages = append(errorsMessages, fmt.Sprintf("user ID: %v", err))
	}

	if err := user.ValidatePassword(command.Password); err != nil {
		errorsMessages = append(errorsMessages, fmt.Sprintf("password: %v", err))
	}

	if len(errorsMessages) > 0 {
		return nil, fmt.Errorf("validation errors: %s", strings.Join(errorsMessages, "; "))
	}

	opts := []user.UserOption{
		user.WithPassword(command.Password),
	}

	if command.Email != nil {
		if *command.Email != "" {
			email, err := valueobject.NewEmail(*command.Email)
			if err != nil {
				errorsMessages = append(errorsMessages, fmt.Sprintf("email: %v", err))
			} else {
				opts = append(opts, user.WithEmail(email))
			}
		}
	}

	if command.PhoneNumber != nil {
		if *command.PhoneNumber != "" {
			phone, err := valueobject.NewPhoneNumber(*command.PhoneNumber)
			if err != nil {
				errorsMessages = append(errorsMessages, fmt.Sprintf("phone: %v", err))
			} else {
				opts = append(opts, user.WithPhoneNumber(phone))
			}
		}
	}

	if len(errorsMessages) > 0 {
		return nil, fmt.Errorf("validation errors: %s", strings.Join(errorsMessages, "; "))
	}

	userEntity, err := user.NewUser(userID, command.Role, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to build user entity: %w", err)
	}

	return userEntity, nil
}

func (h *signupHandler) hashUserPassword(user *user.User) error {
	hashPassword, err := h.passwordEndoder.HashPassword(user.Password())
	if err != nil {
		return err
	}

	user.UpdatePassword(hashPassword)

	return nil
}
