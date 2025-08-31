package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type ChangeEmailCommand struct {
	UserID valueobject.UserID `json:"user_id"`
	Email  valueobject.Email  `json:"email"`
	CTX    context.Context    `json:"-"`
}

type ChangePhoneCommand struct {
	UserID valueobject.UserID      `json:"user_id"`
	Phone  valueobject.PhoneNumber `json:"phone"`
	CTX    context.Context         `json:"-"`
}

type ChangeEmailHandler struct {
	userRepository repository.UserRepository
}

type ChangePhoneHandler struct {
	userRepository repository.UserRepository
}

func NewChangePhoneHandler(userRepository repository.UserRepository) ChangePhoneHandler {
	return ChangePhoneHandler{
		userRepository: userRepository,
	}
}

func NewChangeEmailHandler(userRepository repository.UserRepository) ChangeEmailHandler {
	return ChangeEmailHandler{
		userRepository: userRepository,
	}
}

func (h ChangePhoneHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(ChangePhoneCommand)

	user, err := h.userRepository.GetByID(command.CTX, command.UserID)
	if err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	if err := h.validate(command, user); err != nil {
		return shared.FailureResult("failed to change phone", err)
	}

	user.UpdatePhoneNumber(command.Phone)

	if err := h.userRepository.Save(command.CTX, &user); err != nil {
		return shared.FailureResult("failed to update user", err)
	}

	return shared.SuccessResult(user.ID().String(), "phone changed successfully")
}

func (h ChangeEmailHandler) Handle(cmd any) shared.CommandResult {
	command := cmd.(ChangeEmailCommand)

	user, err := h.userRepository.GetByID(command.CTX, command.UserID)
	if err != nil {
		return shared.FailureResult("failed to find user", err)
	}

	if err := h.validate(command, user); err != nil {
		return shared.FailureResult("failed to change email", err)
	}

	user.UpdateEmail(command.Email)

	if err := h.userRepository.Save(command.CTX, &user); err != nil {
		return shared.FailureResult("failed to update user", err)
	}

	return shared.SuccessResult(user.ID().String(), "email changed successfully")
}

func (h ChangeEmailHandler) validate(command ChangeEmailCommand, user entity.User) error {
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

func (h ChangePhoneHandler) validate(command ChangePhoneCommand, user entity.User) error {
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
