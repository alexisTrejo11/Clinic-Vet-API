package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type VetRepository interface {
	List(ctx context.Context, searchParams interface{}) ([]entity.Veterinarian, error)
	GetByID(ctx context.Context, id valueobject.VetID) (entity.Veterinarian, error)
	Exists(ctx context.Context, id valueobject.VetID) (bool, error)
	GetByUserID(ctx context.Context, userID valueobject.UserID) (entity.Veterinarian, error)
	SoftDelete(ctx context.Context, id valueobject.VetID) error
	Save(ctx context.Context, vet *entity.Veterinarian) error
}
