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

func sqlRowToCustomer(row sqlc.Customer, pets []pet.Pet) customer.Customer {
	userID := valueobject.NewUserID(uint(row.UserID.Int32))
	customerEntity := customer.NewCustomerBuilder().
		WithName(valueobject.NewPersonNameNoErr(row.FirstName, row.LastName)).
		WithDateOfBirth(row.DateOfBirth.Time).
		WithUserID(&userID).
		WithGender(enum.PersonGender(string(row.Gender))).
		WithIsActive(row.IsActive).
		WithPets(pets).
		WithPhoto(row.Photo).
		WithID(valueobject.NewCustomerID(uint(row.ID))).
		Build()

	return *customerEntity
}

func sqlRowsToCustomers(rows []sqlc.Customer) []customer.Customer {
	if len(rows) == 0 {
		return []customer.Customer{}
	}

	customers := make([]customer.Customer, 0, len(rows))
	for i, row := range rows {
		o := sqlRowToCustomer(row, []pet.Pet{})
		customers[i] = o
	}
	return customers
}
