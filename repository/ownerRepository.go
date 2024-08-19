package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type OwnerRepository interface {
	CreateOwner(ctx context.Context, arg sqlc.CreateOwnerParams) (sqlc.Owner, error)
	GetOwnerByID(ctx context.Context, id int32) (sqlc.Owner, error)
	UpdateOwner(ctx context.Context, arg sqlc.UpdateOwnerParams) error
	DeleteOwner(ctx context.Context, id int32) error
	ValidateExistingOwner(ctx context.Context, ownerId int32) bool

	GetAppointmentsByOwner(ctx context.Context, ownerID int32) ([]sqlc.Appointment, error)
}

type OwnerRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewOwnerRepositoryImpl(queries *sqlc.Queries) OwnerRepository {
	return &OwnerRepositoryImpl{queries: queries}
}

func (r *OwnerRepositoryImpl) CreateOwner(ctx context.Context, arg sqlc.CreateOwnerParams) (sqlc.Owner, error) {
	owner, err := r.queries.CreateOwner(ctx, arg)
	if err != nil {
		return sqlc.Owner{}, err
	}
	return owner, nil
}

func (r *OwnerRepositoryImpl) GetOwnerByID(ctx context.Context, ownerId int32) (sqlc.Owner, error) {
	owner, err := r.queries.GetOwnerByID(ctx, ownerId)
	if err != nil {
		return sqlc.Owner{}, err
	}
	return owner, nil
}

func (r *OwnerRepositoryImpl) UpdateOwner(ctx context.Context, arg sqlc.UpdateOwnerParams) error {
	err := r.queries.UpdateOwner(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *OwnerRepositoryImpl) DeleteOwner(ctx context.Context, ownerId int32) error {
	err := r.queries.DeleteOwner(ctx, ownerId)
	if err != nil {
		return err
	}
	return nil
}

func (r *OwnerRepositoryImpl) GetAppointmentsByOwner(ctx context.Context, ownerID int32) ([]sqlc.Appointment, error) {
	return nil, nil
}

func (r *OwnerRepositoryImpl) ValidateExistingOwner(ctx context.Context, ownerId int32) bool {
	_, err := r.queries.GetOwnerByID(ctx, ownerId)
	return err == nil
}
