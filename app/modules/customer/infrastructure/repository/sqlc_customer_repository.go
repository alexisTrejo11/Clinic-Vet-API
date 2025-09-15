package repository

import (
	"clinic-vet-api/app/core/domain/entity/customer"
	"clinic-vet-api/app/core/domain/entity/pet"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
)

type SqlcCustomerRepository struct {
	queries       *sqlc.Queries
	petRepository repository.PetRepository
}

func NewSqlcCustomerRepository(queries *sqlc.Queries, petRepository repository.PetRepository) repository.CustomerRepository {
	return &SqlcCustomerRepository{
		queries:       queries,
		petRepository: petRepository,
	}
}

func (r *SqlcCustomerRepository) FindBySpecification(ctx context.Context, spec specification.CustomerSpecification) (page.Page[customer.Customer], error) {
	return page.Page[customer.Customer]{}, nil
}

// FindByID finds a customer by ID
func (r *SqlcCustomerRepository) FindByID(ctx context.Context, id valueobject.CustomerID) (customer.Customer, error) {
	type petsResult struct {
		pets []pet.Pet
		err  error
	}

	petsChan := make(chan petsResult, ConcurrentOpTimeout)
	var wg sync.WaitGroup

	customerRow, err := r.queries.GetCustomerByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customer.Customer{}, r.notFoundError("id", fmt.Sprintf("%d", id.Value()))
		}
		return customer.Customer{}, r.dbError("select", fmt.Sprintf("failed to get customer with ID %d", id.Value()), err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(petsChan)

		pets, petErr := r.petRepository.FindAllByCustomerID(ctx, id)
		petsChan <- petsResult{pets: pets, err: petErr}
	}()

	wg.Wait()
	petsRes := <-petsChan

	if petsRes.err != nil {
		return customer.Customer{}, fmt.Errorf("failed to get pets for customer ID %d: %w", id.Value(), petsRes.err)
	}

	customerEntity, err := sqlRowToCustomer(customerRow, petsRes.pets)
	if err != nil {
		return customer.Customer{}, r.wrapConversionError(err)
	}

	return customerEntity, nil
}

// FindActive finds active customers with pagination
func (r *SqlcCustomerRepository) FindActive(ctx context.Context, pageInput page.PageInput) (page.Page[customer.Customer], error) {
	offset := (pageInput.Page - 1) * pageInput.Size
	limit := pageInput.Size

	// Get customers
	customerRows, err := r.queries.ListActiveCustomers(ctx, sqlc.ListActiveCustomersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return page.Page[customer.Customer]{}, r.dbError("select", "failed to list active customers", err)
	}

	// Get total count
	total, err := r.queries.CountActiveCustomers(ctx)
	if err != nil {
		return page.Page[customer.Customer]{}, r.dbError("select", "failed to count active customers", err)
	}

	// Convert rows to entities
	var customers []customer.Customer
	for _, row := range customerRows {
		customerEntity, err := sqlRowToCustomer(row, nil)
		if err != nil {
			return page.Page[customer.Customer]{}, r.wrapConversionError(err)
		}
		customers = append(customers, customerEntity)
	}

	paginationMetadata := page.GetPageMetadata(total, pageInput)
	return page.NewPage(customers, *paginationMetadata), nil
}

// ExistsByID checks if a customer exists by ID
func (r *SqlcCustomerRepository) ExistsByID(ctx context.Context, id valueobject.CustomerID) (bool, error) {
	exists, err := r.queries.ExistsCustomerByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check existence by ID %d", id.Value()), err)
	}
	return exists, nil
}

// ExistsByPhone checks if a customer exists by phone number
func (r *SqlcCustomerRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	exists, err := r.queries.ExistsCustomerByPhone(ctx, phone)
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check existence by phone %s", phone), err)
	}
	return exists, nil
}

// ExistsByEmail checks if a customer exists by email
func (r *SqlcCustomerRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := r.queries.ExistsCustomerByEmail(ctx, email)
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check existence by email %s", email), err)
	}
	return exists, nil
}

// Save creates or updates a customer
func (r *SqlcCustomerRepository) Save(ctx context.Context, customer *customer.Customer) error {
	if customer.ID().IsZero() {
		return r.create(ctx, customer)
	}
	return r.update(ctx, customer)
}

// Update updates an existing customer
func (r *SqlcCustomerRepository) Update(ctx context.Context, customer *customer.Customer) error {
	return r.update(ctx, customer)
}

// SoftDelete performs a soft delete of a customer
func (r *SqlcCustomerRepository) SoftDelete(ctx context.Context, id valueobject.CustomerID) error {
	if err := r.queries.SoftDeleteCustomer(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", fmt.Sprintf("failed to soft delete customer with ID %d", id.Value()), err)
	}
	return nil
}

// HardDelete performs a hard delete of a customer
func (r *SqlcCustomerRepository) HardDelete(ctx context.Context, id valueobject.CustomerID) error {
	if err := r.queries.HardDeleteCustomer(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", fmt.Sprintf("failed to hard delete customer with ID %d", id.Value()), err)
	}
	return nil
}

// CountAll counts all non-deleted customers
func (r *SqlcCustomerRepository) CountAll(ctx context.Context) (int64, error) {
	count, err := r.queries.CountAllCustomers(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count all customers", err)
	}
	return count, nil
}

// CountActive counts all active customers
func (r *SqlcCustomerRepository) CountActive(ctx context.Context) (int64, error) {
	count, err := r.queries.CountActiveCustomers(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count active customers", err)
	}
	return count, nil
}

// create inserts a new customer
func (r *SqlcCustomerRepository) create(ctx context.Context, customer *customer.Customer) error {
	createParams := toCreateParams(*customer)

	_, err := r.queries.CreateCustomer(ctx, *createParams)
	if err != nil {
		return r.dbError("insert", "failed to create customer", err)
	}

	return nil
}

// update modifies an existing customer
func (r *SqlcCustomerRepository) update(ctx context.Context, customer *customer.Customer) error {
	params := entityToUpdateParams(*customer)

	if err := r.queries.UpdateCustomer(ctx, *params); err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update customer with ID %d", customer.ID().Value()), err)
	}

	return nil
}
