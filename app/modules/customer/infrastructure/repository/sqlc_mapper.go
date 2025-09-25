package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func toCreateParams(customer customer.Customer) *sqlc.CreateCustomerParams {
	var userID pgtype.Int4
	if customer.UserID() != nil {
		userID = pgtype.Int4{Int32: int32(customer.UserID().Value()), Valid: true}
	} else {
		userID = pgtype.Int4{Valid: false}
	}

	return &sqlc.CreateCustomerParams{
		Photo:       customer.Photo(),
		DateOfBirth: pgtype.Date{Time: customer.DateOfBirth(), Valid: true},
		Gender:      models.PersonGender(customer.Gender().String()),
		FirstName:   customer.Name().FirstName,
		LastName:    customer.Name().LastName,
		IsActive:    customer.IsActive(),
		UserID:      userID,
	}
}

func entityToUpdateParams(customer customer.Customer) *sqlc.UpdateCustomerParams {
	var userID pgtype.Int4
	if customer.UserID() != nil {
		userID = pgtype.Int4{Int32: int32(customer.UserID().Value()), Valid: true}
	} else {
		userID = pgtype.Int4{Valid: false}
	}

	return &sqlc.UpdateCustomerParams{
		ID:          int32(customer.ID().Value()),
		Photo:       customer.Photo(),
		Gender:      models.PersonGender(customer.Gender()),
		DateOfBirth: pgtype.Date{Time: customer.DateOfBirth(), Valid: true},
		FirstName:   customer.Name().FirstName,
		LastName:    customer.Name().LastName,
		IsActive:    customer.IsActive(),
		UserID:      userID,
	}
}

func sqlRowToCustomer(row sqlc.Customer, pets []pet.Pet) (customer.Customer, error) {
	fullName := valueobject.NewPersonNameNoErr(row.FirstName, row.LastName)
	customerID := valueobject.NewCustomerID(uint(row.ID))
	userID := valueobject.NewUserID(uint(row.UserID.Int32))

	customerEntity := customer.NewCustomer(
		customerID,
		customer.WithFullName(fullName),
		customer.WithDateOfBirth(row.DateOfBirth.Time),
		customer.WithUserID(&userID),
		customer.WithGender(enum.PersonGender(string(row.Gender))),
		customer.WithIsActive(row.IsActive),
		customer.WithPets(pets),
	)

	return *customerEntity, nil
}

func sqlRowsToCustomers(rows []sqlc.Customer) ([]customer.Customer, error) {
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
