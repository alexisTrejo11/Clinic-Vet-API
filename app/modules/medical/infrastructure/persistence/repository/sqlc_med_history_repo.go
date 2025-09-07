// Package repositoryimpl implements the MedicalHistoryRepository using SQLC for database operations.
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/medical"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SQLCMedHistRepository struct {
	queries *sqlc.Queries
}

func NewSQLCMedHistRepository(queries *sqlc.Queries) repository.MedicalHistoryRepository {
	return &SQLCMedHistRepository{
		queries: queries,
	}
}

func (r *SQLCMedHistRepository) GetByID(ctx context.Context, medHistID valueobject.MedHistoryID) (medical.MedicalHistory, error) {
	sqlcMedHist, err := r.queries.GetMedicalHistoryByID(ctx, int32(medHistID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return medical.MedicalHistory{}, r.notFoundError("id", fmt.Sprintf("%d", medHistID.Value()))
		}
		return medical.MedicalHistory{}, r.dbError(OpSelect, fmt.Sprintf("failed to get medical history with ID %d", medHistID), err)
	}

	medHist, err := ToDomain(sqlcMedHist)
	if err != nil {
		return medical.MedicalHistory{}, r.wrapConversionError(err)
	}

	return medHist, nil
}

// Search retrieves medical history records based on search criteria with pagination
func (r *SQLCMedHistRepository) Search(ctx context.Context, searchParams any) (page.Page[[]medical.MedicalHistory], error) {
	searchParam, ok := searchParams.(dto.MedHistSearchParams)
	if !ok {
		return page.Page[[]medical.MedicalHistory]{}, r.invalidParamsError(ErrMsgInvalidSearchParams)
	}

	// TODO: Implement actual search parameters in the SQLC query
	queryRows, err := r.queries.SearchMedicalHistory(ctx, sqlc.SearchMedicalHistoryParams{
		// Add search parameters here when SQLC query is updated
		Limit:  int32(searchParam.Page.PageSize),
		Offset: r.calculateOffset(searchParam.Page),
	})
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.dbError(OpSearch, ErrMsgSearchMedicalHistory, err)
	}

	medHistoryList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.wrapConversionError(err)
	}

	// TODO: Get actual total count from database instead of using len(queryRows)
	totalCount, err := r.getSearchResultCount(ctx, searchParam)
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.dbError(OpCount, ErrMsgCountMedicalHistory, err)
	}

	metadata := page.GetPageMetadata(totalCount, searchParam.Page)
	return page.NewPage(medHistoryList, *metadata), nil
}

func (r *SQLCMedHistRepository) ListByVetID(ctx context.Context, vetID valueobject.VetID, pagination page.PageInput) (page.Page[[]medical.MedicalHistory], error) {
	params := sqlc.ListMedicalHistoryByVetParams{
		VeterinarianID: int32(vetID.Value()),
		Limit:          int32(pagination.PageSize),
		Offset:         r.calculateOffset(pagination),
	}

	queryRows, err := r.queries.ListMedicalHistoryByVet(ctx, params)
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list medical history for vet ID %d", vetID), err)
	}

	entityList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.wrapConversionError(err)
	}

	// Get actual total count for proper pagination
	totalCount, err := r.queries.CountMedicalHistoryByVet(ctx, int32(vetID.Value()))
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.dbError(OpCount, fmt.Sprintf("failed to count medical history for vet ID %d", vetID), err)
	}

	metadata := page.GetPageMetadata(int(totalCount), pagination)
	return page.NewPage(entityList, *metadata), nil
}

// TODO: Add pagination parameters to the SQLC query
func (r *SQLCMedHistRepository) ListByPetID(ctx context.Context, petID valueobject.PetID, pagination page.PageInput) (page.Page[[]medical.MedicalHistory], error) {
	// TODO: Update SQLC query to support pagination parameters
	queryRows, err := r.queries.ListMedicalHistoryByPet(ctx, int32(petID.Value()))
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list medical history for pet ID %d", petID), err)
	}

	entityList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.wrapConversionError(err)
	}

	// Apply manual pagination until SQLC query is updated
	paginatedList, totalCount := r.applyManualPagination(entityList, pagination)

	metadata := page.GetPageMetadata(totalCount, pagination)
	return page.NewPage(paginatedList, *metadata), nil
}

// TODO: Create proper SQLC query for owner-based medical history retrieval
func (r *SQLCMedHistRepository) ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID, pagination page.PageInput) (page.Page[[]medical.MedicalHistory], error) {
	// TODO: Create ListMedicalHistoryByOwner SQLC query instead of using pet query
	// This is currently using the wrong query as a placeholder
	queryRows, err := r.queries.ListMedicalHistoryByPet(ctx, int32(ownerID.Value()))
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.dbError(OpSelect, fmt.Sprintf("failed to list medical history for owner ID %d (using placeholder query)", ownerID.Value()), err)
	}

	entityList, err := ToDomainList(queryRows)
	if err != nil {
		return page.Page[[]medical.MedicalHistory]{}, r.wrapConversionError(err)
	}

	// Apply manual pagination until proper SQLC query is created
	paginatedList, totalCount := r.applyManualPagination(entityList, pagination)

	metadata := page.GetPageMetadata(totalCount, pagination)
	return page.NewPage(paginatedList, *metadata), nil
}

func (r *SQLCMedHistRepository) Create(ctx context.Context, medHistory *medical.MedicalHistory) error {
	params := ToCreateParams(*medHistory)
	_, err := r.queries.CreateMedicalHistory(ctx, params)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateMedicalHistory, err)
	}

	return nil
}

func (r *SQLCMedHistRepository) Update(ctx context.Context, medHistory *medical.MedicalHistory) error {
	notes := r.buildNotesParam(medHistory.Notes())
	updateParams := entityToUpdateParam(*medHistory, notes)

	_, err := r.queries.UpdateMedicalHistory(ctx, updateParams)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to update medical history with ID %d", medHistory.ID()), err)
	}

	return nil
}

func (r *SQLCMedHistRepository) Delete(ctx context.Context, medHistID valueobject.MedHistoryID, softDelete bool) error {
	var err error

	if softDelete {
		err = r.queries.SoftDeleteMedicalHistory(ctx, int32(medHistID.Value()))
	} else {
		err = r.queries.HardDeleteMedicalHistory(ctx, int32(medHistID.Value()))
	}

	if err != nil {
		deleteType := "soft"
		if !softDelete {
			deleteType = "hard"
		}
		return r.dbError(OpDelete, fmt.Sprintf("failed to %s delete medical history with ID %d", deleteType, medHistID.Value()), err)
	}

	return nil
}
