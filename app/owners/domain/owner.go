package ownerDomain

import (
	"errors"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

type Owner struct {
	id          int
	photo       string
	fullName    valueObjects.PersonName
	gender      valueObjects.Gender
	dateOfBirth time.Time
	phoneNumber string
	address     *string
	userId      *int
	isActive    bool
	pets        []petDomain.Pet
}

func NewOwner(id int, photo string, fullName valueObjects.PersonName, gender valueObjects.Gender, dateOfBirth time.Time, phoneNumber string) (*Owner, error) {
	if id <= 0 {
		return nil, errors.New("owner ID must be a positive integer")
	}

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
func (o *Owner) Id() int {
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

func (o *Owner) UserId() *int {
	return o.userId
}

func (o *Owner) IsActive() bool {
	return o.isActive
}

func (o *Owner) Pets() []petDomain.Pet {
	return o.pets
}

// --- Setters ---
func (o *Owner) SetId(id int) {
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

func (o *Owner) AddPet(pet petDomain.Pet) {
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

func (o *Owner) SetUserId(userId int) {
	o.userId = &userId
}

func (o *Owner) SetPets(pets []petDomain.Pet) {
	o.pets = pets
}

func (o *Owner) GetId() int {
	return o.id
}

func (o *Owner) SetGender(gender valueObjects.Gender) {
	o.gender = gender
}
