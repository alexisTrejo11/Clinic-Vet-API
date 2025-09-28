package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
	"context"
)

type EmployeeRepository interface {
	FindByID(ctx context.Context, id valueobject.EmployeeID) (employee.Employee, error)
	FindByUserID(ctx context.Context, userID valueobject.UserID) (employee.Employee, error)
	FindActive(ctx context.Context, PaginationRequest page.PaginationRequest) (page.Page[employee.Employee], error)
	FindAll(ctx context.Context, PaginationRequest page.PaginationRequest) (page.Page[employee.Employee], error)
	FindBySpeciality(ctx context.Context, speciality enum.VetSpecialty, PaginationRequest page.PaginationRequest) (page.Page[employee.Employee], error)

	ExistsByID(ctx context.Context, id valueobject.EmployeeID) (bool, error)
	ExistsByUserID(ctx context.Context, userID valueobject.UserID) (bool, error)

	Save(ctx context.Context, employee *employee.Employee) error
	Delete(ctx context.Context, id valueobject.EmployeeID, isHard bool) error

	CountBySpeciality(ctx context.Context, speciality enum.VetSpecialty) (int64, error)
	CountActive(ctx context.Context) (int64, error)
}
