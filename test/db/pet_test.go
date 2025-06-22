package db_test

import (
	"context"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomPet(t *testing.T, ownerID int32) sqlc.Pet {
	random_float := randomFloat(1, 50)
	random_rounded := math.Round(random_float*100) / 100

	arg := sqlc.CreatePetParams{
		Name:       "Test Pet " + randomString(5),
		Photo:      pgtype.Text{String: randomString(10) + ".jpg", Valid: true},
		Species:    "Dog",
		Breed:      pgtype.Text{String: "Labrador", Valid: true},
		Age:        pgtype.Int2{Int16: int16(randomInt(1, 15)), Valid: true},
		Gender:     pgtype.Text{String: "Male", Valid: true},
		Weight:     pgtype.Numeric{Int: big.NewInt(int64(random_rounded)), Valid: true},
		Color:      pgtype.Text{String: "Black", Valid: true},
		Microchip:  pgtype.Text{String: randomString(15), Valid: true},
		IsNeutered: pgtype.Bool{Bool: true, Valid: true},
		OwnerID:    ownerID,
		Allergies:  pgtype.Text{String: "None", Valid: true},
		IsActive:   true,
	}

	pet, err := testQueries.CreatePet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pet)

	require.Equal(t, arg.Name, pet.Name)
	require.Equal(t, arg.Species, pet.Species)
	require.Equal(t, arg.OwnerID, pet.OwnerID)

	require.NotZero(t, pet.ID)
	require.NotZero(t, pet.CreatedAt)

	return pet
}

func deleteTestPet(t *testing.T, id int32) {
	err := testQueries.DeletePet(context.Background(), id)
	require.NoError(t, err)
}

func TestCreatePet(t *testing.T) {
	owner := createRandomOwner(t)
	defer deleteTestOwner(t, owner.ID)

	t.Run("CreatePetWithRequiredFields", func(t *testing.T) {
		arg := sqlc.CreatePetParams{
			Name:     "Required Fields Only",
			Species:  "Cat",
			OwnerID:  owner.ID,
			IsActive: true,
		}

		pet, err := testQueries.CreatePet(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, pet)

		require.Equal(t, arg.Name, pet.Name)
		require.Equal(t, arg.Species, pet.Species)
		require.Equal(t, arg.OwnerID, pet.OwnerID)
		require.Equal(t, arg.IsActive, pet.IsActive)

		// Campos opcionales deben ser nulos
		require.False(t, pet.Photo.Valid)
		require.False(t, pet.Breed.Valid)
		require.False(t, pet.Age.Valid)
		require.False(t, pet.Gender.Valid)
	})

	t.Run("CreatePetWithAllFields", func(t *testing.T) {
		pet := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, pet.ID)

		require.True(t, pet.Photo.Valid)
		require.True(t, pet.Breed.Valid)
		require.True(t, pet.Age.Valid)
		require.True(t, pet.Gender.Valid)
		require.True(t, pet.Weight.Valid)
	})

	// Fix
	t.Run("MissingRequiredFields", func(t *testing.T) {
		testCases := []struct {
			name   string
			arg    sqlc.CreatePetParams
			errMsg string
		}{
			// Missing FIELDS not Nil == ""
			{
				name: "MissingOwnerID",
				arg: sqlc.CreatePetParams{
					Name:     "No Owner",
					Species:  "Dog",
					IsActive: true,
				},
				errMsg: "violates foreign key constraint ",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := testQueries.CreatePet(context.Background(), tc.arg)
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errMsg)
			})
		}
	})

	t.Run("DuplicateMicrochip", func(t *testing.T) {
		// Arrange
		microchip := "CHIP-" + randomString(10)

		pet1 := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, pet1.ID)

		err := testQueries.UpdatePet(context.Background(), sqlc.UpdatePetParams{
			ID:        pet1.ID,
			OwnerID:   owner.ID,
			Microchip: pgtype.Text{String: microchip, Valid: true},
		})
		require.NoError(t, err)

		arg := sqlc.CreatePetParams{
			Name:      "Duplicate Microchip",
			Species:   "Dog",
			OwnerID:   owner.ID,
			Microchip: pgtype.Text{String: microchip, Valid: true},
			IsActive:  true,
		}

		// Act
		_, err = testQueries.CreatePet(context.Background(), arg)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "duplicate key value violates unique constraint")
	})
}

func TestGetPetByID(t *testing.T) {
	owner := createRandomOwner(t)
	defer deleteTestOwner(t, owner.ID)

	t.Run("GetExistingPet", func(t *testing.T) {
		// Arrange
		createdPet := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, createdPet.ID)

		// Act
		retrievedPet, err := testQueries.GetPetByID(context.Background(), createdPet.ID)

		// Assert
		require.NoError(t, err)
		require.NotEmpty(t, retrievedPet)

		require.Equal(t, createdPet.ID, retrievedPet.ID)
		require.Equal(t, createdPet.Name, retrievedPet.Name)
		require.Equal(t, createdPet.Species, retrievedPet.Species)
		require.Equal(t, createdPet.OwnerID, retrievedPet.OwnerID)
		require.Equal(t, createdPet.IsActive, retrievedPet.IsActive)
		require.WithinDuration(t, createdPet.CreatedAt.Time, retrievedPet.CreatedAt.Time, time.Second)
		require.WithinDuration(t, createdPet.UpdatedAt.Time, retrievedPet.UpdatedAt.Time, time.Second)

		if createdPet.Photo.Valid {
			require.Equal(t, createdPet.Photo, retrievedPet.Photo)
		}
		if createdPet.Breed.Valid {
			require.Equal(t, createdPet.Breed, retrievedPet.Breed)
		}
	})

	t.Run("GetNonExistentPet", func(t *testing.T) {
		// Arrange
		nonExistentID := int32(999999)

		// Act
		pet, err := testQueries.GetPetByID(context.Background(), nonExistentID)

		// Assert
		require.Error(t, err)
		require.Empty(t, pet)
		require.Contains(t, err.Error(), "no rows in result set")
	})

	t.Run("GetPetWithAllOptionalFields", func(t *testing.T) {
		// Arrange
		arg := sqlc.CreatePetParams{
			Name:               "Test Pet " + randomString(5),
			Photo:              pgtype.Text{String: "test.jpg", Valid: true},
			Species:            "Dog",
			Breed:              pgtype.Text{String: "Labrador", Valid: true},
			Age:                pgtype.Int2{Int16: 3, Valid: true},
			Gender:             pgtype.Text{String: "Male", Valid: true},
			Weight:             pgtype.Numeric{Int: big.NewInt(16), Exp: -2, Valid: true},
			Color:              pgtype.Text{String: "Black", Valid: true},
			Microchip:          pgtype.Text{String: randomString(9), Valid: true},
			IsNeutered:         pgtype.Bool{Bool: true, Valid: true},
			OwnerID:            owner.ID,
			Allergies:          pgtype.Text{String: "None", Valid: true},
			CurrentMedications: pgtype.Text{String: "None", Valid: true},
			SpecialNeeds:       pgtype.Text{String: "None", Valid: true},
			IsActive:           true,
		}

		createdPet, err := testQueries.CreatePet(context.Background(), arg)
		require.NoError(t, err)
		defer deleteTestPet(t, createdPet.ID)

		// Act
		retrievedPet, err := testQueries.GetPetByID(context.Background(), createdPet.ID)

		// Assert
		require.NoError(t, err)

		require.True(t, retrievedPet.Photo.Valid)
		require.Equal(t, arg.Photo.String, retrievedPet.Photo.String)

		require.True(t, retrievedPet.Breed.Valid)
		require.Equal(t, arg.Breed.String, retrievedPet.Breed.String)

		require.True(t, retrievedPet.Age.Valid)
		require.Equal(t, arg.Age.Int16, retrievedPet.Age.Int16)

		require.True(t, retrievedPet.Weight.Valid)
		require.Equal(t, arg.Weight.Int, retrievedPet.Weight.Int)
	})

	t.Run("GetPetWithSomeOptionalFields", func(t *testing.T) {
		// Arrange
		arg := sqlc.CreatePetParams{
			Name:     "Partial Optional",
			Species:  "Cat",
			Breed:    pgtype.Text{String: "Siamese", Valid: true},
			Age:      pgtype.Int2{Valid: false},
			Gender:   pgtype.Text{String: "Female", Valid: true},
			OwnerID:  owner.ID,
			IsActive: true,
		}
		createdPet, err := testQueries.CreatePet(context.Background(), arg)
		require.NoError(t, err)
		defer deleteTestPet(t, createdPet.ID)

		// Act
		retrievedPet, err := testQueries.GetPetByID(context.Background(), createdPet.ID)

		// Assert
		require.NoError(t, err)

		require.True(t, retrievedPet.Breed.Valid)
		require.Equal(t, arg.Breed.String, retrievedPet.Breed.String)
		require.True(t, retrievedPet.Gender.Valid)
		require.Equal(t, arg.Gender.String, retrievedPet.Gender.String)

		require.False(t, retrievedPet.Age.Valid)
		require.False(t, retrievedPet.Photo.Valid)
	})
}

func TestGetPetsByOwnerID(t *testing.T) {
	owner := createRandomOwner(t)
	defer deleteTestOwner(t, owner.ID)

	otherOwner := createRandomOwner(t)
	defer deleteTestOwner(t, otherOwner.ID)

	t.Run("GetPetsForOwnerWithNoPets", func(t *testing.T) {
		pets, err := testQueries.GetPetsByOwnerID(context.Background(), owner.ID)
		require.NoError(t, err)
		require.Empty(t, pets)
		require.Len(t, pets, 0)
	})

	t.Run("GetPetsForOwnerWithMultiplePets", func(t *testing.T) {
		pet1 := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, pet1.ID)

		pet2 := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, pet2.ID)

		pet3 := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, pet3.ID)

		pets, err := testQueries.GetPetsByOwnerID(context.Background(), owner.ID)
		require.NoError(t, err)
		require.NotEmpty(t, pets)
		require.Len(t, pets, 3)

		for _, pet := range pets {
			require.Equal(t, owner.ID, pet.OwnerID)
		}
	})

	t.Run("GetPetsForOwnerWithSinglePet", func(t *testing.T) {
		pet1 := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, pet1.ID)

		pets, err := testQueries.GetPetsByOwnerID(context.Background(), owner.ID)
		require.NoError(t, err)
		require.NotEmpty(t, pets)
		require.Len(t, pets, 1)

		require.Equal(t, pet1.ID, pets[0].ID)
		require.Equal(t, pet1.Name, pets[0].Name)
		require.Equal(t, pet1.OwnerID, pets[0].OwnerID)
	})

	t.Run("IsolationBetweenOwners", func(t *testing.T) {
		ownerPet := createRandomPet(t, owner.ID)
		defer deleteTestPet(t, ownerPet.ID)

		otherOwnerPet := createRandomPet(t, otherOwner.ID)
		defer deleteTestPet(t, otherOwnerPet.ID)

		pets, err := testQueries.GetPetsByOwnerID(context.Background(), owner.ID)
		require.NoError(t, err)
		require.Len(t, pets, 1)
		require.Equal(t, ownerPet.ID, pets[0].ID)

		otherPets, err := testQueries.GetPetsByOwnerID(context.Background(), otherOwner.ID)
		require.NoError(t, err)
		require.Len(t, otherPets, 1)
		require.Equal(t, otherOwnerPet.ID, otherPets[0].ID)
	})

	t.Run("PetsWithDifferentOptionalFields", func(t *testing.T) {
		pet1 := createRandomPetWithParams(t, owner.ID, sqlc.CreatePetParams{
			Name:    "Pet with all fields",
			Species: "Dog",
			Photo:   pgtype.Text{String: "photo1.jpg", Valid: true},
			Breed:   pgtype.Text{String: "Labrador", Valid: true},
			OwnerID: owner.ID,
		})
		defer deleteTestPet(t, pet1.ID)

		pet2 := createRandomPetWithParams(t, owner.ID, sqlc.CreatePetParams{
			Name:    "Pet with minimal fields",
			Species: "Cat",
			OwnerID: owner.ID,
		})
		defer deleteTestPet(t, pet2.ID)

		pets, err := testQueries.GetPetsByOwnerID(context.Background(), owner.ID)
		require.NoError(t, err)
		require.Len(t, pets, 2)

		var fullFieldsPet, minimalFieldsPet sqlc.Pet
		if pets[0].Name == "Pet with all fields" {
			fullFieldsPet = pets[0]
			minimalFieldsPet = pets[1]
		} else {
			fullFieldsPet = pets[1]
			minimalFieldsPet = pets[0]
		}

		require.True(t, fullFieldsPet.Photo.Valid)
		require.True(t, fullFieldsPet.Breed.Valid)

		require.False(t, minimalFieldsPet.Photo.Valid)
		require.False(t, minimalFieldsPet.Breed.Valid)
	})

}

func createRandomPetWithParams(t *testing.T, ownerID int32, params sqlc.CreatePetParams) sqlc.Pet {
	if params.OwnerID == 0 {
		params.OwnerID = ownerID
	}
	if params.Species == "" {
		params.Species = "Dog"
	}
	if params.Name == "" {
		params.Name = "Test Pet " + randomString(5)
	}

	pet, err := testQueries.CreatePet(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, pet)
	return pet
}

// Check Delete When medical histories exists
func TestDeletePet(t *testing.T) {
	owner := createRandomOwner(t)
	defer deleteTestOwner(t, owner.ID)

	t.Run("DeleteExistingPet", func(t *testing.T) {
		pet := createRandomPet(t, owner.ID)

		err := testQueries.DeletePet(context.Background(), pet.ID)
		require.NoError(t, err)

		deletedPet, err := testQueries.GetPetByID(context.Background(), pet.ID)
		require.Error(t, err)
		require.Empty(t, deletedPet)
		require.Contains(t, err.Error(), "no rows in result set")
	})

	t.Run("DeleteNonExistentPet", func(t *testing.T) {
		nonExistentID := int32(999999)
		err := testQueries.DeletePet(context.Background(), nonExistentID)
		require.NoError(t, err)
	})

	t.Run("DeletePetWithRelatedRecords", func(t *testing.T) {
		pet := createRandomPet(t, owner.ID)

		err := testQueries.DeletePet(context.Background(), pet.ID)

		require.NoError(t, err)

		_, err = testQueries.GetPetByID(context.Background(), pet.ID)
		require.Error(t, err)
	})

	t.Run("ConsecutiveDeletes", func(t *testing.T) {
		pet := createRandomPet(t, owner.ID)

		err := testQueries.DeletePet(context.Background(), pet.ID)
		require.NoError(t, err)

		err = testQueries.DeletePet(context.Background(), pet.ID)
		require.NoError(t, err)
	})

}
