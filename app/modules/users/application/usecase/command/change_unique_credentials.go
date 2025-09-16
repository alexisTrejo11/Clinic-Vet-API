package command

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type ChangeEmailCommand struct {
	userID valueobject.UserID
	email  valueobject.Email
	ctx    context.Context
}

type ChangePhoneCommand struct {
	userID valueobject.UserID
	phone  valueobject.PhoneNumber
	ctx    context.Context
}

func NewChangeEmailCommand(ctx context.Context, userIDInt uint, emailStr string) (ChangeEmailCommand, error) {
	userID := valueobject.NewUserID(userIDInt)
	email, err := valueobject.NewEmail(emailStr)
	if err != nil {
		return ChangeEmailCommand{}, errors.New(ErrInvalidEmail)
	}

	cmd := &ChangeEmailCommand{
		ctx:    ctx,
		userID: userID,
		email:  email,
	}

	return *cmd, nil
}

func NewChangePhoneCommand(ctx context.Context, userIDInt uint, phoneStr string) (ChangePhoneCommand, error) {
	userID := valueobject.NewUserID(userIDInt)
	phone, err := valueobject.NewPhoneNumber(phoneStr)
	if err != nil {
		return ChangePhoneCommand{}, errors.New(ErrInvalidPhone)
	}

	cmd := &ChangePhoneCommand{
		userID: userID,
		phone:  phone,
		ctx:    ctx,
	}

	return *cmd, nil
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

func (h ChangePhoneHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(ChangePhoneCommand)
	if !ok {
		return *cqrs.FailureResult(ErrFailedChangePhone, errors.New("invalid command type"))
	}

	user, err := h.userRepository.FindByID(command.ctx, command.userID)
	if err != nil {
		return *cqrs.FailureResult(ErrFailedFindUser, err)
	}

	if err := h.validate(command, user); err != nil {
		return *cqrs.FailureResult(ErrFailedChangePhone, err)
	}

	user.UpdatePhoneNumber(command.phone)

	if err := h.userRepository.Save(command.ctx, &user); err != nil {
		return *cqrs.FailureResult(ErrFailedUpdateUser, err)
	}

	return *cqrs.SuccessResult(user.ID().String(), "phone changed successfully")
}

func (h ChangeEmailHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(ChangeEmailCommand)
	if !ok {
		return *cqrs.FailureResult(ErrFailedChangeEmail, errors.New("invalid command type"))
	}

	user, err := h.userRepository.FindByID(command.ctx, command.userID)
	if err != nil {
		return *cqrs.FailureResult(ErrFailedFindUser, err)
	}

	if err := h.validate(command, user); err != nil {
		return *cqrs.FailureResult(ErrFailedChangeEmail, err)
	}

	user.UpdateEmail(command.email)

	if err := h.userRepository.Save(command.ctx, &user); err != nil {
		return *cqrs.FailureResult(ErrFailedUpdateUser, err)
	}

	return *cqrs.SuccessResult(user.ID().String(), "email changed successfully")
}

func (h ChangeEmailHandler) validate(command ChangeEmailCommand, user user.User) error {
	if user.Email().String() == command.email.String() {
		return errors.New(ErrEmailUnchanged)
	}

	if exists, err := h.userRepository.ExistsByEmail(command.ctx, command.email.String()); err != nil {
		return err
	} else if exists {
		return errors.New(ErrEmailAlreadyExists)
	}

	return nil
}

func (h ChangePhoneHandler) validate(command ChangePhoneCommand, user user.User) error {
	if user.PhoneNumber().String() == command.phone.String() {
		return errors.New(ErrPhoneUnchanged)
	}

	if exists, err := h.userRepository.ExistsByPhone(command.ctx, command.phone.String()); err != nil {
		return err
	} else if exists {
		return errors.New(ErrPhoneAlreadyTaken)
	}

	return nil
}
