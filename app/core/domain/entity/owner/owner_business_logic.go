package owner

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

func validateDateOfBirth(dob time.Time) error {
	if dob.IsZero() {
		return domainerr.NewValidationError("owner", "date-of-birth", "date of birth is required")
	}
	if dob.After(time.Now()) {
		return domainerr.NewValidationError("owner", "date-of-birth", "date of birth cannot be in the future")
	}
	// Check if owner is at least 18 years old
	minAgeDate := time.Now().AddDate(-18, 0, 0)
	if dob.After(minAgeDate) {
		return domainerr.NewValidationError("owner", "date-of-birth", "owner must be at least 18 years old")
	}
	return nil
}

func validatePhoneNumber(phoneNumber string) error {
	if phoneNumber == "" {
		return domainerr.NewValidationError("owner", "phone-number", "phone number is required")
	}
	if len(phoneNumber) < 10 {
		return domainerr.NewValidationError("owner", "phone-number", "phone number too short")
	}
	if len(phoneNumber) > 20 {
		return domainerr.NewValidationError("owner", "phone-number", "phone number too long")
	}
	return nil
}

func (o *Owner) validate() error {
	if err := validatePhoneNumber(o.phoneNumber); err != nil {
		return err
	}
	if err := validateDateOfBirth(o.Person.DateOfBirth()); err != nil {
		return err
	}
	if !o.Person.Gender().IsValid() {
		return domainerr.NewValidationError("owner", "gender", "gender is required")
	}
	return nil
}

func (o *Owner) UpdatePhoto(newPhoto string) error {
	if newPhoto != "" && len(newPhoto) > 500 {
		return domainerr.NewValidationError("owner", "photo", "photo URL too long")
	}
	o.photo = newPhoto
	o.IncrementVersion()
	return nil
}

func (o *Owner) UpdateFullName(newName valueobject.PersonName) error {
	if err := o.UpdateName(newName); err != nil {
		return err
	}
	o.IncrementVersion()
	return nil
}

func (o *Owner) UpdateGender(newGender enum.PersonGender) error {
	if err := o.Person.UpdateGender(newGender); err != nil {
		return err
	}
	o.IncrementVersion()
	return nil
}

func (o *Owner) UpdatePhoneNumber(newPhone string) error {
	if err := validatePhoneNumber(newPhone); err != nil {
		return err
	}
	o.phoneNumber = newPhone
	o.IncrementVersion()
	return nil
}

func (o *Owner) AssociateWithUser(userID valueobject.UserID) error {
	if o.userID != nil && o.userID.Value() == userID.Value() {
		return nil // Already associated
	}
	o.userID = &userID
	o.IncrementVersion()
	return nil
}

func (o *Owner) RemoveUserAssociation() error {
	if o.userID == nil {
		return nil // Already not associated
	}
	o.userID = nil
	o.IncrementVersion()
	return nil
}

func (o *Owner) Activate() error {
	if o.isActive {
		return nil // Already active
	}
	o.isActive = true
	o.IncrementVersion()
	return nil
}

func (o *Owner) Deactivate() error {
	if !o.isActive {
		return nil // Already inactive
	}
	o.isActive = false
	o.IncrementVersion()
	return nil
}

func (o *Owner) AddPet(newPet *pet.Pet) error {
	if newPet == nil {
		return domainerr.NewValidationError("owner", "pet", "pet cannot be nil")
	}

	for _, existingPet := range o.pets {
		if existingPet.ID().Value() == newPet.ID().Value() {
			return domainerr.NewBusinessRuleError("owner", "add pet", "pet already exists")
		}
	}

	o.pets = append(o.pets, *newPet)
	o.IncrementVersion()
	return nil
}

func (o *Owner) RemovePet(petID valueobject.PetID) error {
	for i, existingPet := range o.pets {
		if existingPet.ID().Value() == petID.Value() {
			// Remove the pet
			o.pets = append(o.pets[:i], o.pets[i+1:]...)
			o.IncrementVersion()
			return nil
		}
	}
	return domainerr.NewBusinessRuleError("owner", "remove Pet", "pet not found")
}

func (o *Owner) HasActivePets() bool {
	for _, pet := range o.pets {
		if pet.IsActive() {
			return true
		}
	}
	return false
}

func (o *Owner) CanBeDeactivated() bool {
	// Cannot  owner with active pets
	return !o.HasActivePets()
}
