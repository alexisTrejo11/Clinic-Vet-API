package veterinarian

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type VeterinarianOption func(*Veterinarian) error

func WithName(name valueobject.PersonName) VeterinarianOption {
	return func(v *Veterinarian) error {
		return v.Person.UpdateName(name)
	}
}

func WithPhoto(photo string) VeterinarianOption {
	return func(v *Veterinarian) error {
		if photo != "" && len(photo) > 500 {
			return errors.New("photo URL too long")
		}
		v.photo = photo
		return nil
	}
}

func WithLicenseNumber(licenseNumber string) VeterinarianOption {
	return func(v *Veterinarian) error {
		if err := validateLicenseNumber(licenseNumber); err != nil {
			return err
		}
		v.licenseNumber = licenseNumber
		return nil
	}
}

func WithSpecialty(specialty enum.VetSpecialty) VeterinarianOption {
	return func(v *Veterinarian) error {
		if !specialty.IsValid() {
			return errors.New("invalid specialty")
		}
		v.specialty = specialty
		return nil
	}
}

func WithYearsExperience(years int) VeterinarianOption {
	return func(v *Veterinarian) error {
		if err := validateYearsExperience(years); err != nil {
			return err
		}
		v.yearsExperience = years
		return nil
	}
}

func WithConsultationFee(fee *valueobject.Money) VeterinarianOption {
	return func(v *Veterinarian) error {
		if fee != nil {
			if fee.Amount() < 0 {
				return errors.New("consultation fee cannot be negative")
			}
		}
		v.consultationFee = fee
		return nil
	}
}

func WithIsActive(isActive bool) VeterinarianOption {
	return func(v *Veterinarian) error {
		v.isActive = isActive
		return nil
	}
}

func WithUserID(userID *valueobject.UserID) VeterinarianOption {
	return func(v *Veterinarian) error {
		v.userID = userID
		return nil
	}
}

func WithSchedule(schedule *valueobject.Schedule) VeterinarianOption {
	return func(v *Veterinarian) error {
		v.schedule = schedule
		return nil
	}
}

func WithTimestamps(createdAt, updatedAt time.Time) VeterinarianOption {
	return func(v *Veterinarian) error {
		if createdAt.IsZero() || updatedAt.IsZero() {
			return errors.New("createdAt and updatedAt cannot be zero")
		}
		v.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

func NewVeterinarian(
	id valueobject.VetID,
	opts ...VeterinarianOption,
) (*Veterinarian, error) {
	vet := &Veterinarian{
		Entity:   base.NewEntity(id, time.Now(), time.Now(), 1),
		isActive: true, // Default to active
	}

	for _, opt := range opts {
		if err := opt(vet); err != nil {
			return nil, fmt.Errorf("invalid veterinarian option: %w", err)
		}
	}
	return vet, nil
}

func CreateVeterinarian(
	opts ...VeterinarianOption,
) (*Veterinarian, error) {
	vet := &Veterinarian{
		Entity:   base.NewEntity(valueobject.VetID{}, time.Now(), time.Now(), 1),
		isActive: true, // Default to active
	}

	for _, opt := range opts {
		if err := opt(vet); err != nil {
			return nil, fmt.Errorf("invalid veterinarian option: %w", err)
		}
	}

	if err := vet.validate(); err != nil {
		return nil, err
	}

	return vet, nil
}
