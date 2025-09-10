package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/employee"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type EmployeeRepository interface {
	Search(ctx context.Context, specification specification.EmployeeSearchSpecification) (page.Page[[]employee.Employee], error)
	GetByID(ctx context.Context, id valueobject.EmployeeID) (employee.Employee, error)
	Exists(ctx context.Context, id valueobject.EmployeeID) (bool, error)
	GetByUserID(ctx context.Context, userID valueobject.UserID) (employee.Employee, error)
	SoftDelete(ctx context.Context, id valueobject.EmployeeID) error
	Save(ctx context.Context, employee *employee.Employee) error
}
