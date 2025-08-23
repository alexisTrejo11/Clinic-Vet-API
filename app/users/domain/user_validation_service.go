package userDomain

import (
	"context"
	"errors"
	"regexp"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

const (
	MIN_PASSWORD_LENGTH = 8
	MAX_PASSWORD_LENGTH = 100
)

type UseSecurityService struct {
	repository UserRepository
}

func NewUseSecurityService(userRepository UserRepository) *UseSecurityService {
	return &UseSecurityService{
		repository: userRepository,
	}
}

func (uc *UseSecurityService) ValidateCreation(ctx context.Context, user User) error {
	if err := uc.ValidatePassword(user.Password()); err != nil {
		return err
	}

	if err := uc.validateUniqueConstraints(ctx, user); err != nil {
		return err
	}
	return nil
}

func (uc *UseSecurityService) ValidatePassword(password string) error {
	if len(password) < MIN_PASSWORD_LENGTH {
		return ErrPasswordTooShort(strconv.Itoa(MIN_PASSWORD_LENGTH))
	}
	if len(password) > MAX_PASSWORD_LENGTH {
		return ErrPasswordTooLong(strconv.Itoa(MAX_PASSWORD_LENGTH))
	}

	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+={}\[\]:;"'<>,.?/\\|-]+$`).MatchString(password)
	if !passwordRegex {
		return ErrPasswordInvalidFormat("Password must contain only alphanumeric characters and special symbols !@#$%^&*()_+={}[]:;\"'<>,.?/\\\\|-")
	}
	return nil
}

func (uc *UseSecurityService) validateUniqueConstraints(ctx context.Context, user User) error {
	exists, err := uc.repository.ExistsByEmail(ctx, user.Email().String())
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	exists, err = uc.repository.ExistsByPhone(ctx, user.PhoneNumber().String())
	if err != nil {
		return err
	}

	if exists {
		return errors.New("phone number already exists")
	}

	return nil
}

func (uc *UseSecurityService) HashPassword(user *User) error {
	hashedPw, err := shared.HashPassword(user.Password())
	if err != nil {
		return err
	}

	user.SetPassword(hashedPw)
	return nil
}
