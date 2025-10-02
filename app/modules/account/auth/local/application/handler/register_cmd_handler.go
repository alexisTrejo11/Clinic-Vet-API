package handler

import (
	c "clinic-vet-api/app/modules/account/auth/local/application/command"
	e "clinic-vet-api/app/modules/core/domain/entity/employee"
	u "clinic-vet-api/app/modules/core/domain/entity/user"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/auth"
	commondto "clinic-vet-api/app/shared/dto"
	"context"
)

func (h *AuthCommandHandler) HandleRegisterCustomer(ctx context.Context, cmd c.RegisterCustomerCommand) auth.AuthCommandResult {
	if err := h.userAuthService.ValidateUserCredentials(
		ctx, cmd.Email(), cmd.PhoneNumber(), cmd.Password(),
	); err != nil {
		return auth.AuthFailure(ErrValidationFailed, err)
	}

	user := cmd.ToEntity()
	if err := h.userAuthService.ProcessUserPersistence(ctx, user); err != nil {
		return auth.AuthFailure(ErrUserCreationFailed, err)
	}

	go h.produceEvent(user, cmd)

	return auth.AuthSuccess(MsgUserCreatedSuccess)
}

func (h *AuthCommandHandler) HandleRegisterEmployee(ctx context.Context, cmd c.RegisterEmployeeCommand) auth.AuthCommandResult {
	if err := cmd.Validate(); err != nil {
		return auth.AuthFailure(ErrValidationFailed, err)
	}

	if err := h.userAuthService.ValidateUserCredentials(ctx, cmd.Email(), cmd.PhoneNumber(), cmd.Password()); err != nil {
		return auth.AuthFailure(ErrValidationFailed, err)
	}

	employee, err := h.userAuthService.ValidateEmployeeAccount(ctx, cmd.EmployeeID())
	if err != nil {
		return auth.AuthFailure(ErrValidationFailed, err)
	}

	user := cmd.ToEntity()
	if err := h.userAuthService.ProcessUserPersistence(ctx, user); err != nil {
		return auth.AuthFailure(ErrUserCreationFailed, err)
	}

	go h.produceRegisterEvent(user, employee)

	return auth.AuthSuccess("user successfully created")
}

func (h *AuthCommandHandler) produceRegisterEvent(user *u.User, employee *e.Employee) {
	event := &event.UserRegisteredEvent{
		UserID:   user.ID(),
		Email:    user.Email(),
		Role:     user.Role(),
		Employee: employee,
	}
	h.userEvent.Registered(*event)
}

func (h *AuthCommandHandler) produceEvent(user *u.User, cmd c.RegisterCustomerCommand) {
	name, _ := valueobject.NewPersonName(cmd.Name().FirstName, cmd.Name().LastName)
	event := event.UserRegisteredEvent{
		UserID: user.ID(),
		Role:   user.Role(),
		Email:  user.Email(),
		PersonalData: commondto.PersonalData{
			Name:        name,
			Gender:      cmd.Gender(),
			DateOfBirth: cmd.DateOfBirth(),
		},
	}

	h.userEvent.Registered(event)
}
