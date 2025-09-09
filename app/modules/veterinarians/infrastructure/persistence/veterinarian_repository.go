package persistence

import (
	"context"
	"database/sql"
	"errors"
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
		Offset:              int32((searchParam.PageNumber - 1) * searchParam.PageSize),
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
			orderParams[0] = true
		} else {
			orderParams[1] = true
		}
	case "specialty":
		if searchParam.SortDirection == page.ASC {
			orderParams[2] = true
		} else {
			orderParams[3] = true
		}
	case "years_experience":
		if searchParam.SortDirection == page.ASC {
			orderParams[4] = true
		} else {
			orderParams[5] = true
		}
	case "created_at":
		if searchParam.SortDirection == page.ASC {
			orderParams[6] = true
		} else {
			orderParams[7] = true
		}
	default:
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
		return nil, r.dbError(OpSelect, ErrMsgListVets, err)
	}

	vets := make([]vet.Veterinarian, len(sqlVets))
	for i, sqlVet := range sqlVets {
		vetEntity, err := SqlcVetToDomain(sqlVet)
		if err != nil {
			return nil, r.wrapConversionError(err)
		}
		vets[i] = *vetEntity
	}

	return vets, nil
}

func (r *SqlcVetRepository) GetByID(ctx context.Context, id valueobject.VetID) (vet.Veterinarian, error) {
	sqlVet, err := r.queries.GetVeterinarianById(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vet.Veterinarian{}, r.notFoundError("id", id.String())
		}
		return vet.Veterinarian{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgGetVet, id.Value()), err)
	}

	veterinarian, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return vet.Veterinarian{}, r.wrapConversionError(err)
	}

	return *veterinarian, nil
}

func (r *SqlcVetRepository) GetByUserID(ctx context.Context, userID valueobject.UserID) (vet.Veterinarian, error) {
	sqlVet, err := r.queries.GetVeterinarianById(ctx, int32(userID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vet.Veterinarian{}, r.notFoundError("user_id", userID.String())
		}
		return vet.Veterinarian{}, r.dbError(OpSelect, fmt.Sprintf("%s with user ID %d", ErrMsgGetVetByUserID, userID.Value()), err)
	}

	veterinarian, err := SqlcVetToDomain(sqlVet)
	if err != nil {
		return vet.Veterinarian{}, r.wrapConversionError(err)
	}
	return *veterinarian, nil
}

func (r *SqlcVetRepository) Save(ctx context.Context, veterinarian *vet.Veterinarian) error {
	if veterinarian.ID().IsZero() {
		return r.create(ctx, veterinarian)
	}
	return r.update(ctx, veterinarian)
}

func (r *SqlcVetRepository) SoftDelete(ctx context.Context, id valueobject.VetID) error {
	if err := r.queries.SoftDeleteVeterinarian(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgSoftDeleteVet, id.Value()), err)
	}
	return nil
}

func (r *SqlcVetRepository) Exists(ctx context.Context, id valueobject.VetID) (bool, error) {
	_, err := r.queries.GetVeterinarianById(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgCheckVetExists, id.Value()), err)
	}
	return true, nil
}

func (r *SqlcVetRepository) create(ctx context.Context, vet *vet.Veterinarian) error {
	createParams := sqlc.CreateVeterinarianParams{
		FirstName:         vet.Name().FirstName,
		LastName:          vet.Name().LastName,
		LicenseNumber:     vet.LicenseNumber(),
		Photo:             vet.Photo(),
		Speciality:        models.VeterinarianSpeciality(vet.Specialty().String()),
		YearsOfExperience: int32(vet.YearsExperience()),
		IsActive:          pgtype.Bool{Bool: vet.IsActive(), Valid: true},
	}

	_, err := r.queries.CreateVeterinarian(ctx, createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateVet, err)
	}

	return nil
}

func (r *SqlcVetRepository) update(ctx context.Context, vet *vet.Veterinarian) error {
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

	_, err := r.queries.UpdateVeterinarian(ctx, updateParams)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s with ID %d", ErrMsgUpdateVet, vet.ID().Value()), err)
	}
	return nil
}
