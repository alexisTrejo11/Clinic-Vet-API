package specification

import (
	"strings"
	"time"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type MedicalHistorySpecification struct {
	PetIDs          []valueobject.PetID
	CustomerID      []valueobject.CustomerID
	EmployeeIDs     []valueobject.EmployeeID
	VisitReasons    []enum.VisitReason
	VisitTypes      []enum.VisitType
	Conditions      []enum.PetCondition
	Diagnosis       *string
	Treatment       *string
	VisitDateFrom   *time.Time
	VisitDateTo     *time.Time
	CreatedDateFrom *time.Time
	CreatedDateTo   *time.Time
	SearchTerm      *string // notes, diagnosis, treatment
	Pagination
}

func (s *MedicalHistorySpecification) IsSatisfiedBy(entity any) bool {
	history, ok := entity.(interface {
		PetID() valueobject.PetID
		customerID() valueobject.CustomerID
		EmployeeID() valueobject.EmployeeID
		VisitReason() enum.VisitReason
		VisitType() enum.VisitType
		VisitDate() time.Time
		Diagnosis() string
		Treatment() string
		Condition() enum.PetCondition
		Notes() *string
		CreatedAt() time.Time
	})
	if !ok {
		return false
	}

	if len(s.PetIDs) > 0 {
		found := false
		for _, petID := range s.PetIDs {
			if petID.Value() == history.PetID().Value() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.CustomerID) > 0 {
		found := false
		for _, customerID := range s.CustomerID {
			if customerID.Value() == history.customerID().Value() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.EmployeeIDs) > 0 {
		found := false
		for _, vetID := range s.EmployeeIDs {
			if vetID.Value() == history.EmployeeID().Value() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.VisitReasons) > 0 {
		found := false
		for _, reason := range s.VisitReasons {
			if reason == history.VisitReason() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.VisitTypes) > 0 {
		found := false
		for _, visitType := range s.VisitTypes {
			if visitType == history.VisitType() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.Conditions) > 0 {
		found := false
		for _, condition := range s.Conditions {
			if condition == history.Condition() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if s.Diagnosis != nil && *s.Diagnosis != "" {
		if !strings.Contains(strings.ToLower(history.Diagnosis()), strings.ToLower(*s.Diagnosis)) {
			return false
		}
	}

	if s.Treatment != nil && *s.Treatment != "" {
		if !strings.Contains(strings.ToLower(history.Treatment()), strings.ToLower(*s.Treatment)) {
			return false
		}
	}

	if s.VisitDateFrom != nil && history.VisitDate().Before(*s.VisitDateFrom) {
		return false
	}
	if s.VisitDateTo != nil && history.VisitDate().After(*s.VisitDateTo) {
		return false
	}

	if s.CreatedDateFrom != nil && history.CreatedAt().Before(*s.CreatedDateFrom) {
		return false
	}
	if s.CreatedDateTo != nil && history.CreatedAt().After(*s.CreatedDateTo) {
		return false
	}

	if s.SearchTerm != nil && *s.SearchTerm != "" {
		searchTerm := strings.ToLower(*s.SearchTerm)
		matches := strings.Contains(strings.ToLower(history.Diagnosis()), searchTerm) ||
			strings.Contains(strings.ToLower(history.Treatment()), searchTerm)

		if history.Notes() != nil {
			matches = matches || strings.Contains(strings.ToLower(*history.Notes()), searchTerm)
		}

		if !matches {
			return false
		}
	}

	return true
}

func (s *MedicalHistorySpecification) ToSQL() (string, []any) {
	var conditions []string
	var args []any

	if len(s.PetIDs) > 0 {
		petIDs := make([]any, len(s.PetIDs))
		for i, petID := range s.PetIDs {
			petIDs[i] = petID.Value()
		}
		conditions = append(conditions, "pet_id IN (?)")
		args = append(args, petIDs)
	}

	if len(s.CustomerID) > 0 {
		customerIDs := make([]any, len(s.CustomerID))
		for i, customerID := range s.CustomerID {
			customerIDs[i] = customerID.Value()
		}
		conditions = append(conditions, "customer_id IN (?)")
		args = append(args, customerIDs)
	}

	if len(s.EmployeeIDs) > 0 {
		vetIDs := make([]any, len(s.EmployeeIDs))
		for i, vetID := range s.EmployeeIDs {
			vetIDs[i] = vetID.Value()
		}
		conditions = append(conditions, "employee_id IN (?)")
		args = append(args, vetIDs)
	}

	if len(s.VisitReasons) > 0 {
		reasons := make([]any, len(s.VisitReasons))
		for i, reason := range s.VisitReasons {
			reasons[i] = reason.String()
		}
		conditions = append(conditions, "visit_reason IN (?)")
		args = append(args, reasons)
	}

	if len(s.VisitTypes) > 0 {
		visitTypes := make([]any, len(s.VisitTypes))
		for i, visitType := range s.VisitTypes {
			visitTypes[i] = visitType.String()
		}
		conditions = append(conditions, "visit_type IN (?)")
		args = append(args, visitTypes)
	}

	if len(s.Conditions) > 0 {
		conditionsList := make([]any, len(s.Conditions))
		for i, condition := range s.Conditions {
			conditionsList[i] = condition.String()
		}
		conditions = append(conditions, "condition IN (?)")
		args = append(args, conditionsList)
	}

	if s.Diagnosis != nil && *s.Diagnosis != "" {
		conditions = append(conditions, "diagnosis ILIKE ?")
		args = append(args, "%"+*s.Diagnosis+"%")
	}

	// Filtro por Treatment
	if s.Treatment != nil && *s.Treatment != "" {
		conditions = append(conditions, "treatment ILIKE ?")
		args = append(args, "%"+*s.Treatment+"%")
	}

	if s.VisitDateFrom != nil {
		conditions = append(conditions, "visit_date >= ?")
		args = append(args, *s.VisitDateFrom)
	}
	if s.VisitDateTo != nil {
		conditions = append(conditions, "visit_date <= ?")
		args = append(args, *s.VisitDateTo)
	}

	if s.CreatedDateFrom != nil {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, *s.CreatedDateFrom)
	}
	if s.CreatedDateTo != nil {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, *s.CreatedDateTo)
	}

	if s.SearchTerm != nil && *s.SearchTerm != "" {
		searchCondition := "(diagnosis ILIKE ? OR treatment ILIKE ?"
		searchArgs := []any{"%" + *s.SearchTerm + "%", "%" + *s.SearchTerm + "%"}

		searchCondition += " OR notes ILIKE ?"
		searchArgs = append(searchArgs, "%"+*s.SearchTerm+"%")

		searchCondition += ")"
		conditions = append(conditions, searchCondition)
		args = append(args, searchArgs...)
	}

	where := ""
	if len(conditions) > 0 {
		where = "WHERE " + strings.Join(conditions, " AND ")
	}

	return where, args
}

func (s *MedicalHistorySpecification) WithPetIDs(petIDs ...valueobject.PetID) *MedicalHistorySpecification {
	s.PetIDs = petIDs
	return s
}

func (s *MedicalHistorySpecification) WithcustomerIDs(customerIDs ...valueobject.CustomerID) *MedicalHistorySpecification {
	s.CustomerID = customerIDs
	return s
}

func (s *MedicalHistorySpecification) WithEmployeeIDs(vetIDs ...valueobject.EmployeeID) *MedicalHistorySpecification {
	s.EmployeeIDs = vetIDs
	return s
}

func (s *MedicalHistorySpecification) WithVisitReasons(reasons ...enum.VisitReason) *MedicalHistorySpecification {
	s.VisitReasons = reasons
	return s
}

func (s *MedicalHistorySpecification) WithVisitTypes(visitTypes ...enum.VisitType) *MedicalHistorySpecification {
	s.VisitTypes = visitTypes
	return s
}

func (s *MedicalHistorySpecification) WithConditions(conditions ...enum.PetCondition) *MedicalHistorySpecification {
	s.Conditions = conditions
	return s
}

func (s *MedicalHistorySpecification) WithDiagnosis(diagnosis string) *MedicalHistorySpecification {
	s.Diagnosis = &diagnosis
	return s
}

func (s *MedicalHistorySpecification) WithTreatment(treatment string) *MedicalHistorySpecification {
	s.Treatment = &treatment
	return s
}

func (s *MedicalHistorySpecification) WithVisitDateRange(from, to *time.Time) *MedicalHistorySpecification {
	s.VisitDateFrom = from
	s.VisitDateTo = to
	return s
}

func (s *MedicalHistorySpecification) WithCreatedDateRange(from, to *time.Time) *MedicalHistorySpecification {
	s.CreatedDateFrom = from
	s.CreatedDateTo = to
	return s
}

func (s *MedicalHistorySpecification) WithSearchTerm(term string) *MedicalHistorySpecification {
	s.SearchTerm = &term
	return s
}

func (s *MedicalHistorySpecification) WithPagination(page, pageSize int, orderBy, sortDir string) *MedicalHistorySpecification {
	s.Pagination = Pagination{
		Page:     page,
		PageSize: pageSize,
		OrderBy:  orderBy,
		SortDir:  sortDir,
	}
	return s
}
