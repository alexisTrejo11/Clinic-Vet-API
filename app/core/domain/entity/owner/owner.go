package owner

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

type Owner struct {
	base.Entity
	photo       string
	fullName    valueobject.PersonName
	gender      enum.PersonGender
	dateOfBirth time.Time
	phoneNumber string
	address     *string
	userID      *valueobject.UserID
	isActive    bool
	pets        []*pet.Pet
}

// OwnerOption defines the functional option type
type OwnerOption func(*Owner) error

// Functional options
func WithPhoto(photo string) OwnerOption {
	return func(o *Owner) error {
		if photo != "" && len(photo) > 500 {
			return domainerr.NewValidationError("owner", "photo", "photo URL too long")
		}
		o.photo = photo
		return nil
	}
}

func WithFullName(fullName valueobject.PersonName) OwnerOption {
	return func(o *Owner) error {
		if fullName.FullName() == "" {
			return domainerr.NewValidationError("owner", "fullName", "full name is required")
		}
		o.fullName = fullName
		return nil
	}
}

func WithGender(gender enum.PersonGender) OwnerOption {
	return func(o *Owner) error {
		if !gender.IsValid() {
			return domainerr.NewValidationError("owner", "gender", "invalid gender")
		}
		o.gender = gender
		return nil
	}
}

func WithDateOfBirth(dob time.Time) OwnerOption {
	return func(o *Owner) error {
		if err := validateDateOfBirth(dob); err != nil {
			return err
		}
		o.dateOfBirth = dob
		return nil
	}
}

func WithPhoneNumber(phoneNumber string) OwnerOption {
	return func(o *Owner) error {
		if err := validatePhoneNumber(phoneNumber); err != nil {
			return err
		}
		o.phoneNumber = phoneNumber
		return nil
	}
}

func WithAddress(address string) OwnerOption {
	return func(o *Owner) error {
		if address != "" && len(address) > 200 {
			return domainerr.NewValidationError("owner", "address", "address too long")
		}
		o.address = &address
		return nil
	}
}

func WithUserID(userID *valueobject.UserID) OwnerOption {
	return func(o *Owner) error {
		o.userID = userID
		return nil
	}
}

func WithIsActive(isActive bool) OwnerOption {
	return func(o *Owner) error {
		o.isActive = isActive
		return nil
	}
}

func WithPets(pets []*pet.Pet) OwnerOption {
	return func(o *Owner) error {
		o.pets = pets
		return nil
	}
}

// NewOwner creates a new Owner with functional options
func NewOwner(
	id valueobject.OwnerID,
	opts ...OwnerOption,
) (*Owner, error) {
	owner := &Owner{
		Entity:   base.NewEntity(id),
		isActive: true, // Default to active
		pets:     []*pet.Pet{},
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(owner); err != nil {
			return nil, err
		}
	}

	// Final validation
	if err := owner.validate(); err != nil {
		return nil, err
	}

	return owner, nil
}

// Validation helpers
func validateDateOfBirth(dob time.Time) error {
	if dob.IsZero() {
		return domainerr.NewValidationError("owner", "dateOfBirth", "date of birth is required")
	}
	if dob.After(time.Now()) {
		return domainerr.NewValidationError("owner", "dateOfBirth", "date of birth cannot be in the future")
	}
	// Check if owner is at least 18 years old
	minAgeDate := time.Now().AddDate(-18, 0, 0)
	if dob.After(minAgeDate) {
		return domainerr.NewValidationError("owner", "dateOfBirth", "owner must be at least 18 years old")
	}
	return nil
}

func validatePhoneNumber(phoneNumber string) error {
	if phoneNumber == "" {
		return domainerr.NewValidationError("owner", "phoneNumber", "phone number is required")
	}
	if len(phoneNumber) < 10 {
		return domainerr.NewValidationError("owner", "phoneNumber", "phone number too short")
	}
	if len(phoneNumber) > 20 {
		return domainerr.NewValidationError("owner", "phoneNumber", "phone number too long")
	}
	return nil
}

func (o *Owner) validate() error {
	if err := validatePhoneNumber(o.phoneNumber); err != nil {
		return err
	}
	if err := validateDateOfBirth(o.dateOfBirth); err != nil {
		return err
	}
	if !o.gender.IsValid() {
		return domainerr.NewValidationError("owner", "gender", "gender is required")
	}
	return nil
}

// Getters
func (o *Owner) ID() valueobject.OwnerID {
	return o.ID()
}

func (o *Owner) Photo() string {
	return o.photo
}

func (o *Owner) FullName() valueobject.PersonName {
	return o.fullName
}

func (o *Owner) Gender() enum.PersonGender {
	return o.gender
}

func (o *Owner) DateOfBirth() time.Time {
	return o.dateOfBirth
}

func (o *Owner) PhoneNumber() string {
	return o.phoneNumber
}

func (o *Owner) Address() *string {
	return o.address
}

func (o *Owner) UserID() *valueobject.UserID {
	return o.userID
}

func (o *Owner) IsActive() bool {
	return o.isActive
}

func (o *Owner) Pets() []*pet.Pet {
	return o.pets
}

// Business logic methods
func (o *Owner) UpdatePhoto(newPhoto string) error {
	if newPhoto != "" && len(newPhoto) > 500 {
		return domainerr.NewValidationError("owner", "photo", "photo URL too long")
	}
	o.photo = newPhoto
	o.IncrementVersion()
	return nil
}

func (o *Owner) UpdateFullName(newName valueobject.PersonName) error {
	o.fullName = newName
	o.IncrementVersion()
	return nil
}

func (o *Owner) UpdateGender(newGender enum.PersonGender) error {
	if !newGender.IsValid() {
		return domainerr.NewValidationError("owner", "gender", "invalid gender")
	}
	o.gender = newGender
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

func (o *Owner) UpdateAddress(newAddress string) error {
	if newAddress != "" && len(newAddress) > 200 {
		return domainerr.NewValidationError("owner", "address", "address too long")
	}
	o.address = &newAddress
	o.IncrementVersion()
	return nil
}

func (o *Owner) AssociateWithUser(userID valueobject.UserID) error {
	if o.userID != nil && o.userID.GetValue() == userID.GetValue() {
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
		if existingPet.ID().GetValue() == newPet.ID().GetValue() {
			return domainerr.NewBusinessRuleError("owner", "addPet", "pet already exists")
		}
	}

	o.pets = append(o.pets, newPet)
	o.IncrementVersion()
	return nil
}

func (o *Owner) RemovePet(petID valueobject.PetID) error {
	for i, existingPet := range o.pets {
		if existingPet.ID().GetValue() == petID.GetValue() {
			// Remove the pet
			o.pets = append(o.pets[:i], o.pets[i+1:]...)
			o.IncrementVersion()
			return nil
		}
	}
	return domainerr.NewBusinessRuleError("owner", "removePet", "pet not found")
}

func (o *Owner) HasActivePets() bool {
	for _, pet := range o.pets {
		if pet.IsActive() {
			return true
		}
	}
	return false
}

func (o *Owner) Age() int {
	return calculateAge(o.dateOfBirth)
}

func (o *Owner) CanBeDeactivated() bool {
	// Cannot  owner with active pets
	return !o.HasActivePets()
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	years := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		years--
	}

	return years
}
