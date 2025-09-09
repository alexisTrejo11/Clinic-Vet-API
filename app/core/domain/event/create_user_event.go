// Package event contains domain events for the Clinic Vet API application.
package event

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type CreateUserEmployeeEvent struct {
	UserID     valueobject.UserID `json:"user_id"`
	Email      valueobject.Email  `json:"email"`
	Role       enum.UserRole      `json:"role"`
	EmployeeID valueobject.VetID  `json:"employee_id"`
}

type CreateUserCustomerEvent struct {
	UserID valueobject.UserID `json:"user_id"`
	Email  valueobject.Email  `json:"email"`
	Role   enum.UserRole      `json:"role"`

	// Personal Data
	Name      valueobject.PersonName `json:"name"`
	Gender    enum.PersonGender      `json:" gender"`
	Birthdate time.Time              `json:"birthdate"`
	Location  string                 `json:"location"`
}
