package db_test

import (
	"context"
	"testing"
	"time"

	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createDummyPet(t *testing.T) sqlc.Pet {
	arg := sqlc.CreatePetParams{
		Name:    "Test Pet",
		Photo:   pgtype.Text{String: "test.jpg", Valid: true},
		Species: "Dog",
		Breed:   pgtype.Text{String: "Labrador", Valid: true},
		Age:     pgtype.Int4{Int32: 3, Valid: true},
		OwnerID: 1,
	}

	pet, err := testQueries.CreatePet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pet)

	require.Equal(t, arg.OwnerID, pet.OwnerID)
	require.Equal(t, arg.Name, pet.Name)
	require.Equal(t, arg.Breed, pet.Breed)
	require.Equal(t, arg.Species, pet.Species)

	require.NotZero(t, pet.ID)
	require.NotZero(t, pet.CreatedAt)

	return pet

}

func TestCreateAccount(t *testing.T) {
	createDummyPet(t)
}

func TestGetPet(t *testing.T) {
	pet1 := createDummyPet(t)
	pet2, err := testQueries.GetPetByID(context.Background(), pet1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pet2)

	require.Equal(t, pet1.Name, pet2.Name)
	require.Equal(t, pet1.OwnerID, pet2.OwnerID)
	require.WithinDuration(t, pet1.CreatedAt.Time, pet1.CreatedAt.Time, time.Second)
}

func TestUpdatePet(t *testing.T) {
	petCreated := createDummyPet(t)

	updateValues := sqlc.UpdatePetParams{
		ID:      petCreated.ID,
		Name:    "Test Pet2",
		Photo:   pgtype.Text{String: "test1.jpg", Valid: true},
		Species: "Cat",
		Breed:   pgtype.Text{String: "Persian", Valid: true},
		Age:     pgtype.Int4{Int32: 4, Valid: true},
		OwnerID: 1,
	}

	err := testQueries.UpdatePet(context.Background(), updateValues)
	require.NoError(t, err)

	updatedPet, err := testQueries.GetPetByID(context.Background(), petCreated.ID)
	require.NoError(t, err)
	require.NotNil(t, updatedPet)

	require.Equal(t, updateValues.Name, updatedPet.Name)
	require.Equal(t, updateValues.Breed, updatedPet.Breed)
	require.Equal(t, updateValues.Species, updatedPet.Species)
	require.Equal(t, updateValues.Age, updatedPet.Age)
	require.Equal(t, petCreated.OwnerID, updatedPet.OwnerID)
	require.Equal(t, petCreated.CreatedAt, updatedPet.CreatedAt)

	require.True(t, updatedPet.UpdatedAt.Time.After(petCreated.UpdatedAt.Time))
}

func TestDeletePet(t *testing.T) {
	petCreated := createDummyPet(t)

	err := testQueries.DeletePet(context.Background(), petCreated.ID)
	require.NoError(t, err)

	_, err = testQueries.GetPetByID(context.Background(), petCreated.ID)
	require.Error(t, err)

}
