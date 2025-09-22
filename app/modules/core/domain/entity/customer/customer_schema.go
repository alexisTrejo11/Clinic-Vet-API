// Package customer defines the Customer entity and its related business logic.
package customer

import (
	"fmt"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type Customer struct {
	base.Entity[valueobject.CustomerID]
	base.Person
	photo    string
	userID   *valueobject.UserID
	isActive bool
	pets     []pet.Pet
}

func (o *Customer) ID() valueobject.CustomerID {
	return o.Entity.ID()
}

func (o *Customer) Photo() string {
	return o.photo
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

func (o *Customer) AssignUser(userID valueobject.UserID) error {
	if o.userID != nil && !o.userID.IsZero() {
		return fmt.Errorf("customer already assigned to a user with ID %s", o.userID.String())
	}
	o.userID = &userID
	return nil
}

func (o *Customer) SetID(id valueobject.CustomerID) {
	o.Entity.SetID(id)
}

func (o *Customer) CreatedAt() time.Time {
	return o.Entity.CreatedAt()
}

func (o *Customer) UpdatedAt() time.Time {
	return o.Entity.UpdatedAt()
}
