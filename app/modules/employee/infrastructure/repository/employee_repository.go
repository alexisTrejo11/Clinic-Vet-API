// Package repository includes the implementation of repositories using SQLC
package repository

import (
	e "clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/mapper"
	p "clinic-vet-api/app/shared/page"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcEmployeeRepository struct {
	queries *sqlc.Queries
	pgMap   *mapper.SqlcFieldMapper
}

func NewSqlcEmployeeRepository(queries *sqlc.Queries, pgMap *mapper.SqlcFieldMapper) repository.EmployeeRepository {
	return &SqlcEmployeeRepository{queries: queries, pgMap: pgMap}
}

func (r *SqlcEmployeeRepository) FindByID(ctx context.Context, id valueobject.EmployeeID) (e.Employee, error) {
	sqlRow, err := r.queries.FindEmployeeByID(ctx, id.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.Employee{}, r.notFoundError("id", id.String())
		}
		return e.Employee{}, r.dbError(OpSelect, "failed to get employee by ID", err)
	}

	return *r.toEntity(sqlRow), nil
}

func (r *SqlcEmployeeRepository) FindByUserID(ctx context.Context, userID valueobject.UserID) (e.Employee, error) {
	sqlRow, err := r.queries.FindEmployeeByUserID(ctx, r.pgMap.UintToPgInt4(userID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.Employee{}, r.notFoundError("user_id", userID.String())
		}
		return e.Employee{}, r.dbError(OpSelect, "failed to get employee by user ID", err)
	}

	return *r.toEntity(sqlRow), nil
}

func (r *SqlcEmployeeRepository) FindActive(ctx context.Context, pagination p.PaginationRequest) (p.Page[e.Employee], error) {
	employeeRows, err := r.queries.FindActiveEmployees(ctx, sqlc.FindActiveEmployeesParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "failed to list active employees", err)
	}

	total, err := r.queries.CountActiveEmployees(ctx)
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "failed to count active employees", err)
	}

	employees := r.toEntities(employeeRows)
	return p.NewPage(employees, total, pagination), nil
}

func (r *SqlcEmployeeRepository) FindAll(ctx context.Context, pagination p.PaginationRequest) (p.Page[e.Employee], error) {
	employeeRows, err := r.queries.FindEmployees(ctx, sqlc.FindEmployeesParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "failed to list employees", err)
	}

	total, err := r.queries.CountAllEmployees(ctx)
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "failed to count employees", err)
	}

	employees := r.toEntities(employeeRows)
	return p.NewPage(employees, total, pagination), nil
}

func (r *SqlcEmployeeRepository) FindBySpeciality(ctx context.Context, speciality enum.VetSpecialty, pagination p.PaginationRequest) (p.Page[e.Employee], error) {
	employeeRows, err := r.queries.FindEmployeesBySpeciality(ctx, sqlc.FindEmployeesBySpecialityParams{
		Speciality: models.VeterinarianSpeciality(speciality.String()),
		Limit:      pagination.Limit(),
		Offset:     pagination.Offset(),
	})
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "failed to list employees by speciality", err)
	}

	total, err := r.queries.CountEmployeesBySpeciality(ctx, models.VeterinarianSpeciality(speciality.DisplayName()))
	if err != nil {
		return p.Page[e.Employee]{}, r.dbError(OpSelect, "failed to count employees by speciality", err)
	}

	employees := r.toEntities(employeeRows)
	return p.NewPage(employees, total, pagination), nil
}

func (r *SqlcEmployeeRepository) ExistsByID(ctx context.Context, id valueobject.EmployeeID) (bool, error) {
	exists, err := r.queries.ExistsEmployeeByID(ctx, id.Int32())
	if err != nil {
		return false, r.dbError(OpSelect, "failed to check employee existence by ID", err)
	}
	return exists, nil
}

func (r *SqlcEmployeeRepository) ExistsByUserID(ctx context.Context, userID valueobject.UserID) (bool, error) {
	exists, err := r.queries.ExistsEmployeeByUserID(ctx, pgtype.Int4{Int32: int32(userID.Value()), Valid: true})
	if err != nil {
		return false, r.dbError(OpSelect, "failed to check employee existence by user ID", err)
	}
	return exists, nil
}

func (r *SqlcEmployeeRepository) Save(ctx context.Context, employee *e.Employee) error {
	if employee.ID().IsZero() {
		return r.create(ctx, employee)
	}
	return r.update(ctx, employee)
}

func (r *SqlcEmployeeRepository) Delete(ctx context.Context, id valueobject.EmployeeID, isHard bool) error {
	if isHard {
		if err := r.queries.HardDeleteEmployee(ctx, id.Int32()); err != nil {
			return r.dbError(OpDelete, "failed to hard delete employee", err)
		}
	}
	if err := r.queries.SoftDeleteEmployee(ctx, id.Int32()); err != nil {
		return r.dbError(OpDelete, "failed to soft delete employee", err)
	}
	return nil
}

func (r *SqlcEmployeeRepository) CountBySpeciality(ctx context.Context, status enum.VetSpecialty) (int64, error) {
	count, err := r.queries.CountEmployeesBySpeciality(ctx, models.VeterinarianSpeciality(status.DisplayName()))
	if err != nil {
		return 0, r.dbError(OpSelect, "failed to count employees by status", err)
	}
	return count, nil
}

func (r *SqlcEmployeeRepository) CountActive(ctx context.Context) (int64, error) {
	count, err := r.queries.CountActiveEmployees(ctx)
	if err != nil {
		return 0, r.dbError(OpSelect, "failed to count active employees", err)
	}
	return count, nil
}

func (r *SqlcEmployeeRepository) create(ctx context.Context, employee *e.Employee) error {
	createParams := r.toCreateParams(*employee)
	employeeCreated, err := r.queries.CreateEmployee(ctx, createParams)
	if err != nil {
		return r.dbError(OpInsert, "failed to create employee", err)
	}

	employee.SetID(valueobject.NewEmployeeID(uint(employeeCreated.ID)))
	return nil
}

func (r *SqlcEmployeeRepository) update(ctx context.Context, employee *e.Employee) error {
	updateParams := r.toUpdateParams(*employee)
	_, err := r.queries.UpdateEmployee(ctx, *updateParams)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to update employee with ID %d", employee.ID().Value()), err)
	}

	return nil
}
