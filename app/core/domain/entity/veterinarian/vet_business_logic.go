package veterinarian

import (
	"errors"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

const (
	MinLicenseLength   = 8
	MaxLicenseLength   = 20
	MaxExperienceYears = 60
)

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
	if newFee != nil && newFee.Amount() < 0 {
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
	if v.userID != nil && v.userID.Value() == userID.Value() {
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
