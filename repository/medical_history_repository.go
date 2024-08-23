package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type MedicalHistoryRepository interface {
	CreateMedicalHistory(params sqlc.CreateMedicalHistoryParams) error
	GetMedicalHistoryByID(MedicalHistoryId int32) (*sqlc.MedicalHistory, error)
	UpdateMedicalHistory(updateParams sqlc.UpdateMedicalHistoryParams) error
	DeleteMedicalHistory(MedicalHistoryId int32) error
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

func (ar MedicalHistoryRepositoryImpl) GetMedicalHistoryByID(MedicalHistoryId int32) (*sqlc.MedicalHistory, error) {
	MedicalHistory, err := ar.queries.GetMedicalHistoryByID(context.Background(), MedicalHistoryId)
	if err != nil {
		return nil, err
	}

	return &MedicalHistory, nil
}

func (ar MedicalHistoryRepositoryImpl) UpdateMedicalHistory(updateParams sqlc.UpdateMedicalHistoryParams) error {
	err := ar.queries.UpdateMedicalHistory(context.Background(), updateParams)
	if err != nil {
		return err
	}

	return nil
}

func (ar MedicalHistoryRepositoryImpl) DeleteMedicalHistory(MedicalHistoryId int32) error {
	if err := ar.queries.DeleteMedicalHistory(context.Background(), MedicalHistoryId); err != nil {
		return err
	}

	return nil
}
