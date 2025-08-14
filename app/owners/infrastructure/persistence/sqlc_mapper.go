package ownerRepository

import (
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func toCreateParams(owner ownerDomain.Owner) *sqlc.CreateOwnerParams {
	createOwnerParam := &sqlc.CreateOwnerParams{
		Photo:       owner.Photo(),
		PhoneNumber: owner.PhoneNumber(),
		Gender:      models.PersonGender(shared.AssertString(owner)),
		DateOfBirth: pgtype.Date{Time: owner.DateOfBirth(), Valid: true},
		FirstName:   owner.FullName().FirstName,
		LastName:    owner.FullName().LastName,
		IsActive:    owner.IsActive(),
	}

	if owner.Address != nil {
		createOwnerParam.Address = pgtype.Text{String: *owner.Address(), Valid: true}
	}

	if owner.UserId() != nil {
		createOwnerParam.UserID = pgtype.Int4{Int32: int32(*owner.UserId())}
	}

	return createOwnerParam
}

func toUpdateParams(owner ownerDomain.Owner) *sqlc.UpdateOwnerParams {
	updateOwnerParam := &sqlc.UpdateOwnerParams{
		ID:          int32(owner.Id()),
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

	if owner.UserId() != nil {
		updateOwnerParam.UserID = pgtype.Int4{Int32: int32(*owner.UserId())}
	}

	return updateOwnerParam
}

func rowToOwner(row sqlc.Owner) ownerDomain.Owner {
	fullName, _ := user.NewPersonName(row.FirstName, row.LastName)

	owner := ownerDomain.NewOwnerBuilder().
		WithId(int(row.ID)).
		WithFullName(fullName).
		WithPhoneNumber(row.PhoneNumber).
		WithAddress(row.Address.String).
		WithDateOfBirth(row.DateOfBirth.Time).
		WithUserId(int(row.UserID.Int32)).
		WithGender(user.Gender(shared.AssertString(row.Gender))).
		WithIsActive(row.IsActive).
		Build()

	return *owner
}

func ListRowToOwner(rows []sqlc.Owner) ([]ownerDomain.Owner, error) {
	owners := make([]ownerDomain.Owner, 0, len(rows))
	for i, row := range rows {
		owner := rowToOwner(row)
		owners[i] = owner
	}
	return owners, nil
}
