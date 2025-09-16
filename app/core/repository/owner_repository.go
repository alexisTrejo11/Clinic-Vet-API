package repository

import (
	"clinic-vet-api/app/core/domain/entity/customer"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
	"context"
)

type CustomerRepository interface {
	FindBySpecification(ctx context.Context, spec specification.CustomerSpecification) (page.Page[customer.Customer], error)
	FindByID(ctx context.Context, id valueobject.CustomerID) (customer.Customer, error)
	FindActive(ctx context.Context, pageInput page.PageInput) (page.Page[customer.Customer], error)

	ExistsByID(ctx context.Context, id valueobject.CustomerID) (bool, error)

	Save(ctx context.Context, customer *customer.Customer) error
	Update(ctx context.Context, customer *customer.Customer) error
	SoftDelete(ctx context.Context, id valueobject.CustomerID) error
	HardDelete(ctx context.Context, id valueobject.CustomerID) error

	CountAll(ctx context.Context) (int64, error)
	CountActive(ctx context.Context) (int64, error)
}
