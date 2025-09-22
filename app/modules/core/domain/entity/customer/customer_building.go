package customer

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type CustomerOption func(context.Context, *Customer) error

func WithPhoto(photo string) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		o.photo = photo
		return nil
	}
}

func WithFullName(fullName valueobject.PersonName) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		return o.UpdateName(ctx, fullName)
	}
}

func WithGender(gender enum.PersonGender) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		return o.Person.UpdateGender(ctx, gender)
	}
}

func WithDateOfBirth(dob time.Time) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		return o.UpdateDateOfBirth(ctx, dob)
	}
}

func WithUserID(userID *valueobject.UserID) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		o.userID = userID
		return nil
	}
}

func WithIsActive(isActive bool) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		o.isActive = isActive
		return nil
	}
}

func WithPets(pets []pet.Pet) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		o.pets = pets
		return nil
	}
}

func WithTimestamp(createdAt, updatedAt time.Time) CustomerOption {
	return func(ctx context.Context, o *Customer) error {
		o.Entity.SetTimeStamps(createdAt, updatedAt)
		return nil
	}
}

func NewCustomer(
	id valueobject.CustomerID,
	opts ...CustomerOption,
) (*Customer, error) {
	ctx := context.Background()
	return NewCustomerWithContext(ctx, id, opts...)
}

func NewCustomerWithContext(
	ctx context.Context,
	id valueobject.CustomerID,
	opts ...CustomerOption,
) (*Customer, error) {
	operation := "NewCustomerWithContext"

	customer := &Customer{
		Entity:   base.NewEntity(id, time.Now(), time.Now(), 1),
		isActive: true,
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		if err := opt(ctx, customer); err != nil {
			return nil, err
		}
	}

	if err := customer.validate(ctx, operation); err != nil {
		return nil, err
	}

	return customer, nil
}

func CreateCustomer(ctx context.Context, opts ...CustomerOption) (*Customer, error) {
	operation := "CreateCustomer"

	customer := &Customer{
		Entity:   base.CreateEntity(valueobject.CustomerID{}),
		isActive: true,
		pets:     []pet.Pet{},
	}

	for _, opt := range opts {
		if err := opt(ctx, customer); err != nil {
			return nil, err
		}
	}

	if err := customer.validate(ctx, operation); err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *Customer) validate(ctx context.Context, operation string) error {
	if err := c.Person.Validate(ctx, "create customer"); err != nil {
		return err
	}

	return nil
}
