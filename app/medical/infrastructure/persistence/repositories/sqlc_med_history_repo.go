package sqlcMedHistoryRepo

import (
	"context"

	mhDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/dtos"
	mhDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain"
	medHistRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain/repositories"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCMedHistRepository struct {
	queries *sqlc.Queries
}

func NewSQLCMedHistRepository(queries *sqlc.Queries) medHistRepo.MedicalHistoryRepository {
	return &SQLCMedHistRepository{
		queries: queries,
	}
}

func (r *SQLCMedHistRepository) GetById(ctx context.Context, medicalHistoryId int) (*mhDomain.MedicalHistory, error) {
	sqlcMedHist, err := r.queries.GetMedicalHistoryByID(ctx, int32(medicalHistoryId))
	if err != nil {
		return nil, err
	}

	medHist, err := ToDomain(sqlcMedHist)
	if err != nil {
		return nil, err
	}

	return &medHist, nil
}

func (r *SQLCMedHistRepository) Search(ctx context.Context, searchParams mhDTOs.MedHistSearchParams) (*page.Page[[]mhDomain.MedicalHistory], error) {
	queryRows, err := r.queries.SearchMedicalHistory(ctx, sqlc.SearchMedicalHistoryParams{})
	if err != nil {
		return nil, err
	}

	entities, err := ToDomainList(queryRows)
	if err != nil {
		return nil, err
	}

	medicalHistPage := page.NewPage(entities, *page.GetPageMetadata(len(queryRows), searchParams.Page))
	return medicalHistPage, nil

}

func (r *SQLCMedHistRepository) ListByVetId(ctx context.Context, vetId int, pagination page.PageData) (*page.Page[[]mhDomain.MedicalHistory], error) {
	arg := sqlc.ListMedicalHistoryByVetParams{
		VeterinarianID: int32(vetId),
		Limit:          int32(pagination.PageNumber),
		Offset:         int32(pagination.PageSize * (pagination.PageNumber - 1)),
	}

	queryRows, err := r.queries.ListMedicalHistoryByVet(ctx, arg)
	if err != nil {
		return nil, err
	}

	entitiyList, err := ToDomainList(queryRows)
	if err != nil {
		return nil, err
	}

	medicalHistPage := page.NewPage(entitiyList, *page.GetPageMetadata(len(queryRows), pagination))
	return medicalHistPage, nil

}

// TODO: Paginate SQLC
func (r *SQLCMedHistRepository) ListByPetId(ctx context.Context, petId int, pagination page.PageData) (*page.Page[[]mhDomain.MedicalHistory], error) {
	queryRows, err := r.queries.ListMedicalHistoryByPet(ctx, int32(petId))
	if err != nil {
		return nil, err
	}

	entitiyList, err := ToDomainList(queryRows)
	if err != nil {
		return nil, err
	}

	medicalHistPage := page.NewPage(entitiyList, *page.GetPageMetadata(len(queryRows), pagination))
	return medicalHistPage, nil
}

// TODO: Create QUERY
func (r *SQLCMedHistRepository) ListByOwnerId(ctx context.Context, ownerId int, pagination page.PageData) (*page.Page[[]mhDomain.MedicalHistory], error) {
	queryRows, err := r.queries.ListMedicalHistoryByPet(ctx, int32(ownerId))
	if err != nil {
		return nil, err
	}

	entitiyList, err := ToDomainList(queryRows)
	if err != nil {
		return nil, err
	}

	medicalHistPage := page.NewPage(entitiyList, *page.GetPageMetadata(len(queryRows), pagination))
	return medicalHistPage, nil
}

func (r *SQLCMedHistRepository) Create(ctx context.Context, medHistory *mhDomain.MedicalHistory) error {
	createdRow, err := r.queries.CreateMedicalHistory(ctx, sqlc.CreateMedicalHistoryParams{
		PetID:          int32(medHistory.PetId.GetValue()),
		OwnerID:        int32(medHistory.OwnerId),
		VeterinarianID: int32(medHistory.VetId.GetValue()),
		VisitDate:      pgtype.Timestamptz{Time: medHistory.VisitDate, Valid: true},
		Diagnosis:      pgtype.Text{String: medHistory.Diagnosis, Valid: true},
		Treatment:      pgtype.Text{String: medHistory.Treatment, Valid: true},
		Notes:          pgtype.Text{String: medHistory.Notes, Valid: true},
		Condition:      pgtype.Text{String: medHistory.Condition, Valid: true},
	})

	if err != nil {
		return err
	}

	medHistory.SetId(int(createdRow.ID))
	return nil
}

func (r *SQLCMedHistRepository) Update(ctx context.Context, medHistory *mhDomain.MedicalHistory) error {
	if _, err := r.queries.UpdateMedicalHistory(ctx, sqlc.UpdateMedicalHistoryParams{
		ID:             int32(medHistory.Id.GetValue()),
		PetID:          int32(medHistory.PetId.GetValue()),
		OwnerID:        int32(medHistory.OwnerId),
		VeterinarianID: int32(medHistory.VetId.GetValue()),
		VisitDate:      pgtype.Timestamptz{Time: medHistory.VisitDate, Valid: true},
		Diagnosis:      pgtype.Text{String: medHistory.Diagnosis, Valid: true},
		Treatment:      pgtype.Text{String: medHistory.Treatment, Valid: true},
		Notes:          pgtype.Text{String: medHistory.Notes, Valid: true},
		Condition:      pgtype.Text{String: medHistory.Condition, Valid: true},
	}); err != nil {
		return err
	}
	return nil
}

func (r *SQLCMedHistRepository) Delete(ctx context.Context, medicalHistoryId int, softDelete bool) error {
	if softDelete {
		return r.queries.SoftDeleteMedicalHistory(ctx, int32(medicalHistoryId))
	}
	return r.queries.HardDeleteMedicalHistory(ctx, int32(medicalHistoryId))
}

func ToDomain(medHist sqlc.MedicalHistory) (mhDomain.MedicalHistory, error) {
	medHistId, err := mhDomain.NewMedHistoryId(medHist.ID)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	petId, err := petDomain.NewPetId(medHist.PetID)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	vetId, err := vetDomain.NewVeterinarianId(medHist.VeterinarianID)
	if err != nil {
		return mhDomain.MedicalHistory{}, err
	}

	return mhDomain.MedicalHistory{
		Id:        medHistId,
		PetId:     petId,
		OwnerId:   int(medHist.OwnerID),
		VetId:     vetId,
		VisitDate: medHist.VisitDate.Time,
		Diagnosis: medHist.Diagnosis.String,
		Treatment: medHist.Treatment.String,
		Notes:     medHist.Notes.String,
		Condition: medHist.Condition.String,
	}, nil
}

func ToDomainList(medHistList []sqlc.MedicalHistory) ([]mhDomain.MedicalHistory, error) {
	domainList := make([]mhDomain.MedicalHistory, len(medHistList))

	for i, medHist := range medHistList {
		domainMedHist, err := ToDomain(medHist)
		if err != nil {
			return nil, err
		}
		domainList[i] = domainMedHist
	}

	return domainList, nil
}
