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

func (id baseID) Int32() int32 {
	return int32(id.value)
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
	VaccinationID struct{ baseID }
	DewormID      struct{ baseID }
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

func NewVaccinationID(value uint) VaccinationID {
	return VaccinationID{baseID{value}}
}

func NewDewormID(value uint) DewormID {
	return DewormID{baseID{value}}
}
