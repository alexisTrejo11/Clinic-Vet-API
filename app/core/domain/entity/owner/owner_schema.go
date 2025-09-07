// Package owner defines the Owner entity and its related business logic.
package owner

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type Owner struct {
	base.Entity[valueobject.OwnerID]
	photo       string
	fullName    valueobject.PersonName
	gender      enum.PersonGender
	dateOfBirth time.Time
	phoneNumber string
	userID      *valueobject.UserID
	isActive    bool
	pets        []pet.Pet
}

func (o *Owner) ID() valueobject.OwnerID {
	return o.Entity.ID()
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

func (o *Owner) UserID() *valueobject.UserID {
	return o.userID
}

func (o *Owner) IsActive() bool {
	return o.isActive
}

func (o *Owner) Pets() []pet.Pet {
	return o.pets
}

func (o *Owner) Age() int {
	return calculateAge(o.dateOfBirth)
}
