package repository

import (
	"context"

	vet "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/veterinarian"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type VetRepository interface {
	List(ctx context.Context, searchParams any) ([]vet.Veterinarian, error)
	GetByID(ctx context.Context, id valueobject.VetID) (vet.Veterinarian, error)
	Exists(ctx context.Context, id valueobject.VetID) (bool, error)
	GetByUserID(ctx context.Context, userID valueobject.UserID) (vet.Veterinarian, error)
	SoftDelete(ctx context.Context, id valueobject.VetID) error
	Save(ctx context.Context, vet *vet.Veterinarian) error
}
