package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type MedicalHistoryRepository interface {
	CreateMedicalHistory(params sqlc.CreateMedicalHistoryParams) error
	GetMedicalHistoryByID(medicalHistoryId int32) (*sqlc.MedicalHistory, error)
	GetMedicalHistoryByVetID(vetID int32) ([]sqlc.MedicalHistory, error)
	GetMedicalHistoryByPetID(petID int32) ([]sqlc.MedicalHistory, error)
	UpdateMedicalHistory(updateParams sqlc.UpdateMedicalHistoryParams) error
	DeleteMedicalHistory(medicalHistoryId int32) error
}

type MedicalHistoryRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewMedicalHistoryRepository(queries *sqlc.Queries) MedicalHistoryRepository {
	return &MedicalHistoryRepositoryImpl{
		queries: queries,
	}
}

func (ar MedicalHistoryRepositoryImpl) CreateMedicalHistory(params sqlc.CreateMedicalHistoryParams) error {
	if _, err := ar.queries.CreateMedicalHistory(context.Background(), params); err != nil {
		return err
	}

	return nil
}

func (ar MedicalHistoryRepositoryImpl) GetMedicalHistoryByID(medicalHistoryId int32) (*sqlc.MedicalHistory, error) {
	medicalHistory, err := ar.queries.GetMedicalHistoryByID(context.Background(), medicalHistoryId)
	if err != nil {
		return nil, err
	}

	return &medicalHistory, nil
}

func (ar MedicalHistoryRepositoryImpl) GetMedicalHistoryByVetID(vetID int32) ([]sqlc.MedicalHistory, error) {
	medicalHistory, err := ar.queries.ListMedicalHistoriesByVetID(context.Background(), vetID)
	if err != nil {
		return nil, err
	}

	return medicalHistory, nil
}

func (ar MedicalHistoryRepositoryImpl) GetMedicalHistoryByPetID(petID int32) ([]sqlc.MedicalHistory, error) {
	medicalHistory, err := ar.queries.ListMedicalHistoriesByPetID(context.Background(), petID)
	if err != nil {
		return nil, err
	}

	return medicalHistory, nil
}

func (ar MedicalHistoryRepositoryImpl) UpdateMedicalHistory(updateParams sqlc.UpdateMedicalHistoryParams) error {
	err := ar.queries.UpdateMedicalHistory(context.Background(), updateParams)
	if err != nil {
		return err
	}

	return nil
}

func (ar MedicalHistoryRepositoryImpl) DeleteMedicalHistory(medicalHistoryId int32) error {
	if err := ar.queries.DeleteMedicalHistory(context.Background(), medicalHistoryId); err != nil {
		return err
	}

	return nil
}
