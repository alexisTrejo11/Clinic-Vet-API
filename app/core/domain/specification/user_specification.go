package specification

import (
	"strings"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type UserSpecification struct {
	IDs             []valueobject.UserID
	Emails          []valueobject.Email
	Roles           []enum.UserRole
	Statuses        []enum.UserStatus
	IsActive        *bool
	CreatedAfter    *time.Time
	CreatedBefore   *time.Time
	LastLoginAfter  *time.Time
	LastLoginBefore *time.Time
	SearchTerm      *string // email, nombre, etc.
	HasTwoFactor    *bool
	Pagination
}

func (s *UserSpecification) IsSatisfiedBy(entity any) bool {
	user, ok := entity.(interface {
		ID() valueobject.UserID
		Email() valueobject.Email
		Role() enum.UserRole
		Status() enum.UserStatus
		IsActive() bool
		CreatedAt() time.Time
		LastLoginAt() *time.Time
		TwoFactorAuth() valueobject.TwoFactorAuth
	})
	if !ok {
		return false
	}

	if len(s.IDs) > 0 {
		found := false
		for _, id := range s.IDs {
			if id.Value() == user.ID().Value() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.Emails) > 0 {
		found := false
		for _, email := range s.Emails {
			if email.String() == user.Email().String() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if len(s.Roles) > 0 {
		found := false
		for _, role := range s.Roles {
			if role == user.Role() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// Filtro por Statuses
	if len(s.Statuses) > 0 {
		found := false
		for _, status := range s.Statuses {
			if status == user.Status() {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	if s.IsActive != nil && user.IsActive() != *s.IsActive {
		return false
	}

	if s.CreatedAfter != nil && user.CreatedAt().Before(*s.CreatedAfter) {
		return false
	}
	if s.CreatedBefore != nil && user.CreatedAt().After(*s.CreatedBefore) {
		return false
	}

	if s.LastLoginAfter != nil {
		if user.LastLoginAt() == nil || user.LastLoginAt().Before(*s.LastLoginAfter) {
			return false
		}
	}
	if s.LastLoginBefore != nil {
		if user.LastLoginAt() == nil || user.LastLoginAt().After(*s.LastLoginBefore) {
			return false
		}
	}

	if s.HasTwoFactor != nil {
		has2FA := user.TwoFactorAuth().IsEnabled
		if has2FA != *s.HasTwoFactor {
			return false
		}
	}

	if s.SearchTerm != nil && *s.SearchTerm != "" {
		if userEntity, ok := user.(interface {
			FirstName() string
			LastName() string
		}); ok {
			searchTerm := strings.ToLower(*s.SearchTerm)
			matches := strings.Contains(strings.ToLower(user.Email().String()), searchTerm) ||
				strings.Contains(strings.ToLower(userEntity.FirstName()), searchTerm) ||
				strings.Contains(strings.ToLower(userEntity.LastName()), searchTerm)

			if !matches {
				return false
			}
		}
	}

	return true
}

func (s *UserSpecification) ToSQL() (string, []any) {
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

	// Filtro por Emails
	if len(s.Emails) > 0 {
		emails := make([]any, len(s.Emails))
		for i, email := range s.Emails {
			emails[i] = email.String()
		}
		conditions = append(conditions, "email IN (?)")
		args = append(args, emails)
	}

	if len(s.Roles) > 0 {
		roles := make([]any, len(s.Roles))
		for i, role := range s.Roles {
			roles[i] = role.String()
		}
		conditions = append(conditions, "role IN (?)")
		args = append(args, roles)
	}

	if len(s.Statuses) > 0 {
		statuses := make([]any, len(s.Statuses))
		for i, status := range s.Statuses {
			statuses[i] = status.String()
		}
		conditions = append(conditions, "status IN (?)")
		args = append(args, statuses)
	}

	if s.IsActive != nil {
		conditions = append(conditions, "is_active = ?")
		args = append(args, *s.IsActive)
	}

	if s.CreatedAfter != nil {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, *s.CreatedAfter)
	}
	if s.CreatedBefore != nil {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, *s.CreatedBefore)
	}

	if s.LastLoginAfter != nil {
		conditions = append(conditions, "last_login_at >= ?")
		args = append(args, *s.LastLoginAfter)
	}
	if s.LastLoginBefore != nil {
		conditions = append(conditions, "last_login_at <= ?")
		args = append(args, *s.LastLoginBefore)
	}

	if s.HasTwoFactor != nil {
		conditions = append(conditions, "two_factor_enabled = ?")
		args = append(args, *s.HasTwoFactor)
	}

	// Búsqueda por término
	if s.SearchTerm != nil && *s.SearchTerm != "" {
		searchCondition := "(email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ?)"
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

func (s *UserSpecification) WithIDs(ids ...valueobject.UserID) *UserSpecification {
	s.IDs = ids
	return s
}

func (s *UserSpecification) WithEmails(emails ...valueobject.Email) *UserSpecification {
	s.Emails = emails
	return s
}

func (s *UserSpecification) WithRoles(roles ...enum.UserRole) *UserSpecification {
	s.Roles = roles
	return s
}

func (s *UserSpecification) WithStatuses(statuses ...enum.UserStatus) *UserSpecification {
	s.Statuses = statuses
	return s
}

func (s *UserSpecification) WithIsActive(isActive bool) *UserSpecification {
	s.IsActive = &isActive
	return s
}

func (s *UserSpecification) WithCreatedDateRange(from, to *time.Time) *UserSpecification {
	s.CreatedAfter = from
	s.CreatedBefore = to
	return s
}

func (s *UserSpecification) WithLastLoginRange(from, to *time.Time) *UserSpecification {
	s.LastLoginAfter = from
	s.LastLoginBefore = to
	return s
}

func (s *UserSpecification) WithSearchTerm(term string) *UserSpecification {
	s.SearchTerm = &term
	return s
}

func (s *UserSpecification) WithTwoFactor(hasTwoFactor bool) *UserSpecification {
	s.HasTwoFactor = &hasTwoFactor
	return s
}

func (s *UserSpecification) WithPagination(page, pageSize int, orderBy, sortDir string) *UserSpecification {
	s.Pagination = Pagination{
		Page:     page,
		PageSize: pageSize,
		OrderBy:  orderBy,
		SortDir:  sortDir,
	}
	return s
}
