package valueobject

import (
	"errors"
	"fmt"
)

var (
	ErrNegativeID     = errors.New("ID cannot be negative")
	ErrEntityNotFound = errors.New("entity not supported")
)

type IntegerID interface {
	Value() int
	Equals(number int) bool
	String() string
	IsZero() bool
}

type baseID struct {
	value int
}

func (id baseID) Value() int {
	return id.value
}

func (id baseID) Equals(number int) bool {
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

func NewPetID(value int) (PetID, error) {
	if err := validateID(value); err != nil {
		return PetID{}, err
	}
	return PetID{baseID{value}}, nil
}

func NewPaymentID(value int) (PaymentID, error) {
	if err := validateID(value); err != nil {
		return PaymentID{}, err
	}
	return PaymentID{baseID{value}}, nil
}

func NewVetID(value int) (VetID, error) {
	if err := validateID(value); err != nil {
		return VetID{}, err
	}
	return VetID{baseID{value}}, nil
}

func NewUserID(value int) (UserID, error) {
	if err := validateID(value); err != nil {
		return UserID{}, err
	}
	return UserID{baseID{value}}, nil
}

func NewOwnerID(value int) (OwnerID, error) {
	if err := validateID(value); err != nil {
		return OwnerID{}, err
	}
	return OwnerID{baseID{value}}, nil
}

func NewMedHistoryID(value int) (MedHistoryID, error) {
	if err := validateID(value); err != nil {
		return MedHistoryID{}, err
	}
	return MedHistoryID{baseID{value}}, nil
}

func NewAppointmentID(value int) (AppointmentID, error) {
	if err := validateID(value); err != nil {
		return AppointmentID{}, err
	}
	return AppointmentID{baseID{value}}, nil
}

func NewIDFactory(value int, entity string) (IntegerID, error) {
	if err := validateID(value); err != nil {
		return nil, err
	}
	switch entity {
	case "payment":
		return NewPaymentID(value)
	case "veterinarian", "vet":
		return NewVetID(value)
	case "user":
		return NewUserID(value)
	case "owner":
		return NewOwnerID(value)
	case "medical_history", "medicalhistory":
		return NewMedHistoryID(value)
	case "appointment":
		return NewAppointmentID(value)
	case "pet":
		return NewVetID(value)
	default:
		return nil, fmt.Errorf("%w: %s", ErrEntityNotFound, entity)
	}
}

func validateID(value int) error {
	if value < 0 {
		return fmt.Errorf("%w: %d", ErrNegativeID, value)
	}
	return nil
}

func (id PaymentID) IsPayment() bool         { return true }
func (id VetID) IsVet() bool                 { return true }
func (id UserID) IsUser() bool               { return true }
func (id OwnerID) IsOwner() bool             { return true }
func (id AppointmentID) IsAppointment() bool { return true }
