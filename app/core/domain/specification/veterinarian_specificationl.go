package specification

import (
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type VeterinarianSpecification struct {
	IDs                []valueobject.VetID
	UserIDs            []valueobject.UserID
	LicenseNumbers     []string
	Specialties        []enum.VetSpecialty
	YearsExperienceMin *int
	YearsExperienceMax *int
	ConsultationFeeMin *valueobject.Money
	ConsultationFeeMax *valueobject.Money
	IsActive           *bool
	HasUserAccount     *bool
	IsAvailable        *bool
	AvailableDate      *time.Time
	SearchTerm         *string // Búsqueda en nombre, licencia, etc.
	CreatedAfter       *time.Time
	CreatedBefore      *time.Time
	Pagination
}

func (s *VeterinarianSpecification) IsSatisfiedBy(entity any) bool {
	vet, ok := entity.(interface {
		ID() valueobject.VetID
		UserID() *valueobject.UserID
		Name() valueobject.PersonName
		LicenseNumber() string
		Specialty() enum.VetSpecialty
		YearsExperience() int
		ConsultationFee() *valueobject.Money
		IsActive() bool
		CreatedAt() time.Time
		Schedule() *valueobject.Schedule
	})
	if !ok {
		return false
	}

	if len(s.IDs) > 0 {
		found := false
		for _, id := range s.IDs {
			if id.Value() == vet.ID().Value() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.UserIDs) > 0 {
		if vet.UserID() == nil {
			return false
		}
		found := false
		for _, userID := range s.UserIDs {
			if userID.Value() == vet.UserID().Value() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.LicenseNumbers) > 0 {
		found := false
		for _, license := range s.LicenseNumbers {
			if license == vet.LicenseNumber() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.Specialties) > 0 {
		found := false
		for _, specialty := range s.Specialties {
			if specialty == vet.Specialty() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if s.YearsExperienceMin != nil && vet.YearsExperience() < *s.YearsExperienceMin {
		return false
	}
	if s.YearsExperienceMax != nil && vet.YearsExperience() > *s.YearsExperienceMax {
		return false
	}

	if s.ConsultationFeeMin != nil {
		if vet.ConsultationFee() == nil || vet.ConsultationFee().Amount() < s.ConsultationFeeMin.Amount() {
			return false
		}
	}
	if s.ConsultationFeeMax != nil {
		if vet.ConsultationFee() == nil || vet.ConsultationFee().Amount() > s.ConsultationFeeMax.Amount() {
			return false
		}
	}

	if s.IsActive != nil && vet.IsActive() != *s.IsActive {
		return false
	}

	if s.HasUserAccount != nil {
		hasAccount := vet.UserID() != nil
		if hasAccount != *s.HasUserAccount {
			return false
		}
	}

	if s.IsAvailable != nil {
		isAvailable := s.checkAvailability(vet)
		if isAvailable != *s.IsAvailable {
			return false
		}
	}

	if s.CreatedAfter != nil && vet.CreatedAt().Before(*s.CreatedAfter) {
		return false
	}
	if s.CreatedBefore != nil && vet.CreatedAt().After(*s.CreatedBefore) {
		return false
	}

	// Búsqueda por término
	if s.SearchTerm != nil && *s.SearchTerm != "" {
		searchTerm := strings.ToLower(*s.SearchTerm)
		matches := strings.Contains(strings.ToLower(vet.Name().FullName()), searchTerm) ||
			strings.Contains(strings.ToLower(vet.LicenseNumber()), searchTerm) ||
			strings.Contains(strings.ToLower(vet.Specialty().String()), searchTerm)

		if !matches {
			return false
		}
	}

	return true
}

func (s *VeterinarianSpecification) ToSQL() (string, []any) {
	var conditions []string
	var args []any

	if len(s.IDs) > 0 {
		ids := make([]any, len(s.IDs))
		for i, id := range s.IDs {
			ids[i] = id.Value()
		}
		conditions = append(conditions, "id IN (?)")
		args = append(args, ids)
	}

	if len(s.UserIDs) > 0 {
		userIDs := make([]any, len(s.UserIDs))
		for i, userID := range s.UserIDs {
			userIDs[i] = userID.Value()
		}
		conditions = append(conditions, "user_id IN (?)")
		args = append(args, userIDs)
	}

	if len(s.LicenseNumbers) > 0 {
		licenses := make([]any, len(s.LicenseNumbers))
		for i, license := range s.LicenseNumbers {
			licenses[i] = license
		}
		conditions = append(conditions, "license_number IN (?)")
		args = append(args, licenses)
	}

	// Filtro por Specialties
	if len(s.Specialties) > 0 {
		specialties := make([]any, len(s.Specialties))
		for i, specialty := range s.Specialties {
			specialties[i] = specialty.String()
		}
		conditions = append(conditions, "specialty IN (?)")
		args = append(args, specialties)
	}

	if s.YearsExperienceMin != nil {
		conditions = append(conditions, "years_experience >= ?")
		args = append(args, *s.YearsExperienceMin)
	}
	if s.YearsExperienceMax != nil {
		conditions = append(conditions, "years_experience <= ?")
		args = append(args, *s.YearsExperienceMax)
	}

	if s.ConsultationFeeMin != nil {
		conditions = append(conditions, "consultation_fee >= ?")
		args = append(args, s.ConsultationFeeMin.Amount())
	}
	if s.ConsultationFeeMax != nil {
		conditions = append(conditions, "consultation_fee <= ?")
		args = append(args, s.ConsultationFeeMax.Amount())
	}

	if s.IsActive != nil {
		conditions = append(conditions, "is_active = ?")
		args = append(args, *s.IsActive)
	}

	if s.HasUserAccount != nil {
		if *s.HasUserAccount {
			conditions = append(conditions, "user_id IS NOT NULL")
		} else {
			conditions = append(conditions, "user_id IS NULL")
		}
	}

	if s.CreatedAfter != nil {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, *s.CreatedAfter)
	}
	if s.CreatedBefore != nil {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, *s.CreatedBefore)
	}

	// Búsqueda por término
	if s.SearchTerm != nil && *s.SearchTerm != "" {
		searchCondition := "(name ILIKE ? OR license_number ILIKE ? OR specialty ILIKE ?)"
		conditions = append(conditions, searchCondition)
		searchArg := "%" + *s.SearchTerm + "%"
		args = append(args, searchArg, searchArg, searchArg)
	}

	where := ""
	if len(conditions) > 0 {
		where = "WHERE " + strings.Join(conditions, " AND ")
	}

	return where, args
}

// Helper methods para disponibilidad
func (s *VeterinarianSpecification) checkAvailability(vet interface {
	Schedule() *valueobject.Schedule
	IsActive() bool
},
) bool {
	if !vet.IsActive() {
		return false
	}

	schedule := vet.Schedule()
	if schedule == nil {
		return false
	}

	return true
}

func (s *VeterinarianSpecification) WithIDs(ids ...valueobject.VetID) *VeterinarianSpecification {
	s.IDs = ids
	return s
}

func (s *VeterinarianSpecification) WithUserIDs(userIDs ...valueobject.UserID) *VeterinarianSpecification {
	s.UserIDs = userIDs
	return s
}

func (s *VeterinarianSpecification) WithLicenseNumbers(licenses ...string) *VeterinarianSpecification {
	s.LicenseNumbers = licenses
	return s
}

func (s *VeterinarianSpecification) WithSpecialties(specialties ...enum.VetSpecialty) *VeterinarianSpecification {
	s.Specialties = specialties
	return s
}

func (s *VeterinarianSpecification) WithYearsExperienceRange(min, max *int) *VeterinarianSpecification {
	s.YearsExperienceMin = min
	s.YearsExperienceMax = max
	return s
}

func (s *VeterinarianSpecification) WithConsultationFeeRange(min, max *valueobject.Money) *VeterinarianSpecification {
	s.ConsultationFeeMin = min
	s.ConsultationFeeMax = max
	return s
}

func (s *VeterinarianSpecification) WithIsActive(isActive bool) *VeterinarianSpecification {
	s.IsActive = &isActive
	return s
}

func (s *VeterinarianSpecification) WithHasUserAccount(hasAccount bool) *VeterinarianSpecification {
	s.HasUserAccount = &hasAccount
	return s
}

func (s *VeterinarianSpecification) WithIsAvailable(isAvailable bool) *VeterinarianSpecification {
	s.IsAvailable = &isAvailable
	return s
}

func (s *VeterinarianSpecification) WithAvailableDate(date time.Time) *VeterinarianSpecification {
	s.AvailableDate = &date
	return s
}

func (s *VeterinarianSpecification) WithSearchTerm(term string) *VeterinarianSpecification {
	s.SearchTerm = &term
	return s
}

func (s *VeterinarianSpecification) WithCreatedDateRange(from, to *time.Time) *VeterinarianSpecification {
	s.CreatedAfter = from
	s.CreatedBefore = to
	return s
}

func (s *VeterinarianSpecification) WithPagination(page, pageSize int, orderBy, sortDir string) *VeterinarianSpecification {
	s.Pagination = Pagination{
		Page:     page,
		PageSize: pageSize,
		OrderBy:  orderBy,
		SortDir:  sortDir,
	}
	return s
}
