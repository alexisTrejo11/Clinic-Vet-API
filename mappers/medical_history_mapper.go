package mappers

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type MedicalHistoryMappers struct {
}

func (MedicalHistoryMappers) InsertDtoCreateParams(medicalHistoryInsertDTO DTOs.MedicalHistoryInsertDTO) sqlc.CreateMedicalHistoryParams {
	return sqlc.CreateMedicalHistoryParams{
		PetID:       medicalHistoryInsertDTO.PetID,
		Date:        pgtype.Timestamp{Time: medicalHistoryInsertDTO.Date, Valid: true},
		Description: pgtype.Text{String: medicalHistoryInsertDTO.Description, Valid: true},
		VetID:       medicalHistoryInsertDTO.VetID,
	}
}

func (MedicalHistoryMappers) SqlcEntityToDTO(medicalHistory sqlc.MedicalHistory) DTOs.MedicalHistoryDTO {
	return DTOs.MedicalHistoryDTO{
		PetID:       medicalHistory.PetID,
		Date:        medicalHistory.Date.Time,
		Description: medicalHistory.Description.String,
		VetID:       medicalHistory.VetID,
	}
}

func (MedicalHistoryMappers) UpdateDtoUpdateParams(medicalHistoryUpdateDTO DTOs.MedicalHistoryUpdateDTO) sqlc.UpdateMedicalHistoryParams {
	return sqlc.UpdateMedicalHistoryParams{
		PetID:       medicalHistoryUpdateDTO.PetID,
		Date:        pgtype.Timestamp{Time: medicalHistoryUpdateDTO.Date, Valid: true},
		Description: pgtype.Text{String: medicalHistoryUpdateDTO.Description, Valid: true},
		VetID:       medicalHistoryUpdateDTO.VetID,
	}
}
