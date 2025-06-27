package sqlcVetRepo

import (
	"context"

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

func (c *SqlcVetRepository) List(ctx context.Context, searchParams map[string]interface{}) ([]vetDomain.Veterinarian, error) {
	listParams := sqlc.ListVeterinariansParams{Limit: 10, Offset: 0}
	sqlVetList, err := c.queries.ListVeterinarians(ctx, listParams)
	if err != nil {
		return []vetDomain.Veterinarian{}, err
	}

	entitiesList := make([]vetDomain.Veterinarian, len(sqlVetList))
	for i, sqlVet := range sqlVetList {
		entitiesList[i] = *SqlcVetToDomain(sqlVet)
	}

	return entitiesList, nil
}

func (c *SqlcVetRepository) GetByID(ctx context.Context, id uint) (vetDomain.Veterinarian, error) {
	sqlVet, err := c.queries.GetVeterinarianById(ctx, int32(id))
	if err != nil {
		return vetDomain.Veterinarian{}, err
	}
	return *SqlcVetToDomain(sqlVet), nil
}

func (c *SqlcVetRepository) GetByUserID(ctx context.Context, id uint) (vetDomain.Veterinarian, error) {
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

func (c *SqlcVetRepository) Delete(ctx context.Context, id uint, isSoftDelete bool) error {
	if err := c.queries.SoftDeleteVeterinarian(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (c *SqlcVetRepository) Exists(ctx context.Context, vetId uint) (bool, error) {
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
		Speciality:        sqlc.VeterinarianSpeciality(vet.Specialty.String()),
		YearsOfExperience: int32(vet.YearsExperience),
		IsActive:          pgtype.Bool{Bool: vet.IsActive, Valid: true},
	}

	sqlVet, err := c.queries.CreateVeterinarian(ctx, createParams)
	if err != nil {
		return err
	}

	vet.ID = uint(sqlVet.ID)
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
		Speciality:        sqlc.VeterinarianSpeciality(vet.Specialty.String()),
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
