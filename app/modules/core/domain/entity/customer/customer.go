// Package customer defines the Customer entity and its related business logic.
package customer

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
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

type CustomerBuilder struct{ customer *Customer }

func NewCustomerBuilder() *CustomerBuilder {
	return &CustomerBuilder{customer: &Customer{}}
}

func (cb *CustomerBuilder) WithID(id valueobject.CustomerID) *CustomerBuilder {
	cb.customer.Entity.SetID(id)
	return cb
}

func (cb *CustomerBuilder) WithPhoto(photo string) *CustomerBuilder {
	cb.customer.photo = photo
	return cb
}

func (cb *CustomerBuilder) WithName(fullName valueobject.PersonName) *CustomerBuilder {
	cb.customer.Person.UpdateName(context.Background(), fullName)
	return cb
}

func (cb *CustomerBuilder) WithGender(gender enum.PersonGender) *CustomerBuilder {
	cb.customer.Person.UpdateGender(context.Background(), gender)
	return cb
}

func (cb *CustomerBuilder) WithDateOfBirth(dob time.Time) *CustomerBuilder {
	cb.customer.Person.UpdateDateOfBirth(context.Background(), dob)
	return cb
}

func (cb *CustomerBuilder) WithUserID(userID *valueobject.UserID) *CustomerBuilder {
	cb.customer.userID = userID
	return cb
}

func (cb *CustomerBuilder) WithIsActive(isActive bool) *CustomerBuilder {
	cb.customer.isActive = isActive
	return cb
}

func (cb *CustomerBuilder) WithPets(pets []pet.Pet) *CustomerBuilder {
	cb.customer.pets = pets
	return cb
}

func (cb *CustomerBuilder) WithTimestamp(createdAt, updatedAt time.Time) *CustomerBuilder {
	cb.customer.SetTimeStamps(createdAt, updatedAt)
	return cb
}

func (cb *CustomerBuilder) Build() *Customer { return cb.customer }

func (o *Customer) ID() valueobject.CustomerID      { return o.Entity.ID() }
func (o *Customer) Photo() string                   { return o.photo }
func (o *Customer) UserID() *valueobject.UserID     { return o.userID }
func (o *Customer) IsActive() bool                  { return o.isActive }
func (o *Customer) Pets() []pet.Pet                 { return o.pets }
func (o *Customer) SetID(id valueobject.CustomerID) { o.Entity.SetID(id) }
func (o *Customer) CreatedAt() time.Time            { return o.Entity.CreatedAt() }
func (o *Customer) UpdatedAt() time.Time            { return o.Entity.UpdatedAt() }
