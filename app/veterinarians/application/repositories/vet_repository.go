package vetRepo

import (
	"context"

	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type VeterinarianRepository interface {
	List(ctx context.Context, searchParams vetDtos.VetSearchParams) ([]vetDomain.Veterinarian, error)
	GetByID(ctx context.Context, id int) (vetDomain.Veterinarian, error)
	GetByUserID(ctx context.Context, id int) (vetDomain.Veterinarian, error)
	Save(ctx context.Context, pet *vetDomain.Veterinarian) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, vetId int) (bool, error)
}
