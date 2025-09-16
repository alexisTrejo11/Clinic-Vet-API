package repository

import (
	"context"

	"clinic-vet-api/app/core/domain/entity/pet"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type PetRepository interface {
	FindBySpecification(ctx context.Context, spec specification.PetSpecification) (page.Page[pet.Pet], error)
	FindAllByCustomerID(ctx context.Context, customerID valueobject.CustomerID) ([]pet.Pet, error)
	FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pageInput page.PageInput) (page.Page[pet.Pet], error)
	FindByID(ctx context.Context, petID valueobject.PetID) (pet.Pet, error)
	FindByIDAndCustomerID(ctx context.Context, id valueobject.PetID, customerID valueobject.CustomerID) (pet.Pet, error)

	ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error)
	ExistsByMicrochip(ctx context.Context, microchip string) (bool, error)

	Save(ctx context.Context, pet *pet.Pet) error
	Update(ctx context.Context, pet *pet.Pet) error
	Delete(ctx context.Context, petID valueobject.PetID) error

	CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error)
}
