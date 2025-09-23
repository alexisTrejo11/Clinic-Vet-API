// Package repository defines the persistence layer implementations.
package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/sqlc"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcPetRepository struct {
	queries *sqlc.Queries
}

func NewSqlcPetRepository(queries *sqlc.Queries) repository.PetRepository {
	return &SqlcPetRepository{
		queries: queries,
	}
}

func (r *SqlcPetRepository) FindBySpecification(ctx context.Context, spec specification.PetSpecification) (page.Page[pet.Pet], error) {
	return page.Page[pet.Pet]{}, errors.New("FindBySpecification not implemented")
}

func (r *SqlcPetRepository) FindBySpecies(ctx context.Context, petSpecies enum.PetSpecies, pageInput page.PageInput) (page.Page[pet.Pet], error) {
	offset := (pageInput.Offset) * pageInput.Limit
	petRows, err := r.queries.FindPetsBySpecies(ctx, sqlc.FindPetsBySpeciesParams{
		Species: petSpecies.String(),
		Limit:   int32(pageInput.Limit),
		Offset:  int32(offset),
	})

	if err != nil {
		return page.Page[pet.Pet]{}, r.dbError("select", fmt.Sprintf("failed to find pets for species %s", petSpecies.String()), err)
	}

	total, err := r.queries.CountPetsBySpecies(ctx, petSpecies.String())
	if err != nil {
		return page.Page[pet.Pet]{}, r.dbError("select", fmt.Sprintf("failed to count pets for species %s", petSpecies.String()), err)
	}

	pets, err := sqlcRowsToEntities(petRows)
	if err != nil {
		return page.Page[pet.Pet]{}, r.wrapConversionError(err)
	}

	return page.NewPage(pets, *page.GetPageMetadata(int(total), pageInput)), nil
}

func (r *SqlcPetRepository) FindAllByCustomerID(ctx context.Context, customerID valueobject.CustomerID) ([]pet.Pet, error) {
	petRows, err := r.queries.FindPetsByCustomerID(ctx, sqlc.FindPetsByCustomerIDParams{
		CustomerID: int32(customerID.Value()),
		Limit:      1000, // Large limit to get all pets
		Offset:     0,
	})
	if err != nil {
		return nil, r.dbError("select", fmt.Sprintf("failed to find all pets for customer ID %d", customerID.Value()), err)
	}

	// Convert rows to entities
	var pets []pet.Pet
	for _, row := range petRows {
		petEntity, err := sqlcRowToEntity(row)
		if err != nil {
			return nil, r.wrapConversionError(err)
		}
		pets = append(pets, *petEntity)
	}

	return pets, nil
}

func (r *SqlcPetRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pageInput page.PageInput) (page.Page[pet.Pet], error) {
	offset := (pageInput.Offset) * pageInput.Limit
	petRows, err := r.queries.FindPetsByCustomerID(ctx, sqlc.FindPetsByCustomerIDParams{
		CustomerID: int32(customerID.Value()),
		Limit:      int32(pageInput.Limit),
		Offset:     int32(offset),
	})
	if err != nil {
		return page.Page[pet.Pet]{}, r.dbError("select", fmt.Sprintf("failed to find pets for customer ID %d", customerID.Value()), err)
	}

	total, err := r.queries.CountPetsByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return page.Page[pet.Pet]{}, r.dbError("select", fmt.Sprintf("failed to count pets for customer ID %d", customerID.Value()), err)
	}

	pets, _ := sqlcRowsToEntities(petRows)

	return page.NewPage(pets, *page.GetPageMetadata(int(total), pageInput)), nil
}

func (r *SqlcPetRepository) FindByID(ctx context.Context, petID valueobject.PetID) (pet.Pet, error) {
	sqlPet, err := r.queries.FindPetByID(ctx, int32(petID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pet.Pet{}, r.notFoundError("id", petID.String())
		}
		return pet.Pet{}, r.dbError("select", fmt.Sprintf("failed to find pet with ID %d", petID.Value()), err)
	}

	domainPet, err := sqlcRowToEntity(sqlPet)
	if err != nil {
		return pet.Pet{}, r.wrapConversionError(err)
	}

	return *domainPet, nil
}

func (r *SqlcPetRepository) FindByIDAndCustomerID(ctx context.Context, id valueobject.PetID, customerID valueobject.CustomerID) (pet.Pet, error) {
	sqlPet, err := r.queries.FindPetByIDAndCustomerID(ctx, sqlc.FindPetByIDAndCustomerIDParams{
		ID:         int32(id.Value()),
		CustomerID: int32(customerID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pet.Pet{}, r.notFoundError("id and customer_id", fmt.Sprintf("pet %d for customer %d", id.Value(), customerID.Value()))
		}
		return pet.Pet{}, r.dbError("select", fmt.Sprintf("failed to find pet with ID %d and customer ID %d", id.Value(), customerID.Value()), err)
	}

	domainPet, err := sqlcRowToEntity(sqlPet)
	if err != nil {
		return pet.Pet{}, r.wrapConversionError(err)
	}

	return *domainPet, nil
}

func (r *SqlcPetRepository) ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error) {
	exists, err := r.queries.ExistsPetByID(ctx, int32(petID.Value()))
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check pet existence by ID %d", petID.Value()), err)
	}
	return exists, nil
}

func (r *SqlcPetRepository) ExistsByMicrochip(ctx context.Context, microchip string) (bool, error) {
	exists, err := r.queries.ExistsPetByMicrochip(ctx, pgtype.Text{String: microchip, Valid: true})
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check pet existence by microchip %s", microchip), err)
	}
	return exists, nil
}

func (r *SqlcPetRepository) Save(ctx context.Context, pet pet.Pet) (pet.Pet, error) {
	if pet.ID().IsZero() {
		return r.create(ctx, pet)
	}
	return r.update(ctx, pet)
}

func (r *SqlcPetRepository) Delete(ctx context.Context, petID valueobject.PetID, isHard bool) error {
	if isHard {
		return r.hardDelete(ctx, petID)
	}
	return r.softDelete(ctx, petID)
}

func (r *SqlcPetRepository) CountByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (int64, error) {
	count, err := r.queries.CountPetsByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return 0, r.dbError("select", fmt.Sprintf("failed to count pets for customer ID %d", customerID.Value()), err)
	}
	return count, nil
}

// create inserts a new pet
func (r *SqlcPetRepository) create(ctx context.Context, entity pet.Pet) (pet.Pet, error) {
	params := ToSqlCreateParam(entity)

	petCreated, err := r.queries.CreatePet(ctx, *params)
	if err != nil {
		return pet.Pet{}, r.dbError("insert", "failed to create pet", err)
	}

	petEntity, err := sqlcRowToEntity(petCreated)
	if err != nil {
		return pet.Pet{}, r.wrapConversionError(err)
	}

	return *petEntity, nil
}

// update modifies an existing pet
func (r *SqlcPetRepository) update(ctx context.Context, entity pet.Pet) (pet.Pet, error) {
	params := ToSqlUpdateParam(&entity)

	petUpdated, err := r.queries.UpdatePet(ctx, *params)
	if err != nil {
		return pet.Pet{}, r.dbError("update", fmt.Sprintf("failed to update pet with ID %d", entity.ID().Value()), err)
	}

	petEntity, err := sqlcRowToEntity(petUpdated)
	if err != nil {
		return pet.Pet{}, r.wrapConversionError(err)
	}

	return *petEntity, nil
}

func (r *SqlcPetRepository) Restore(ctx context.Context, petID valueobject.PetID) error {
	if err := r.queries.RestorePet(ctx, int32(petID.Value())); err != nil {
		return r.dbError("update", fmt.Sprintf("failed to restore pet with ID %d", petID.Value()), err)
	}
	return nil
}

func (r *SqlcPetRepository) softDelete(ctx context.Context, petID valueobject.PetID) error {
	if err := r.queries.SoftDeletePet(ctx, int32(petID.Value())); err != nil {
		return r.dbError("delete", fmt.Sprintf("failed to soft delete pet with ID %d", petID.Value()), err)
	}
	return nil
}

func (r *SqlcPetRepository) hardDelete(ctx context.Context, petID valueobject.PetID) error {
	if err := r.queries.HardDeletePet(ctx, int32(petID.Value())); err != nil {
		return r.dbError("delete", fmt.Sprintf("failed to hard delete pet with ID %d", petID.Value()), err)
	}
	return nil
}
