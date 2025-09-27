package query

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/repository"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type MedSessionQueryHandler struct {
	repo repository.MedicalSessionRepository
}

func NewMedicalSessionQueryHandler(repo repository.MedicalSessionRepository) *MedSessionQueryHandler {
	return &MedSessionQueryHandler{repo: repo}
}

func (h *MedSessionQueryHandler) FindMedSessionByID(ctx context.Context, query FindMedSessionByIDQuery) (*MedSessionResult, error) {
	var (
		medSession *medical.MedicalSession
		err        error
	)

	switch {
	case query.optCustomerID != nil:
		medSession, err = h.repo.FindByIDAndCustomerID(ctx, query.ID, *query.optCustomerID)
	case query.optPetID != nil:
		medSession, err = h.repo.FindByIDAndPetID(ctx, query.ID, *query.optPetID)
	case query.optEmployeeID != nil:
		medSession, err = h.repo.FindByIDAndEmployeeID(ctx, query.ID, *query.optEmployeeID)
	default:
		medSession, err = h.repo.FindByID(ctx, query.ID)
	}

	if err != nil {
		return nil, err
	}

	result := toResult(*medSession)
	return &result, nil
}

func (h *MedSessionQueryHandler) FindMedSessionBySpec(ctx context.Context, query FindMedSessionBySpecQuery) (*p.Page[MedSessionResult], error) {
	medSessionpage, err := h.repo.FindBySpecification(ctx, query.Spec)
	if err != nil {
		return nil, err
	}

	result := p.MapItems(medSessionpage, toResult)
	return &result, nil
}

func (h *MedSessionQueryHandler) FindMedSessionByEmployeeID(ctx context.Context, query FindMedSessionByEmployeeIDQuery) (*p.Page[MedSessionResult], error) {
	medSessionpage, err := h.repo.FindByEmployeeID(ctx, query.EmployeeID, query.PaginationRequest)
	if err != nil {
		return nil, err
	}

	result := p.MapItems(medSessionpage, toResult)
	return &result, nil
}

func (h *MedSessionQueryHandler) FindMedSessionByPetID(ctx context.Context, query FindMedSessionByPetIDQuery) (*p.Page[MedSessionResult], error) {
	medSessionpage, err := h.repo.FindByPetID(ctx, query.petID, query.PaginationRequest)
	if err != nil {
		return nil, err
	}

	result := p.MapItems(medSessionpage, toResult)
	return &result, nil
}

func (h *MedSessionQueryHandler) FindMedSessionByCustomerID(ctx context.Context, query FindMedSessionByCustomerIDQuery) (*p.Page[MedSessionResult], error) {
	medSessionpage, err := h.repo.FindByCustomerID(ctx, query.CustomerID, query.PaginationRequest)
	if err != nil {
		return nil, err
	}

	result := p.MapItems(medSessionpage, toResult)
	return &result, nil
}

func (h *MedSessionQueryHandler) FindRecentMedSessionByPetID(ctx context.Context, query FindRecentMedSessionByPetIDQuery) ([]MedSessionResult, error) {
	medSession, err := h.repo.FindRecentByPetID(ctx, query.PetID, query.Limit)
	if err != nil {
		return nil, err
	}

	return toResultList(medSession), nil
}

func (h *MedSessionQueryHandler) FindMedSessionByDateRange(ctx context.Context, query FindMedSessionByDateRangeQuery) (*p.Page[MedSessionResult], error) {
	medSessionpage, err := h.repo.FindByDateRange(ctx, query.StartDate, query.EndDate, query.PaginationRequest)
	if err != nil {
		return nil, err
	}

	result := p.MapItems(medSessionpage, toResult)
	return &result, nil
}

func (h *MedSessionQueryHandler) FindMedSessionByPetAndDateRange(ctx context.Context, query FindMedSessionByPetAndDateRangeQuery) ([]MedSessionResult, error) {
	medSession, err := h.repo.FindByPetAndDateRange(ctx, query.PetID, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	return toResultList(medSession), nil
}

func (h *MedSessionQueryHandler) FindMedSessionByDiagnosis(ctx context.Context, query FindMedSessionByDiagnosisQuery) (*p.Page[MedSessionResult], error) {
	medSessionpage, err := h.repo.FindByDiagnosis(ctx, query.Diagnosis, query.PaginationRequest)
	if err != nil {
		return nil, err
	}

	result := p.MapItems(medSessionpage, toResult)
	return &result, nil
}
