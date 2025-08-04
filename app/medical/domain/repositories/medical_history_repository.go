package medHistRepo

import (
	"context"

	mhDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/dtos"
	mhDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MedicalHistoryRepository interface {
	GetById(ctx context.Context, medicalHistoryId int) (*mhDomain.MedicalHistory, error)
	Search(ctx context.Context, searchParams mhDTOs.MedHistSearchParams) (*page.Page[[]mhDomain.MedicalHistory], error)
	ListByVetId(ctx context.Context, vetId int, pagintation page.PageData) (*page.Page[[]mhDomain.MedicalHistory], error)
	ListByPetId(ctx context.Context, petId int, pagintation page.PageData) (*page.Page[[]mhDomain.MedicalHistory], error)
	ListByOwnerId(ctx context.Context, ownerId int, pagintation page.PageData) (*page.Page[[]mhDomain.MedicalHistory], error)
	Create(ctx context.Context, medHistory *mhDomain.MedicalHistory) error
	Update(ctx context.Context, medHistory *mhDomain.MedicalHistory) error
	Delete(ctx context.Context, medicalHistoryId int, softDelete bool) error
}
