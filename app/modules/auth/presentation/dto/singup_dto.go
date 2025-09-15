package dto

import (
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/modules/auth/application/command"
	"context"
	"time"

	apperror "clinic-vet-api/app/shared/error/application"
)

// UserCredentials represents user login credentials
// @Description Basic user credentials for authentication
type UserCredentials struct {
	// User email address
	// Required: true
	// Format: email
	Email string `json:"email" binding:"required,email" example:"user@example.com"`

	// User phone number in E.164 format
	// Required: false
	// Format: e164
	PhoneNumber string `json:"phone_number" binding:"omitempty,e164" example:"+1234567890"`

	// User password
	// Required: true
	// Minimum length: 8 characters
	Password string `json:"password" binding:"required,min=8" example:"SecurePass123!"`
}

// EmployeeRequestSingup represents employee registration request
// @Description Request body for employee registration/signup
type EmployeeRequestRegister struct {
	UserCredentials

	// Employee identification number
	// Required: true
	EmployeeID uint `json:"employee_id" binding:"required" example:"1001"`
}

// CustomerRequestSingup represents customer registration request
// @Description Request body for customer registration/signup
type CustomerRequestSingup struct {
	UserCredentials

	// Customer's first name
	// Required: true
	// Minimum length: 2, Maximum length: 50
	FirstName string `json:"first_name" binding:"required,min=2,max=50" example:"John"`

	// Customer's last name
	// Required: true
	// Minimum length: 2, Maximum length: 50
	LastName string `json:"last_name" binding:"required,min=2,max=50" example:"Doe"`

	// Customer's gender
	// Required: true
	// Enum: male, female, other
	Gender string `json:"gender" binding:"required" example:"male"`

	// Customer's date of birth
	// Required: true
	// Format: date
	DateOfBirth time.Time `json:"date_of_birth" binding:"required" example:"1990-01-15T00:00:00Z"`

	// Customer's location or address
	// Required: false
	Location string `json:"location" example:"123 Main St, City, State"`
}

func (r *EmployeeRequestRegister) ToCommand(ctx context.Context) (command.EmployeeRegisterCommand, error) {
	errorMessages := make([]string, 0)
	emailVo, err := valueobject.NewEmail(r.Email)
	if err != nil {
		errorMessages = append(errorMessages, "email: "+err.Error())
	}

	phoneVo, err := valueobject.NewPhoneNumber(r.PhoneNumber)
	if err != nil {
		errorMessages = append(errorMessages, "phone: "+err.Error())
	}

	if len(errorMessages) > 0 {
		return command.EmployeeRegisterCommand{}, apperror.MappingError(errorMessages, "request", "SignupRequest", "userSingup")
	}

	cmd := command.NewEmployeeRegisterCommand(ctx, emailVo, r.Password, &phoneVo, valueobject.NewEmployeeID(r.EmployeeID))
	return *cmd, nil
}

func (r *CustomerRequestSingup) ToCommand() (command.CustomerRegisterCommand, error) {
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
		return command.CustomerRegisterCommand{}, apperror.MappingError(errorMessages, "request", "SignupRequest", "userSingup")
	}

	cmd := &command.CustomerRegisterCommand{
		Email:       emailVo,
		Password:    r.Password,
		PhoneNumber: &phoneVo,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Gender:      gender,
		DateOfBirth: r.DateOfBirth,
	}
	return *cmd, nil
}
