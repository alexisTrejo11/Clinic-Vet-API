package repository

import (
	"fmt"

	"clinic-vet-api/app/core/domain/entity/pet"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcRowToEntity(sqlPet sqlc.Pet) (*pet.Pet, error) {
	petID := valueobject.NewPetID(uint(sqlPet.ID))
	customerID := valueobject.NewCustomerID(uint(sqlPet.CustomerID))

	opts := []pet.PetOption{
		pet.WithName(sqlPet.Name),
		pet.WithSpecies(sqlPet.Species),
		pet.WithIsActive(sqlPet.IsActive),
	}

	if sqlPet.Photo.Valid {
		opts = append(opts, pet.WithPhoto(&sqlPet.Photo.String))
	}

	if sqlPet.Breed.Valid && sqlPet.Breed.String != "" {
		opts = append(opts, pet.WithBreed(&sqlPet.Breed.String))
	}

	if sqlPet.Age.Valid {
		age := int(sqlPet.Age.Int16)
		opts = append(opts, pet.WithAge(&age))
	}

	if sqlPet.Gender.Valid && sqlPet.Gender.String != "" {
		gender, err := enum.ParsePetGender(sqlPet.Gender.String)
		if err != nil {
			return nil, fmt.Errorf("invalid gender '%s': %w", sqlPet.Gender.String, err)
		}
		opts = append(opts, pet.WithGender(&gender))
	}

	if sqlPet.Weight.Valid {
		weight := float64(sqlPet.Weight.Int.Int64()) / 100.0 // Ajustar según la precisión de tu BD
		opts = append(opts, pet.WithWeight(&weight))
	}

	if sqlPet.Color.Valid && sqlPet.Color.String != "" {
		opts = append(opts, pet.WithColor(&sqlPet.Color.String))
	}

	if sqlPet.Microchip.Valid && sqlPet.Microchip.String != "" {
		opts = append(opts, pet.WithMicrochip(&sqlPet.Microchip.String))
	}

	if sqlPet.IsNeutered.Valid {
		opts = append(opts, pet.WithIsNeutered(&sqlPet.IsNeutered.Bool))
	}

	if sqlPet.Allergies.Valid && sqlPet.Allergies.String != "" {
		opts = append(opts, pet.WithAllergies(&sqlPet.Allergies.String))
	}

	if sqlPet.CurrentMedications.Valid && sqlPet.CurrentMedications.String != "" {
		opts = append(opts, pet.WithCurrentMedications(&sqlPet.CurrentMedications.String))
	}

	if sqlPet.SpecialNeeds.Valid && sqlPet.SpecialNeeds.String != "" {
		opts = append(opts, pet.WithSpecialNeeds(&sqlPet.SpecialNeeds.String))
	}

	petEntity, err := pet.NewPet(
		petID,
		customerID,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create pet from SQL: %w", err)
	}

	return petEntity, nil
}

func sqlcRowsToEntities(rows []sqlc.Pet) ([]pet.Pet, error) {
	var pets []pet.Pet

	if len(rows) == 0 {
		return []pet.Pet{}, nil
	}

	for _, row := range rows {
		petEntity, err := sqlcRowToEntity(row)
		if err != nil {
			return nil, err
		}
		pets = append(pets, *petEntity)
	}
	return pets, nil
}

func ToSqlCreateParam(pet *pet.Pet) *sqlc.CreatePetParams {
	return &sqlc.CreatePetParams{
		Name:               pet.Name(),
		Photo:              toPgTypeText(pet.Photo()),
		Species:            pet.Species(),
		Breed:              toPgTypeText(pet.Breed()),
		Age:                toPgTypeInt2(pet.Age()),
		Gender:             toPgTypeText((*string)(pet.Gender())),
		Weight:             toPgTypeNumeric(pet.Weight()),
		Color:              toPgTypeText(pet.Color()),
		Microchip:          toPgTypeText(pet.Microchip()),
		IsNeutered:         toPgTypeBool(pet.IsNeutered()),
		CustomerID:         int32(pet.CustomerID().Value()),
		Allergies:          toPgTypeText(pet.Allergies()),
		CurrentMedications: toPgTypeText(pet.CurrentMedications()),
		SpecialNeeds:       toPgTypeText(pet.SpecialNeeds()),
		IsActive:           pet.IsActive(),
	}
}

func ToSqlUpdateParam(pet *pet.Pet) *sqlc.UpdatePetParams {
	return &sqlc.UpdatePetParams{
		ID:                 int32(pet.ID().Value()),
		Name:               pet.Name(),
		Photo:              toPgTypeText(pet.Photo()),
		Species:            pet.Species(),
		Breed:              toPgTypeText(pet.Breed()),
		Age:                toPgTypeInt2(pet.Age()),
		Gender:             toPgTypeText((*string)(pet.Gender())),
		Weight:             toPgTypeNumeric(pet.Weight()),
		Color:              toPgTypeText(pet.Color()),
		Microchip:          toPgTypeText(pet.Microchip()),
		IsNeutered:         toPgTypeBool(pet.IsNeutered()),
		CustomerID:         int32(pet.CustomerID().Value()),
		Allergies:          toPgTypeText(pet.Allergies()),
		CurrentMedications: toPgTypeText(pet.CurrentMedications()),
		SpecialNeeds:       toPgTypeText(pet.SpecialNeeds()),
		IsActive:           pet.IsActive(),
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
