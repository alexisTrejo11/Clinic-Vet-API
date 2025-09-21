package valueobject

import (
	"errors"
	"fmt"
)

var ErrEntityNotFound = errors.New("entity not supported")

type IntegerID interface {
	Value() uint
	Equals(number uint) bool
	String() string
	IsZero() bool
}

type baseID struct {
	value uint
}

func (id baseID) Value() uint {
	return id.value
}

func (id baseID) Equals(number uint) bool {
	return id.value == number
}

func (id baseID) String() string {
	return fmt.Sprintf("%d", id.value)
}

func (id baseID) IsZero() bool {
	return id.value == 0
}

type (
	PaymentID     struct{ baseID }
	EmployeeID    struct{ baseID }
	PetID         struct{ baseID }
	AppointmentID struct{ baseID }
	UserID        struct{ baseID }
	CustomerID    struct{ baseID }
	MedSessionID  struct{ baseID }
)

func NewPetID(value uint) PetID {
	return PetID{baseID{value}}
}

func NewPaymentID(value uint) PaymentID {
	return PaymentID{baseID{value}}
}

func NewEmployeeID(value uint) EmployeeID {
	return EmployeeID{baseID{value}}
}

func NewUserID(value uint) UserID {
	return UserID{baseID{value}}
}

func NewCustomerID(value uint) CustomerID {
	return CustomerID{baseID{value}}
}

func NewMedSessionID(value uint) MedSessionID {
	return MedSessionID{baseID{value}}
}

func NewAppointmentID(value uint) AppointmentID {
	return AppointmentID{baseID{value}}
}

func NewIDFactory(value uint, entity string) (IntegerID, error) {
	switch entity {
	case "payment":
		return NewPaymentID(value), nil
	case "veterinarian", "employee", "vet":
		return NewEmployeeID(value), nil
	case "user":
		return NewUserID(value), nil
	case "owner", "customer":
		return NewCustomerID(value), nil
	case "medical_sessions", "medicalhistory":
		return NewMedSessionID(value), nil
	case "appointment":
		return NewAppointmentID(value), nil
	case "pet":
		return NewPetID(value), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrEntityNotFound, entity)
	}
}
