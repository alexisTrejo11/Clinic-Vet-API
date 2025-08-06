package sqlcVetRepo

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcVetRepository struct {
	queries sqlc.Querier
}

func NewSqlcVetRepository(queries sqlc.Querier) vetRepo.VeterinarianRepository {
	return &SqlcVetRepository{queries: queries}
}

func (r *SqlcVetRepository) List(ctx context.Context, searchParams vetDtos.VetSearchParams) ([]vetDomain.Veterinarian, error) {
	params := sqlc.ListVeterinariansParams{
		FirstName:           "%",
		LastName:            "%",
		LicenseNumber:       "%",
		Speciality:          "",
		YearsOfExperience:   0,
		YearsOfExperience_2: 0,
		IsActive:            pgtype.Bool{Bool: false, Valid: false},
		Limit:               int32(searchParams.PageSize),
		Offset:              int32(searchParams.PageNumber - 1),
	}

	// Apply Filters
	if searchParams.Filters.Name != nil {
		name := "%" + *searchParams.Filters.Name + "%"
		params.FirstName = name
		params.LastName = name
	}

	if searchParams.Filters.LicenseNumber != nil {
		params.LicenseNumber = "%" + *searchParams.Filters.LicenseNumber + "%"
	}

	if searchParams.Filters.Specialty != nil {
		params.Speciality = searchParams.Filters.Specialty.String()
	}

	if searchParams.Filters.YearsExperience != nil {
		if searchParams.Filters.YearsExperience.Min != nil {
			params.YearsOfExperience = int32(*searchParams.Filters.YearsExperience.Min)
		}
		if searchParams.Filters.YearsExperience.Max != nil {
			params.YearsOfExperience_2 = int32(*searchParams.Filters.YearsExperience.Max)
		}
	}

	if searchParams.Filters.IsActive != nil {
		params.IsActive = pgtype.Bool{Bool: *searchParams.Filters.IsActive, Valid: true}
	}

	// Ordering Config
	orderParams := make([]interface{}, 8)
	for i := range orderParams {
		orderParams[i] = false
	}

	switch searchParams.OrderBy {
	case "name":
		if searchParams.SortDirection == page.ASC {
			orderParams[0] = true // $8: Order by first_name ASC
		} else {
			orderParams[1] = true // $9: Order by first_name DESC
		}
	case "specialty":
		if searchParams.SortDirection == page.ASC {
			orderParams[2] = true // $10: Order by speciality ASC
		} else {
			orderParams[3] = true // $11: Order by speciality DESC
		}
	case "years_experience":
		if searchParams.SortDirection == page.ASC {
			orderParams[4] = true // $12: Order by years_of_experience ASC
		} else {
			orderParams[5] = true // $13: Order by years_of_experience DESC
		}
	case "created_at":
		if searchParams.SortDirection == page.ASC {
			orderParams[6] = true // $14: Order by created_at ASC
		} else {
			orderParams[7] = true // $15: Order by created_at DESC
		}
	default:
		// Default ordering: for recents first
		orderParams[7] = true
	}

	params.Column8 = orderParams[0]
	params.Column9 = orderParams[1]
	params.Column10 = orderParams[2]
	params.Column11 = orderParams[3]
	params.Column12 = orderParams[4]
	params.Column13 = orderParams[5]
	params.Column14 = orderParams[6]
	params.Column15 = orderParams[7]

	sqlVets, err := r.queries.ListVeterinarians(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list veterinarians: %w", err)
	}

	vets := make([]vetDomain.Veterinarian, len(sqlVets))
	for i, sqlVet := range sqlVets {
		vets[i] = *SqlcVetToDomain(sqlVet)
	}

	return vets, nil
}

func (c *SqlcVetRepository) GetById(ctx context.Context, id int) (vetDomain.Veterinarian, error) {
	sqlVet, err := c.queries.GetVeterinarianById(ctx, int32(id))
	if err != nil {
		return vetDomain.Veterinarian{}, err
	}
	return *SqlcVetToDomain(sqlVet), nil
}

func (c *SqlcVetRepository) GetByUserId(ctx context.Context, id int) (vetDomain.Veterinarian, error) {
	sqlVet, err := c.queries.GetVeterinarianById(ctx, int32(id))
	if err != nil {
		return vetDomain.Veterinarian{}, err
	}

	return *SqlcVetToDomain(sqlVet), nil
}

func (c *SqlcVetRepository) Save(ctx context.Context, vet *vetDomain.Veterinarian) error {
	if vet.ID == 0 {
		if err := c.create(ctx, vet); err != nil {
			return err
		}
	} else {
		if err := c.update(ctx, vet); err != nil {
			return err
		}
	}
	return nil
}

func (c *SqlcVetRepository) Delete(ctx context.Context, id int) error {
	if err := c.queries.SoftDeleteVeterinarian(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (c *SqlcVetRepository) Exists(ctx context.Context, vetId int) (bool, error) {
	_, err := c.queries.GetVeterinarianById(ctx, int32(vetId))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (c *SqlcVetRepository) create(ctx context.Context, vet *vetDomain.Veterinarian) error {
	createParams := sqlc.CreateVeterinarianParams{
		FirstName:         vet.Name.FirstName(),
		LastName:          vet.Name.LastName(),
		LicenseNumber:     vet.LicenseNumber,
		Photo:             vet.Photo,
		Speciality:        vet.Specialty.String(),
		YearsOfExperience: int32(vet.YearsExperience),
		IsActive:          pgtype.Bool{Bool: vet.IsActive, Valid: true},
	}

	sqlVet, err := c.queries.CreateVeterinarian(ctx, createParams)
	if err != nil {
		return err
	}

	vet.ID = int(sqlVet.ID)
	vet.CreatedAt = sqlVet.CreatedAt.Time
	vet.UpdatedAt = sqlVet.UpdatedAt.Time

	return nil
}

func (c *SqlcVetRepository) update(ctx context.Context, vet *vetDomain.Veterinarian) error {
	updateParams := sqlc.UpdateVeterinarianParams{
		ID:                int32(vet.ID),
		FirstName:         vet.Name.FirstName(),
		LastName:          vet.Name.LastName(),
		LicenseNumber:     vet.LicenseNumber,
		Photo:             vet.Photo,
		Speciality:        vet.Specialty.String(),
		YearsOfExperience: int32(vet.YearsExperience),
		IsActive:          pgtype.Bool{Bool: vet.IsActive, Valid: true},
	}

	sqlVet, err := c.queries.UpdateVeterinarian(ctx, updateParams)
	if err != nil {
		return err
	}

	vet.UpdatedAt = sqlVet.UpdatedAt.Time

	return nil
}
