// Package customer defines the Customer entity and its related business logic.
package customer

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type Customer struct {
	base.Entity[valueobject.CustomerID]
	base.Person
	photo     string
	gender    enum.PersonGender
	birthDate time.Time
	userID    *valueobject.UserID
	isActive  bool
	pets      []pet.Pet
}

func (o *Customer) ID() valueobject.CustomerID {
	return o.Entity.ID()
}

func (o *Customer) Photo() string {
	return o.photo
}

func (o *Customer) FullName() valueobject.PersonName {
	return o.Name()
}

func (o *Customer) DateOfBirth() time.Time {
	return o.Person.DateOfBirth()
}

func (o *Customer) UserID() *valueobject.UserID {
	return o.userID
}

func (o *Customer) IsActive() bool {
	return o.isActive
}

func (o *Customer) Pets() []pet.Pet {
	return o.pets
}
