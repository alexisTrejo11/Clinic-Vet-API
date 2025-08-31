package entity

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

type Owner struct {
	id          valueobject.OwnerID
	photo       string
	fullName    valueObjects.PersonName
	gender      valueObjects.Gender
	dateOfBirth time.Time
	phoneNumber string
	address     *string
	userID      *valueobject.UserID
	isActive    bool
	pets        []Pet
}

func NewOwner(
	id valueobject.OwnerID,
	photo string,
	fullName valueObjects.PersonName,
	gender valueObjects.Gender,
	dateOfBirth time.Time,
	phoneNumber string,
) (*Owner, error) {
	return &Owner{
		id:          id,
		photo:       photo,
		fullName:    fullName,
		gender:      gender,
		dateOfBirth: dateOfBirth,
		phoneNumber: phoneNumber,
		isActive:    true,
	}, nil
}

// --- Getters ---
func (o *Owner) ID() valueobject.OwnerID {
	return o.id
}

func (o *Owner) Photo() string {
	return o.photo
}

func (o *Owner) FullName() valueObjects.PersonName {
	return o.fullName
}

func (o *Owner) Gender() valueObjects.Gender {
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

func (o *Owner) Pets() []Pet {
	return o.pets
}

// --- Setters ---
func (o *Owner) SetID(id valueobject.OwnerID) {
	o.id = id
}

func (o *Owner) SetPhoto(photo string) {
	o.photo = photo
}

func (o *Owner) SetPhoneNumber(phoneNumber string) {
	o.phoneNumber = phoneNumber
}

func (o *Owner) SetAddress(address string) {
	o.address = &address
}

func (o *Owner) AddPet(pet Pet) {
	o.pets = append(o.pets, pet)
}

func (o *Owner) SetActive(status bool) {
	o.isActive = status
}

func (o *Owner) SetFullName(fullName valueObjects.PersonName) {
	o.fullName = fullName
}

func (o *Owner) SetDateOfBirth(dateOfBirth time.Time) {
	o.dateOfBirth = dateOfBirth
}

func (o *Owner) SetUserID(userID valueobject.UserID) {
	o.userID = &userID
}

func (o *Owner) SetPets(pets []Pet) {
	o.pets = pets
}

func (o *Owner) GetID() valueobject.OwnerID {
	return o.id
}

func (o *Owner) SetGender(gender valueObjects.Gender) {
	o.gender = gender
}

type OwnerBuilder struct {
	owner *Owner
}

func NewOwnerBuilder() *OwnerBuilder {
	owner := &Owner{
		pets: []Pet{},
	}
	return &OwnerBuilder{owner: owner}
}

func (b *OwnerBuilder) WithID(id valueobject.OwnerID) *OwnerBuilder {
	b.owner.id = id
	return b
}

func (b *OwnerBuilder) WithFullName(fullName valueObjects.PersonName) *OwnerBuilder {
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

func (b *OwnerBuilder) WithGender(gender valueObjects.Gender) *OwnerBuilder {
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

func (b *OwnerBuilder) WithPets(pets []Pet) *OwnerBuilder {
	b.owner.pets = pets
	return b
}

func (b *OwnerBuilder) WithUserID(userID valueobject.UserID) *OwnerBuilder {
	b.owner.userID = &userID
	return b
}

func (b *OwnerBuilder) WithIsActive(isActive bool) *OwnerBuilder {
	b.owner.isActive = isActive
	return b
}

func (b *OwnerBuilder) Build() *Owner {
	return b.owner
}
