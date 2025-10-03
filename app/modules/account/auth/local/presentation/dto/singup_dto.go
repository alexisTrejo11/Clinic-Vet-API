package dto

import (
	"fmt"
	"time"

	registerCmd "clinic-vet-api/app/modules/account/auth/local/application/command"
	"clinic-vet-api/app/modules/core/domain/valueobject"

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
	Password string `json:"password" binding:"required,min=6" example:"SecurePass123!"`
}

// EmployeeRequestSingup represents employee registration request
// @Description Request body for employee registration/signup
type EmployeeRequestRegister struct {
	UserCredentials

	// Employee's role in the system
	// Required: true
	// Minimum length: 2, Maximum length: 100
	Role string `json:"role" binding:"min=2,max=100" example:"vet"`

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
}

func (r *EmployeeRequestRegister) ToCommand() (registerCmd.RegisterEmployeeCommand, error) {
	errorMessages := make([]string, 0)
	emailVo, err := valueobject.NewEmail(r.Email)
	if err != nil {
		errorMessages = append(errorMessages, "email: "+err.Error())
	}

	var phone *valueobject.PhoneNumber
	if r.PhoneNumber != "" {
		phoneVo, err := valueobject.NewPhoneNumber(r.PhoneNumber)
		if err != nil {
			errorMessages = append(errorMessages, "phone: "+err.Error())
		}
		phone = &phoneVo
	}

	if len(errorMessages) > 0 {
		return registerCmd.RegisterEmployeeCommand{}, apperror.MappingError(errorMessages, "request", "SignupRequest", "userSingup")
	}

	cmd := registerCmd.NewRegisterEmployeeCommand(
		emailVo,
		r.Password,
		phone,
		valueobject.NewEmployeeID(r.EmployeeID),
		r.Role,
	)
	return cmd, nil
}

func (r *CustomerRequestSingup) ToCommand() (registerCmd.RegisterCustomerCommand, error) {
	fmt.Println("date of birth:", r.DateOfBirth)
	return registerCmd.NewRegisterCustomerCommand(
		r.Email,
		&r.PhoneNumber,
		r.Password,
		r.FirstName,
		r.LastName,
		r.Gender,
		r.DateOfBirth,
	)

}
