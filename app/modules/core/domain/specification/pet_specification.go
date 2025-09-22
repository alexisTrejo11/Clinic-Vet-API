package specification

import (
	"fmt"
	"strings"

	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

// PetSpecification implementa la interfaz Specification para la entidad Pet
type PetSpecification struct {
	filters    []Filter
	pagination Pagination
}

// Filter representa un filtro individual para la consulta
type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}

// NewPetSpecification crea una nueva especificación para mascotas
func NewPetSpecification() *PetSpecification {
	return &PetSpecification{
		filters: make([]Filter, 0),
	}
}

// IsSatisfiedBy verifica si una mascota cumple con los criterios de la especificación
func (s *PetSpecification) IsSatisfiedBy(candidate interface{}) bool {
	pet, ok := candidate.(*pet.Pet)
	if !ok {
		return false
	}

	for _, filter := range s.filters {
		if !s.applyFilter(pet, filter) {
			return false
		}
	}

	return true
}

// applyFilter aplica un filtro individual a una mascota
func (s *PetSpecification) applyFilter(pet *pet.Pet, filter Filter) bool {
	switch filter.Field {
	case "name":
		return strings.Contains(strings.ToLower(pet.Name()), strings.ToLower(filter.Value.(string)))
	case "species":
		return pet.Species().String() == filter.Value.(string)
	case "breed":
		if pet.Breed() == nil {
			return false
		}
		return *pet.Breed() == filter.Value.(string)
	case "age":
		if pet.Age() == nil {
			return false
		}
		return *pet.Age() == filter.Value.(int)
	case "gender":
		if pet.Gender() == nil {
			return false
		}
		return *pet.Gender() == filter.Value.(enum.PetGender)
	case "customer_id":
		return pet.CustomerID().String() == filter.Value.(string)
	case "is_active":
		return pet.IsActive() == filter.Value.(bool)
	case "is_neutered":
		if pet.IsNeutered() == nil {
			return false
		}
		return *pet.IsNeutered() == filter.Value.(bool)
	default:
		return false
	}
}

// ToSQL convierte la especificación a una consulta SQL con parámetros
func (s *PetSpecification) ToSQL() (string, []interface{}) {
	whereClauses := make([]string, 0)
	params := make([]interface{}, 0)
	paramCount := 1

	// Añadir filtros básicos
	for _, filter := range s.filters {
		clause, newParams := s.filterToSQL(filter, &paramCount)
		if clause != "" {
			whereClauses = append(whereClauses, clause)
			params = append(params, newParams...)
		}
	}

	// Construir la consulta base
	query := "SELECT * FROM pets WHERE deleted_at IS NULL"

	// Añadir condiciones WHERE si existen
	if len(whereClauses) > 0 {
		query += " AND " + strings.Join(whereClauses, " AND ")
	}

	// Añadir paginación y ordenamiento
	if s.pagination.OrderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s %s", s.pagination.OrderBy, s.pagination.SortDir)
	}

	if s.pagination.PageSize > 0 {
		offset := (s.pagination.Page - 1) * s.pagination.PageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", s.pagination.PageSize, offset)
	}

	return query, params
}

// filterToSQL convierte un filtro individual a SQL
func (s *PetSpecification) filterToSQL(filter Filter, paramCount *int) (string, []interface{}) {
	switch filter.Operator {
	case "=", "!=", "<", ">", "<=", ">=":
		clause := fmt.Sprintf("%s %s $%d", filter.Field, filter.Operator, *paramCount)
		*paramCount++
		return clause, []interface{}{filter.Value}
	case "LIKE":
		clause := fmt.Sprintf("%s ILIKE $%d", filter.Field, *paramCount)
		*paramCount++
		return clause, []interface{}{"%" + filter.Value.(string) + "%"}
	case "IN":
		values := filter.Value.([]interface{})
		placeholders := make([]string, len(values))
		for i := range values {
			placeholders[i] = fmt.Sprintf("$%d", *paramCount)
			*paramCount++
		}
		clause := fmt.Sprintf("%s IN (%s)", filter.Field, strings.Join(placeholders, ", "))
		return clause, values
	default:
		return "", nil
	}
}

// Métodos de construcción para la especificación

func (s *PetSpecification) WithName(name string) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "name", Operator: "LIKE", Value: name})
	return s
}

func (s *PetSpecification) WithSpecies(species string) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "species", Operator: "=", Value: species})
	return s
}

func (s *PetSpecification) WithBreed(breed string) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "breed", Operator: "=", Value: breed})
	return s
}

func (s *PetSpecification) WithAge(age int) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "age", Operator: "=", Value: age})
	return s
}

func (s *PetSpecification) WithAgeRange(minAge, maxAge int) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "age", Operator: ">=", Value: minAge})
	s.filters = append(s.filters, Filter{Field: "age", Operator: "<=", Value: maxAge})
	return s
}

func (s *PetSpecification) WithGender(gender enum.PetGender) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "gender", Operator: "=", Value: gender})
	return s
}

func (s *PetSpecification) WithCustomerID(customerID valueobject.CustomerID) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "customer_id", Operator: "=", Value: customerID.String()})
	return s
}

func (s *PetSpecification) WithIsActive(isActive bool) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "is_active", Operator: "=", Value: isActive})
	return s
}

func (s *PetSpecification) WithIsNeutered(isNeutered bool) *PetSpecification {
	s.filters = append(s.filters, Filter{Field: "is_neutered", Operator: "=", Value: isNeutered})
	return s
}

func (s *PetSpecification) WithPagination(pagination Pagination) *PetSpecification {
	s.pagination = pagination
	return s
}

// Builder para crear especificaciones de manera fluida
type PetSpecificationBuilder struct {
	spec *PetSpecification
}

func NewPetSpecificationBuilder() *PetSpecificationBuilder {
	return &PetSpecificationBuilder{
		spec: NewPetSpecification(),
	}
}

func (b *PetSpecificationBuilder) Name(name string) *PetSpecificationBuilder {
	b.spec.WithName(name)
	return b
}

func (b *PetSpecificationBuilder) Species(species string) *PetSpecificationBuilder {
	b.spec.WithSpecies(species)
	return b
}

func (b *PetSpecificationBuilder) Breed(breed string) *PetSpecificationBuilder {
	b.spec.WithBreed(breed)
	return b
}

func (b *PetSpecificationBuilder) Age(age int) *PetSpecificationBuilder {
	b.spec.WithAge(age)
	return b
}

func (b *PetSpecificationBuilder) AgeRange(minAge, maxAge int) *PetSpecificationBuilder {
	b.spec.WithAgeRange(minAge, maxAge)
	return b
}

func (b *PetSpecificationBuilder) IsNeutered(isNeutered bool) *PetSpecificationBuilder {
	b.spec.WithIsNeutered(isNeutered)
	return b
}

func (b *PetSpecificationBuilder) Gender(gender enum.PetGender) *PetSpecificationBuilder {
	b.spec.WithGender(gender)
	return b
}

func (b *PetSpecificationBuilder) CustomerID(customerID valueobject.CustomerID) *PetSpecificationBuilder {
	b.spec.WithCustomerID(customerID)
	return b
}

func (b *PetSpecificationBuilder) IsActive(isActive bool) *PetSpecificationBuilder {
	b.spec.WithIsActive(isActive)
	return b
}

func (b *PetSpecificationBuilder) Pagination(pagination Pagination) *PetSpecificationBuilder {
	b.spec.WithPagination(pagination)
	return b
}

func (b *PetSpecificationBuilder) Build() *PetSpecification {
	return b.spec
}
