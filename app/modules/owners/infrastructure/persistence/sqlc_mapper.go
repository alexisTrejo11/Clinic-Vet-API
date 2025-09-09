package persistence

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/owner"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"

	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func toCreateParams(owner owner.Owner) *sqlc.CreateOwnerParams {
	createOwnerParam := &sqlc.CreateOwnerParams{
		Photo:       owner.Photo(),
		PhoneNumber: owner.PhoneNumber(),
		DateOfBirth: pgtype.Date{Time: owner.DateOfBirth(), Valid: true},
		Gender:      models.PersonGender(owner.Gender().String()),
		FirstName:   owner.FullName().FirstName,
		LastName:    owner.FullName().LastName,
		IsActive:    owner.IsActive(),
	}

	if owner.UserID() != nil {
		createOwnerParam.UserID = pgtype.Int4{Int32: int32(owner.UserID().Value())}
	}

	return createOwnerParam
}

func entityToUpdateParams(owner owner.Owner) *sqlc.UpdateOwnerParams {
	updateOwnerParam := &sqlc.UpdateOwnerParams{
		ID:          int32(owner.ID().Value()),
		Photo:       owner.Photo(),
		PhoneNumber: owner.PhoneNumber(),
		Gender:      models.PersonGender(owner.Gender()),
		DateOfBirth: pgtype.Date{Time: owner.DateOfBirth(), Valid: true},
		FirstName:   owner.FullName().FirstName,
		LastName:    owner.FullName().LastName,
		IsActive:    owner.IsActive(),
	}

	if owner.UserID() != nil {
		updateOwnerParam.UserID = pgtype.Int4{Int32: int32(owner.UserID().Value())}
	}

	return updateOwnerParam
}

func sqlRowToOwner(row sqlc.Owner, pets []pet.Pet) (owner.Owner, error) {
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
		return owner.Owner{}, appError.MappingError(errorMessages, "sql", "domain-entity", "owner")
	}

	gender, err := enum.ParseGender(string(row.Gender))
	if err != nil {
		return owner.Owner{}, err
	}
	ownerEntity, err := owner.NewOwner(
		ownerID,
		owner.WithFullName(fullName),
		owner.WithPhoneNumber(row.PhoneNumber),
		owner.WithDateOfBirth(row.DateOfBirth.Time),
		owner.WithUserID(&userID),
		owner.WithGender(gender),
		owner.WithIsActive(row.IsActive),
		owner.WithPets(pets),
	)
	if err != nil {
		return owner.Owner{}, err
	}

	return *ownerEntity, nil
}

func ListRowToOwner(rows []sqlc.Owner) ([]owner.Owner, error) {
	if len(rows) == 0 {
		return []owner.Owner{}, nil
	}

	owners := make([]owner.Owner, 0, len(rows))
	for i, row := range rows {
		o, err := sqlRowToOwner(row, []pet.Pet{})
		if err != nil {
			return []owner.Owner{}, err
		}
		owners[i] = o
	}
	return owners, nil
}
