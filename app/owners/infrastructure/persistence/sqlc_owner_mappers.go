package sqlcOwnerRepository

import (
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	userEnums "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/enum"
	userValueObjects "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/valueobjects"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToCreateParams(owner ownerDomain.Owner) *sqlc.CreateOwnerParams {
	return &sqlc.CreateOwnerParams{
		Photo:       owner.Photo,
		Phonenumber: owner.PhoneNumber,
		Gender:      sqlc.PersonGender(owner.Gender),
		Firstname:   owner.FullName.FirstName,
		Lastname:    owner.FullName.LastName,
		Address:     pgtype.Text{String: *owner.Address, Valid: true},
		UserID:      pgtype.Int4{Int32: int32(*owner.UserId)},
		Isactive:    owner.IsActive,
	}
}

func ToUpdateParams(owner ownerDomain.Owner) *sqlc.UpdateOwnerParams {
	return &sqlc.UpdateOwnerParams{
		ID:          int32(owner.Id),
		Photo:       owner.Photo,
		Phonenumber: owner.PhoneNumber,
		Gender:      sqlc.PersonGender(owner.Gender),
		Firstname:   owner.FullName.FirstName,
		Lastname:    owner.FullName.LastName,
		Address:     pgtype.Text{String: *owner.Address, Valid: true},
		UserID:      pgtype.Int4{Int32: int32(*owner.UserId)},
		Isactive:    owner.IsActive,
	}
}

func FromCreateToDomain(row sqlc.CreateOwnerRow) *ownerDomain.Owner {
	ownerName, err := userValueObjects.NewPersonName(row.Firstname, row.Lastname)
	if err != nil {
		return &ownerDomain.Owner{}
	}

	return &ownerDomain.Owner{
		Id:          uint(row.ID),
		Photo:       row.Phonenumber,
		FullName:    ownerName,
		Gender:      userEnums.Gender(row.Gender),
		PhoneNumber: row.Phonenumber,
		Address:     &row.Address.String,
		IsActive:    row.Isactive,
	}
}

func RowToOwner(ownerRow sqlc.GetOwnerByIDRow) (ownerDomain.Owner, error) {
	ownerName, err := userValueObjects.NewPersonName(ownerRow.Firstname, ownerRow.Lastname)
	if err != nil {
		return ownerDomain.Owner{}, nil
	}

	return ownerDomain.Owner{
		Id:          uint(ownerRow.ID),
		Photo:       ownerRow.Phonenumber,
		FullName:    ownerName,
		Gender:      userEnums.Gender(ownerRow.Gender),
		PhoneNumber: ownerRow.Phonenumber,
		Address:     &ownerRow.Address.String,
		IsActive:    ownerRow.Isactive,
	}, nil
}

func ListRowToOwner(rows []sqlc.ListOwnersRow) ([]ownerDomain.Owner, error) {
	owners := make([]ownerDomain.Owner, len(rows))
	if len(rows) == 0 {
		return []ownerDomain.Owner{}, nil

	}

	for i, row := range rows {
		ownerName, err := userValueObjects.NewPersonName(row.Firstname, row.Lastname)
		if err != nil {
			return []ownerDomain.Owner{}, err
		}
		owner := ownerDomain.Owner{
			Id:          uint(row.ID),
			Photo:       row.Phonenumber,
			FullName:    ownerName,
			Gender:      userEnums.Gender(row.Gender),
			PhoneNumber: row.Phonenumber,
			Address:     &row.Address.String,
			IsActive:    row.Isactive,
		}

		owners[i] = owner
	}

	return owners, nil
}
