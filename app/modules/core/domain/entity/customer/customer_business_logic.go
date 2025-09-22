package customer

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	domainerr "clinic-vet-api/app/modules/core/error"
)

func validateDateOfBirth(ctx context.Context, dob time.Time, operation string) error {
	if dob.IsZero() {
		return DateOfBirthRequiredError(ctx, operation)
	}
	if dob.After(time.Now()) {
		return DateOfBirthFutureError(ctx, dob, operation)
	}

	minAgeDate := time.Now().AddDate(-18, 0, 0)
	if dob.After(minAgeDate) {
		return UnderageCustomerError(ctx, dob, operation)
	}
	return nil
}

func (o *Customer) UpdatePhoto(ctx context.Context, newPhoto string) error {
	const operation = "update_photo"

	if newPhoto != "" && len(newPhoto) > 500 {
		return PhotoURLLongError(ctx, len(newPhoto), operation)
	}

	o.photo = newPhoto
	o.IncrementVersion()
	return nil
}

func (o *Customer) UpdateFullName(ctx context.Context, newName valueobject.PersonName) error {
	const operation = "update_full_name"

	if err := o.UpdateName(ctx, newName); err != nil {
		return domainerr.WrapError(ctx, err, "Failed to update name", "customer", "name", operation)
	}

	o.IncrementVersion()
	return nil
}

func (o *Customer) UpdateGender(ctx context.Context, newGender enum.PersonGender) error {
	const operation = "update_gender"

	if err := o.Person.UpdateGender(ctx, newGender); err != nil {
		return domainerr.WrapError(ctx, err, "Failed to update gender", "customer", "gender", operation)
	}

	o.IncrementVersion()
	return nil
}

func (o *Customer) AssociateWithUser(ctx context.Context, userID valueobject.UserID) error {
	const operation = "associate_user"

	if o.userID != nil && o.userID.Value() == userID.Value() {
		return UserAlreadyAssociatedError(ctx, userID, operation)
	}

	o.userID = &userID
	o.IncrementVersion()
	return nil
}

func (o *Customer) RemoveUserAssociation(ctx context.Context) error {
	const operation = "remove_user_association"

	if o.userID == nil {
		return UserNotAssociatedError(ctx, operation)
	}

	o.userID = nil
	o.IncrementVersion()
	return nil
}

func (o *Customer) Activate(ctx context.Context) error {
	const operation = "activate_customer"

	if o.isActive {
		return nil // Already active
	}

	o.isActive = true
	o.IncrementVersion()
	return nil
}

func (o *Customer) Deactivate(ctx context.Context) error {
	const operation = "deactivate_customer"

	if !o.isActive {
		return nil // Already inactive
	}

	if !o.CanBeDeactivated() {
		activePets := 0
		for _, pet := range o.pets {
			if pet.IsActive() {
				activePets++
			}
		}
		return CannotDeactivateWithPetsError(ctx, activePets, operation)
	}

	o.isActive = false
	o.IncrementVersion()
	return nil
}

func (o *Customer) AddPet(ctx context.Context, newPet *pet.Pet) error {
	const operation = "add_pet"

	if newPet == nil {
		return domainerr.MissingEntity(ctx, "pet", "Pet cannot be nil", operation)
	}

	for _, existingPet := range o.pets {
		if existingPet.ID().Value() == newPet.ID().Value() {
			return PetAlreadyExistsError(ctx, newPet.ID(), operation)
		}
	}

	o.pets = append(o.pets, *newPet)
	o.IncrementVersion()
	return nil
}

func (o *Customer) RemovePet(ctx context.Context, petID valueobject.PetID) error {
	const operation = "remove_pet"

	for i, existingPet := range o.pets {
		if existingPet.ID().Value() == petID.Value() {
			// Remove the pet
			o.pets = append(o.pets[:i], o.pets[i+1:]...)
			o.IncrementVersion()
			return nil
		}
	}

	return PetNotFoundError(ctx, petID, operation)
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
	return !o.HasActivePets()
}
