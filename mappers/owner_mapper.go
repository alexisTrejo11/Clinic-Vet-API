package mappers

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapOwnerUpdateDtoToEntity(ownerUpdateDTO *DTOs.OwnerUpdateDTO, existingOwner sqlc.Owner) sqlc.UpdateOwnerParams {
	params := sqlc.UpdateOwnerParams{
		ID: ownerUpdateDTO.Id,
	}

	params.Name = coalesceString(ownerUpdateDTO.Name, existingOwner.Name)
	params.Photo = coalescePgText(ownerUpdateDTO.Photo, existingOwner.Photo)

	return params
}

func coalesceString(newVal, existingVal string) string {
	if newVal != "" {
		return newVal
	}
	return existingVal
}

func coalescePgText(newVal string, existingVal pgtype.Text) pgtype.Text {
	if newVal != "" {
		return pgtype.Text{String: newVal, Valid: true}
	}
	return existingVal
}
