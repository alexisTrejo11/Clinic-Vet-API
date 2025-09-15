package query

import (
	"clinic-vet-api/app/core/repository"
	p "clinic-vet-api/app/shared/page"
)

type MedicalHistoryQueryHandlers interface {
	GetMedHistByID(query GetMedHistByIDQuery) (*MedHistoryResult, error)
	GetMedHistBySpec(query GetMedHistBySpecQuery) (*p.Page[MedHistoryResult], error)
	GetAllMedHist(query GetAllMedHistQuery) (*p.Page[MedHistoryResult], error)
	GetMedHistByEmployeeID(query GetMedHistByEmployeeIDQuery) (*p.Page[MedHistoryResult], error)
	GetMedHistByPetID(query GetMedHistByPetIDQuery) (*p.Page[MedHistoryResult], error)
	GetMedHistByCustomerID(query GetMedHistByCustomerIDQuery) (*p.Page[MedHistoryResult], error)
	GetRecentMedHistByPetID(query GetRecentMedHistByPetIDQuery) (*MedHistoryResult, error)
	GetMedHistByDateRange(query GetMedHistByDateRangeQuery) (*p.Page[MedHistoryResult], error)
	GetMedHistByPetAndDateRange(query GetMedHistByPetAndDateRangeQuery) (*MedHistoryResult, error)
	GetMedHistByDiagnosis(query GetMedHistByDiagnosisQuery) (*p.Page[MedHistoryResult], error)
	ExistsMedHistByID(query ExistsMedHistByIDQuery) (bool, error)
	ExistsMedHistByPetAndDate(query ExistsMedHistByPetAndDateQuery) (bool, error)
}

type medHistQueryHandler struct {
	repo repository.MedicalHistoryRepository
}

func NewMedicalHistoryQueryHandlers(repo repository.MedicalHistoryRepository) MedicalHistoryQueryHandlers {
	return &medHistQueryHandler{repo: repo}
}

func (h *medHistQueryHandler) GetMedHistByID(query GetMedHistByIDQuery) (*MedHistoryResult, error) {
	medHistory, err := h.repo.FindByID(query.CTX, query.ID)
	if err != nil {
		return nil, err
	}

	result := toResult(medHistory)
	return &result, nil
}

func (h *medHistQueryHandler) GetMedHistBySpec(query GetMedHistBySpecQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindBySpecification(query.CTX, query.Spec)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) GetAllMedHist(query GetAllMedHistQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindAll(query.CTX, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) GetMedHistByEmployeeID(query GetMedHistByEmployeeIDQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByEmployeeID(query.CTX, query.EmployeeID, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) GetMedHistByPetID(query GetMedHistByPetIDQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByPetID(query.CTX, query.PetID, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) GetMedHistByCustomerID(query GetMedHistByCustomerIDQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByCustomerID(query.CTX, query.CustomerID, query.PageInput)
	if err != nil {
		return nil, err
	}

	result := toResultPage(page)
	return &result, nil
}

func (h *medHistQueryHandler) GetRecentMedHistByPetID(query GetRecentMedHistByPetIDQuery) (*MedHistoryResult, error) {
	medHistory, err := h.repo.FindRecentByPetID(query.CTX, query.PetID, query.Limit)
	if err != nil {
		return nil, err
	}

	result := toResult(medHistory)
	return &result, nil
}

func (h *medHistQueryHandler) GetMedHistByDateRange(query GetMedHistByDateRangeQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByDateRange(query.CTX, query.StartDate, query.EndDate, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) GetMedHistByPetAndDateRange(query GetMedHistByPetAndDateRangeQuery) (*MedHistoryResult, error) {
	medHistory, err := h.repo.FindByPetAndDateRange(query.CTX, query.PetID, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	result := toResult(medHistory)
	return &result, nil
}

func (h *medHistQueryHandler) GetMedHistByDiagnosis(query GetMedHistByDiagnosisQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByDiagnosis(query.CTX, query.Diagnosis, query.PageInput)
	if err != nil {
		return nil, err
	}

	result := toResultPage(page)
	return &result, nil
}

func (h *medHistQueryHandler) ExistsMedHistByID(query ExistsMedHistByIDQuery) (bool, error) {
	exists, err := h.repo.ExistsByID(query.CTX, query.ID)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (h *medHistQueryHandler) ExistsMedHistByPetAndDate(query ExistsMedHistByPetAndDateQuery) (bool, error) {
	exists, err := h.repo.ExistsByPetAndDate(query.CTX, query.PetID, query.Date)
	if err != nil {
		return false, err
	}

	return exists, nil
}
