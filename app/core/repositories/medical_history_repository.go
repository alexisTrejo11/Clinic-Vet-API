package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MedicalHistoryRepository interface {
	GetByID(ctx context.Context, medicalHistoryID int) (entity.MedicalHistory, error)
	Search(ctx context.Context, searchParams interface{}) (page.Page[[]entity.MedicalHistory], error)
	ListByVetID(ctx context.Context, vetID int, pagintation page.PageInput) (page.Page[[]entity.MedicalHistory], error)
	ListByPetID(ctx context.Context, petID int, pagintation page.PageInput) (page.Page[[]entity.MedicalHistory], error)
	ListByOwnerID(ctx context.Context, ownerID int, pagintation page.PageInput) (page.Page[[]entity.MedicalHistory], error)
	Create(ctx context.Context, medHistory *entity.MedicalHistory) error
	Update(ctx context.Context, medHistory *entity.MedicalHistory) error
	Delete(ctx context.Context, medicalHistoryID int, softDelete bool) error
}
