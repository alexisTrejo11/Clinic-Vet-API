package persistence

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
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

func toUpdateParams(owner entity.Owner) *sqlc.UpdateOwnerParams {
	updateOwnerParam := &sqlc.UpdateOwnerParams{
		ID:          int32(owner.ID().GetValue()),
		Photo:       owner.Photo(),
		PhoneNumber: owner.PhoneNumber(),
		Gender:      models.PersonGender(shared.AssertString(owner)),
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

func rowToOwner(row sqlc.Owner) entity.Owner {
	fullName, _ := valueObjects.NewPersonName(row.FirstName, row.LastName)
	ownerID, _ := valueobject.NewOwnerID(int(row.ID))
	userID, _ := valueobject.NewUserID(int(row.UserID.Int32))

	owner := entity.NewOwnerBuilder().
		WithID(ownerID).
		WithFullName(fullName).
		WithPhoneNumber(row.PhoneNumber).
		WithAddress(row.Address.String).
		WithDateOfBirth(row.DateOfBirth.Time).
		WithUserID(userID).
		WithGender(valueObjects.Gender(shared.AssertString(row.Gender))).
		WithIsActive(row.IsActive).
		Build()

	return *owner
}

func ListRowToOwner(rows []sqlc.Owner) ([]entity.Owner, error) {
	owners := make([]entity.Owner, 0, len(rows))
	for i, row := range rows {
		owner := rowToOwner(row)
		owners[i] = owner
	}
	return owners, nil
}
