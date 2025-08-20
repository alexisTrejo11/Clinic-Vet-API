package petDomain

import (
	"context"
)

type PetRepository interface {
	List(ctx context.Context) ([]Pet, error)
	ListByOwnerId(ctx context.Context, petId int) ([]Pet, error)
	GetById(ctx context.Context, petId int) (Pet, error)
	ExistsById(ctx context.Context, petId int) (bool, error)
	Save(ctx context.Context, pet *Pet) error
	Delete(ctx context.Context, petId int) error
}
