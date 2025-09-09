package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type PetRepository interface {
	Search(ctx context.Context, searchParmas any) ([]pet.Pet, error)
	ListByOwnerID(ctx context.Context, petID valueobject.OwnerID) ([]pet.Pet, error)
	GetByIDAndOwnerID(ctx context.Context, id valueobject.PetID, petID valueobject.OwnerID) (pet.Pet, error)
	GetByID(ctx context.Context, petID valueobject.PetID) (pet.Pet, error)
	ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error)
	Save(ctx context.Context, pet *pet.Pet) error
	Delete(ctx context.Context, petID valueobject.PetID) error
}
