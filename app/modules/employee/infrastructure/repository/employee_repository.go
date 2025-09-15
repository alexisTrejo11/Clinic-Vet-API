// Package repository includes the implementation of repositories using SQLC
package repository

import (
	"clinic-vet-api/app/core/domain/entity/customer"
	"clinic-vet-api/app/core/domain/entity/employee"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SqlcEmployeeRepository struct {
	queries *sqlc.Queries
	pool    *pgxpool.Pool
}

func NewSqlcEmployeeRepository(queries *sqlc.Queries, pool *pgxpool.Pool) repository.Emplocurl -L https://github.com/sqlc-dev/sqlc/releases/download/v1.26.0/sqlc_1.26.0_linux_amd64.tar.gz -o sqlc.tar.gzyeeRepository {
	return &SqlcEmployeeRepository{
		queries: queries,
		pool:    pool,
	}
}

func (r *SqlcEmployeeRepository) Search(ctx context.Context, spec specification.EmployeeSearchSpecification) (page.Page[employee.Employee], error) {
	query, params := spec.ToSQL()

	// Execute the query
	rows, err := r.pool.Query(ctx, query, params...)
	if err != nil {
		return page.Page[[]customer.Customer]{}, r.dbError(OpSelect, "failed to search veterinarians", err)
	}
	defer rows.Close()

	// Iterate through the rows and scan into Customer structs
	var vets []customer.Customer
	for rows.Next() {
		var veterinarian customer.Customer
		err := r.scanEmployeeFromRow(rows, &veterinarian)
		if err != nil {
			return page.Page[[]customer.Customer]{}, r.wrapConversionError(err)
		}
		vets = append(vets, veterinarian)
	}

	if err := rows.Err(); err != nil {
		return page.Page[[]customer.Customer]{}, r.dbError(OpSelect, "error iterating search results", err)
	}

	// Get total count for pagination
	totalCount, err := r.getTotalCountWithFilters(ctx, &spec)
	if err != nil {
		return page.Page[[]customer.Customer]{}, err
	}

	// Handle pagination
	pageMetadata := page.GetPageMetadata(totalCount, page.PageInput{
		PageNumber: spec.GetPagination().Page,
		PageSize:   spec.GetPagination().PageSize,
	})

	return page.NewPage(vets, *pageMetadata), nil
}

func (r *SqlcEmployeeRepository) FindByID(ctx context.Context, id valueobject.EmployeeID) (employee.Employee, error) {
	sqlRow, err := r.queries.GetEmployeeByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return employee.Employee{}, r.notFoundError("id", id.String())
		}
		return employee.Employee{}, r.dbError("select", "failed to get employee by ID", err)
	}

	employeeEntity, err := sqlRowToEmployee(sqlRow)
	if err != nil {
		return employee.Employee{}, r.wrapConversionError(err)
	}

	return employeeEntity, nil
}

// FindByUserID finds an employee by user ID
func (r *SqlcEmployeeRepository) FindByUserID(ctx context.Context, userID valueobject.UserID) (employee.Employee, error) {
	sqlRow, err := r.queries.GetEmployeeByUserID(ctx, int32(userID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return employee.Employee{}, r.notFoundError("user_id", userID.String())
		}
		return employee.Employee{}, r.dbError("select", "failed to get employee by user ID", err)
	}

	employeeEntity, err := sqlRowToEmployee(sqlRow)
	if err != nil {
		return employee.Employee{}, r.wrapConversionError(err)
	}

	return employeeEntity, nil
}

// FindActive finds active employees with pagination
func (r *SqlcEmployeeRepository) FindActive(ctx context.Context, pageInput page.PageInput) (page.Page[employee.Employee], error) {
	offset := (pageInput.Page - 1) * pageInput.Size
	limit := pageInput.Size

	// Get employees
	employeeRows, err := r.queries.ListActiveEmployees(ctx, sqlc.ListActiveEmployeesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to list active employees", err)
	}

	// Get total count
	total, err := r.queries.CountActiveEmployees(ctx)
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to count active employees", err)
	}

	// Convert rows to entities
	var employees []employee.Employee
	for _, row := range employeeRows {
		employeeEntity, err := sqlRowToEmployee(row)
		if err != nil {
			return page.Page[employee.Employee]{}, r.wrapConversionError(err)
		}
		employees = append(employees, employeeEntity)
	}

	return page.Page[employee.Employee]{
		Items:      employees,
		Page:       pageInput.Page,
		Size:       pageInput.Size,
		Total:      total,
		TotalPages: (total + int64(pageInput.Size) - 1) / int64(pageInput.Size),
	}, nil
}

// FindAll finds all employees with pagination
func (r *SqlcEmployeeRepository) FindAll(ctx context.Context, pageInput page.PageInput) (page.Page[employee.Employee], error) {
	offset := (pageInput.Page - 1) * pageInput.Size
	limit := pageInput.Size

	// Get employees
	employeeRows, err := r.queries.ListEmployees(ctx, sqlc.ListEmployeesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to list employees", err)
	}

	// Get total count
	total, err := r.queries.CountAllEmployees(ctx)
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to count employees", err)
	}

	// Convert rows to entities
	var employees []employee.Employee
	for _, row := range employeeRows {
		employeeEntity, err := sqlRowToEmployee(row)
		if err != nil {
			return page.Page[employee.Employee]{}, r.wrapConversionError(err)
		}
		employees = append(employees, employeeEntity)
	}

	return page.Page[employee.Employee]{
		Items:      employees,
		Page:       pageInput.Page,
		Size:       pageInput.Size,
		Total:      total,
		TotalPages: (total + int64(pageInput.Size) - 1) / int64(pageInput.Size),
	}, nil
}

// FindByRole finds employees by role with pagination
func (r *SqlcEmployeeRepository) FindByRole(ctx context.Context, role string, pageInput page.PageInput) (page.Page[employee.Employee], error) {
	offset := (pageInput.Page - 1) * pageInput.Size
	limit := pageInput.Size

	// Get employees
	employeeRows, err := r.queries.ListEmployeesByRole(ctx, sqlc.ListEmployeesByRoleParams{
		Role:   role,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to list employees by role", err)
	}

	// Get total count
	total, err := r.queries.CountEmployeesByRole(ctx, role)
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to count employees by role", err)
	}

	// Convert rows to entities
	var employees []employee.Employee
	for _, row := range employeeRows {
		employeeEntity, err := sqlRowToEmployee(row)
		if err != nil {
			return page.Page[employee.Employee]{}, r.wrapConversionError(err)
		}
		employees = append(employees, employeeEntity)
	}

	return page.Page[employee.Employee]{
		Items:      employees,
		Page:       pageInput.Page,
		Size:       pageInput.Size,
		Total:      total,
		TotalPages: (total + int64(pageInput.Size) - 1) / int64(pageInput.Size),
	}, nil
}

// FindByStatus finds employees by status with pagination
func (r *SqlcEmployeeRepository) FindByStatus(ctx context.Context, status string, pageInput page.PageInput) (page.Page[employee.Employee], error) {
	offset := (pageInput.Page - 1) * pageInput.Size
	limit := pageInput.Size

	// Get employees
	employeeRows, err := r.queries.ListEmployeesByStatus(ctx, sqlc.ListEmployeesByStatusParams{
		Status: status,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to list employees by status", err)
	}

	// Get total count
	total, err := r.queries.CountEmployeesByStatus(ctx, status)
	if err != nil {
		return page.Page[employee.Employee]{}, r.dbError("select", "failed to count employees by status", err)
	}

	// Convert rows to entities
	var employees []employee.Employee
	for _, row := range employeeRows {
		employeeEntity, err := sqlRowToEmployee(row)
		if err != nil {
			return page.Page[employee.Employee]{}, r.wrapConversionError(err)
		}
		employees = append(employees, employeeEntity)
	}

	return page.Page[employee.Employee]{
		Items:      employees,
		Page:       pageInput.Page,
		Size:       pageInput.Size,
		Total:      total,
		TotalPages: (total + int64(pageInput.Size) - 1) / int64(pageInput.Size),
	}, nil
}

// ExistsByID checks if an employee exists by ID
func (r *SqlcEmployeeRepository) ExistsByID(ctx context.Context, id valueobject.EmployeeID) (bool, error) {
	exists, err := r.queries.ExistsEmployeeByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError("select", "failed to check employee existence by ID", err)
	}
	return exists, nil
}

// ExistsByUserID checks if an employee exists by user ID
func (r *SqlcEmployeeRepository) ExistsByUserID(ctx context.Context, userID valueobject.UserID) (bool, error) {
	exists, err := r.queries.ExistsEmployeeByUserID(ctx, int32(userID.Value()))
	if err != nil {
		return false, r.dbError("select", "failed to check employee existence by user ID", err)
	}
	return exists, nil
}

// ExistsByEmail checks if an employee exists by email
func (r *SqlcEmployeeRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := r.queries.ExistsEmployeeByEmail(ctx, email)
	if err != nil {
		return false, r.dbError("select", "failed to check employee existence by email", err)
	}
	return exists, nil
}

// Save creates or updates an employee
func (r *SqlcEmployeeRepository) Save(ctx context.Context, employee *employee.Employee) error {
	if employee.ID().IsZero() {
		return r.create(ctx, employee)
	}
	return r.update(ctx, employee)
}

// Update updates an existing employee
func (r *SqlcEmployeeRepository) Update(ctx context.Context, employee *employee.Employee) error {
	return r.update(ctx, employee)
}

// SoftDelete performs a soft delete of an employee
func (r *SqlcEmployeeRepository) SoftDelete(ctx context.Context, id valueobject.EmployeeID) error {
	if err := r.queries.SoftDeleteEmployee(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", "failed to soft delete employee", err)
	}
	return nil
}

// HardDelete performs a hard delete of an employee
func (r *SqlcEmployeeRepository) HardDelete(ctx context.Context, id valueobject.EmployeeID) error {
	if err := r.queries.HardDeleteEmployee(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", "failed to hard delete employee", err)
	}
	return nil
}

// CountByRole counts employees by role
func (r *SqlcEmployeeRepository) CountByRole(ctx context.Context, role string) (int64, error) {
	count, err := r.queries.CountEmployeesByRole(ctx, role)
	if err != nil {
		return 0, r.dbError("select", "failed to count employees by role", err)
	}
	return count, nil
}

// CountByStatus counts employees by status
func (r *SqlcEmployeeRepository) CountByStatus(ctx context.Context, status string) (int64, error) {
	count, err := r.queries.CountEmployeesByStatus(ctx, status)
	if err != nil {
		return 0, r.dbError("select", "failed to count employees by status", err)
	}
	return count, nil
}

// CountActive counts active employees
func (r *SqlcEmployeeRepository) CountActive(ctx context.Context) (int64, error) {
	count, err := r.queries.CountActiveEmployees(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count active employees", err)
	}
	return count, nil
}

// Helper methods
func (r *SqlcEmployeeRepository) notFoundError(field, value string) error {
	return fmt.Errorf("employee with %s %s not found", field, value)
}

func (r *SqlcEmployeeRepository) dbError(operation, message string, err error) error {
	return fmt.Errorf("database error during %s: %s: %w", operation, message, err)
}

func (r *SqlcEmployeeRepository) wrapConversionError(err error) error {
	return fmt.Errorf("conversion error: %w", err)
}

// create inserts a new employee
func (r *SqlcEmployeeRepository) create(ctx context.Context, employee *employee.Employee) error {
	params := toCreateEmployeeParams(employee)

	_, err := r.queries.CreateEmployee(ctx, *params)
	if err != nil {
		return r.dbError("insert", "failed to create employee", err)
	}

	return nil
}

// update modifies an existing employee
func (r *SqlcEmployeeRepository) update(ctx context.Context, employee *employee.Employee) error {
	params := employeeToUpdateParams(employee)

	_, err := r.queries.UpdateEmployee(ctx, *params)
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update employee with ID %d", employee.ID().Value()), err)
	}

	return nil
}
