package sqlcOwnerRepository

import (
	"time"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToCreateParams(owner ownerDomain.Owner) *sqlc.CreateOwnerParams {
	createOwnerParam := &sqlc.CreateOwnerParams{
		Photo:       owner.Photo,
		PhoneNumber: owner.PhoneNumber,
		Gender:      owner.Gender,
		DateOfBirth: pgtype.Date{Time: time.Date(owner.DateOfBirth.Year(), owner.DateOfBirth.Month(), owner.DateOfBirth.Day(), 0, 0, 0, 0, time.UTC), Valid: true},
		FirstName:   owner.FullName.FirstName,
		LastName:    owner.FullName.LastName,
		IsActive:    owner.IsActive,
	}

	if owner.Address != nil {
		createOwnerParam.Address = pgtype.Text{String: *owner.Address, Valid: true}
	}

	if owner.UserId != nil {
		createOwnerParam.UserID = pgtype.Int4{Int32: int32(*owner.UserId)}
	}

	return createOwnerParam
}

func ToUpdateParams(owner ownerDomain.Owner) *sqlc.UpdateOwnerParams {
	updateParams := &sqlc.UpdateOwnerParams{
		ID:          int32(owner.Id),
		Photo:       owner.Photo,
		PhoneNumber: owner.PhoneNumber,
		Gender:      owner.Gender,
		FirstName:   owner.FullName.FirstName,
		DateOfBirth: pgtype.Date{Time: time.Date(owner.DateOfBirth.Year(), owner.DateOfBirth.Month(), owner.DateOfBirth.Day(), 0, 0, 0, 0, time.UTC), Valid: true},
		LastName:    owner.FullName.LastName,
		IsActive:    owner.IsActive,
	}
	if owner.Address != nil {
		updateParams.Address = pgtype.Text{String: *owner.Address, Valid: true}
	}

	if owner.UserId != nil {
		updateParams.UserID = pgtype.Int4{Int32: int32(*owner.UserId)}
	}

	return updateParams
}

func FromCreateToDomain(row sqlc.CreateOwnerRow) *ownerDomain.Owner {
	ownerName, err := userDomain.NewPersonName(row.FirstName, row.LastName)
	if err != nil {
		return &ownerDomain.Owner{}
	}

	return &ownerDomain.Owner{
		Id:          int(row.ID),
		Photo:       row.Photo,
		FullName:    ownerName,
		Gender:      userDomain.Gender(shared.AssertString(row.Gender)),
		PhoneNumber: row.PhoneNumber,
		DateOfBirth: row.DateOfBirth.Time,
		Address:     &row.Address.String,
		IsActive:    row.IsActive,
	}
}

func ListRowToOwner(rows []sqlc.ListOwnersRow) ([]ownerDomain.Owner, error) {
	owners := make([]ownerDomain.Owner, len(rows))
	if len(rows) == 0 {
		return []ownerDomain.Owner{}, nil

	}

	for i, row := range rows {
		ownerName, err := userDomain.NewPersonName(row.FirstName, row.LastName)
		if err != nil {
			return []ownerDomain.Owner{}, err
		}
		owner := ownerDomain.Owner{
			Id:          int(row.ID),
			Photo:       row.Photo,
			FullName:    ownerName,
			Gender:      userDomain.Gender(shared.AssertString(row.Gender)),
			PhoneNumber: row.PhoneNumber,
			Address:     &row.Address.String,
			IsActive:    row.IsActive,
		}

		owners[i] = owner
	}

	return owners, nil
}
