package owner_repository

import (
	"context"

	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repositories"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type OwnerRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewOwnerRepositoryImpl(queries *sqlc.Queries) ownerRepository.OwnerRepository {
	return &OwnerRepositoryImpl{queries: queries}
}

func (r *OwnerRepositoryImpl) Save(owner *ownerDomain.Owner) error {
	return nil
}

func (r *OwnerRepositoryImpl) GetByID(ctx context.Context, ownerId uint) (ownerDomain.Owner, error) {
	_, err := r.queries.GetOwnerByID(ctx, int32(ownerId))
	if err != nil {
		return ownerDomain.Owner{}, err
	}
	return ownerDomain.Owner{}, nil
}

func (r *OwnerRepositoryImpl) GetByUserID(ownerId uint) (ownerDomain.Owner, error) {
	_, err := r.queries.GetOwnerByUserID(context.Background(), pgtype.Int4{Int32: int32(ownerId), Valid: true})
	if err != nil {
		return ownerDomain.Owner{}, err
	}
	return ownerDomain.Owner{}, nil
}

func (r *OwnerRepositoryImpl) Delete(ctx context.Context, ownerId uint) error {
	err := r.queries.DeleteOwner(ctx, int32(ownerId))
	if err != nil {
		return err
	}
	return nil
}

func (r *OwnerRepositoryImpl) Exists(ctx context.Context, ownerId uint) (bool, error) {
	_, err := r.queries.GetOwnerByID(ctx, int32(ownerId))
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, nil
		}
		return false, err
	}
	return true, err
}
