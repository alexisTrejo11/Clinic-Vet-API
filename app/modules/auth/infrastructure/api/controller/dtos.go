package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type RequestSignup struct {
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required,min=8"`
	Phone       string    `json:"phone"`
	FirstName   string    `json:"first_name" binding:"required,min=2,max=50"`
	LastName    string    `json:"last_name" binding:"required,min=2,max=50"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
}

func (r *RequestSignup) ToCommand() (command.SignupCommand, error) {
	errorMessages := make([]string, 0)
	gender, err := enum.ParseGender(r.Gender)
	if err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if r.DateOfBirth.After(time.Now()) {
		errorMessages = append(errorMessages, "date_of_birth cannot be in the future")
	}

	emailVo, err := valueobject.NewEmail(r.Email)
	if err != nil {
		errorMessages = append(errorMessages, "email: "+err.Error())
	}

	phoneVo, err := valueobject.NewPhoneNumber(r.Phone)
	if err != nil {
		errorMessages = append(errorMessages, "phone: "+err.Error())
	}

	if len(errorMessages) > 0 {
		return command.SignupCommand{}, apperror.MappingError(errorMessages, "request", "SignupRequest", "userSingup")
	}

	cmd := &command.SignupCommand{
		Email:       emailVo,
		Password:    r.Password,
		PhoneNumber: phoneVo,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Address:     r.Address,
		Gender:      gender,
		DateOfBirth: r.DateOfBirth,
	}
	return *cmd, nil
}

type RequestLogin struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}

func (r *RequestLogin) ToCommand() *command.LoginCommand {
	return &command.LoginCommand{
		Identifier: r.Identifier,
		Password:   r.Password,
	}
}

type RequestLogout struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r *RequestLogout) ToCommand(userIdInt int) (command.LogoutCommand, error) {
	userId, err := valueobject.NewUserID(userIdInt)
	if err != nil {
		return command.LogoutCommand{}, apperror.FieldValidationError("userID", strconv.Itoa(userIdInt), err.Error())
	}

	cmd := &command.LogoutCommand{
		RefreshToken: r.RefreshToken,
		UserID:       userId,
		CTX:          context.Background(),
	}

	return *cmd, nil
}

type RefreshSessionRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TwoFactorAuthRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}
