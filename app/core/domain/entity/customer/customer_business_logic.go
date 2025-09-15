package customer

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/pet"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	domainerr "clinic-vet-api/app/core/error"
)

func validateDateOfBirth(dob time.Time) error {
	if dob.IsZero() {
		return domainerr.NewValidationError("customer", "date-of-birth", "date of birth is required")
	}
	if dob.After(time.Now()) {
		return domainerr.NewValidationError("customer", "date-of-birth", "date of birth cannot be in the future")
	}
	// Check if customer is at least 18 years old
	minAgeDate := time.Now().AddDate(-18, 0, 0)
	if dob.After(minAgeDate) {
		return domainerr.NewValidationError("customer", "date-of-birth", "owner must be at least 18 years old")
	}
	return nil
}

func (o *Customer) validate() error {
	if err := validateDateOfBirth(o.Person.DateOfBirth()); err != nil {
		return err
	}
	if !o.Person.Gender().IsValid() {
		return domainerr.NewValidationError("customer", "gender", "gender is required")
	}
	return nil
}

func (o *Customer) UpdatePhoto(newPhoto string) error {
	if newPhoto != "" && len(newPhoto) > 500 {
		return domainerr.NewValidationError("customer", "photo", "photo URL too long")
	}
	o.photo = newPhoto
	o.IncrementVersion()
	return nil
}

func (o *Customer) UpdateFullName(newName valueobject.PersonName) error {
	if err := o.UpdateName(newName); err != nil {
		return err
	}
	o.IncrementVersion()
	return nil
}

func (o *Customer) UpdateGender(newGender enum.PersonGender) error {
	if err := o.Person.UpdateGender(newGender); err != nil {
		return err
	}
	o.IncrementVersion()
	return nil
}

func (o *Customer) AssociateWithUser(userID valueobject.UserID) error {
	if o.userID != nil && o.userID.Value() == userID.Value() {
		return nil // Already associated
	}
	o.userID = &userID
	o.IncrementVersion()
	return nil
}

func (o *Customer) RemoveUserAssociation() error {
	if o.userID == nil {
		return nil // Already not associated
	}
	o.userID = nil
	o.IncrementVersion()
	return nil
}

func (o *Customer) Activate() error {
	if o.isActive {
		return nil // Already active
	}
	o.isActive = true
	o.IncrementVersion()
	return nil
}

func (o *Customer) Deactivate() error {
	if !o.isActive {
		return nil // Already inactive
	}
	o.isActive = false
	o.IncrementVersion()
	return nil
}

func (o *Customer) AddPet(newPet *pet.Pet) error {
	if newPet == nil {
		return domainerr.NewValidationError("customer", "pet", "pet cannot be nil")
	}

	for _, existingPet := range o.pets {
		if existingPet.ID().Value() == newPet.ID().Value() {
			return domainerr.NewBusinessRuleError("customer", "add pet", "pet already exists")
		}
	}

	o.pets = append(o.pets, *newPet)
	o.IncrementVersion()
	return nil
}

func (o *Customer) RemovePet(petID valueobject.PetID) error {
	for i, existingPet := range o.pets {
		if existingPet.ID().Value() == petID.Value() {
			// Remove the pet
			o.pets = append(o.pets[:i], o.pets[i+1:]...)
			o.IncrementVersion()
			return nil
		}
	}
	return domainerr.NewBusinessRuleError("customer", "remove Pet", "pet not found")
}

func (o *Customer) HasActivePets() bool {
	for _, pet := range o.pets {
		if pet.IsActive() {
			return true
		}
	}
	return false
}

func (o *Customer) CanBeDeactivated() bool {
	// Cannot  customer with active pets
	return !o.HasActivePets()
}
