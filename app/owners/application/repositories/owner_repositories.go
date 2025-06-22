package ownerRepository

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
)

type OwnerRepository interface {
	Save(arg *ownerDomain.Owner) error
	GetByID(ctx context.Context, id uint) (ownerDomain.Owner, error)
	GetByUserID(ownerId uint) (ownerDomain.Owner, error)
	Delete(ctx context.Context, id uint) error
	Exists(ctx context.Context, ownerId uint) (bool, error)

	//ListAppointmentsByOwner(ctx context.Context, ownerID uint) ([]sqlc.Appointment, error)
}
