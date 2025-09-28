package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/sqlc"
)

func (r *SqlcPetRepository) toEntity(sqlPet sqlc.Pet) pet.Pet {
	petID := valueobject.NewPetID(uint(sqlPet.ID))
	customerID := valueobject.NewCustomerID(uint(sqlPet.CustomerID))
	photo := r.pgMap.PgText.ToStringPtr(sqlPet.Photo)
	breed := r.pgMap.PgText.ToStringPtr(sqlPet.Breed)
	color := r.pgMap.PgText.ToStringPtr(sqlPet.Color)
	microchip := r.pgMap.PgText.ToStringPtr(sqlPet.Microchip)
	isNeutered := r.pgMap.PgBool.ToBoolPtr(sqlPet.IsNeutered)
	age := r.pgMap.PgInt2.ToIntPtr(sqlPet.Age)

	return *pet.NewPetBuilder().
		WithID(petID).
		WithCustomerID(customerID).
		WithName(sqlPet.Name).
		WithSpecies(enum.PetSpecies(sqlPet.Species)).
		WithIsActive(sqlPet.IsActive).
		WithPhoto(photo).
		WithBreed(breed).
		WithAge(age).
		WithGender(enum.PetGender(sqlPet.Gender.String)).
		WithColor(color).
		WithMicrochip(microchip).
		WithIsNeutered(isNeutered).
		Build()
}

func (r *SqlcPetRepository) toEntities(rows []sqlc.Pet) []pet.Pet {
	var pets []pet.Pet

	for _, row := range rows {
		petEntity := r.toEntity(row)
		pets = append(pets, petEntity)
	}

	return pets
}

func (r *SqlcPetRepository) toCreateParams(pet pet.Pet) *sqlc.CreatePetParams {
	return &sqlc.CreatePetParams{
		Name:       pet.Name(),
		Photo:      r.pgMap.PgText.FromStringPtr(pet.Photo()),
		Species:    pet.Species().String(),
		Breed:      r.pgMap.PgText.FromStringPtr(pet.Breed()),
		Age:        r.pgMap.PgInt2.FromInt(pet.Age()),
		Gender:     r.pgMap.PgText.FromString(pet.Gender().String()),
		Color:      r.pgMap.PgText.FromStringPtr(pet.Color()),
		Microchip:  r.pgMap.PgText.FromStringPtr(pet.Microchip()),
		IsNeutered: r.pgMap.PgBool.FromBool(pet.IsNeutered()),
		CustomerID: pet.CustomerID().Int32(),
		IsActive:   pet.IsActive(),
	}
}

func (r *SqlcPetRepository) toUpdateParams(pet *pet.Pet) *sqlc.UpdatePetParams {
	return &sqlc.UpdatePetParams{
		ID:         pet.ID().Int32(),
		Name:       pet.Name(),
		Photo:      r.pgMap.PgText.FromStringPtr(pet.Photo()),
		Species:    pet.Species().String(),
		Breed:      r.pgMap.PgText.FromStringPtr(pet.Breed()),
		Age:        r.pgMap.PgInt2.FromInt(pet.Age()),
		Gender:     r.pgMap.PgText.FromString(pet.Gender().String()),
		Color:      r.pgMap.PgText.FromStringPtr(pet.Color()),
		Microchip:  r.pgMap.PgText.FromStringPtr(pet.Microchip()),
		IsNeutered: r.pgMap.PgBool.FromBool(pet.IsNeutered()),
		CustomerID: pet.CustomerID().Int32(),
		IsActive:   pet.IsActive(),
	}
}
