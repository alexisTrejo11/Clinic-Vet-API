package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/medical"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	p "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MedicalHistoryRepository interface {
	Search(ctx context.Context, searchCriteria any) (p.Page[[]medical.MedicalHistory], error)
	GetByID(ctx context.Context, medicalHistoryID valueobject.MedHistoryID) (medical.MedicalHistory, error)
	ListByVetID(ctx context.Context, vetID valueobject.VetID, pagintation p.PageInput) (p.Page[[]medical.MedicalHistory], error)
	ListByPetID(ctx context.Context, petID valueobject.PetID, pagintation p.PageInput) (p.Page[[]medical.MedicalHistory], error)
	ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID, pagintation p.PageInput) (p.Page[[]medical.MedicalHistory], error)
	Create(ctx context.Context, medHistory *medical.MedicalHistory) error
	Update(ctx context.Context, medHistory *medical.MedicalHistory) error
	Delete(ctx context.Context, medicalHistoryID valueobject.MedHistoryID, softDelete bool) error
}
