package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type PetRepository interface {
	FindByCriteria(ctx context.Context, criteria map[string]any) (page.Page[[]pet.Pet], error)
	FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (page.Page[[]pet.Pet], error)
	FindByID(ctx context.Context, petID valueobject.PetID) (pet.Pet, error)
	FindByIDAndCustomerID(ctx context.Context, id valueobject.PetID, customerID valueobject.CustomerID) (pet.Pet, error)

	ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error)
	ExistsByMicrochip(ctx context.Context, microchip string) (bool, error)

	Save(ctx context.Context, pet *pet.Pet) error
	Update(ctx context.Context, pet *pet.Pet) error
	Delete(ctx context.Context, petID valueobject.PetID) error

	CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error)
}
