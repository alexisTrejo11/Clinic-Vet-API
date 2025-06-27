package vetRepo

import (
	"context"

	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type VeterinarianRepository interface {
	List(ctx context.Context, searchParams map[string]interface{}) ([]vetDomain.Veterinarian, error)
	GetByID(ctx context.Context, id uint) (vetDomain.Veterinarian, error)
	GetByUserID(ctx context.Context, id uint) (vetDomain.Veterinarian, error)
	Save(ctx context.Context, pet *vetDomain.Veterinarian) error
	Delete(ctx context.Context, id uint, isSoftDelete bool) error
	Exists(ctx context.Context, vetId uint) (bool, error)
}
