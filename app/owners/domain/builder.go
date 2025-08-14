package ownerDomain

import (
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type OwnerBuilder struct {
	owner *Owner
}

func NewOwnerBuilder() *OwnerBuilder {
	owner := &Owner{
		pets: []petDomain.Pet{},
	}
	return &OwnerBuilder{owner: owner}
}

func (b *OwnerBuilder) WithId(id int) *OwnerBuilder {
	b.owner.id = id
	return b
}

func (b *OwnerBuilder) WithFullName(fullName user.PersonName) *OwnerBuilder {
	b.owner.fullName = fullName
	return b
}

func (b *OwnerBuilder) WithPhoneNumber(phoneNumber string) *OwnerBuilder {
	b.owner.phoneNumber = phoneNumber
	return b
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

func (b *OwnerBuilder) WithUserId(userId int) *OwnerBuilder {
	b.owner.userId = &userId
	return b
}

func (b *OwnerBuilder) WithIsActive(isActive bool) *OwnerBuilder {
	b.owner.isActive = isActive
	return b
}

func (b *OwnerBuilder) Build() *Owner {
	return b.owner
}
