package specification

import (
	"fmt"
	"strings"

	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

// EmployeeSearchSpecification implementa la interfaz Specification para búsqueda de veterinarios
type EmployeeSearchSpecification struct {
	Name           *string
	LicenseNumber  *string
	Specialty      *enum.VetSpecialty
	MinExperience  *int
	MaxExperience  *int
	IsActive       *bool
	MinFee         *valueobject.Money
	MaxFee         *valueobject.Money
	HasUserAccount *bool
	Pagination     Pagination
}

// NewEmployeeSearchSpecification crea una nueva especificación de búsqueda
func NewEmployeeSearchSpecification() *EmployeeSearchSpecification {
	return &EmployeeSearchSpecification{
		Pagination: Pagination{
			Offset:  0,
			Limit:   10,
			OrderBy: "created_at",
			SortDir: "DESC",
		},
	}
}

// WithName filtra por nombre (búsqueda parcial case-insensitive)
func (v *EmployeeSearchSpecification) WithName(name string) *EmployeeSearchSpecification {
	v.Name = &name
	return v
}

// WithLicenseNumber filtra por número de licencia (búsqueda exacta)
func (v *EmployeeSearchSpecification) WithLicenseNumber(licenseNumber string) *EmployeeSearchSpecification {
	v.LicenseNumber = &licenseNumber
	return v
}

// WithSpecialty filtra por especialidad
func (v *EmployeeSearchSpecification) WithSpecialty(specialty enum.VetSpecialty) *EmployeeSearchSpecification {
	v.Specialty = &specialty
	return v
}

// WithExperienceRange filtra por rango de años de experiencia
func (v *EmployeeSearchSpecification) WithExperienceRange(min, max int) *EmployeeSearchSpecification {
	v.MinExperience = &min
	v.MaxExperience = &max
	return v
}

// WithMinExperience filtra por años de experiencia mínimos
func (v *EmployeeSearchSpecification) WithMinExperience(min int) *EmployeeSearchSpecification {
	v.MinExperience = &min
	return v
}

// WithMaxExperience filtra por años de experiencia máximos
func (v *EmployeeSearchSpecification) WithMaxExperience(max int) *EmployeeSearchSpecification {
	v.MaxExperience = &max
	return v
}

// WithActiveStatus filtra por estado activo/inactivo
func (v *EmployeeSearchSpecification) WithActiveStatus(isActive bool) *EmployeeSearchSpecification {
	v.IsActive = &isActive
	return v
}

// WithFeeRange filtra por rango de tarifas de consulta
func (v *EmployeeSearchSpecification) WithFeeRange(min, max valueobject.Money) *EmployeeSearchSpecification {
	v.MinFee = &min
	v.MaxFee = &max
	return v
}

// WithMinFee filtra por tarifa mínima de consulta
func (v *EmployeeSearchSpecification) WithMinFee(min valueobject.Money) *EmployeeSearchSpecification {
	v.MinFee = &min
	return v
}

// WithMaxFee filtra por tarifa máxima de consulta
func (v *EmployeeSearchSpecification) WithMaxFee(max valueobject.Money) *EmployeeSearchSpecification {
	v.MaxFee = &max
	return v
}

// WithUserAccount filtra por veterinarios con/without cuenta de usuario
func (v *EmployeeSearchSpecification) WithUserAccount(hasUserAccount bool) *EmployeeSearchSpecification {
	v.HasUserAccount = &hasUserAccount
	return v
}

// WithPagination configura la paginación
func (v *EmployeeSearchSpecification) WithPagination(offset, limit int, orderBy, sortDir string) *EmployeeSearchSpecification {
	v.Pagination = Pagination{
		Offset:  offset,
		Limit:   limit,
		OrderBy: orderBy,
		SortDir: strings.ToUpper(sortDir),
	}
	return v
}

// IsSatisfiedBy verifica si un veterinario cumple con los criterios de búsqueda
func (v *EmployeeSearchSpecification) IsSatisfiedBy(candidate any) bool {
	employee, ok := candidate.(employee.Employee)
	if !ok {
		return false
	}

	// Filtro por nombre
	if v.Name != nil {
		fullName := strings.ToLower(employee.Name().FirstName + " " + employee.Name().LastName)
		searchName := strings.ToLower(*v.Name)
		if !strings.Contains(fullName, searchName) {
			return false
		}
	}

	// Filtro por número de licencia
	if v.LicenseNumber != nil && employee.LicenseNumber() != *v.LicenseNumber {
		return false
	}

	// Filtro por especialidad
	if v.Specialty != nil && employee.Specialty() != *v.Specialty {
		return false
	}

	// Filtro por experiencia mínima
	if v.MinExperience != nil && employee.YearsExperience() < *v.MinExperience {
		return false
	}

	// Filtro por experiencia máxima
	if v.MaxExperience != nil && employee.YearsExperience() > *v.MaxExperience {
		return false
	}

	// Filtro por estado activo
	if v.IsActive != nil && employee.IsActive() != *v.IsActive {
		return false
	}

	// Filtro por tarifa mínima
	if v.MinFee != nil {
		if employee.ConsultationFee() == nil {
			return false
		}
		if employee.ConsultationFee().Amount().LessThan(v.MinFee.Amount()) {
			return false
		}
	}

	// Filtro por tarifa máxima
	if v.MaxFee != nil {
		if employee.ConsultationFee() == nil {
			return false
		}
		if employee.ConsultationFee().Amount().GreaterThan(v.MaxFee.Amount()) {
			return false
		}
	}

	// Filtro por cuenta de usuario
	if v.HasUserAccount != nil {
		hasUser := employee.UserID() != nil
		if hasUser != *v.HasUserAccount {
			return false
		}
	}

	return true
}

// ToSQL convierte la especificación a consulta SQL y parámetros
func (v *EmployeeSearchSpecification) ToSQL() (string, []any) {
	var conditions []string
	var params []any
	paramCount := 1

	// Filtro por nombre (búsqueda ILIKE en first_name y last_name)
	if v.Name != nil {
		searchPattern := fmt.Sprintf("%%%s%%", *v.Name)
		conditions = append(conditions,
			fmt.Sprintf("(first_name ILIKE $%d OR last_name ILIKE $%d)", paramCount, paramCount))
		params = append(params, searchPattern)
		paramCount++
	}

	// Filtro por número de licencia (búsqueda exacta)
	if v.LicenseNumber != nil {
		conditions = append(conditions, fmt.Sprintf("license_number = $%d", paramCount))
		params = append(params, *v.LicenseNumber)
		paramCount++
	}

	// Filtro por especialidad
	if v.Specialty != nil {
		conditions = append(conditions, fmt.Sprintf("speciality = $%d", paramCount))
		params = append(params, v.Specialty.String())
		paramCount++
	}

	// Filtro por experiencia mínima
	if v.MinExperience != nil {
		conditions = append(conditions, fmt.Sprintf("years_of_experience >= $%d", paramCount))
		params = append(params, *v.MinExperience)
		paramCount++
	}

	// Filtro por experiencia máxima
	if v.MaxExperience != nil {
		conditions = append(conditions, fmt.Sprintf("years_of_experience <= $%d", paramCount))
		params = append(params, *v.MaxExperience)
		paramCount++
	}

	// Filtro por estado activo
	if v.IsActive != nil {
		conditions = append(conditions, fmt.Sprintf("is_active = $%d", paramCount))
		params = append(params, *v.IsActive)
		paramCount++
	}

	// Filtro por tarifa mínima
	if v.MinFee != nil {
		conditions = append(conditions, fmt.Sprintf("consultation_fee >= $%d", paramCount))
		params = append(params, v.MinFee.Amount())
		paramCount++
	}

	// Filtro por tarifa máxima
	if v.MaxFee != nil {
		conditions = append(conditions, fmt.Sprintf("consultation_fee <= $%d", paramCount))
		params = append(params, v.MaxFee.Amount())
		paramCount++
	}

	// Filtro por cuenta de usuario
	if v.HasUserAccount != nil {
		if *v.HasUserAccount {
			conditions = append(conditions, "user_id IS NOT NULL")
		} else {
			conditions = append(conditions, "user_id IS NULL")
		}
	}

	// Construir la consulta WHERE
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Ordenamiento
	orderBy := v.getOrderByClause()

	// Paginación
	limitOffset := fmt.Sprintf("LIMIT $%d OFFSET $%d", paramCount, paramCount+1)
	params = append(params, v.Pagination.GetLimit(), v.Pagination.GetOffset())

	// Consulta final
	query := fmt.Sprintf(`
		SELECT id, first_name, last_name, license_number, photo, speciality, 
			   years_of_experience, consultation_fee, is_active, user_id, schedule,
			   created_at, updated_at, deleted_at
		FROM employees 
		%s 
		%s 
		%s`,
		whereClause, orderBy, limitOffset)

	return query, params
}

// getOrderByClause devuelve la cláusula ORDER BY basada en la paginación
func (v *EmployeeSearchSpecification) getOrderByClause() string {
	orderBy := v.Pagination.OrderBy
	sortDir := v.Pagination.SortDir

	if sortDir != "ASC" && sortDir != "DESC" {
		sortDir = "DESC"
	}

	// Mapear campos de ordenamiento
	switch orderBy {
	case "name":
		return fmt.Sprintf("ORDER BY first_name %s, last_name %s", sortDir, sortDir)
	case "specialty":
		return fmt.Sprintf("ORDER BY speciality %s", sortDir)
	case "experience":
		return fmt.Sprintf("ORDER BY years_of_experience %s", sortDir)
	case "fee":
		return fmt.Sprintf("ORDER BY consultation_fee %s", sortDir)
	case "created_at", "updated_at":
		return fmt.Sprintf("ORDER BY %s %s", orderBy, sortDir)
	default:
		return "ORDER BY created_at DESC"
	}
}

// GetPagination retorna la información de paginación
func (v *EmployeeSearchSpecification) GetPagination() Pagination {
	return v.Pagination
}
