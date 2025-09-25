package customer

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type CustomerOption func(context.Context, *Customer)

func WithPhoto(photo string) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.photo = photo
	}
}

func WithFullName(fullName valueobject.PersonName) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.UpdateName(ctx, fullName)
	}
}

func WithGender(gender enum.PersonGender) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.Person.UpdateGender(ctx, gender)
	}
}

func WithDateOfBirth(dob time.Time) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.UpdateDateOfBirth(ctx, dob)
	}
}

func WithUserID(userID *valueobject.UserID) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.userID = userID
	}
}

func WithIsActive(isActive bool) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.isActive = isActive
	}
}

func WithPets(pets []pet.Pet) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.pets = pets
	}
}

func WithTimestamp(createdAt, updatedAt time.Time) CustomerOption {
	return func(ctx context.Context, o *Customer) {
		o.Entity.SetTimeStamps(createdAt, updatedAt)
	}
}

func NewCustomer(
	id valueobject.CustomerID,
	opts ...CustomerOption,
) *Customer {
	ctx := context.Background()
	return NewCustomerWithContext(ctx, id, opts...)
}

func NewCustomerWithContext(
	ctx context.Context,
	id valueobject.CustomerID,
	opts ...CustomerOption,
) *Customer {
	customer := &Customer{
		Entity:   base.NewEntity(id, nil, nil, 1),
		isActive: true,
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		opt(ctx, customer)

	}

	return customer
}

func CreateCustomer(ctx context.Context, opts ...CustomerOption) (*Customer, error) {
	operation := "CreateCustomer"

	customer := &Customer{
		Entity:   base.CreateEntity(valueobject.CustomerID{}),
		isActive: true,
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		opt(ctx, customer)
	}

	if err := customer.validate(ctx, operation); err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) validate(ctx context.Context, operation string) error {
	return c.Person.Validate(ctx, operation)
}
