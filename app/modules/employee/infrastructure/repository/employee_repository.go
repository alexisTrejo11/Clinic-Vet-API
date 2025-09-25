// Package repository includes the implementation of repositories using SQLC
package repository

import (
	e "clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	p "clinic-vet-api/app/shared/page"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SqlcEmployeeRepository struct {
	queries *sqlc.Queries
	pool    *pgxpool.Pool
}

func NewSqlcEmployeeRepository(queries *sqlc.Queries, pool *pgxpool.Pool) repository.EmployeeRepository {
	return &SqlcEmployeeRepository{
		queries: queries,
		pool:    pool,
	}
}

func (r *SqlcEmployeeRepository) FindBySpecification(ctx context.Context, spec specification.EmployeeSearchSpecification) (p.Page[e.Employee], error) {
	query, params := spec.ToSQL()

	rows, err := r.pool.Query(ctx, query, params...)
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "failed to search veterinarians", err)
	}
	defer rows.Close()

	var employees []e.Employee
	for rows.Next() {
		var employee e.Employee
		err := r.scanEmployeeFromRow(rows, &employee)
		if err != nil {
			return p.Page[e.Employee]{}, r.wrapConversionError(err)
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "error iterating search results", err)
	}

	totalCount, err := r.getTotalCountWithFilters(ctx, &spec)
	if err != nil {
		return p.Page[e.Employee]{}, err
	}

	pagination := p.FromSpecPagination(spec.Pagination)
	return p.NewPage(employees, totalCount, pagination), nil
}

func (r *SqlcEmployeeRepository) FindByID(ctx context.Context, id valueobject.EmployeeID) (e.Employee, error) {
	sqlRow, err := r.queries.FindEmployeeByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.Employee{}, r.notFoundError("id", id.String())
		}
		return e.Employee{}, r.dbError("select", "failed to get employee by ID", err)
	}

	employeeEntity := sqlcRowToEntity(sqlRow)
	return *employeeEntity, nil
}

func (r *SqlcEmployeeRepository) FindByUserID(ctx context.Context, userID valueobject.UserID) (e.Employee, error) {
	sqlRow, err := r.queries.FindEmployeeByUserID(ctx, pgtype.Int4{Int32: int32(userID.Value()), Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.Employee{}, r.notFoundError("user_id", userID.String())
		}
		return e.Employee{}, r.dbError("select", "failed to get employee by user ID", err)
	}

	employee := sqlcRowToEntity(sqlRow)
	return *employee, nil
}

func (r *SqlcEmployeeRepository) FindActive(ctx context.Context, pagination p.PaginationRequest) (p.Page[e.Employee], error) {
	employeeRows, err := r.queries.FindActiveEmployees(ctx, sqlc.FindActiveEmployeesParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError("select", "failed to list active employees", err)
	}

	total, err := r.queries.CountActiveEmployees(ctx)
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError("select", "failed to count active employees", err)
	}
	return p.NewPage(sqlcRowsToEntities(employeeRows), int(total), pagination), nil
}

func (r *SqlcEmployeeRepository) FindAll(ctx context.Context, pagination p.PaginationRequest) (p.Page[e.Employee], error) {
	employeeRows, err := r.queries.FindEmployees(ctx, sqlc.FindEmployeesParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError("select", "failed to list employees", err)
	}

	total, err := r.queries.CountAllEmployees(ctx)
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError("select", "failed to count employees", err)
	}

	return p.NewPage(sqlcRowsToEntities(employeeRows), int(total), pagination), nil
}

func (r *SqlcEmployeeRepository) FindBySpeciality(ctx context.Context, speciality enum.VetSpecialty, pagination p.PaginationRequest) (p.Page[e.Employee], error) {
	employeeRows, err := r.queries.FindEmployeesBySpeciality(ctx, sqlc.FindEmployeesBySpecialityParams{
		Speciality: models.VeterinarianSpeciality(speciality.DisplayName()),
		Limit:      int32(pagination.Limit()),
		Offset:     int32(pagination.Offset()),
	})
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError("select", "failed to list employees by speciality", err)
	}

	total, err := r.queries.CountEmployeesBySpeciality(ctx, models.VeterinarianSpeciality(speciality.DisplayName()))
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError("select", "failed to count employees by speciality", err)
	}

	employees := sqlcRowsToEntities(employeeRows)

	return p.NewPage(employees, int(total), pagination), nil
}

func (r *SqlcEmployeeRepository) ExistsByID(ctx context.Context, id valueobject.EmployeeID) (bool, error) {
	exists, err := r.queries.ExistsEmployeeByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError("select", "failed to check employee existence by ID", err)
	}
	return exists, nil
}

func (r *SqlcEmployeeRepository) ExistsByUserID(ctx context.Context, userID valueobject.UserID) (bool, error) {
	exists, err := r.queries.ExistsEmployeeByUserID(ctx, pgtype.Int4{Int32: int32(userID.Value()), Valid: true})
	if err != nil {
		return false, r.dbError("select", "failed to check employee existence by user ID", err)
	}
	return exists, nil
}

func (r *SqlcEmployeeRepository) Save(ctx context.Context, employee *e.Employee) error {
	if employee.ID().IsZero() {
		return r.create(ctx, employee)
	}
	return r.update(ctx, employee)
}

func (r *SqlcEmployeeRepository) Update(ctx context.Context, employee *e.Employee) error {
	return r.update(ctx, employee)
}

func (r *SqlcEmployeeRepository) SoftDelete(ctx context.Context, id valueobject.EmployeeID) error {
	if err := r.queries.SoftDeleteEmployee(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", "failed to soft delete employee", err)
	}
	return nil
}

func (r *SqlcEmployeeRepository) HardDelete(ctx context.Context, id valueobject.EmployeeID) error {
	if err := r.queries.HardDeleteEmployee(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", "failed to hard delete employee", err)
	}
	return nil
}

func (r *SqlcEmployeeRepository) CountBySpeciality(ctx context.Context, status enum.VetSpecialty) (int64, error) {
	count, err := r.queries.CountEmployeesBySpeciality(ctx, models.VeterinarianSpeciality(status.DisplayName()))
	if err != nil {
		return 0, r.dbError("select", "failed to count employees by status", err)
	}
	return count, nil
}

func (r *SqlcEmployeeRepository) CountActive(ctx context.Context) (int64, error) {
	count, err := r.queries.CountActiveEmployees(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count active employees", err)
	}
	return count, nil
}

func (r *SqlcEmployeeRepository) create(ctx context.Context, employee *e.Employee) error {
	params := entityToCreateParams(employee)

	employeeCreated, err := r.queries.CreateEmployee(ctx, *params)
	if err != nil {
		return r.dbError("insert", "failed to create employee", err)
	}

	employee.SetID(valueobject.NewEmployeeID(uint(employeeCreated.ID)))

	return nil
}

func (r *SqlcEmployeeRepository) update(ctx context.Context, employee *e.Employee) error {
	params := entityToUpdateParams(employee)

	_, err := r.queries.UpdateEmployee(ctx, *params)
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update employee with ID %d", employee.ID().Value()), err)
	}

	return nil
}
