package employee

import (
	"errors"
	"fmt"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

const (
	MinLicenseLength   = 8
	MaxLicenseLength   = 20
	MaxExperienceYears = 60
)

func (v *Employee) UpdatePhoto(newPhoto string) error {
	if newPhoto != "" && len(newPhoto) > 500 {
		return errors.New("photo URL too long")
	}

	v.photo = newPhoto
	v.IncrementVersion()
	return nil
}

func (v *Employee) UpdateName(newName valueobject.PersonName) error {
	if err := v.Person.UpdateName(newName); err != nil {
		return err
	}
	v.IncrementVersion()
	return nil
}

func (v *Employee) UpdateLicenseNumber(newLicense string) error {
	if err := validateLicenseNumber(newLicense); err != nil {
		return err
	}
	v.licenseNumber = newLicense
	v.IncrementVersion()
	return nil
}

func (v *Employee) UpdateSpecialty(newSpecialty enum.VetSpecialty) error {
	if !newSpecialty.IsValid() {
		return errors.New("invalid specialty")
	}
	v.specialty = newSpecialty
	v.IncrementVersion()
	return nil
}

func (v *Employee) UpdateYearsExperience(newYears int) error {
	if err := validateYearsExperience(newYears); err != nil {
		return err
	}
	v.yearsExperience = newYears
	v.IncrementVersion()
	return nil
}

func (v *Employee) UpdateConsultationFee(newFee *valueobject.Money) error {
	if newFee != nil && newFee.Amount() < 0 {
		return errors.New("consultation fee cannot be negative")
	}
	v.consultationFee = newFee
	v.IncrementVersion()
	return nil
}

func (v *Employee) Activate() error {
	if v.isActive {
		return nil // Already active
	}
	v.isActive = true
	v.IncrementVersion()
	return nil
}

func (v *Employee) Deactivate() error {
	if !v.isActive {
		return nil // Already inactive
	}
	v.isActive = false
	v.IncrementVersion()
	return nil
}

func (v *Employee) UpdateSchedule(newSchedule *valueobject.Schedule) error {
	if newSchedule != nil {
		if err := newSchedule.ValidateBuissnessLogic(); err != nil {
			return fmt.Errorf("invalid schedule: %w", err)
		}
	}
	v.schedule = newSchedule
	v.IncrementVersion()
	return nil
}

func (v *Employee) AssociateWithUser(userID valueobject.UserID) error {
	if v.userID != nil && v.userID.Value() == userID.Value() {
		return nil // Already associated
	}
	v.userID = &userID
	v.IncrementVersion()
	return nil
}

func (v *Employee) RemoveUserAssociation() error {
	if v.userID == nil {
		return nil // Already not associated
	}
	v.userID = nil
	v.IncrementVersion()
	return nil
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

func (v *Employee) validate() error {
	if err := validateLicenseNumber(v.licenseNumber); err != nil {
		return err
	}
	if err := validateYearsExperience(v.yearsExperience); err != nil {
		return err
	}

	if v.schedule != nil {
		if err := v.schedule.ValidateBuissnessLogic(); err != nil {
			return fmt.Errorf("invalid schedule: %w", err)
		}
	}

	if !v.specialty.IsValid() {
		return errors.New("specialty is required")
	}
	return nil
}
