package persistence

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCMedHistRepository struct {
	queries *sqlc.Queries
}

func NewSQLCMedHistRepository(queries *sqlc.Queries) repository.MedicalHistoryRepository {
	return &SQLCMedHistRepository{
		queries: queries,
	}
}

func (r *SQLCMedHistRepository) GetByID(ctx context.Context, medicalHistoryID int) (*entity.MedicalHistory, error) {
	sqlcMedHist, err := r.queries.GetMedicalHistoryByID(ctx, int32(medicalHistoryID))
	if err != nil {
		return nil, err
	}

	fmt.Printf("SQLCMedHist: %+v\n", sqlcMedHist)

	medHist, err := ToDomain(sqlcMedHist)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Domain MedHist: %+v\n", medHist)

	return &medHist, nil
}

func (r *SQLCMedHistRepository) Search(ctx context.Context, searchParams interface{}) (page.Page[[]entity.MedicalHistory], error) {
	searchParam := searchParams.(dto.MedHistSearchParams)
	queryRows, err := r.queries.SearchMedicalHistory(ctx, sqlc.SearchMedicalHistoryParams{})
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	medHistoryList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	metadata := *page.GetPageMetadata(len(queryRows), searchParam.Page)
	return page.NewPage(medHistoryList, metadata), nil
}

func (r *SQLCMedHistRepository) ListByVetID(ctx context.Context, vetID int, pagination page.PageData) (page.Page[[]entity.MedicalHistory], error) {
	arg := sqlc.ListMedicalHistoryByVetParams{
		VeterinarianID: int32(vetID),
		Limit:          int32(pagination.PageNumber),
		Offset:         int32(pagination.PageSize * (pagination.PageNumber - 1)),
	}

	queryRows, err := r.queries.ListMedicalHistoryByVet(ctx, arg)
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	entitiyList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	medicalHistPage := page.NewPage(entitiyList, *page.GetPageMetadata(len(queryRows), pagination))
	return medicalHistPage, nil
}

// TODO: Paginate SQLC
func (r *SQLCMedHistRepository) ListByPetID(ctx context.Context, petID int, pagination page.PageData) (page.Page[[]entity.MedicalHistory], error) {
	queryRows, err := r.queries.ListMedicalHistoryByPet(ctx, int32(petID))
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	entitiyList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	metadata := *page.GetPageMetadata(len(queryRows), pagination)
	return page.NewPage(entitiyList, metadata), nil
}

// TODO: Create QUERY
func (r *SQLCMedHistRepository) ListByOwnerID(ctx context.Context, ownerID int, pagination page.PageData) (page.Page[[]entity.MedicalHistory], error) {
	queryRows, err := r.queries.ListMedicalHistoryByPet(ctx, int32(ownerID))
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	entitiyList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]entity.MedicalHistory]{}, err
	}

	medicalHistPage := page.NewPage(entitiyList, *page.GetPageMetadata(len(queryRows), pagination))
	return medicalHistPage, nil
}

func (r *SQLCMedHistRepository) Create(ctx context.Context, medHistory *entity.MedicalHistory) error {
	params := ToCreateParams(*medHistory)
	createdRow, err := r.queries.CreateMedicalHistory(ctx, params)
	if err != nil {
		return err
	}

	medHistory.SetID(int(createdRow.ID))
	return nil
}

func (r *SQLCMedHistRepository) Update(ctx context.Context, medHistory *entity.MedicalHistory) error {
	var notes pgtype.Text
	if medHistory.Notes != nil {
		notes = pgtype.Text{String: *medHistory.Notes(), Valid: true}
	} else {
		notes = pgtype.Text{Valid: false}
	}

	params := sqlc.UpdateMedicalHistoryParams{
		ID:             int32(medHistory.ID().GetValue()),
		PetID:          int32(medHistory.PetID().GetValue()),
		OwnerID:        int32(medHistory.OwnerID()),
		VeterinarianID: int32(medHistory.VetID().GetValue()),
		VisitDate:      pgtype.Timestamptz{Time: medHistory.VisitDate(), Valid: true},
		Diagnosis:      pgtype.Text{String: medHistory.Diagnosis(), Valid: true},
		Treatment:      pgtype.Text{String: medHistory.Treatment(), Valid: true},
		Notes:          notes,
		VisitType:      medHistory.VisitType().ToString(),
		Condition:      pgtype.Text{String: medHistory.Condition().ToString(), Valid: true},
	}

	if _, err := r.queries.UpdateMedicalHistory(ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *SQLCMedHistRepository) Delete(ctx context.Context, medicalHistoryID int, softDelete bool) error {
	if softDelete {
		return r.queries.SoftDeleteMedicalHistory(ctx, int32(medicalHistoryID))
	}
	return r.queries.HardDeleteMedicalHistory(ctx, int32(medicalHistoryID))
}
