package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type PetRepository interface {
	List(ctx context.Context) ([]entity.Pet, error)
	ListByOwnerID(ctx context.Context, petID valueobject.OwnerID) ([]entity.Pet, error)
	GetByID(ctx context.Context, petID valueobject.PetID) (entity.Pet, error)
	ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error)
	Save(ctx context.Context, pet *entity.Pet) error
	Delete(ctx context.Context, petID valueobject.PetID) error
}
