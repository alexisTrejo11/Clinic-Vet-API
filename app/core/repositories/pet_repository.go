package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
)

type PetRepository interface {
	List(ctx context.Context) ([]entity.Pet, error)
	ListByOwnerID(ctx context.Context, petID int) ([]entity.Pet, error)
	GetByID(ctx context.Context, petID int) (entity.Pet, error)
	ExistsByID(ctx context.Context, petID int) (bool, error)
	Save(ctx context.Context, pet *entity.Pet) error
	Delete(ctx context.Context, petID int) error
}
