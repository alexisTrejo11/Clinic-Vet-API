// Package event contains domain events for the Clinic Vet API application.
package event

import (
	"time"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type CreateUserEmployeeEvent struct {
	UserID     valueobject.UserID     `json:"user_id"`
	Email      valueobject.Email      `json:"email"`
	Role       enum.UserRole          `json:"role"`
	EmployeeID valueobject.EmployeeID `json:"employee_id"`
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
