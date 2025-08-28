package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
)

type PetRepository interface {
	List(ctx context.Context) ([]entity.Pet, error)
	ListByOwnerId(ctx context.Context, petId int) ([]entity.Pet, error)
	GetById(ctx context.Context, petId int) (entity.Pet, error)
	ExistsByID(ctx context.Context, petId int) (bool, error)
	Save(ctx context.Context, pet *entity.Pet) error
	Delete(ctx context.Context, petId int) error
}
