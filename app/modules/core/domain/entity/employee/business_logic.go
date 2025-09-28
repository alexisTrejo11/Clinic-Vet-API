package employee

import (
	"context"
	"errors"
	"fmt"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	domainErr "clinic-vet-api/app/modules/core/error"
	domainerr "clinic-vet-api/app/modules/core/error"
)

const (
	MinLicenseLength   = 8
	MaxLicenseLength   = 20
	MaxExperienceYears = 60
)

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

func validateYearsExperience(years int32) error {
	if years < 0 {
		return errors.New("years of experience cannot be negative")
	}
	if years > MaxExperienceYears {
		return fmt.Errorf("years of experience cannot exceed %d", MaxExperienceYears)
	}
	return nil
}

func (v *Employee) validate(ctx context.Context) error {
	if err := v.Person.Validate(ctx); err != nil {
		return fmt.Errorf("invalid person data: %w", err)
	}

	if err := validateLicenseNumber(v.licenseNumber); err != nil {
		return err
	}
	if err := validateYearsExperience(v.yearsExperience); err != nil {
		return err
	}

	if v.schedule != nil {
		if err := v.schedule.ValidateBuissnessLogic(ctx); err != nil {
			return err
		}
	}

	if !v.specialty.IsValid() {
		return domainerr.InvalidFieldValue(ctx, "specialty", string(v.specialty), "invalid specialty", "validate employee")
	}
	return nil
}

func (v *Employee) AssignUser(ctx context.Context, userID vo.UserID) error {
	if v.userID != nil {
		return domainErr.ConflictError(ctx, "userID", fmt.Sprintf("employee %s is already assigned to a user", v.ID().String()), "assinging user to employee")
	}

	v.userID = &userID
	return nil
}
