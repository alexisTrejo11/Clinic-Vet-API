package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/customer"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *customer.Customer) error
	Search(ctx context.Context, search any) (page.Page[[]customer.Customer], error)
	GetByID(ctx context.Context, id valueobject.CustomerID) (customer.Customer, error)
	GetByPhone(ctx context.Context, phone string) (customer.Customer, error)
	SoftDelete(ctx context.Context, id valueobject.CustomerID) error

	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id valueobject.CustomerID) (bool, error)
}
