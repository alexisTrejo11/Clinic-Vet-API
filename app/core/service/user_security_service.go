// Package service contains the business logic for entities services.
package service

import (
	"context"
	"errors"
	"regexp"
	"sync"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/event"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/password"
)

type EventPublisher interface {
	Publish(event any) error
}

type UserSecurityService struct {
	userRepo        repository.UserRepository
	employeeRepo    repository.VetRepository
	passwordEncoder password.PasswordEncoder
	eventPublisher  EventPublisher
}

func NewUserSecurityService(
	userRepo repository.UserRepository,
	employeeRepo repository.VetRepository,
	passwordEncoder password.PasswordEncoder,
	eventPublisher EventPublisher,
) *UserSecurityService {
	return &UserSecurityService{
		userRepo:        userRepo,
		employeeRepo:    employeeRepo,
		passwordEncoder: passwordEncoder,
	}
}

func (s *UserSecurityService) ProcessUserCreation(ctx context.Context, user *user.User) error {
	hashedPassword, err := s.passwordEncoder.HashPassword(user.Password())
	if err != nil {
		return err
	}
	user.SetHashedPassword(hashedPassword)

	if err := s.userRepo.Save(ctx, user); err != nil {
		return err
	}

	go s.ProduceUserCreatedEvent(*user)

	return nil
}

func (s *UserSecurityService) AuthenticateUser(ctx context.Context, email string, password string) (user.User, error) {
	return user.User{}, nil
}

func (s *UserSecurityService) ValidateUserCredentials(ctx context.Context, email valueobject.Email, phone *valueobject.PhoneNumber, rawPassword string) error {
	var wg sync.WaitGroup
	errorChannel := make(chan error, 3)
	var validationErrors []error

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.ValidateEmail(ctx, email); err != nil {
			errorChannel <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.ValidatePhoneNumber(ctx, phone); err != nil {
			errorChannel <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.ValidateRawPassword(rawPassword); err != nil {
			errorChannel <- err
		}
	}()

	go func() {
		wg.Wait()
		close(errorChannel)
	}()

	for err := range errorChannel {
		if err != nil {
			validationErrors = append(validationErrors, err)
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors[0]
	}

	return nil
}

func (s *UserSecurityService) ValidateEmployeeAccount(ctx context.Context, employeeID valueobject.VetID) error {
	var wg sync.WaitGroup
	errorChannel := make(chan error, 3)
	var validationErrors []error
	wg.Add(1)
	go func() {
		defer wg.Done()
		exists, err := s.employeeRepo.Exists(ctx, employeeID)
		if err != nil {
			errorChannel <- err
			return
		}
		if !exists {
			errorChannel <- errors.New("a valid employee is required to create an employee account")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		exists, err := s.userRepo.ExistsByEmployeeID(ctx, employeeID)
		if err != nil {
			errorChannel <- err
			return
		}
		if exists {
			errorChannel <- errors.New("employee already has an account")
		}
	}()

	go func() {
		wg.Wait()
		close(errorChannel)
	}()

	for err := range errorChannel {
		if err != nil {
			validationErrors = append(validationErrors, err)
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors[0]
	}

	return nil
}

func (s *UserSecurityService) ValidateEmail(ctx context.Context, email valueobject.Email) error {
	exists, err := s.userRepo.ExistsByEmail(ctx, email.String())
	if err != nil {
		return err
	}

	if exists {
		return apperror.ConflictError("email", "email already taken")
	}
	return nil
}

func (s *UserSecurityService) ValidatePhoneNumber(ctx context.Context, phoneNumber *valueobject.PhoneNumber) error {
	if phoneNumber == nil {
		return nil
	}

	exists, err := s.userRepo.ExistsByPhone(ctx, phoneNumber.String())
	if err != nil {
		return err
	}

	if exists {
		return apperror.ConflictError("phone  number", "phone number already taken")
	}
	return nil
}

func (s *UserSecurityService) ValidateRawPassword(rawPassword string) error {
	if rawPassword == "" {
		return errors.New("password cannot be empty")
	}

	if len(rawPassword) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	regex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	if !regex.MatchString(rawPassword) {
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	return nil
}

func (s *UserSecurityService) ProduceUserCreatedEvent(user user.User) error {
	if user.IsEmployee() {
		event := &event.CreateUserEmployeeEvent{
			UserID:     user.ID(),
			Email:      user.Email(),
			Role:       user.Role(),
			EmployeeID: *user.EmployeeID(),
		}
		return s.eventPublisher.Publish(event)
	} else {
		event := &event.CreateUserCustomerEvent{
			UserID: user.ID(),
			Email:  user.Email(),
			Role:   user.Role(),
		}
		return s.eventPublisher.Publish(event)
	}
}
