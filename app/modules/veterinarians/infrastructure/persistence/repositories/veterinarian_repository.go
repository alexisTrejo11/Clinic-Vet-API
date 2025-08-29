package sqlcVetRepo

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
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

func (r *SqlcVetRepository) List(ctx context.Context, searchParams interface{}) ([]entity.Veterinarian, error) {
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

	vets := make([]entity.Veterinarian, len(sqlVets))
	for i, sqlVet := range sqlVets {
		vet, err := SqlcVetToDomain(sqlVet)
		if err != nil {
			return nil, err
		}
		vets[i] = *vet
	}

	return vets, nil
}

func (c *SqlcVetRepository) GetByID(ctx context.Context, id valueobject.VetID) (entity.Veterinarian, error) {
	sqlVet, err := c.queries.GetVeterinarianById(ctx, int32(id.GetValue()))
	if err != nil {
		return entity.Veterinarian{}, err
	}

	vet, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return entity.Veterinarian{}, err
	}

	return *vet, nil
}

func (c *SqlcVetRepository) GetByUserID(ctx context.Context, userID valueobject.UserID) (entity.Veterinarian, error) {
	sqlVet, err := c.queries.GetVeterinarianById(ctx, int32(userID.GetValue()))
	if err != nil {
		return entity.Veterinarian{}, err
	}

	vet, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return entity.Veterinarian{}, err
	}
	return *vet, nil
}

func (c *SqlcVetRepository) Save(ctx context.Context, vet *entity.Veterinarian) error {
	if vet.GetID() == 0 {
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

func (c *SqlcVetRepository) SoftDelete(ctx context.Context, id valueobject.VetID) error {
	if err := c.queries.SoftDeleteVeterinarian(ctx, int32(id.GetValue())); err != nil {
		return err
	}
	return nil
}

func (c *SqlcVetRepository) Exists(ctx context.Context, id valueobject.VetID) (bool, error) {
	_, err := c.queries.GetVeterinarianById(ctx, int32(id.GetValue()))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func (c *SqlcVetRepository) create(ctx context.Context, vet *entity.Veterinarian) error {
	createParams := sqlc.CreateVeterinarianParams{
		FirstName:         vet.GetName().FirstName,
		LastName:          vet.GetName().LastName,
		LicenseNumber:     vet.GetLicenseNumber(),
		Photo:             vet.GetPhoto(),
		Speciality:        models.VeterinarianSpeciality(vet.GetSpecialty().String()),
		YearsOfExperience: int32(vet.GetYearsExperience()),
		IsActive:          pgtype.Bool{Bool: vet.GetIsActive(), Valid: true},
	}

	sqlVet, err := c.queries.CreateVeterinarian(ctx, createParams)
	if err != nil {
		return err
	}

	vet.SetID(int(sqlVet.ID))
	vet.SetCreatedAt(sqlVet.CreatedAt.Time)
	vet.SetUpdatedAt(sqlVet.UpdatedAt.Time)

	return nil
}

func (c *SqlcVetRepository) update(ctx context.Context, vet *entity.Veterinarian) error {
	updateParams := sqlc.UpdateVeterinarianParams{
		ID:                int32(vet.GetID()),
		FirstName:         vet.GetName().FirstName,
		LastName:          vet.GetName().LastName,
		LicenseNumber:     vet.GetLicenseNumber(),
		Photo:             vet.GetPhoto(),
		Speciality:        models.VeterinarianSpeciality(vet.GetSpecialty().String()),
		YearsOfExperience: int32(vet.GetYearsExperience()),
		IsActive:          pgtype.Bool{Bool: vet.GetIsActive(), Valid: true},
	}

	sqlVet, err := c.queries.UpdateVeterinarian(ctx, updateParams)
	if err != nil {
		return err
	}

	vet.SetUpdatedAt(sqlVet.UpdatedAt.Time)

	return nil
}
