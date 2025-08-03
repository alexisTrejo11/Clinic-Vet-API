package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	ActivateUser(ctx context.Context, id int32) error
	CreateOwner(ctx context.Context, arg CreateOwnerParams) (CreateOwnerRow, error)
	CreatePet(ctx context.Context, arg CreatePetParams) (Pet, error)
	CreateVeterinarian(ctx context.Context, arg CreateVeterinarianParams) (Veterinarian, error)
	DeactivateUser(ctx context.Context, id int32) error
	DeleteOwner(ctx context.Context, id int32) error
	DeletePet(ctx context.Context, id int32) error
	ExistByID(ctx context.Context, id int32) (bool, error)
	ExistByPhoneNumber(ctx context.Context, phoneNumber string) (bool, error)
	GetOwnerByID(ctx context.Context, id int32) (GetOwnerByIDRow, error)
	GetOwnerByPhone(ctx context.Context, phoneNumber string) (GetOwnerByPhoneRow, error)
	GetOwnerByUserID(ctx context.Context, userID pgtype.Int4) (GetOwnerByUserIDRow, error)
	GetPetByID(ctx context.Context, id int32) (Pet, error)
	GetPetsByOwnerID(ctx context.Context, ownerID int32) ([]Pet, error)
	GetVeterinarianById(ctx context.Context, id int32) (Veterinarian, error)
	ListOwners(ctx context.Context, arg ListOwnersParams) ([]ListOwnersRow, error)
	ListPets(ctx context.Context) ([]Pet, error)
	ListVeterinarians(ctx context.Context, arg ListVeterinariansParams) ([]Veterinarian, error)
	SoftDeleteVeterinarian(ctx context.Context, id int32) error
	UpdateOwner(ctx context.Context, arg UpdateOwnerParams) error
	UpdatePet(ctx context.Context, arg UpdatePetParams) error
	UpdateVeterinarian(ctx context.Context, arg UpdateVeterinarianParams) (Veterinarian, error)
	WithTx(tx pgx.Tx) *Queries
}

var _ Querier = (*Queries)(nil)
