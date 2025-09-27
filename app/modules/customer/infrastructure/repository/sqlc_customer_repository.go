package repository

import (
	c "clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
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

func (r *SqlcCustomerRepository) FindBySpecification(ctx context.Context, spec specification.CustomerSpecification) (page.Page[c.Customer], error) {
	return page.Page[c.Customer]{}, nil
}

func (r *SqlcCustomerRepository) FindByID(ctx context.Context, id valueobject.CustomerID) (c.Customer, error) {
	type petsResult struct {
		pets []pet.Pet
		err  error
	}

	petsChan := make(chan petsResult, ConcurrentOpTimeout)
	var wg sync.WaitGroup

	customerRow, err := r.queries.GetCustomerByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Customer{}, r.notFoundError("id", fmt.Sprintf("%d", id.Value()))
		}
		return c.Customer{}, r.dbError("select", fmt.Sprintf("failed to get customer with ID %d", id.Value()), err)
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
		return c.Customer{}, fmt.Errorf("failed to get pets for customer ID %d: %w", id.Value(), petsRes.err)
	}

	customerEntity := sqlRowToCustomer(customerRow, petsRes.pets)
	return customerEntity, nil
}

func (r *SqlcCustomerRepository) FindActive(ctx context.Context, pagination page.PaginationRequest) (page.Page[c.Customer], error) {
	customerRows, err := r.queries.FindActiveCustomers(ctx, sqlc.FindActiveCustomersParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[c.Customer]{}, r.dbError("select", "failed to list active customers", err)
	}

	total, err := r.queries.CountActiveCustomers(ctx)
	if err != nil {
		return page.Page[c.Customer]{}, r.dbError("select", "failed to count active customers", err)
	}

	customers := sqlRowsToCustomers(customerRows)
	return page.NewPage(customers, int(total), pagination), nil
}

func (r *SqlcCustomerRepository) ExistsByID(ctx context.Context, id valueobject.CustomerID) (bool, error) {
	exists, err := r.queries.ExistsCustomerByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check existence by ID %d", id.Value()), err)
	}
	return exists, nil
}

func (r *SqlcCustomerRepository) Save(ctx context.Context, customer *c.Customer) error {
	if customer.ID().IsZero() {
		return r.create(ctx, customer)
	}
	return r.update(ctx, customer)
}

func (r *SqlcCustomerRepository) Update(ctx context.Context, customer *c.Customer) error {
	return r.update(ctx, customer)
}

func (r *SqlcCustomerRepository) SoftDelete(ctx context.Context, id valueobject.CustomerID) error {
	if err := r.queries.SoftDeleteCustomer(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", fmt.Sprintf("failed to soft delete customer with ID %d", id.Value()), err)
	}
	return nil
}

func (r *SqlcCustomerRepository) HardDelete(ctx context.Context, id valueobject.CustomerID) error {
	if err := r.queries.HardDeleteCustomer(ctx, int32(id.Value())); err != nil {
		return r.dbError("delete", fmt.Sprintf("failed to hard delete customer with ID %d", id.Value()), err)
	}
	return nil
}

func (r *SqlcCustomerRepository) CountAll(ctx context.Context) (int64, error) {
	count, err := r.queries.CountAllCustomers(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count all customers", err)
	}
	return count, nil
}

func (r *SqlcCustomerRepository) CountActive(ctx context.Context) (int64, error) {
	count, err := r.queries.CountActiveCustomers(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count active customers", err)
	}
	return count, nil
}

func (r *SqlcCustomerRepository) create(ctx context.Context, customer *c.Customer) error {
	createParams := toCreateParams(*customer)

	customerCreated, err := r.queries.CreateCustomer(ctx, *createParams)
	if err != nil {
		return r.dbError("insert", "failed to create customer", err)
	}

	customer.SetID(valueobject.NewCustomerID(uint(customerCreated.ID)))

	return nil
}

func (r *SqlcCustomerRepository) update(ctx context.Context, customer *c.Customer) error {
	params := entityToUpdateParams(*customer)
	if err := r.queries.UpdateCustomer(ctx, *params); err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update customer with ID %d", customer.ID().Value()), err)
	}

	return nil
}
