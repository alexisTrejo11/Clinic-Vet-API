package ownerDomain

import (
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type OwnerBuilder struct {
	owner *Owner
}

func NewOwnerBuilder(id int, fullName user.PersonName, phoneNumber string) *OwnerBuilder {
	owner := &Owner{
		id:          id,
		fullName:    fullName,
		phoneNumber: phoneNumber,
		isActive:    true,
		pets:        []petDomain.Pet{},
	}
	return &OwnerBuilder{owner: owner}
}

func (b *OwnerBuilder) WithPhoto(photo string) *OwnerBuilder {
	b.owner.photo = photo
	return b
}

func (b *OwnerBuilder) WithGender(gender user.Gender) *OwnerBuilder {
	b.owner.gender = gender
	return b
}

func (b *OwnerBuilder) WithDateOfBirth(dob time.Time) *OwnerBuilder {
	b.owner.dateOfBirth = dob
	return b
}

func (b *OwnerBuilder) WithAddress(address string) *OwnerBuilder {
	b.owner.address = &address
	return b
}

func (b *OwnerBuilder) WithPets(pets []petDomain.Pet) *OwnerBuilder {
	b.owner.pets = pets
	return b
}

func (b *OwnerBuilder) Build() *Owner {
	return b.owner
}
