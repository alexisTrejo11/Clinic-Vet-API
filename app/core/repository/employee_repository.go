package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/employee"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type EmployeeRepository interface {
	FindBySpecification(ctx context.Context, spec specification.EmployeeSearchSpecification) (page.Page[employee.Employee], error)
	FindByID(ctx context.Context, id valueobject.EmployeeID) (employee.Employee, error)
	FindByUserID(ctx context.Context, userID valueobject.UserID) (employee.Employee, error)

	FindAll(ctx context.Context, pageInput page.PageInput) (page.Page[employee.Employee], error)
	FindByRole(ctx context.Context, role string, pageInput page.PageInput) (page.Page[employee.Employee], error)
	FindByStatus(ctx context.Context, string, pageInput page.PageInput) (page.Page[employee.Employee], error)

	ExistsByID(ctx context.Context, id valueobject.EmployeeID) (bool, error)
	ExistsByUserID(ctx context.Context, userID valueobject.UserID) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	Save(ctx context.Context, employee *employee.Employee) error
	Update(ctx context.Context, employee *employee.Employee) error
	SoftDelete(ctx context.Context, id valueobject.EmployeeID) error
	HardDelete(ctx context.Context, id valueobject.EmployeeID) error

	CountByRole(ctx context.Context, role string) (int64, error)
	CountByStatus(ctx context.Context, status string) (int64, error)
	CountActive(ctx context.Context) (int64, error)
}
