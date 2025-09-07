package persistence

import (
	"context"
	"fmt"

	vet "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/veterinarian"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcVetRepository struct {
	queries *sqlc.Queries
}

func NewSqlcVetRepository(queries *sqlc.Queries) repository.VetRepository {
	return &SqlcVetRepository{queries: queries}
}

func (r *SqlcVetRepository) List(ctx context.Context, searchParams interface{}) ([]vet.Veterinarian, error) {
	searchParam := searchParams.(dto.VetSearchParams)

	params := sqlc.ListVeterinariansParams{
		FirstName:           "%",
		LastName:            "%",
		LicenseNumber:       "%",
		Speciality:          "",
		YearsOfExperience:   0,
		YearsOfExperience_2: 0,
		IsActive:            pgtype.Bool{Bool: false, Valid: false},
		Limit:               int32(searchParam.PageSize),
		Offset:              int32(searchParam.PageNumber - 1),
	}

	// Apply Filters
	if searchParam.Filters.Name != nil {
		name := "%" + *searchParam.Filters.Name + "%"
		params.FirstName = name
		params.LastName = name
	}

	if searchParam.Filters.LicenseNumber != nil {
		params.LicenseNumber = "%" + *searchParam.Filters.LicenseNumber + "%"
	}

	if searchParam.Filters.Specialty != nil {
		params.Speciality = models.VeterinarianSpeciality(searchParam.Filters.Specialty.String())
	}

	if searchParam.Filters.YearsExperience != nil {
		if searchParam.Filters.YearsExperience.Min != nil {
			params.YearsOfExperience = int32(*searchParam.Filters.YearsExperience.Min)
		}
		if searchParam.Filters.YearsExperience.Max != nil {
			params.YearsOfExperience_2 = int32(*searchParam.Filters.YearsExperience.Max)
		}
	}

	if searchParam.Filters.IsActive != nil {
		params.IsActive = pgtype.Bool{Bool: *searchParam.Filters.IsActive, Valid: true}
	}

	// Ordering Config
	orderParams := make([]interface{}, 8)
	for i := range orderParams {
		orderParams[i] = false
	}

	switch searchParam.OrderBy {
	case "name":
		if searchParam.SortDirection == page.ASC {
			orderParams[0] = true // $8: Order by first_name ASC
		} else {
			orderParams[1] = true // $9: Order by first_name DESC
		}
	case "specialty":
		if searchParam.SortDirection == page.ASC {
			orderParams[2] = true // $10: Order by speciality ASC
		} else {
			orderParams[3] = true // $11: Order by speciality DESC
		}
	case "years_experience":
		if searchParam.SortDirection == page.ASC {
			orderParams[4] = true // $12: Order by years_of_experience ASC
		} else {
			orderParams[5] = true // $13: Order by years_of_experience DESC
		}
	case "created_at":
		if searchParam.SortDirection == page.ASC {
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

	vets := make([]vet.Veterinarian, len(sqlVets))
	for i, sqlVet := range sqlVets {
		vet, err := SqlcVetToDomain(sqlVet)
		if err != nil {
			return nil, err
		}
		vets[i] = *vet
	}

	return vets, nil
}

func (c *SqlcVetRepository) GetByID(ctx context.Context, id valueobject.VetID) (vet.Veterinarian, error) {
	sqlVet, err := c.queries.GetVeterinarianById(ctx, int32(id.Value()))
	if err != nil {
		return vet.Veterinarian{}, err
	}

	veterinarian, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return vet.Veterinarian{}, err
	}

	return *veterinarian, nil
}

func (c *SqlcVetRepository) GetByUserID(ctx context.Context, userID valueobject.UserID) (vet.Veterinarian, error) {
	sqlVet, err := c.queries.GetVeterinarianById(ctx, int32(userID.Value()))
	if err != nil {
		return vet.Veterinarian{}, err
	}

	veterinarian, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return vet.Veterinarian{}, err
	}
	return *veterinarian, nil
}

func (c *SqlcVetRepository) Save(ctx context.Context, veterinarian *vet.Veterinarian) error {
	if veterinarian.ID().IsZero() {
		if err := c.create(ctx, veterinarian); err != nil {
			return err
		}
	} else {
		if err := c.update(ctx, veterinarian); err != nil {
			return err
		}
	}
	return nil
}

func (c *SqlcVetRepository) SoftDelete(ctx context.Context, id valueobject.VetID) error {
	if err := c.queries.SoftDeleteVeterinarian(ctx, int32(id.Value())); err != nil {
		return err
	}
	return nil
}

func (c *SqlcVetRepository) Exists(ctx context.Context, id valueobject.VetID) (bool, error) {
	_, err := c.queries.GetVeterinarianById(ctx, int32(id.Value()))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (c *SqlcVetRepository) create(ctx context.Context, vet *vet.Veterinarian) error {
	createParams := sqlc.CreateVeterinarianParams{
		FirstName:         vet.Name().FirstName,
		LastName:          vet.Name().LastName,
		LicenseNumber:     vet.LicenseNumber(),
		Photo:             vet.Photo(),
		Speciality:        models.VeterinarianSpeciality(vet.Specialty().String()),
		YearsOfExperience: int32(vet.YearsExperience()),
		IsActive:          pgtype.Bool{Bool: vet.IsActive(), Valid: true},
	}

	_, err := c.queries.CreateVeterinarian(ctx, createParams)
	if err != nil {
		return err
	}

	return nil
}

func (c *SqlcVetRepository) update(ctx context.Context, vet *vet.Veterinarian) error {
	updateParams := sqlc.UpdateVeterinarianParams{
		ID:                int32(vet.ID().Value()),
		FirstName:         vet.Name().FirstName,
		LastName:          vet.Name().LastName,
		LicenseNumber:     vet.LicenseNumber(),
		Photo:             vet.Photo(),
		Speciality:        models.VeterinarianSpeciality(vet.Specialty().String()),
		YearsOfExperience: int32(vet.YearsExperience()),
		IsActive:          pgtype.Bool{Bool: vet.IsActive(), Valid: true},
	}

	_, err := c.queries.UpdateVeterinarian(ctx, updateParams)
	if err != nil {
		return err
	}
	return nil
}
