package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type PetRepository interface {
	Search(ctx context.Context, searchParmas any) (page.Page[[]pet.Pet], error)
	ListByCustomerID(ctx context.Context, petID valueobject.CustomerID) (page.Page[[]pet.Pet], error)
	GetByIDAndCustomerID(ctx context.Context, id valueobject.PetID, petID valueobject.CustomerID) (pet.Pet, error)
	GetByID(ctx context.Context, petID valueobject.PetID) (pet.Pet, error)
	ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error)
	Save(ctx context.Context, pet *pet.Pet) error
	Delete(ctx context.Context, petID valueobject.PetID) error
}
