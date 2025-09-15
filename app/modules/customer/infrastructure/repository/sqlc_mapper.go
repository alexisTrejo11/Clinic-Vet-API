package repository

import (
	"clinic-vet-api/app/core/domain/entity/customer"
	"clinic-vet-api/app/core/domain/entity/pet"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	appError "clinic-vet-api/app/shared/error/application"

	"github.com/jackc/pgx/v5/pgtype"
)

func toCreateParams(customer customer.Customer) *sqlc.CreateCustomerParams {
	createcustomerParam := &sqlc.CreateCustomerParams{
		Photo:       customer.Photo(),
		DateOfBirth: pgtype.Date{Time: customer.DateOfBirth(), Valid: true},
		Gender:      models.PersonGender(customer.Gender().String()),
		FirstName:   customer.FullName().FirstName,
		LastName:    customer.FullName().LastName,
		IsActive:    customer.IsActive(),
	}

	if customer.UserID() != nil {
		createcustomerParam.UserID = pgtype.Int4{Int32: int32(customer.UserID().Value())}
	}

	return createcustomerParam
}

func entityToUpdateParams(customer customer.Customer) *sqlc.UpdateCustomerParams {
	updatecustomerParam := &sqlc.UpdateCustomerParams{
		ID:          int32(customer.ID().Value()),
		Photo:       customer.Photo(),
		Gender:      models.PersonGender(customer.Gender()),
		DateOfBirth: pgtype.Date{Time: customer.DateOfBirth(), Valid: true},
		FirstName:   customer.FullName().FirstName,
		LastName:    customer.FullName().LastName,
		IsActive:    customer.IsActive(),
	}

	if customer.UserID() != nil {
		updatecustomerParam.UserID = pgtype.Int4{Int32: int32(customer.UserID().Value())}
	}

	return updatecustomerParam
}

func sqlRowToCustomer(row sqlc.Customer, pets []pet.Pet) (customer.Customer, error) {
	errorMessages := make([]string, 0)

	fullName, err := valueobject.NewPersonName(row.FirstName, row.LastName)
	if err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	customerID := valueobject.NewCustomerID(uint(row.ID))
	userID := valueobject.NewUserID(uint(row.UserID.Int32))

	if len(errorMessages) > 0 {
		return customer.Customer{}, appError.MappingError(errorMessages, "sql", "domain-entity", "customer")
	}

	gender, err := enum.ParseGender(string(row.Gender))
	if err != nil {
		return customer.Customer{}, err
	}
	customerEntity, err := customer.NewCustomer(
		customerID,
		customer.WithFullName(fullName),
		customer.WithDateOfBirth(row.DateOfBirth.Time),
		customer.WithUserID(&userID),
		customer.WithGender(gender),
		customer.WithIsActive(row.IsActive),
		customer.WithPets(pets),
	)
	if err != nil {
		return customer.Customer{}, err
	}

	return *customerEntity, nil
}

func ListRowToCustomer(rows []sqlc.Customer) ([]customer.Customer, error) {
	if len(rows) == 0 {
		return []customer.Customer{}, nil
	}

	customers := make([]customer.Customer, 0, len(rows))
	for i, row := range rows {
		o, err := sqlRowToCustomer(row, []pet.Pet{})
		if err != nil {
			return []customer.Customer{}, err
		}
		customers[i] = o
	}
	return customers, nil
}
