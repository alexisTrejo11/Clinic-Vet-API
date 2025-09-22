package query

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/repository"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type MedicalSessionQueryHandlers interface {
	FindMedSessionByID(ctx context.Context, query FindMedSessionByIDQuery) (*MedSessionResult, error)
	FindMedSessionBySpec(ctx context.Context, query FindMedSessionBySpecQuery) (*p.Page[MedSessionResult], error)
	FindMedSessionByEmployeeID(ctx context.Context, query FindMedSessionByEmployeeIDQuery) (*p.Page[MedSessionResult], error)
	FindMedSessionByPetID(ctx context.Context, query FindMedSessionByPetIDQuery) (*p.Page[MedSessionResult], error)
	FindMedSessionByCustomerID(ctx context.Context, query FindMedSessionByCustomerIDQuery) (*p.Page[MedSessionResult], error)
	FindRecentMedSessionByPetID(ctx context.Context, query FindRecentMedSessionByPetIDQuery) ([]MedSessionResult, error)
	FindMedSessionByDateRange(ctx context.Context, query FindMedSessionByDateRangeQuery) (*p.Page[MedSessionResult], error)
	FindMedSessionByPetAndDateRange(ctx context.Context, query FindMedSessionByPetAndDateRangeQuery) ([]MedSessionResult, error)
	FindMedSessionByDiagnosis(ctx context.Context, query FindMedSessionByDiagnosisQuery) (*p.Page[MedSessionResult], error)
}

type medSessionQueryHandler struct {
	repo repository.MedicalSessionRepository
}

func NewMedicalSessionQueryHandlers(repo repository.MedicalSessionRepository) MedicalSessionQueryHandlers {
	return &medSessionQueryHandler{repo: repo}
}

func (h *medSessionQueryHandler) FindMedSessionByID(ctx context.Context, query FindMedSessionByIDQuery) (*MedSessionResult, error) {
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

func (h *medSessionQueryHandler) FindMedSessionBySpec(ctx context.Context, query FindMedSessionBySpecQuery) (*p.Page[MedSessionResult], error) {
	page, err := h.repo.FindBySpecification(ctx, query.Spec)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medSessionQueryHandler) FindMedSessionByEmployeeID(ctx context.Context, query FindMedSessionByEmployeeIDQuery) (*p.Page[MedSessionResult], error) {
	page, err := h.repo.FindByEmployeeID(ctx, query.EmployeeID, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medSessionQueryHandler) FindMedSessionByPetID(ctx context.Context, query FindMedSessionByPetIDQuery) (*p.Page[MedSessionResult], error) {
	page, err := h.repo.FindByPetID(ctx, query.petID, query.pageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medSessionQueryHandler) FindMedSessionByCustomerID(ctx context.Context, query FindMedSessionByCustomerIDQuery) (*p.Page[MedSessionResult], error) {
	page, err := h.repo.FindByCustomerID(ctx, query.CustomerID, query.PageInput)
	if err != nil {
		return nil, err
	}

	result := toResultPage(page)
	return &result, nil
}

func (h *medSessionQueryHandler) FindRecentMedSessionByPetID(ctx context.Context, query FindRecentMedSessionByPetIDQuery) ([]MedSessionResult, error) {
	medSession, err := h.repo.FindRecentByPetID(ctx, query.PetID, query.Limit)
	if err != nil {
		return nil, err
	}

	return toResultList(medSession), nil
}

func (h *medSessionQueryHandler) FindMedSessionByDateRange(ctx context.Context, query FindMedSessionByDateRangeQuery) (*p.Page[MedSessionResult], error) {
	page, err := h.repo.FindByDateRange(ctx, query.StartDate, query.EndDate, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medSessionQueryHandler) FindMedSessionByPetAndDateRange(ctx context.Context, query FindMedSessionByPetAndDateRangeQuery) ([]MedSessionResult, error) {
	medSession, err := h.repo.FindByPetAndDateRange(ctx, query.PetID, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	return toResultList(medSession), nil
}

func (h *medSessionQueryHandler) FindMedSessionByDiagnosis(ctx context.Context, query FindMedSessionByDiagnosisQuery) (*p.Page[MedSessionResult], error) {
	page, err := h.repo.FindByDiagnosis(ctx, query.Diagnosis, query.PageInput)
	if err != nil {
		return nil, err
	}

	result := toResultPage(page)
	return &result, nil
}
