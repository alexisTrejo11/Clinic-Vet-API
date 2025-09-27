package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcRowToEntity(sqlPet sqlc.Pet) pet.Pet {
	petBuilder := pet.NewPetBuilder().
		WithID(valueobject.NewPetID(uint(sqlPet.ID))).
		WithCustomerID(valueobject.NewCustomerID(uint(sqlPet.CustomerID))).
		WithName(sqlPet.Name).
		WithSpecies(enum.PetSpecies(sqlPet.Species)).
		WithIsActive(sqlPet.IsActive)

	if sqlPet.Photo.Valid {
		petBuilder = petBuilder.WithPhoto(&sqlPet.Photo.String)
	}

	if sqlPet.Breed.Valid {
		petBuilder = petBuilder.WithBreed(&sqlPet.Breed.String)
	}

	if sqlPet.Age.Valid {
		age := int(sqlPet.Age.Int16)
		petBuilder = petBuilder.WithAge(&age)
	}

	if sqlPet.Gender.Valid && sqlPet.Gender.String != "" {
		petBuilder = petBuilder.WithGender(enum.PetGender(sqlPet.Gender.String))
	}

	if sqlPet.Color.Valid && sqlPet.Color.String != "" {
		petBuilder = petBuilder.WithColor(&sqlPet.Color.String)
	}

	if sqlPet.Microchip.Valid && sqlPet.Microchip.String != "" {
		petBuilder = petBuilder.WithMicrochip(&sqlPet.Microchip.String)
	}

	if sqlPet.IsNeutered.Valid {
		petBuilder = petBuilder.WithIsNeutered(&sqlPet.IsNeutered.Bool)
	}

	return *petBuilder.Build()
}

func sqlcRowsToEntities(rows []sqlc.Pet) []pet.Pet {
	var pets []pet.Pet

	for _, row := range rows {
		petEntity := sqlcRowToEntity(row)
		pets = append(pets, petEntity)
	}

	return pets
}

func ToSqlCreateParam(pet pet.Pet) *sqlc.CreatePetParams {
	return &sqlc.CreatePetParams{
		Name:       pet.Name(),
		Photo:      toPgTypeText(pet.Photo()),
		Species:    pet.Species().String(),
		Breed:      toPgTypeText(pet.Breed()),
		Age:        toPgTypeInt2(pet.Age()),
		Gender:     pgtype.Text{String: pet.Gender().String(), Valid: true},
		Color:      toPgTypeText(pet.Color()),
		Microchip:  toPgTypeText(pet.Microchip()),
		IsNeutered: toPgTypeBool(pet.IsNeutered()),
		CustomerID: int32(pet.CustomerID().Value()),
		IsActive:   pet.IsActive(),
	}
}

func ToSqlUpdateParam(pet *pet.Pet) *sqlc.UpdatePetParams {
	return &sqlc.UpdatePetParams{
		ID:         int32(pet.ID().Value()),
		Name:       pet.Name(),
		Photo:      toPgTypeText(pet.Photo()),
		Species:    pet.Species().String(),
		Breed:      toPgTypeText(pet.Breed()),
		Age:        toPgTypeInt2(pet.Age()),
		Gender:     pgtype.Text{String: pet.Gender().String(), Valid: true},
		Color:      toPgTypeText(pet.Color()),
		Microchip:  toPgTypeText(pet.Microchip()),
		IsNeutered: toPgTypeBool(pet.IsNeutered()),
		CustomerID: int32(pet.CustomerID().Value()),
		IsActive:   pet.IsActive(),
	}
}

func toPgTypeText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func toPgTypeInt2(i *int) pgtype.Int2 {
	if i == nil {
		return pgtype.Int2{Valid: false}
	}
	return pgtype.Int2{Int16: int16(*i), Valid: true}
}

func toPgTypeBool(b *bool) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{Valid: false}
	}
	return pgtype.Bool{Bool: *b, Valid: true}
}

func toPgTypeNumeric(f *float64) pgtype.Numeric {
	var num pgtype.Numeric
	if f == nil {
		num.Valid = false
		return num
	}
	err := num.Scan(*f)
	if err != nil {
		num.Valid = false
		return num
	}
	num.Valid = true
	return num
}
