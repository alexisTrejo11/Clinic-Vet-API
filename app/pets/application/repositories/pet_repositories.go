package petRepo

import (
	"context"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type PetRepository interface {
	List(ctx context.Context) ([]petDomain.Pet, error)
	ListByOwnerId(ctx context.Context, petId uint) ([]petDomain.Pet, error)
	GetById(ctx context.Context, petId uint) (petDomain.Pet, error)
	Save(ctx context.Context, pet *petDomain.Pet) error
	Delete(ctx context.Context, petId uint) error
}
