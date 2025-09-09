package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/application/command"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type UserCredentials struct {
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,e164"`
	Password    string `json:"password" binding:"required,min=8"`
}

type EmployeeRequestSingup struct {
	UserCredentials
	EmployeeID uint `json:"employee_id" binding:"required"`
}

type CustomerRequestSingup struct {
	UserCredentials
	FirstName   string    `json:"first_name" binding:"required,min=2,max=50"`
	LastName    string    `json:"last_name" binding:"required,min=2,max=50"`
	Gender      string    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	Location    string    `json:"location"`
}

func (r *CustomerRequestSingup) ToCommand() (command.SignupCommand, error) {
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

	phoneVo, err := valueobject.NewPhoneNumber(r.PhoneNumber)
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
		Gender:      gender,
		DateOfBirth: r.DateOfBirth,
	}
	return *cmd, nil
}
