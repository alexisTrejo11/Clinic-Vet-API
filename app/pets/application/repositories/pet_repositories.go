package petRepo

import (
	"context"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type PetRepository interface {
	List(ctx context.Context) ([]petDomain.Pet, error)
	ListByOwnerId(ctx context.Context, petId int) ([]petDomain.Pet, error)
	GetById(ctx context.Context, petId int) (petDomain.Pet, error)
	ExistsById(ctx context.Context, petId int) (bool, error)
	Save(ctx context.Context, pet *petDomain.Pet) error
	Delete(ctx context.Context, petId int) error
}
