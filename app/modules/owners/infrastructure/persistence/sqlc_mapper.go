package persistence

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func toCreateParams(owner entity.Owner) *sqlc.CreateOwnerParams {
	createOwnerParam := &sqlc.CreateOwnerParams{
		Photo:       owner.Photo(),
		PhoneNumber: owner.PhoneNumber(),
		DateOfBirth: pgtype.Date{Time: owner.DateOfBirth(), Valid: true},
		Gender:      models.PersonGender(owner.Gender().String()),
		FirstName:   owner.FullName().FirstName,
		LastName:    owner.FullName().LastName,
		IsActive:    owner.IsActive(),
	}

	if owner.Address() != nil {
		createOwnerParam.Address = pgtype.Text{String: *owner.Address(), Valid: true}
	}

	if owner.UserID() != nil {
		createOwnerParam.UserID = pgtype.Int4{Int32: int32(owner.UserID().GetValue())}
	}

	return createOwnerParam
}

func entityToUpdateParams(owner entity.Owner) *sqlc.UpdateOwnerParams {
	updateOwnerParam := &sqlc.UpdateOwnerParams{
		ID:          int32(owner.ID().GetValue()),
		Photo:       owner.Photo(),
		PhoneNumber: owner.PhoneNumber(),
		Gender:      models.PersonGender(owner.Gender()),
		DateOfBirth: pgtype.Date{Time: owner.DateOfBirth(), Valid: true},
		FirstName:   owner.FullName().FirstName,
		LastName:    owner.FullName().LastName,
		IsActive:    owner.IsActive(),
	}

	if owner.Address() != nil {
		updateOwnerParam.Address = pgtype.Text{String: *owner.Address(), Valid: true}
	}

	if owner.UserID() != nil {
		updateOwnerParam.UserID = pgtype.Int4{Int32: int32(owner.UserID().GetValue())}
	}

	return updateOwnerParam
}

func sqlRowToOwner(row sqlc.Owner) (entity.Owner, error) {
	errorMessages := make([]string, 0)

	fullName, err := valueobject.NewPersonName(row.FirstName, row.LastName)
	if err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	ownerID, err := valueobject.NewOwnerID(int(row.ID))
	if err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	userID, err := valueobject.NewUserID(int(row.UserID.Int32))
	if err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if len(errorMessages) > 0 {
		return entity.Owner{}, appError.MappingError(errorMessages, "sql", "domainEntity", "owner")
	}

	gender := enum.NewGender(string(row.Gender))

	owner := entity.NewOwnerBuilder().
		WithID(ownerID).
		WithFullName(fullName).
		WithPhoneNumber(row.PhoneNumber).
		WithAddress(row.Address.String).
		WithDateOfBirth(row.DateOfBirth.Time).
		WithUserID(userID).
		WithGender(gender).
		WithIsActive(row.IsActive).
		Build()

	return *owner, nil
}

func ListRowToOwner(rows []sqlc.Owner) ([]entity.Owner, error) {
	owners := make([]entity.Owner, 0, len(rows))
	for i, row := range rows {
		owner, err := sqlRowToOwner(row)
		if err != nil {
			return []entity.Owner{}, err
		}
		owners[i] = owner
	}
	return owners, nil
}
