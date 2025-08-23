package authCmd

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"

	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type SignupCommand struct {
	// User credentials
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Password    string  `json:"password"`
	UserId      int     `json:"user_id"`

	// Personal details
	FirstName      string              `json:"first_name"`
	LastName       string              `json:"last_name"`
	Gender         valueObjects.Gender `json:"gender"`
	DateOfBirth    time.Time           `json:"date_of_birth"`
	Location       string              `json:"location"`
	Address        string              `json:"address"`
	Role           userDomain.UserRole `json:"role"`
	ProfilePicture string              `json:"profile_picture"`
	Bio            string              `json:"bio"`

	// Veterinarian details
	LicenseNumber   *string
	Specialty       *vetDomain.VetSpecialty
	YearsExperience *int

	CTX context.Context `json:"-"`

	// Metadata
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Source    string `json:"source"`
}

type signupHandler struct {
	userRepo userDomain.UserRepository
}

func NewSignupCommandHandler(
	userRepo userDomain.UserRepository,
) *signupHandler {
	return &signupHandler{
		userRepo: userRepo,
	}
}

func (h *signupHandler) Handle(cmd SignupCommand) AuthCommandResult {
	if err := h.validateCredentials(&cmd); err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("an conflict ocurred", err)}
	}

	user, err := toDomain(cmd)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("an error ocurred while converting to domain", err)}
	}

	if err := h.userRepo.Save(cmd.CTX, &user); err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("an error ocurred while creating user", err)}
	}

	return AuthCommandResult{CommandResult: shared.SuccessResult(user.Id().String(), "User created successfully")}

}

func (h *signupHandler) validateCredentials(cmd *SignupCommand) error {
	if cmd.Email == nil && cmd.PhoneNumber == nil {
		return errors.New("email or phone number must be provided")
	}

	if err := h.validateUniqueCredentials(cmd); err != nil {
		return err
	}

	return nil
}

func (h *signupHandler) validateUniqueCredentials(cmd *SignupCommand) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	if cmd.Email != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exists, err := h.userRepo.ExistsByEmail(cmd.CTX, *cmd.Email)
			if err != nil {
				errChan <- err
				return
			}
			if exists {
				errChan <- errors.New("email already exists")
			}
		}()
	}

	if cmd.PhoneNumber != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exists, err := h.userRepo.ExistsByPhone(cmd.CTX, *cmd.PhoneNumber)
			if err != nil {
				errChan <- err
				return
			}
			if exists {
				errChan <- errors.New("phone number already exists")
			}
		}()
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func toDomain(cmd SignupCommand) (userDomain.User, error) {
	user, err := userDomain.NewUserBuilder().
		WithId(cmd.UserId).
		WithPassword(cmd.Password).
		WithRole(cmd.Role).
		WithPhoneNumber(*cmd.PhoneNumber).
		WithEmail(*cmd.Email).
		Build()

	if err != nil {
		return userDomain.User{}, err
	}

	return user, nil
}
