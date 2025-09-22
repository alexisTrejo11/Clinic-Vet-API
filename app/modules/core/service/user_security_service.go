// Package service contains the business logic for entities services.
package service

import (
	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	apperror "clinic-vet-api/app/shared/error/application"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/password"
	"context"
	"errors"
	"sync"
	"unicode"
)

type UserSecurityService struct {
	userRepo        repository.UserRepository
	employeeRepo    repository.EmployeeRepository
	passwordEncoder password.PasswordEncoder
}

func NewUserSecurityService(
	userRepo repository.UserRepository,
	employeeRepo repository.EmployeeRepository,
	passwordEncoder password.PasswordEncoder,
) *UserSecurityService {
	return &UserSecurityService{
		userRepo:        userRepo,
		employeeRepo:    employeeRepo,
		passwordEncoder: passwordEncoder,
	}
}

// ProcessUserPersistence handles hashing the user's password and saving the user to the repository.
func (s *UserSecurityService) ProcessUserPersistence(ctx context.Context, user *u.User) error {
	hashedPassword, err := s.passwordEncoder.HashPassword(user.HashedPassword())
	if err != nil {
		return err
	}
	user.SetHashedPassword(hashedPassword)

	if err := s.userRepo.Save(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserSecurityService) AuthenticateUser(ctx context.Context, emailOrPhone string, plainPassword string) (u.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, emailOrPhone)
	if err == nil {
		return s.authenticateWithPassword(user, plainPassword)
	}

	user, err = s.userRepo.FindByPhone(ctx, emailOrPhone)
	if err == nil {
		return s.authenticateWithPassword(user, plainPassword)
	}

	return u.User{}, autherror.InvalidCredentialsError(err, emailOrPhone)
}

func (s *UserSecurityService) authenticateWithPassword(user u.User, plainPassword string) (u.User, error) {
	valid := s.passwordEncoder.CheckPassword(user.HashedPassword(), plainPassword)
	if !valid {
		return u.User{}, autherror.InvalidCredentialsError(errors.New("invalid password"), "")
	}

	return user, nil
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

func (s *UserSecurityService) ValidateEmployeeAccount(ctx context.Context, employeeID valueobject.EmployeeID) error {
	var wg sync.WaitGroup
	errorChannel := make(chan error, 3)
	var validationErrors []error
	wg.Add(1)
	go func() {
		defer wg.Done()
		exists, err := s.employeeRepo.ExistsByID(ctx, employeeID)
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

	if len(rawPassword) > 72 {
		return errors.New("password must be less than 72 characters")
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range rawPassword {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case !unicode.IsLetter(char) && !unicode.IsDigit(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (s *UserSecurityService) ValidateTwoFactorToken(user u.User, token string) error {
	if !user.IsTwoFactorEnabled() {
		return errors.New("two-factor authentication is not enabled for this user")
	}

	return nil
}

func (s *UserSecurityService) AuthenticateOAuthUser(ctx context.Context, provider string, oauthToken string) (u.User, error) {
	user, err := s.userRepo.FindByOAuthProvider(ctx, provider, oauthToken)
	if err != nil {
		return u.User{}, autherror.InvalidCredentialsError(err, "")
	}

	return user, nil
}
