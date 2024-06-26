package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type OwnerRepositoryI interface {
	CreateOwner(ctx context.Context, arg sqlc.CreateOwnerParams) (sqlc.Owner, error)
	GetOwnerByID(ctx context.Context, id int32) (sqlc.Owner, error)
	ListOwners(ctx context.Context) ([]sqlc.Owner, error)
	UpdateOwner(ctx context.Context, arg sqlc.UpdateOwnerParams) error
	DeleteOwner(ctx context.Context, id int32) error

	GetAppointmentsByOwner(ctx context.Context, ownerID int32) ([]sqlc.Appointment, error)
}

type OwnerRepository struct {
	queries *sqlc.Queries
}

func NewOwnerRepository(queries *sqlc.Queries) *OwnerRepository {
	return &OwnerRepository{queries: queries}
}

func (r *OwnerRepository) CreateOwner(ctx context.Context, arg sqlc.CreateOwnerParams) (sqlc.Owner, error) {
	owner, err := r.queries.CreateOwner(ctx, arg)
	if err != nil {
		return sqlc.Owner{}, err
	}
	return owner, nil
}

func (r *OwnerRepository) GetOwnerById(ctx context.Context, id int32) (sqlc.Owner, error) {
	owner, err := r.queries.GetOwnerByID(ctx, id)
	if err != nil {
		return sqlc.Owner{}, err
	}
	return owner, nil
}
