package persistence

/*
import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type OwnerRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewOwnerRepositoryImpl(queries *sqlc.Queries) OwnerRepository {
	return &OwnerRepositoryImpl{queries: queries}
}

func (r *OwnerRepositoryImpl) Create(arg sqlc.CreateOwnerParams) (sqlc.Owner, error) {
	owner, err := r.queries.CreateOwner(context.Background(), arg)
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

func (r *OwnerRepositoryImpl) GetOwnerByUserID(ownerId int32) (*sqlc.Owner, error) {
	owner, err := r.queries.GetOwnerByUserID(context.Background(), pgtype.Int4{Int32: ownerId, Valid: true})
	if err != nil {
		return nil, err
	}
	return &owner, nil
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
*/
