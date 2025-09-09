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
	VetID         struct{ baseID }
	PetID         struct{ baseID }
	AppointmentID struct{ baseID }
	UserID        struct{ baseID }
	OwnerID       struct{ baseID }
	MedHistoryID  struct{ baseID }
)

func NewPetID(value uint) PetID {
	return PetID{baseID{value}}
}

func NewPaymentID(value uint) PaymentID {
	return PaymentID{baseID{value}}
}

func NewVetID(value uint) VetID {
	return VetID{baseID{value}}
}

func NewUserID(value uint) UserID {
	return UserID{baseID{value}}
}

func NewOwnerID(value uint) OwnerID {
	return OwnerID{baseID{value}}
}

func NewMedHistoryID(value uint) MedHistoryID {
	return MedHistoryID{baseID{value}}
}

func NewAppointmentID(value uint) AppointmentID {
	return AppointmentID{baseID{value}}
}

func NewIDFactory(value uint, entity string) (IntegerID, error) {
	switch entity {
	case "payment":
		return NewPaymentID(value), nil
	case "veterinarian", "vet":
		return NewVetID(value), nil
	case "user":
		return NewUserID(value), nil
	case "owner":
		return NewOwnerID(value), nil
	case "medical_history", "medicalhistory":
		return NewMedHistoryID(value), nil
	case "appointment":
		return NewAppointmentID(value), nil
	case "pet":
		return NewVetID(value), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrEntityNotFound, entity)
	}
}

func (id PaymentID) IsPayment() bool         { return true }
func (id VetID) IsVet() bool                 { return true }
func (id UserID) IsUser() bool               { return true }
func (id OwnerID) IsOwner() bool             { return true }
func (id AppointmentID) IsAppointment() bool { return true }
