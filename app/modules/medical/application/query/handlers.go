package query

import (
	"clinic-vet-api/app/core/repository"
	p "clinic-vet-api/app/shared/page"
)

type MedicalHistoryQueryHandlers interface {
	FindMedHistByID(query FindMedHistByIDQuery) (*MedHistoryResult, error)
	FindMedHistBySpec(query FindMedHistBySpecQuery) (*p.Page[MedHistoryResult], error)
	FindAllMedHist(query FindAllMedHistQuery) (*p.Page[MedHistoryResult], error)
	FindMedHistByEmployeeID(query FindMedHistByEmployeeIDQuery) (*p.Page[MedHistoryResult], error)
	FindMedHistByPetID(query FindMedHistByPetIDQuery) (*p.Page[MedHistoryResult], error)
	FindMedHistByCustomerID(query FindMedHistByCustomerIDQuery) (*p.Page[MedHistoryResult], error)
	FindRecentMedHistByPetID(query FindRecentMedHistByPetIDQuery) ([]MedHistoryResult, error)
	FindMedHistByDateRange(query FindMedHistByDateRangeQuery) (*p.Page[MedHistoryResult], error)
	FindMedHistByPetAndDateRange(query FindMedHistByPetAndDateRangeQuery) ([]MedHistoryResult, error)
	FindMedHistByDiagnosis(query FindMedHistByDiagnosisQuery) (*p.Page[MedHistoryResult], error)
	ExistsMedHistByID(query ExistsMedHistByIDQuery) (bool, error)
	ExistsMedHistByPetAndDate(query ExistsMedHistByPetAndDateQuery) (bool, error)
}

type medHistQueryHandler struct {
	repo repository.MedicalHistoryRepository
}

func NewMedicalHistoryQueryHandlers(repo repository.MedicalHistoryRepository) MedicalHistoryQueryHandlers {
	return &medHistQueryHandler{repo: repo}
}

func (h *medHistQueryHandler) FindMedHistByID(query FindMedHistByIDQuery) (*MedHistoryResult, error) {
	medHistory, err := h.repo.FindByID(query.CTX, query.ID)
	if err != nil {
		return nil, err
	}

	result := toResult(medHistory)
	return &result, nil
}

func (h *medHistQueryHandler) FindMedHistBySpec(query FindMedHistBySpecQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindBySpecification(query.CTX, query.Spec)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) FindAllMedHist(query FindAllMedHistQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindAll(query.CTX, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) FindMedHistByEmployeeID(query FindMedHistByEmployeeIDQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByEmployeeID(query.CTX, query.EmployeeID, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) FindMedHistByPetID(query FindMedHistByPetIDQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByPetID(query.CTX, query.PetID, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) FindMedHistByCustomerID(query FindMedHistByCustomerIDQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByCustomerID(query.CTX, query.CustomerID, query.PageInput)
	if err != nil {
		return nil, err
	}

	result := toResultPage(page)
	return &result, nil
}

func (h *medHistQueryHandler) FindRecentMedHistByPetID(query FindRecentMedHistByPetIDQuery) ([]MedHistoryResult, error) {
	medHistory, err := h.repo.FindRecentByPetID(query.CTX, query.PetID, query.Limit)
	if err != nil {
		return nil, err
	}

	return toResultList(medHistory), nil
}

func (h *medHistQueryHandler) FindMedHistByDateRange(query FindMedHistByDateRangeQuery) (*p.Page[MedHistoryResult], error) {
	page, err := h.repo.FindByDateRange(query.CTX, query.StartDate, query.EndDate, query.PageInput)
	if err != nil {
		return nil, err
	}

	age := toResultPage(page)
	return &age, nil
}

func (h *medHistQueryHandler) FindMedHistByPetAndDateRange(query FindMedHistByPetAndDateRangeQuery) ([]MedHistoryResult, error) {
	medHistory, err := h.repo.FindByPetAndDateRange(query.CTX, query.PetID, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	return toResultList(medHistory), nil
}

func (h *medHistQueryHandler) FindMedHistByDiagnosis(query FindMedHistByDiagnosisQuery) (*p.Page[MedHistoryResult], error) {
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
