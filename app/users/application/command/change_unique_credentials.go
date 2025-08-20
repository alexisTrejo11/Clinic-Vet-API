package userCommand

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type ChangeEmailCommand struct {
	UserId int                `json:"user_id"`
	Email  valueObjects.Email `json:"email"`
	CTX    context.Context    `json:"-"`
}

type ChangePhoneCommand struct {
	UserId int                      `json:"user_id"`
	Phone  valueObjects.PhoneNumber `json:"phone"`
	CTX    context.Context          `json:"-"`
}

type ChangeEmailHandler struct {
	userRepository userRepository.UserRepository
}

type ChangePhoneHandler struct {
	userRepository userRepository.UserRepository
}

func NewChangePhoneHandler(userRepository userRepository.UserRepository) ChangePhoneHandler {
	return ChangePhoneHandler{
		userRepository: userRepository,
	}
}

func NewChangeEmailHandler(userRepository userRepository.UserRepository) ChangeEmailHandler {
	return ChangeEmailHandler{
		userRepository: userRepository,
	}
}

func (h ChangePhoneHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(ChangePhoneCommand)

	user, err := h.userRepository.GetByIdWithProfile(command.CTX, command.UserId)
	if err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	if err := h.validate(command, *user); err != nil {
		return shared.FailureResult("failed to change phone", err)
	}

	user.UpdatePhoneNumber(command.Phone)

	if err := h.userRepository.Save(command.CTX, user); err != nil {
		return shared.FailureResult("failed to update user", err)
	}

	return shared.SuccessResult(user.Id().String(), "phone changed successfully")
}

func (h ChangeEmailHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(ChangeEmailCommand)

	user, err := h.userRepository.GetByIdWithProfile(command.CTX, command.UserId)
	if err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	if err := h.validate(command, *user); err != nil {
		return shared.FailureResult("failed to change email", err)
	}

	user.UpdateEmail(command.Email)

	if err := h.userRepository.Save(command.CTX, user); err != nil {
		return shared.FailureResult("failed to update user", err)
	}

	return shared.SuccessResult(user.Id().String(), "email changed successfully")
}

func (h ChangeEmailHandler) validate(command ChangeEmailCommand, user user.User) error {
	if user.Email().String() == command.Email.String() {
		return nil
	}

	if exists, err := h.userRepository.ExistsByEmail(command.CTX, command.Email.String()); err != nil {
		return err
	} else if exists {
		return errors.New("email already exists")
	}

	return nil
}

func (h ChangePhoneHandler) validate(command ChangePhoneCommand, user user.User) error {
	if user.PhoneNumber().String() == command.Phone.String() {
		return nil
	}

	if exists, err := h.userRepository.ExistsByPhone(command.CTX, command.Phone.String()); err != nil {
		return err
	} else if exists {
		return errors.New("phone already taken")
	}

	return nil
}
