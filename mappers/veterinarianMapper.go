package mappers

import (
	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapVetInsertDtoToVetInsertParams(vetInsertDTO dtos.VetInsertDTO) sqlc.CreateVeterinarianParams {
	return sqlc.CreateVeterinarianParams{
		Name:      vetInsertDTO.Name,
		Photo:     pgtype.Text{String: vetInsertDTO.Photo, Valid: true},
		Email:     vetInsertDTO.Email,
		Specialty: pgtype.Text{String: vetInsertDTO.Specialty, Valid: true},
	}
}

func MapSqlcEntityToDTO(veterinarian sqlc.Veterinarian) dtos.VetDTO {
	return dtos.VetDTO{
		Id:        veterinarian.ID,
		Name:      veterinarian.Name,
		Photo:     veterinarian.Photo.String,
		Email:     veterinarian.Email,
		Specialty: veterinarian.Specialty.String,
	}
}

func MapVetUpdateDtoToEntity(vetUpdateDTO *dtos.VetUpdateDTO, existingVet sqlc.Veterinarian) sqlc.UpdateVeterinarianParams {
	params := sqlc.UpdateVeterinarianParams{
		ID: vetUpdateDTO.Id,
	}

	params.Name = coalesceString(vetUpdateDTO.Name, existingVet.Name)
	params.Photo = coalescePgText(vetUpdateDTO.Photo, existingVet.Photo)
	params.Email = coalesceString(vetUpdateDTO.Email, existingVet.Email)
	params.Specialty = coalescePgText(vetUpdateDTO.Specialty, existingVet.Specialty)

	return params
}
