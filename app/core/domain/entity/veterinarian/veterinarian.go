package veterinarian

import (
	"errors"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

const (
	MinLicenseLength   = 8
	MaxLicenseLength   = 20
	MaxExperienceYears = 60
)

type Veterinarian struct {
	base.Entity
	name            valueobject.PersonName
	photo           string
	licenseNumber   string
	specialty       enum.VetSpecialty
	yearsExperience int
	consultationFee *valueobject.Money
	isActive        bool
	userID          *valueobject.UserID
	schedule        *valueobject.Schedule
}

// VeterinarianOption defines the functional option type
type VeterinarianOption func(*Veterinarian) error

func WithName(name valueobject.PersonName) VeterinarianOption {
	return func(v *Veterinarian) error {
		v.name = name
		return nil
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
			if fee.Amount < 0 {
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
		if schedule != nil {
			if err := schedule.Validate(); err != nil {
				return fmt.Errorf("invalid schedule: %w", err)
			}
		}
		v.schedule = schedule
		return nil
	}
}

// NewVeterinarian creates a new Veterinarian with functional options
func NewVeterinarian(
	id valueobject.VetID,
	opts ...VeterinarianOption,
) (*Veterinarian, error) {
	vet := &Veterinarian{
		Entity:   base.NewEntity(id),
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

func validateLicenseNumber(licenseNumber string) error {
	if licenseNumber == "" {
		return errors.New("license number is required")
	}
	if len(licenseNumber) < MinLicenseLength || len(licenseNumber) > MaxLicenseLength {
		return fmt.Errorf("license number must be between %d and %d characters", MinLicenseLength, MaxLicenseLength)
	}
	return nil
}

func validateYearsExperience(years int) error {
	if years < 0 {
		return errors.New("years of experience cannot be negative")
	}
	if years > MaxExperienceYears {
		return fmt.Errorf("years of experience cannot exceed %d", MaxExperienceYears)
	}
	return nil
}

func (v *Veterinarian) validate() error {
	if err := validateLicenseNumber(v.licenseNumber); err != nil {
		return err
	}
	if err := validateYearsExperience(v.yearsExperience); err != nil {
		return err
	}
	if !v.specialty.IsValid() {
		return errors.New("specialty is required")
	}
	return nil
}

func (v *Veterinarian) ID() valueobject.VetID {
	return v.ID()
}

func (v *Veterinarian) Name() valueobject.PersonName {
	return v.name
}

func (v *Veterinarian) Photo() string {
	return v.photo
}

func (v *Veterinarian) LicenseNumber() string {
	return v.licenseNumber
}

func (v *Veterinarian) Specialty() enum.VetSpecialty {
	return v.specialty
}

func (v *Veterinarian) YearsExperience() int {
	return v.yearsExperience
}

func (v *Veterinarian) ConsultationFee() *valueobject.Money {
	return v.consultationFee
}

func (v *Veterinarian) IsActive() bool {
	return v.isActive
}

func (v *Veterinarian) UserID() *valueobject.UserID {
	return v.userID
}

func (v *Veterinarian) Schedule() *valueobject.Schedule {
	return v.schedule
}

// Business logic methods
func (v *Veterinarian) UpdateName(newName valueobject.PersonName) error {
	v.name = newName
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) UpdateLicenseNumber(newLicense string) error {
	if err := validateLicenseNumber(newLicense); err != nil {
		return err
	}
	v.licenseNumber = newLicense
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) UpdateSpecialty(newSpecialty enum.VetSpecialty) error {
	if !newSpecialty.IsValid() {
		return errors.New("invalid specialty")
	}
	v.specialty = newSpecialty
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) UpdateYearsExperience(newYears int) error {
	if err := validateYearsExperience(newYears); err != nil {
		return err
	}
	v.yearsExperience = newYears
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) UpdateConsultationFee(newFee *valueobject.Money) error {
	if newFee != nil && newFee.Amount < 0 {
		return errors.New("consultation fee cannot be negative")
	}
	v.consultationFee = newFee
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) Activate() error {
	if v.isActive {
		return nil // Already active
	}
	v.isActive = true
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) Deactivate() error {
	if !v.isActive {
		return nil // Already inactive
	}
	v.isActive = false
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) UpdateSchedule(newSchedule *valueobject.Schedule) error {
	if newSchedule != nil {
		if err := newSchedule.Validate(); err != nil {
			return fmt.Errorf("invalid schedule: %w", err)
		}
	}
	v.schedule = newSchedule
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) AssociateWithUser(userID valueobject.UserID) error {
	if v.userID != nil && v.userID.GetValue() == userID.GetValue() {
		return nil // Already associated
	}
	v.userID = &userID
	v.IncrementVersion()
	return nil
}

func (v *Veterinarian) RemoveUserAssociation() error {
	if v.userID == nil {
		return nil // Already not associated
	}
	v.userID = nil
	v.IncrementVersion()
	return nil
}
