package mappers

/*
import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func MapVetInsertDtoToVetInsertParams(vetInsertDTO DTOs.VetInsertDTO) sqlc.CreateVeterinarianParams {
	return sqlc.CreateVeterinarianParams{
		Name:      vetInsertDTO.Name,
		Photo:     pgtype.Text{String: vetInsertDTO.Photo, Valid: true},
		Specialty: pgtype.Text{String: vetInsertDTO.Specialty, Valid: true},
	}
}

func MapSqlcEntityToDTO(veterinarian sqlc.Veterinarian) DTOs.VetDTO {
	return DTOs.VetDTO{
		Id:        veterinarian.ID,
		UserId:    &veterinarian.UserID.Int32,
		Name:      veterinarian.Name,
		Photo:     veterinarian.Photo.String,
		Specialty: veterinarian.Specialty.String,
	}
}

func MapVetUpdateDtoToEntity(vetUpdateDTO *DTOs.VetUpdateDTO, existingVet sqlc.Veterinarian) sqlc.UpdateVeterinarianParams {
	params := sqlc.UpdateVeterinarianParams{
		ID: vetUpdateDTO.Id,
	}

	params.Name = coalesceString(vetUpdateDTO.Name, existingVet.Name)
	params.Photo = coalescePgText(vetUpdateDTO.Photo, existingVet.Photo)
	params.Specialty = coalescePgText(vetUpdateDTO.Specialty, existingVet.Specialty)

	return params
}
*/
