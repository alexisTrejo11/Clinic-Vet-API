package authController

import (
	"time"

	authCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/command"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type RequestSignup struct {
	Email       *string   `json:"email"`
	Password    string    `json:"password" binding:"required,min=8"`
	Phone       *string   `json:"phone"`
	FirstName   string    `json:"first_name" binding:"required,min=2,max=50"`
	LastName    string    `json:"last_name" binding:"required,min=2,max=50"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
}

func (r *RequestSignup) ToCommand() *authCmd.SignupCommand {
	gender := user.NewGender(r.Gender)

	return &authCmd.SignupCommand{
		Email:       r.Email,
		Password:    r.Password,
		PhoneNumber: r.Phone,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Address:     r.Address,
		Gender:      gender,
		DateOfBirth: r.DateOfBirth,
	}
}

type RequestLogin struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}

func (r *RequestLogin) ToCommand() *authCmd.LoginCommand {
	return &authCmd.LoginCommand{
		Identifier: r.Identifier,
		Password:   r.Password,
	}
}

type RequestLogout struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r *RequestLogout) ToCommand() *authCmd.LogoutCommand {
	return &authCmd.LogoutCommand{
		RefreshToken: r.RefreshToken,
	}
}

type RefreshSessionRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TwoFactorAuthRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}
