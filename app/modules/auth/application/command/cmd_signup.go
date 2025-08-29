package command

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
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
	Role           enum.UserRole       `json:"role"`
	ProfilePicture string              `json:"profile_picture"`
	Bio            string              `json:"bio"`

	// Veterinarian details
	LicenseNumber   *string
	Specialty       *enum.VetSpecialty
	YearsExperience *int

	CTX context.Context `json:"-"`

	// Metadata
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Source    string `json:"source"`
}

type signupHandler struct {
	userRepo repository.UserRepository
}

func NewSignupCommandHandler(
	userRepo repository.UserRepository,
) AuthCommandHandler {
	return &signupHandler{
		userRepo: userRepo,
	}
}

func (h *signupHandler) Handle(cmd any) AuthCommandResult {
	command := cmd.(SignupCommand)

	if err := h.validateCredentials(&command); err != nil {
		return FailureAuthResult("an error or coflinct ocurred validting unique credentials", err)
	}

	user, err := toDomain(command)
	if err != nil {
		return FailureAuthResult("an error ocurred parsing data", err)
	}

	if err := h.userRepo.Save(command.CTX, &user); err != nil {
		return FailureAuthResult("an error ocurred saving user", err)
	}

	return SuccessAuthResult(nil, user.Id().String(), "User Successfully Created")
}

func (h *signupHandler) validateCredentials(command *SignupCommand) error {
	if command.Email == nil && command.PhoneNumber == nil {
		return errors.New("email or phone number must be provided")
	}

	if err := h.validateUniqueCredentials(command); err != nil {
		return err
	}

	return nil
}

func (h *signupHandler) validateUniqueCredentials(command *SignupCommand) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	if command.Email != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exists, err := h.userRepo.ExistsByEmail(command.CTX, *command.Email)
			if err != nil {
				errChan <- err
				return
			}
			if exists {
				errChan <- errors.New("email already exists")
			}
		}()
	}

	if command.PhoneNumber != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exists, err := h.userRepo.ExistsByPhone(command.CTX, *command.PhoneNumber)
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

func toDomain(command SignupCommand) (entity.User, error) {
	user, err := entity.NewUserBuilder().
		WithId(command.UserId).
		WithPassword(command.Password).
		WithRole(command.Role).
		WithPhoneNumber(*command.PhoneNumber).
		WithEmail(*command.Email).
		Build()
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
