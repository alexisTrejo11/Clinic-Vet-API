package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
)

func (r *SqlcCustomerRepository) toCreateParams(customer customer.Customer) *sqlc.CreateCustomerParams {
	return &sqlc.CreateCustomerParams{
		Photo:       customer.Photo(),
		DateOfBirth: r.mapper.TimeToPgDate(customer.DateOfBirth()),
		Gender:      models.PersonGender(customer.Gender().String()),
		FirstName:   customer.FirstName(),
		LastName:    customer.LastName(),
		IsActive:    customer.IsActive(),
		UserID:      r.mapper.UserIDPtrToInt32(customer.UserID()),
	}
}

func (r *SqlcCustomerRepository) ToUpdateParams(customer customer.Customer) *sqlc.UpdateCustomerParams {
	return &sqlc.UpdateCustomerParams{
		ID:          customer.ID().Int32(),
		Photo:       customer.Photo(),
		Gender:      models.PersonGender(customer.Gender()),
		DateOfBirth: r.mapper.TimeToPgDate(customer.DateOfBirth()),
		FirstName:   customer.FirstName(),
		LastName:    customer.LastName(),
		IsActive:    customer.IsActive(),
		UserID:      r.mapper.UserIDPtrToInt32(customer.UserID()),
	}
}

func (r *SqlcCustomerRepository) sqlcRowToEntityWithPets(row sqlc.Customer, pets []pet.Pet) customer.Customer {
	return *customer.NewCustomerBuilder().
		WithName(valueobject.NewPersonName(row.FirstName, row.LastName)).
		WithDateOfBirth(row.DateOfBirth.Time).
		WithUserID(r.mapper.Int32ToUserIDPtr(row.UserID.Int32)).
		WithGender(enum.PersonGender(string(row.Gender))).
		WithIsActive(row.IsActive).
		WithPets(pets).
		WithPhoto(row.Photo).
		WithID(valueobject.NewCustomerID(uint(row.ID))).
		Build()
}

func (r *SqlcCustomerRepository) sqlRowsToEntities(rows []sqlc.Customer) []customer.Customer {
	if len(rows) == 0 {
		return []customer.Customer{}
	}

	customers := make([]customer.Customer, 0, len(rows))
	for i, row := range rows {
		o := r.sqlcRowToEntityWithPets(row, []pet.Pet{})
		customers[i] = o
	}
	return customers
}
